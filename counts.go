package util

func Counts[S []E, E comparable](s S, k E) int {
	h := map[E]int{}
	for _, v := range s {
		h[v] = h[v] + 1
	}
	return h[k]
}
