package string

func KnuthMorrisPrattMatch(a, b []rune, n, m int) int {
	next := getNext(b, m)
	j := 0
	for i := 0; i < n; i++ {
		for j > 0 && j < m && a[i] != b[j] {
			j = next[j-1] + 1
		}
		if a[i] == b[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}

func getNext(b []rune, m int) []int {
	next := make([]int, m, m)
	next[0] = -1
	index := -1
	for i := 1; i < m; i++ {
		for index != -1 && b[index+1] != b[i] {
			index = next[index]
		}
		if b[index+1] == b[i] {
			index++
		}
		next[i] = index
	}
	return next
}
