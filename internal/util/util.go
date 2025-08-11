package util

import "strconv"

func Map[M, V any](models []M, fn func(M) V) []V {
	results := make([]V, len(models))
	for i, m := range models {
		results[i] = fn(m)
	}
	return results
}

func StringToUint(s string) (uint, error) {
	val, err := strconv.ParseUint(s, 10, 32) // base 10, 32 bits
	if err != nil {
		return 0, err
	}
	return uint(val), nil
}
