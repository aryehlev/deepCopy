package deepCopy

func CopyPointer[T any](original *T, copier ...func(T) T) *T {
	if original == nil {
		return nil
	}

	var copyOfValue T
	if len(copier) > 0 {
		copyOfValue = copier[0](*original)
	} else {
		copyOfValue = *original
	}

	return &copyOfValue
}

func CopyMap[K comparable, V any](original map[K]V, copier ...func(V) V) map[K]V {
	if original == nil {
		return nil
	}

	copyOfMap := make(map[K]V)

	for key, value := range original {
		if len(copier) > 0 {
			copyOfMap[key] = copier[0](value)
		} else {
			copyOfMap[key] = value
		}
	}

	return copyOfMap
}

func CopySlice[T any](original []T, copier ...func(T) T) []T {
	if original == nil {
		return nil
	}

	var copyOfList = make([]T, len(original), len(original))

	for i := 0; i < len(original); i++ {
		if len(copier) > 0 {
			copyOfList[i] = copier[0](original[i])
		} else {
			copyOfList[i] = original[i]
		}

	}

	return copyOfList
}
