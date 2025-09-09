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
				BgColor(ColorNeutral).
				PX("1rem").PY("0.5rem")
			instance02.Build().
				Flex().
				TextColor(ColorNeutral).
				PX("1rem").PY("1rem")
			instance01.Join(instance02)
			assert.Equal(
				t,
				`.inline-flex{display:inline-flex;}.bg-neutral{background-color:var(--color-neutral);}.px-1rem{padding-right:1rem;padding-left:1rem;}.py-0\.5rem{padding-top:0.5rem;padding-bottom:0.5rem;}.flex{display:flex;}.text-neutral{color:var(--color-neutral);}.py-1rem{padding-top:1rem;padding-bottom:1rem;}`,
				instance01.String(),
			)
		},
	)
}
