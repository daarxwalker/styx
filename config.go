package styx

type Config struct {
	FontSize    string
	FontFamily  string
	Colors      map[Color]string
	Breakpoints map[Modifier]string
	Shadows     map[Modifier]string
}

const (
	BaseFontSize   = "16px"
	BaseFontFamily = `ui-sans-serif, system-ui, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol', 'Noto Color Emoji'`
)

var (
	BaseBreakpoints = map[Modifier]string{
		Sm: "640px",
		Md: "768px",
		Lg: "1024px",
		Xl: "1280px",
	}
	BaseColors = map[Color]string{
		ColorInherit:      "inherit",
		ColorWhite:        "#ffffff",
		ColorNeutral:      "#34d399",
		ColorSuccess:      "#34d399",
		ColorSuccessHover: "#34d399",
	}
	BaseBoxShadows = map[Modifier]string{
		Sm: "0 1px 3px 0 rgb(0 0 0 / 0.1), 0 1px 2px -1px rgb(0 0 0 / 0.1)",
		Md: "0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1)",
		Lg: "0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1)",
		Xl: "0 20px 25px -5px rgb(0 0 0 / 0.1), 0 8px 10px -6px rgb(0 0 0 / 0.1)",
	}
)

func CreateBaseConfig() *Config {
	return &Config{
		FontSize:    BaseFontSize,
		FontFamily:  BaseFontFamily,
		Colors:      BaseColors,
		Breakpoints: BaseBreakpoints,
		Shadows:     BaseBoxShadows,
	}
}

func (c *Config) Merge(other *Config) {
	c.FontSize = other.FontSize
	for k, v := range other.Colors {
		c.Colors[k] = v
	}
	for k, v := range other.Breakpoints {
		c.Breakpoints[k] = v
	}
	for k, v := range other.Shadows {
		c.Shadows[k] = v
	}
}
