package styx

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	Full       = "100%"
	Auto       = "auto"
	Zero       = "0"
	Minus      = "-"
	None       = "none"
	Current    = "currentColor"
	MaxContent = "max-content"
	Inherit    = "inherit"
	From       = "from"
	To         = "to"
)

const (
	AutoFit  = "auto-fit"
	AutoFill = "auto-fill"
)

const (
	WeightThin       = "100"
	WeightExtraLight = "200"
	WeightLight      = "300"
	WeightNormal     = "400"
	WeightMedium     = "500"
	WeightSemiBold   = "600"
	WeightBold       = "700"
	WeightExtraBold  = "800"
	WeightBlack      = "900"
)

const (
	OutlineSolid  = "solid"
	OutlineDashed = "dashed"
	OutlineDotted = "dotted"
	OutlineDouble = "double"
	OutlineNone   = "none"
	OutlineHidden = "hidden"
)

func Rotate(deg float64) string {
	return fmt.Sprintf("rotate(%gdeg)", deg)
}

func Translate(value string) string {
	return "translate(" + value + ")"
}

func TranslateX(value string) string {
	return "translateX(" + value + ")"
}

func TranslateY(value string) string {
	return "translateY(" + value + ")"
}

func TranslateZ(value string) string {
	return "translateZ(" + value + ")"
}

func Translate3D(x, y, z string) string {
	return "translate3D(" + x + ", " + y + ", " + z + ")"
}

func Scale(v string) string {
	return "scale(" + v + ")"
}

func Scale3D(x, y, z string) string {
	return "scale3d(" + x + ", " + y + ", " + z + ")"
}

func Skew(x, y string) string {
	return "skew(" + x + ", " + y + ")"
}

func SkewX(v string) string {
	return "skewX(" + v + ")"
}

func SkewY(v string) string {
	return "skewY(" + v + ")"
}

func Blur(v string) string {
	return "blur(" + v + ")"
}

func Brightness(v string) string {
	return "brightness(" + v + ")"
}

func Contrast(v string) string {
	return "contrast(" + v + ")"
}

func DropShadow(v string) string {
	return "drop-shadow(" + v + ")"
}

func Grayscale(v string) string {
	return "grayscale(" + v + ")"
}

func HueRotate(v string) string {
	return "hue-rotate(" + v + ")"
}

func Invert(v string) string {
	return "invert(" + v + ")"
}

func OpacityFilter(v string) string {
	return "opacity(" + v + ")"
}

func Saturate(v string) string {
	return "saturate(" + v + ")"
}

func Sepia(v string) string {
	return "sepia(" + v + ")"
}

func ColorMix(interpolationMethod string, color1, color2 string) string {
	return fmt.Sprintf("color-mix(%s, %s, %s)", interpolationMethod, color1, color2)
}

func Calc(expr string) string {
	return "calc(" + expr + ")"
}

func Min(values ...string) string {
	return "min(" + strings.Join(values, ",") + ")"
}

func Max(values ...string) string {
	return "max(" + strings.Join(values, ",") + ")"
}

func Clamp(min, val, max string) string {
	return "clamp(" + min + "," + val + "," + max + ")"
}

func Format(value string) string {
	return "format(" + value + ")"
}

func URL(path string) string {
	return `url("` + path + `")`
}

func LocalFont(name string) string {
	return "local(" + name + ")"
}

func Rem(value float64) string {
	return fmt.Sprintf("%g", value) + "rem"
}

func Px(value float64) string {
	return fmt.Sprintf("%g", value) + "px"
}

func Percent(value float64) string {
	return fmt.Sprintf("%g", value) + "%"
}

func Var(name string) string {
	return "var(--" + name + ")"
}

func Em(value int) string {
	return strconv.Itoa(value) + "em"
}

func VW(value int) string {
	return strconv.Itoa(value) + "vw"
}

func VH(value int) string {
	return strconv.Itoa(value) + "vh"
}

func Vmin(value int) string {
	return strconv.Itoa(value) + "vmin"
}

func Vmax(value int) string {
	return strconv.Itoa(value) + "vmax"
}

func Fr(value int) string {
	return strconv.Itoa(value) + "fr"
}

func RGB(r, g, b int) string {
	return fmt.Sprintf("rgb(%d %d %d)", r, g, b)
}

func RGBA(r, g, b int, a float64) string {
	return fmt.Sprintf("rgb(%d %d %d / %.2f)", r, g, b, a)
}

func HEX(hex string) string {
	if !strings.HasPrefix(hex, "#") {
		return "#" + hex
	}
	return hex
}

func OKLCH(lightness, chroma float64, hue float64, alpha ...float64) string {
	if len(alpha) > 0 {
		return fmt.Sprintf("oklch(%.2f%% %.3f %.1f / %.2f)", lightness, chroma, hue, alpha[0])
	}
	return fmt.Sprintf("oklch(%.2f%% %.3f %.1f)", lightness, chroma, hue)
}

func Minmax(from, to string) string {
	return fmt.Sprintf("minmax(%s,%s)", from, to)
}

func FitContent(value string) string {
	return fmt.Sprintf("fit-content(%s)", value)
}

func Repeat(count string, values ...string) string {
	return fmt.Sprintf("repeat(%s,%s)", count, strings.Join(values, " "))
}
