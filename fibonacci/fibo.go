package fibbonaci

func NthFibo() func(n int) []int {
	m := make([]int, 2)
	m[0] = 0
	m[1] = 1
	i := 2
	return func(n int) []int {
		if n == 0 {
			return make([]int, 0)
		}
		for ; i < n; i++ {
			m = append(m, (m[i-2] + m[i-1]))

		}
		return m[0:n]
	}

}
