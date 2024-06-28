package arrays

import (
	"math/rand"
	"slices"
	"time"
)

func Reverse[T any](elements []T) []T {
	result := slices.Clone(elements)

	slices.Reverse(result)

	return result
}

func Shuffle[T any](elements ...T) []T {
	rand.New(rand.NewSource(time.Now().UnixNano())).
		Shuffle(len(elements), func(i int, j int) {
			elements[i], elements[j] = elements[j], elements[i]
		})

	return elements
}

func OneOf[T comparable](target T, elements ...T) bool {
	return slices.Contains(elements, target)
}

// TODO - check this out :)
func Sort[T any](elements []T, comparator func(T, T) int) []T {
	result := slices.Clone(elements)

	slices.SortFunc(result, comparator)

	return result
}

func All[T any](elements []T, predicate func(T) bool) bool {
	return !slices.ContainsFunc(elements, func(element T) bool { return !predicate(element) })
}

func ToMap[T any, K comparable, V any](elements []T, mapper func(element T) (K, V)) map[K]V {
	result := make(map[K]V, len(elements))

	for _, element := range elements {
		key, value := mapper(element)
		result[key] = value
	}

	return result
}

func Map[I, O any](elements []I, transformer func(index int, element I) O) []O {
	result := make([]O, len(elements))

	for index, element := range elements {
		result[index] = transformer(index, element)
	}

	return result
}

func Filter[T any](elements []T, predicate func(T) bool) []T {
	result := make([]T, 0, len(elements))

	for _, element := range elements {
		if predicate(element) {
			result = append(result, element)
		}
	}

	return result
}

func Reduce[T, M any](
	elements []T,
	initial M,
	reducer func(M, T) M,
	abort ...func(M) bool,
) M {
	result := initial

	for i := 0; i < len(elements) && (len(abort) == 0 || !abort[0](result)); i++ {
		result = reducer(result, elements[i])
	}

	return result
}

func Split[T any](elements []T, predicate func(T) bool) ([]T, []T) {
	matched := make([]T, 0, len(elements))
	unmatched := make([]T, 0, len(elements))

	for _, element := range elements {
		if predicate(element) {
			matched = append(matched, element)
		} else {
			unmatched = append(unmatched, element)
		}
	}

	return matched, unmatched
}
