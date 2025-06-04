package utils

import (
	"cmp"
	"slices"
)

func IsEqual[T cmp.Ordered](a, b []T, sorted bool) bool {
	if len(a) != len(b) {
		return false
	}
	if sorted {
		slices.Sort(a)
		slices.Sort(b)
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
