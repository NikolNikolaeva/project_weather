package utils

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))

	for _, v := range m {
		values = append(values, v)
	}

	return values
}

func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	result := make(map[K]V, len(maps))

	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}

	return result
}
