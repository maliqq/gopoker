package util

func Fact(n int64) int64 {
	if n < 0 || n == 0 {
		return 1
	}
	return n * Fact(n-1)
}

func Comb(n int, m int) [][]int {
	if n < m {
		return [][]int{}
	}

	size := Fact(int64(n)) / Fact(int64(n-m)) / Fact(int64(m))
	result := make([][]int, size)

	index := make([]int, m)
	for i := range index {
		index[i] = i
	}

	k := 0

	for {
		result[k] = make([]int, m)
		copy(result[k], index)
		k++

		i := m - 1
		for ; i >= 0 && index[i] == i+n-m; i -= 1 {
		}

		if i < 0 {
			break
		}

		index[i] += 1
		for j := i + 1; j < m; j += 1 {
			index[j] = index[j-1] + 1
		}
	}

	return result
}
