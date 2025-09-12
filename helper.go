package styx

import (
	"strconv"
)

func Rem(value int) string {
	return strconv.Itoa(value) + "rem"
}

func Px(value int) string {
	return strconv.Itoa(value) + "px"
}

func Percent(value int) string {
	return strconv.Itoa(value) + "%"
}

func Var(name string) string {
	return "var(--" + name + ")"
}
