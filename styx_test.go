package styx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStyx(t *testing.T) {
	t.Run(
		"join instances", func(t *testing.T) {
			instance01 := New()
			instance02 := New()
			instance01.Build().
				InlineFlex().
				BgColor(ColorNeutral400).
				PX("1rem").PY("0.5rem")
			instance02.Build().
				Flex().
				TextColor(ColorNeutral400).
				PX("1rem").PY("1rem")
			instance01.Join(instance02)
			assert.Equal(
				t,
				`.inline-flex{display:inline-flex;}.bg-neutral-400{background-color:var(--color-neutral-400);}.px-1rem{padding-right:1rem;padding-left:1rem;}.py-0\.5rem{padding-top:0.5rem;padding-bottom:0.5rem;}.flex{display:flex;}.text-neutral-400{color:var(--color-neutral-400);}.py-1rem{padding-top:1rem;padding-bottom:1rem;}`,
				instance01.String(),
			)
		},
	)
}
