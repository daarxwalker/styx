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

func Webkit(prop string, values ...string) Node {
	return WithProp("-webkit-"+prop, values...)
}

func Moz(prop string, values ...string) Node {
	return WithProp("-moz-"+prop, values...)
}

func Ms(prop string, values ...string) Node {
	return WithProp("-ms-"+prop, values...)
}

func O(prop string, values ...string) Node {
	return WithProp("-o-"+prop, values...)
}
