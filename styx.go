package styx

import (
	"bytes"
)

type Styx struct {
	cfg        *Config
	styles     []*bytes.Buffer
	builders   []*Builder
	seenStyles map[uint64]struct{}
}

func New(cfg ...*Config) *Styx {
	instance := &Styx{
		cfg:        CreateBaseConfig(),
		styles:     make([]*bytes.Buffer, 0),
		seenStyles: make(map[uint64]struct{}),
	}
	if len(cfg) > 0 {
		instance.cfg.Merge(cfg[0])
	}
	return instance
}

func (s *Styx) Build(selector ...string) *Builder {
	var customSelector string
	if len(selector) > 0 {
		customSelector = selector[0]
	}
	builder := &Builder{
		cfg:               s.cfg,
		customSelector:    customSelector,
		hasCustomSelector: len(customSelector) > 0,
		classes:           make([]string, 0),
		styles:            &s.styles,
		seenClasses:       make(map[string]struct{}),
		seenStyles:        s.seenStyles,
	}
	s.builders = append(s.builders, builder)
	return builder
}

func (s *Styx) Join(other *Styx) *Styx {
	s.cfg.Merge(other.cfg)
	for _, style := range other.styles {
		styleBytes := style.Bytes()
		hash := hashBytes(styleBytes)
		if _, ok := s.seenStyles[hash]; !ok {
			s.seenStyles[hash] = struct{}{}
			s.styles = append(s.styles, style)
		}
	}
	return s
}

func (s *Styx) Bytes() []byte {
	return s.createResult().Bytes()
}

func (s *Styx) String() string {
	return s.createResult().String()
}

func (s *Styx) createResult() *bytes.Buffer {
	result := new(bytes.Buffer)
	for _, style := range s.styles {
		result.Write(style.Bytes())
	}
	return result
}
