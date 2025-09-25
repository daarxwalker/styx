package styx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStyx(t *testing.T) {
	t.Run(
		"single property", func(t *testing.T) {
			styler := New()
			selector := styler.Rule(
				WithSelector(Class("test")),
				WithProp("background", "red"),
			)
			assert.Equal(
				t,
				"test",
				selector,
			)
			assert.Equal(
				t,
				".test{background:red;}",
				styler.String(),
			)
		},
	)
	t.Run(
		"two classes", func(t *testing.T) {
			styler := New()
			styler.Rule(
				WithSelector(Class("test01")),
				WithProp("background", "red"),
			)
			styler.Rule(
				WithSelector(Class("test02")),
				WithProp("background", "blue"),
			)
			assert.Equal(
				t,
				".test01{background:red;}.test02{background:blue;}",
				styler.String(),
			)
		},
	)
	t.Run(
		"fragment", func(t *testing.T) {
			styler := New()
			styler.Rule(
				WithSelector(Class("test")),
				WithFragment(
					WithProp("background", "red"),
					WithProp("padding", "1rem"),
				),
			)
			assert.Equal(
				t,
				".test{background:red;padding:1rem;}",
				styler.String(),
			)
		},
	)
	t.Run(
		"nested rule", func(t *testing.T) {
			styler := New()
			styler.Rule(
				WithSelector(Class("test")),
				WithRule(
					WithSelector(Self),
					WithProp("background", "red"),
					WithProp("padding", "1rem"),
				),
			)
			assert.Equal(
				t,
				".test{&{background:red;padding:1rem;}}",
				styler.String(),
			)
		},
	)
	t.Run(
		"@media", func(t *testing.T) {
			styler := New()
			styler.Rule(
				WithSelector(Class("test")),
				WithMediaRule(
					"(min-width:1200px)",
					WithProp("background", "red"),
					WithProp("padding", "1rem"),
				),
			)
			assert.Equal(
				t,
				".test{@media (min-width:1200px){background:red;padding:1rem;}}",
				styler.String(),
			)
		},
	)
	t.Run(
		"@media", func(t *testing.T) {
			styler := New()
			styler.Rule(
				WithSelector(Class("test")),
				WithMediaRule(
					"(min-width:1200px)",
					WithProp("background", "red"),
					WithProp("padding", "1rem"),
				),
			)
			assert.Equal(
				t,
				".test{@media (min-width:1200px){background:red;padding:1rem;}}",
				styler.String(),
			)
		},
	)
	t.Run(
		"hover pseudo selector", func(t *testing.T) {
			styler := New()
			styler.Rule(
				WithSelector(Class("test")),
				WithProp("background", "red"),
				WithRule(
					WithSelector(Self+Hover),
					WithProp("background", "blue"),
				),
			)
			assert.Equal(
				t,
				".test{background:red;&:hover{background:blue;}}",
				styler.String(),
			)
		},
	)
	t.Run(
		"color rgb", func(t *testing.T) {
			styler := New()
			styler.Rule(
				WithSelector(Class("test")),
				WithProp(Color, RGB(255, 0, 0)),
			)
			assert.Equal(
				t,
				".test{color:rgb(255 0 0);}",
				styler.String(),
			)
		},
	)
	t.Run(
		":root", func(t *testing.T) {
			styler := New()
			styler.Global(
				WithRule(
					WithSelector(Root),
					DefineVar("test", "red"),
				),
			)
			assert.Equal(
				t,
				":root{--test:red;}",
				styler.String(),
			)
		},
	)
	t.Run(
		"@font-face", func(t *testing.T) {
			styler := New()
			styler.FontFace(
				"Inter",
				"100 900",
				"normal",
				"/fonts/Inter-VariableFont_opsz,wght.ttf",
				"truetype-variations",
			)
			assert.Equal(
				t,
				`@font-face{font-family:Inter;font-weight:100 900;font-style:normal;src:url("/fonts/Inter-VariableFont_opsz,wght.ttf") format(truetype-variations);}`,
				styler.String(),
			)
		},
	)
	t.Run(
		".btn", func(t *testing.T) {
			styler := New()
			styler.Rule(
				WithSelector(Class("btn")),
				WithProp(BackgroundColor, "blue"),
				WithProp(Border, Px(1), "solid", "gray"),
				WithRule(
					WithSelector(Self+Hover),
					WithProp(BackgroundColor, "red"),
				),
				WithRule(
					WithSelector(Self+Focus),
					WithProp(BorderColor, "black"),
				),
			)
			assert.Equal(
				t,
				".btn{background-color:blue;border:1px solid gray;&:hover{background-color:red;}&:focus{border-color:black;}}",
				styler.String(),
			)
		},
	)
	t.Run(
		"get selector", func(t *testing.T) {
			justifyBetween := GetSelector(
				WithRule(
					WithSelector(Class("is-justify-between")),
					WithProp(JustifyContent, "space-between"),
				),
			)
			assert.Equal(
				t,
				"is-justify-between",
				justifyBetween,
			)
		},
	)
	t.Run(
		"selector modifier", func(t *testing.T) {
			styler := New()
			styler.Rule(
				WithSelector(Class("btn")),
				WithRule(
					WithSelector(Self+Class("btn-primary")),
				),
			)
			assert.Equal(
				t,
				".btn{&.btn-primary{}}",
				styler.String(),
			)
		},
	)
}
