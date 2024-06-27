package maps

import (
	"reflect"
)

func Keys[K comparable, V any](m map[K]V) []K {
	result := make([]K, 0, len(m))

	for key := range m {
		result = append(result, key)
	}

	return result
}

func Values[K comparable, V any](m map[K]V) []V {
	result := make([]V, 0, len(m))

	for _, value := range m {
		result = append(result, value)
	}

	return result
}

func ContainsKey[K comparable, V any](m map[K]V, k K) bool {
	_, ok := m[k]

	return ok
}

func Remap[K1, K2 comparable, V1, V2 any](m map[K1]V1, remapper func(k K1, v V1) (K2, V2)) map[K2]V2 {
	result := make(map[K2]V2, len(m))

	for key, value := range m {
		if rKey, rValue := remapper(key, value); !reflect.ValueOf(rKey).IsZero() {
			result[rKey] = rValue
		}
	}

	return result
}

func ComputeIfAbsent[K comparable, V any](m map[K]V, k K, supplier func(k K) V) V {
	if _, ok := m[k]; !ok {
		m[k] = supplier(k)
	}

	return m[k]
}

func Merge[K comparable, V any](parts ...map[K]V) map[K]V {
	result := make(map[K]V)

	for _, part := range parts {
		for key, value := range part {
			result[key] = value
		}
	}

	return result
}
