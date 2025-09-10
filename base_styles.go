package styx

import "fmt"

func (s *Styx) CreateBaseStyles() {
	{
		root := s.Build(":root")
		root.DefineVar("font-size", s.cfg.FontSize)
		root.DefineVar("base-font", string(s.cfg.Font))
		for k, v := range s.cfg.Colors {
			root.DefineVar("color-"+string(k), v)
		}
		for k, v := range s.cfg.Breakpoints {
			root.DefineVar("breakpoint-"+string(k), v)
		}
		for k, v := range s.cfg.Shadows {
			root.DefineVar("shadow-"+string(k), v)
		}
		for k, v := range s.cfg.FontFaces {
			root.DefineVar("font-"+string(k), v.Family)
		}
		root.Finish()
	}
	{
		for _, font := range s.cfg.FontFaces {
			s.Build("@font-face").
				Raw(
					fmt.Sprintf(
						`font-family:"%s";src:url("%s") format("%s");font-weight:%s;font-style:%s;`,
						font.Family,
						font.Path,
						font.Format,
						font.Weight,
						font.Style,
					),
				).
				Finish()
		}
	}
	{
		s.Build("html").
			FontSize(Var("font-size")).
			Finish()
	}
	{
		s.Build("body").
			Raw(`font-family:var(--base-font);`).
			Finish()
	}
}
