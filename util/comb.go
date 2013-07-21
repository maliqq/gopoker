package util

func Fact(n int) int {
	fact := 1
	for i := 1; i <= n; i++ {
  	fact *= i
  }

	return fact
}

func Combinations(n, m int) int {
	return Fact(n) / Fact(n-m) / Fact(m)
}

// abc(2) = ab, ac, bc
func Combine(n, m int) [][]int {
	if n < m {
		return [][]int{}
	}

	size := Combinations(n, m)
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
