package styx

import "hash/fnv"

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
