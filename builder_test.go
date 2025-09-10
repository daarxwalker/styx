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
				BgColor(ColorSuccess400).
				PX("1rem").PY("0.5rem")
			assert.Equal(
				t,
				`lg:inline-flex bg-success-400 px-1rem py-0.5rem`,
				button.Class(),
			)
			assert.Equal(
				t,
				`@media (min-width:1024px) {.lg\:inline-flex{display:inline-flex;}}.bg-success-400{background-color:var(--color-success-400);}.px-1rem{padding-right:1rem;padding-left:1rem;}.py-0\.5rem{padding-top:0.5rem;padding-bottom:0.5rem;}`,
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
				BgColor(ColorNeutral400).
				MX("1rem").MY("0.5rem")
			part01.Join(part02)
			assert.Equal(
				t,
				`flex px-1rem py-0.5rem bg-neutral-400 mx-1rem my-0.5rem`,
				part01.Class(),
			)
			assert.Equal(
				t,
				`.flex{display:flex;}.px-1rem{padding-right:1rem;padding-left:1rem;}.py-0\.5rem{padding-top:0.5rem;padding-bottom:0.5rem;}.bg-neutral-400{background-color:var(--color-neutral-400);}.mx-1rem{margin-right:1rem;margin-left:1rem;}.my-0\.5rem{margin-top:0.5rem;margin-bottom:0.5rem;}`,
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
				BgColor(ColorNeutral400).
				Finish()
			assert.Equal(
				t,
				`*{padding:0;margin:0;box-sizing:border-box;}.test{background-color:var(--color-neutral-400);}`,
				instance.String(),
			)
		},
	)
	t.Run(
		"example navbar", func(t *testing.T) {
			styler := New()
			styler.Build().
				Flex().ItemsCenter().JustifyBetween().Gap("1rem").
				BgColor(ColorWhite).
				PX("1rem").
				H("3rem")
			assert.Equal(
				t,
				`.flex{display:flex;}.items-center{align-items:center;}.justify-between{justify-content:space-between;}.gap-1rem{gap:1rem;}.bg-white{background-color:var(--color-white);}.px-1rem{padding-right:1rem;padding-left:1rem;}.h-3rem{height:3rem;}`,
				styler.String(),
			)
		},
	)
	t.Run(
		"example table row", func(t *testing.T) {
			styler := New()
			styler.Build().
				BorderBottom("1px").
				BorderColor(ColorNeutral400).
				Border("0", LastOfType)
			assert.Equal(
				t,
				`.border-b-1px{border-bottom-width:1px;}.border-neutral-400{border-color:var(--color-neutral-400);}.last-of-type\:border-0{&:last-of-type{border-width:0;}}`,
				styler.String(),
			)
		},
	)
	t.Run(
		"example form field", func(t *testing.T) {
			styler := New()
			styler.Build().
				Flex().FlexCol().Gap("1rem").MT("1rem")
			assert.Equal(
				t,
				`.flex{display:flex;}.flex-col{flex-direction:column;}.gap-1rem{gap:1rem;}.mt-1rem{margin-top:1rem;}`,
				styler.String(),
			)
		},
	)
	t.Run(
		"example checkbox", func(t *testing.T) {
			styler := New()
			{
				wrapperStyle := styler.Build().
					Relative().Grid().PlaceItemsCenter().Cursor("pointer")
				assert.Equal(
					t,
					"relative grid place-items-center cursor-pointer",
					wrapperStyle.Class(),
				)
				assert.Equal(
					t,
					`.relative{position:relative;}.grid{display:grid;}.place-items-center{place-items:center;}.cursor-pointer{cursor:pointer;}`,
					wrapperStyle.Style(),
				)
			}
			{
				iconWrapperStyle := styler.Build().
					Absolute().Inset("0").M("auto").Grid().PlaceItemsCenter()
				assert.Equal(
					t,
					"absolute inset-0 m-auto grid place-items-center",
					iconWrapperStyle.Class(),
				)
				assert.Equal(
					t,
					`.absolute{position:absolute;}.inset-0{top:0;right:0;bottom:0;left:0;}.m-auto{margin:auto;}.grid{display:grid;}.place-items-center{place-items:center;}`,
					iconWrapperStyle.Style(),
				)
			}
			{
				checkboxStyle := styler.Build().
					Appearance("none").
					W("1rem").H("1rem").
					Border("1px").BorderColor(ColorNeutral400).Rounded("1rem").
					BgColor(ColorWhite).
					Cursor("pointer").
					BgColor(ColorNeutral900, Checked).
					BorderColor(ColorNeutral900, Checked)
				assert.Equal(
					t,
					"appearance-none w-1rem h-1rem border-1px border-neutral-400 rounded-1rem bg-white cursor-pointer checked:bg-neutral-900 checked:border-neutral-900",
					checkboxStyle.Class(),
				)
				assert.Equal(
					t,
					`.appearance-none{appearance:none;}.w-1rem{width:1rem;}.h-1rem{height:1rem;}.border-1px{border-width:1px;}.border-neutral-400{border-color:var(--color-neutral-400);}.rounded-1rem{border-radius:1rem;}.bg-white{background-color:var(--color-white);}.cursor-pointer{cursor:pointer;}.checked\:bg-neutral-900{&:checked{background-color:var(--color-neutral-900);}}.checked\:border-neutral-900{&:checked{border-color:var(--color-neutral-900);}}`,
					checkboxStyle.Style(),
				)
			}
			{
				assert.Equal(
					t,
					`.relative{position:relative;}.grid{display:grid;}.place-items-center{place-items:center;}.cursor-pointer{cursor:pointer;}.absolute{position:absolute;}.inset-0{top:0;right:0;bottom:0;left:0;}.m-auto{margin:auto;}.appearance-none{appearance:none;}.w-1rem{width:1rem;}.h-1rem{height:1rem;}.border-1px{border-width:1px;}.border-neutral-400{border-color:var(--color-neutral-400);}.rounded-1rem{border-radius:1rem;}.bg-white{background-color:var(--color-white);}.checked\:bg-neutral-900{&:checked{background-color:var(--color-neutral-900);}}.checked\:border-neutral-900{&:checked{border-color:var(--color-neutral-900);}}`,
					styler.String(),
				)
			}
		},
	)
}
