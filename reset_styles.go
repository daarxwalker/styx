package styx

func (s *Styx) ResetStyles() {
	s.Rule(
		WithSelector("*,::before,::after"),
		WithProp(Margin, Zero),
		WithProp(Padding, Zero),
		WithProp(BoxSizing, "border-box"),
	)
	s.Rule(
		WithSelector("html,body"),
		WithProp(Height, Full),
		WithProp(LineHeight, "1.5"),
		WithProp(
			FontFamily,
			`ui-sans-serif,system-ui,sans-serif,'Apple Color Emoji','Segoe UI Emoji','Segoe UI Symbol','Noto Color Emoji'`,
		),
	)
	s.Rule(
		WithSelector("body,h1,h2,h3,h4,p,figure,blockquote,dl,dd"),
		WithProp(Margin, "0"),
	)
	s.Rule(
		WithSelector("ul,ol"),
		WithProp(ListStyle, None),
		WithProp(Margin, Zero),
		WithProp(Padding, Zero),
	)
	s.Rule(
		WithSelector("a"),
		WithProp(TextDecoration, None),
		WithProp(Color, Inherit),
	)
	s.Rule(
		WithSelector("img,video,svg,canvas"),
		WithProp(Display, "block"),
		WithProp(MaxWidth, Full),
	)
	s.Rule(
		WithSelector("button,input,select,textarea"),
		WithProp(Font, Inherit),
		WithProp(Color, Inherit),
		WithProp(Border, None),
		WithProp(Margin, Zero),
		WithProp(Padding, Zero),
	)
	s.Rule(
		WithSelector("button"),
		WithProp(MinWidth, MaxContent),
		WithProp(BackgroundColor, "transparent"),
		WithProp(WhiteSpace, "nowrap"),
		WithProp(Cursor, "pointer"),
	)
	s.Rule(
		WithSelector("table"),
		WithProp(TableLayout, Auto),
		WithProp(TextIndent, Zero),
		WithProp(BorderColor, Inherit),
		WithProp(BorderCollapse, "collapse"),
		WithProp(BorderSpacing, Zero),
	)
	s.Rule(
		WithSelector("textarea"),
		WithProp(Resize, "vertical"),
	)
}
