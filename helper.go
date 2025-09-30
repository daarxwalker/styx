package styx

import (
	"strings"
)

func DefineVar(name, value string) Node {
	return WithProp("--"+name, value)
}

func Comma(values ...string) string {
	return strings.Join(values, ", ")
}

func Negate(value string) string {
	return Minus + value
}

func Important(values ...string) []string {
	return append(values, "!important")
}

func OklabOpacity(color string, opacity float64) string {
	return ColorMix("in oklab", color+" "+Percent(opacity), "transparent")
}

func Webkit(prop string) string {
	return "-webkit-" + prop
}

func Moz(prop string) string {
	return "-moz-" + prop
}

func Ms(prop string) string {
	return "-ms-" + prop
}

func O(prop string) string {
	return "-o-" + prop
}
