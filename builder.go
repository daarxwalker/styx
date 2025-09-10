package styx

import (
	"bytes"
	"strings"
)

type Builder struct {
	cfg                 *Config
	classes             []string
	globalStyles        *[]*bytes.Buffer
	styles              []*bytes.Buffer
	seenClasses         map[string]struct{}
	seenStyles          map[uint64]struct{}
	customSelector      string
	hasCustomSelector   bool
	customSelectorOpen  bool
	customSelectorClose bool
}

var (
	classReplaceSymbols = []string{
		" ", `_`,
	}
)

var (
	classReplacer = strings.NewReplacer(classReplaceSymbols...)
)

var (
	styleSelectorEscaper = strings.NewReplacer(
		append(
			classReplaceSymbols,
			".", `\.`,
			"/", `\/`,
		)...,
	)
)

func (b *Builder) DefineVar(name, value string) *Builder {
	tmpStyle := new(bytes.Buffer)
	if b.hasCustomSelector && !b.customSelectorOpen {
		b.openCustomSelector(tmpStyle)
	}
	tmpStyle.WriteString("--")
	tmpStyle.WriteString(name)
	tmpStyle.WriteString(":")
	tmpStyle.WriteString(value)
	tmpStyle.WriteString(";")
	*b.globalStyles = append(*b.globalStyles, tmpStyle)
	b.styles = append(b.styles, tmpStyle)
	return b
}

func (b *Builder) Join(other *Builder) *Builder {
	for _, class := range other.classes {
		b.rememberClass(class)
	}
	for _, style := range other.styles {
		b.rememberStyle(style)
	}
	return b
}

func (b *Builder) Class() string {
	if len(b.customSelector) > 0 {
		return b.customSelector
	}
	if b.hasCustomSelector && !b.customSelectorClose {
		b.closeCustomSelector()
	}
	return strings.TrimSpace(strings.Join(b.classes, " "))
}

func (b *Builder) Finish() {
	if b.hasCustomSelector && !b.customSelectorClose {
		b.closeCustomSelector()
	}
}

func (b *Builder) Style() string {
	result := new(bytes.Buffer)
	if b.hasCustomSelector && !b.customSelectorClose {
		b.closeCustomSelector()
	}
	for _, style := range b.styles {
		result.WriteString(style.String())
	}
	return result.String()
}

func (b *Builder) createClass(class string, modifiers ...Modifier) {
	tmpClass := new(bytes.Buffer)
	for _, modifier := range modifiers {
		tmpClass.WriteString(string(modifier))
		tmpClass.WriteString(":")
	}
	tmpClass.WriteString(classReplacer.Replace(class))
	b.rememberClass(tmpClass.String())
}

func (b *Builder) createStyle(class, style string, modifiers ...Modifier) {
	tmpStyle := new(bytes.Buffer)
	if b.hasCustomSelector && !b.customSelectorOpen {
		b.openCustomSelector(tmpStyle)
	}
	var hasBreakpoint, hasPseudoelement, hasPseudoselector bool
	var pseudoelement, pseudoselector Modifier
	for _, modifier := range modifiers {
		if _, ok := b.cfg.Breakpoints[modifier]; !ok {
			continue
		}
		hasBreakpoint = true
		tmpStyle.WriteString("@media (min-width:")
		tmpStyle.WriteString(b.cfg.Breakpoints[modifier])
		tmpStyle.WriteString(") {")
		break
	}
	if !b.hasCustomSelector {
		tmpStyle.WriteString(".")
		for _, modifier := range modifiers {
			if _, ok := Pseudoselectors[modifier]; ok {
				pseudoselector = modifier
				hasPseudoselector = true
			}
			if _, ok := Pseudoelements[modifier]; ok {
				pseudoelement = modifier
				hasPseudoelement = true
			}
			tmpStyle.WriteString(string(modifier))
			tmpStyle.WriteString(`\:`)
		}
		tmpStyle.WriteString(styleSelectorEscaper.Replace(class))
		tmpStyle.WriteString("{")
	}
	if hasPseudoselector {
		tmpStyle.WriteString("&:")
		tmpStyle.WriteString(string(pseudoselector))
		tmpStyle.WriteString("{")
	}
	if hasPseudoelement {
		tmpStyle.WriteString("&::")
		tmpStyle.WriteString(string(pseudoelement))
		tmpStyle.WriteString("{")
	}
	tmpStyle.WriteString(style)
	if hasPseudoelement {
		tmpStyle.WriteString("}")
	}
	if hasPseudoselector {
		tmpStyle.WriteString("}")
	}
	if !b.hasCustomSelector {
		tmpStyle.WriteString("}")
	}
	if hasBreakpoint {
		tmpStyle.WriteString("}")
	}
	if b.hasCustomSelector {
		*b.globalStyles = append(*b.globalStyles, tmpStyle)
	}
	if !b.hasCustomSelector {
		b.rememberStyle(tmpStyle)
	}
}

func (b *Builder) rememberClass(className string) {
	if _, ok := b.seenClasses[className]; !ok {
		b.seenClasses[className] = struct{}{}
		b.classes = append(b.classes, className)
	}
}

func (b *Builder) rememberStyle(buf *bytes.Buffer) {
	styleBytes := buf.Bytes()
	styleHash := hashBytes(styleBytes)
	if _, ok := b.seenStyles[styleHash]; !ok {
		b.seenStyles[styleHash] = struct{}{}
		*b.globalStyles = append(*b.globalStyles, buf)
	}
	b.styles = append(b.styles, buf)
}

func (b *Builder) openCustomSelector(buf *bytes.Buffer) {
	b.customSelectorOpen = true
	buf.WriteString(b.customSelector)
	buf.WriteString("{")
}

func (b *Builder) closeCustomSelector() {
	b.customSelectorClose = true
	tmpStyle := new(bytes.Buffer)
	tmpStyle.WriteString("}")
	*b.globalStyles = append(*b.globalStyles, tmpStyle)
}
