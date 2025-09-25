package styx

func (s *Styx) FontFace(fontFamily, fontWeight, fontStyle, url, format string) {
	s.Global(
		WithFontFaceRule(
			WithProp(FontFamily, fontFamily),
			WithProp(FontWeight, fontWeight),
			WithProp(FontStyle, fontStyle),
			WithProp(Src, URL(url), Format(format)),
		),
	)
}
