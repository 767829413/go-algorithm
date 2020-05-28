package string

func BruteForceMatch(primary, secondary string) int {
	pLen := len(primary)
	sLen := len(secondary)
	if pLen < 1 || sLen < 1 || pLen < sLen {
		return -1
	}
	num := pLen - sLen
	for i := 0; i < num; i++ {
		subStr := primary[i : i+sLen]
		if subStr == secondary {
			return i
		}
	}
	return -1
}
