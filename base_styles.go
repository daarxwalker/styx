package styx

func (s *Styx) CreateBaseStyles() {
	{
		root := s.Build(":root")
		root.DefineVar("font-size", s.cfg.FontSize)
		root.DefineVar("font-family", s.cfg.FontFamily)
		for k, v := range s.cfg.Colors {
			root.DefineVar("color-"+string(k), v)
		}
		for k, v := range s.cfg.Breakpoints {
			root.DefineVar("breakpoint-"+string(k), v)
		}
		for k, v := range s.cfg.Shadows {
			root.DefineVar("shadow-"+string(k), v)
		}
		root.Finish()
	}
	{
		s.Build("html").
			FontSize(Var("font-size")).
			Finish()
	}
	{
		s.Build("body").
			FontFamily(Var("font-family")).
			Finish()
	}
}
