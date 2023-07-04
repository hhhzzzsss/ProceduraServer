package util

import (
	"math/rand"

	"golang.org/x/exp/constraints"
)

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

func Clamp(x, lower, upper float64) float64 {
	if x < lower {
		return lower
	} else if x > upper {
		return upper
	} else {
		return x
	}
}

func RemoveRandomFromSlice[T any](s *[]T) T {
	selectedIdx := rand.Intn(len(*s))
	lastIdx := len(*s) - 1
	selectedElem := (*s)[selectedIdx]
	(*s)[selectedIdx] = (*s)[lastIdx]
	(*s) = (*s)[:lastIdx]
	return selectedElem
}

func RemoveWeightedRandomFromSlice[T any](s *[]T, weights *[]float32) T {
	if len(*s) != len(*weights) {
		panic("Element slice must have same length as weight slice")
	}

	lastIdx := len(*s) - 1

	var totalWeight float32 = 0
	for _, weight := range *weights {
		totalWeight += weight
	}

	rval := rand.Float32() * totalWeight
	var cumWeight float32 = 0
	for i, weight := range *weights {
		cumWeight += weight
		if cumWeight >= rval {
			selectedElem := (*s)[i]
			(*s)[i] = (*s)[lastIdx]
			(*s) = (*s)[:lastIdx]
			(*weights)[i] = (*weights)[lastIdx]
			(*weights) = (*weights)[:lastIdx]
			return selectedElem
		}
	}
	selectedElem := (*s)[lastIdx]
	(*s) = (*s)[:lastIdx]
	(*weights) = (*weights)[:lastIdx]
	return selectedElem
}

func RemoveFromUnorderedSlice[T any](s *[]T, i int) T {
	lastIdx := len(*s) - 1
	selectedElem := (*s)[i]
	(*s)[i] = (*s)[lastIdx]
	(*s) = (*s)[:lastIdx]
	return selectedElem
}
