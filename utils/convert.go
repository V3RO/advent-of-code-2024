package utils

import "strconv"

func StringSliceToIntSlice(strings []string) []int {
	ints := make([]int, len(strings))

	for i, s := range strings {
		ints[i], _ = strconv.Atoi(s)
	}

	return ints
}
