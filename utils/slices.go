package utils

func Reduce[R any, E any](initial R, elements []E, reducer func(R, E) R) R {
	for _, element := range elements {
		initial = reducer(initial, element)
	}

	return initial
}
