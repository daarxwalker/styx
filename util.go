package styx

import (
	"hash/fnv"
	"strings"
)

func Var(name string) string {
	return "var(--" + name + ")"
}

func hashBytes(data []byte) uint64 {
	h := fnv.New64a()
	_, err := h.Write(data)
	if err != nil {
		return 0
	}
	return h.Sum64()
}

func detectNegativeValue(value string) bool {
	return strings.HasPrefix(value, "-")
}

func negateClass(className, value string) string {
	return "-" + strings.TrimSuffix(className, "-") + value
}
