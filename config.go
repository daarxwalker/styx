package styx

type Config struct {
	FontSize    string
	Font        Font
	FontFaces   map[Font]FontFace
	Colors      map[Color]string
	Breakpoints map[Modifier]string
	Shadows     map[Modifier]string
}

const (
	BaseFontSize = "16px"
	BaseFont     = Font(`ui-sans-serif, system-ui, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol', 'Noto Color Emoji'`)
)

var (
	BaseBreakpoints = map[Modifier]string{
		Sm: "640px",
		Md: "768px",
		Lg: "1024px",
		Xl: "1280px",
	}
	BaseColors = map[Color]string{
		ColorInherit: "inherit",
		ColorWhite:   "#ffffff",

		ColorNeutral50:  "#f8fafc",
		ColorNeutral100: "#f1f5f9",
		ColorNeutral200: "#e2e8f0",
		ColorNeutral300: "#cbd5e1",
		ColorNeutral400: "#94a3b8",
		ColorNeutral500: "#64748b",
		ColorNeutral600: "#475569",
		ColorNeutral700: "#334155",
		ColorNeutral800: "#1e293b",
		ColorNeutral900: "#0f172a",

		ColorSuccess50:  "#ecfdf5",
		ColorSuccess100: "#d1fae5",
		ColorSuccess200: "#a7f3d0",
		ColorSuccess300: "#6ee7b7",
		ColorSuccess400: "#34d399",
		ColorSuccess500: "#10b981",
		ColorSuccess600: "#059669",
		ColorSuccess700: "#047857",
		ColorSuccess800: "#065f46",
		ColorSuccess900: "#064e3b",

		ColorDanger50:  "#fef2f2",
		ColorDanger100: "#fee2e2",
		ColorDanger200: "#fecaca",
		ColorDanger300: "#fca5a5",
		ColorDanger400: "#f87171",
		ColorDanger500: "#ef4444",
		ColorDanger600: "#dc2626",
		ColorDanger700: "#b91c1c",
		ColorDanger800: "#991b1b",
		ColorDanger900: "#7f1d1d",
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
		Font:        BaseFont,
		Colors:      BaseColors,
		Breakpoints: BaseBreakpoints,
		Shadows:     BaseBoxShadows,
		FontFaces:   make(map[Font]FontFace),
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
