package string

func RabinKarpMatch_1(p, s string) int {
	pL := len(p)
	sL := len(s)
	num := pL - sL
	strMap := make(map[string]int, num+1)

	for i := 0; i < num; i++ {
		subStr := p[i : i+sL]
		strMap[subStr] = i
	}
	if index, ok := strMap[s]; ok {
		return index
	}
	return -1
}