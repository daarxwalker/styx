package styx

import "strings"

const (
	Active               = ":active"
	ActiveViewTransition = ":active-view-transition"
	AnyLink              = ":any-link"
	Autofill             = ":autofill"
	Blank                = ":blank"
	Buffering            = ":buffering"
	Checked              = ":checked"
	Default              = ":default"
	Defined              = ":defined"
	Disabled             = ":disabled"
	Empty                = ":empty"
	Enabled              = ":enabled"
	First                = ":first"
	FirstChild           = ":first-child"
	FirstOfType          = ":first-of-type"
	Focus                = ":focus"
	FocusVisible         = ":focus-visible"
	FocusWithin          = ":focus-within"
	Fullscreen           = ":fullscreen"
	Future               = ":future"
	Hover                = ":hover"
	InRange              = ":in-range"
	Indeterminate        = ":indeterminate"
	Invalid              = ":invalid"
	LastChild            = ":last-child"
	LastOfType           = ":last-of-type"
	LeftPseudo           = ":left"
	Link                 = ":link"
	Modal                = ":modal"
	Muted                = ":muted"
	OnlyChild            = ":only-child"
	OnlyOfType           = ":only-of-type"
	Open                 = ":open"
	Optional             = ":optional"
	OutOfRange           = ":out-of-range"
	Past                 = ":past"
	Paused               = ":paused"
	PictureInPicture     = ":picture-in-picture"
	PlaceholderShown     = ":placeholder-shown"
	Playing              = ":playing"
	PopoverOpen          = ":popover-open"
	ReadOnly             = ":read-only"
	ReadWrite            = ":read-write"
	Required             = ":required"
	RightPseudo          = ":right"
	Root                 = ":root"
	Scope                = ":scope"
	Seeking              = ":seeking"
	Stalled              = ":stalled"
	Target               = ":target"
	TargetCurrent        = ":target-current"
	UserInvalid          = ":user-invalid"
	UserValid            = ":user-valid"
	Valid                = ":valid"
	Visited              = ":visited"
	VolumeLocked         = ":volume-locked"
)

func ActiveViewTransitionType(t string) string {
	return ":active-view-transition-type(" + t + ")"
}

func Dir(direction string) string {
	return ":dir(" + direction + ")"
}

func Has(selector string) string {
	return ":has(" + selector + ")"
}

func Heading(level string) string {
	return ":heading(" + level + ")"
}

func Host(sel ...string) string {
	if len(sel) == 0 {
		return ":host"
	}
	return ":host(" + strings.Join(sel, ",") + ")"
}

func HostContext(sel string) string {
	return ":host-context(" + sel + ")"
}

func Is(selectors ...string) string {
	return ":is(" + strings.Join(selectors, ",") + ")"
}

func Lang(lang string) string {
	return ":lang(" + lang + ")"
}

func Not(selectors ...string) string {
	return ":not(" + strings.Join(selectors, ",") + ")"
}

func NthChild(expr string) string {
	return ":nth-child(" + expr + ")"
}

func NthLastChild(expr string) string {
	return ":nth-last-child(" + expr + ")"
}

func NthLastOfType(expr string) string {
	return ":nth-last-of-type(" + expr + ")"
}

func NthOfType(expr string) string {
	return ":nth-of-type(" + expr + ")"
}

func State(name string) string {
	return ":state(" + name + ")"
}

func Where(selectors ...string) string {
	return ":where(" + strings.Join(selectors, ",") + ")"
}
