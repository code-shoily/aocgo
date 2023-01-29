package seq

// Permutations returns a list of all permutations of an array
func Permutations[T any](array []T) (result [][]T) {
	var permutationUtil func(array []T, n int, result *[][]T)

	permutationUtil = func(array []T, n int, result *[][]T) {
		if n == 1 {
			dst := make([]T, len(array))
			copy(dst, array[:])
			*result = append(*result, dst)
		} else {
			for i := 0; i < n; i++ {
				permutationUtil(array, n-1, result)
				if n%2 == 0 {
					array[0], array[n-1] = array[n-1], array[0]
				} else {
					array[i], array[n-1] = array[n-1], array[i]
				}
			}
		}
	}

	permutationUtil(array, len(array), &result)

	return result
}
