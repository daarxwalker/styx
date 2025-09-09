package styx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	t.Run(
		"create styles", func(t *testing.T) {
			css := New()
			button := css.Build().
				InlineFlex(Lg).
				BgColor(ColorSuccess).
				PX("1rem").PY("0.5rem")
			assert.Equal(
				t,
				`lg:inline-flex bg-success px-1rem py-0.5rem`,
				button.Class(),
			)
			assert.Equal(
				t,
				`@media (min-width:1024px) {.lg\:inline-flex{display:inline-flex;}}.bg-success{background-color:var(--color-success);}.px-1rem{padding-right:1rem;padding-left:1rem;}.py-0\.5rem{padding-top:0.5rem;padding-bottom:0.5rem;}`,
				button.Style(),
			)
		},
	)
	t.Run(
		"join styles", func(t *testing.T) {
			instance := New()
			part01 := instance.Build().
				Flex().
				PX("1rem").PY("0.5rem")
			part02 := instance.Build().
				BgColor(ColorNeutral).
				MX("1rem").MY("0.5rem")
			part01.Join(part02)
			assert.Equal(
				t,
				`flex px-1rem py-0.5rem bg-neutral mx-1rem my-0.5rem`,
				part01.Class(),
			)
			assert.Equal(
				t,
				`.flex{display:flex;}.px-1rem{padding-right:1rem;padding-left:1rem;}.py-0\.5rem{padding-top:0.5rem;padding-bottom:0.5rem;}.bg-neutral{background-color:var(--color-neutral);}.mx-1rem{margin-right:1rem;margin-left:1rem;}.my-0\.5rem{margin-top:0.5rem;margin-bottom:0.5rem;}`,
				instance.String(),
			)
		},
	)
	t.Run(
		"custom selector style", func(t *testing.T) {
			instance := New()
			wildcardStyle := instance.Build("*").
				P("0").
				M("0").
				BoxBorder().
				Style()
			assert.Equal(
				t,
				`*{padding:0;margin:0;box-sizing:border-box;}`,
				wildcardStyle,
			)
		},
	)
	t.Run(
		"multiple custom selector style", func(t *testing.T) {
			instance := New()
			instance.Build("*").
				P("0").
				M("0").
				BoxBorder().
				Finish()
			instance.Build(".test").
				BgColor(ColorNeutral).
				Finish()
			assert.Equal(
				t,
				`*{padding:0;margin:0;box-sizing:border-box;}.test{background-color:var(--color-neutral);}`,
				instance.String(),
			)
		},
	)
}
