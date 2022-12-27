package util

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](elems ...T) T {
	if len(elems) == 0 {
		panic("Cannot take the max of 0 values")
	}
	maxElem := elems[0]
	for i := 1; i < len(elems); i++ {
		if elems[i] > maxElem {
			maxElem = elems[i]
		}
	}
	return maxElem
}
