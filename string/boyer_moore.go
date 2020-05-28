package string

import (
	"math"
)

func BoyerMooreMatch(a, b []rune, n, m int) int {
	//记录模式串中每个字符最后出现的位置
	hash := generateHash(b, m)
	//后缀匹配预处理
	suffix, prefix := generatePre(b, m)
	//i表示主串与模式串对齐的第一个字符
	i := 0
	num := n - m
	for i <= num {
		j := m - 1
		for ; j >= 0; j-- {
			if a[i+j] != b[j] {
				break
			}
		}
		if j < 0 {
			return i
		}
		//这里等同于将模式串往后滑动j-hash[a[i+j]]位
		x := j - hash[a[i+j]]
		y := 0
		if j < m-1 {
			y = moveByPre(j, m, suffix, prefix)
		}
		ll := int(math.Max(float64(x), float64(y)))
		if ll == 0 {
			i++
		}else{
			i = i + ll
		}

	}
	return -1
}

//构建坏字符规则的hash表,记录模式串中字符位置
func generateHash(b []rune, m int) map[rune]int {
	hash := make(map[rune]int, m)
	for k, v := range b {
		hash[v] = k
	}
	return hash
}

//预处理好后缀规则,返回移动距离
func generatePre(b []rune, m int) (suffix []int, prefix []bool) {
	suffix = make([]int, m, m)
	prefix = make([]bool, m, m)
	for i := 0; i < m; i++ {
		suffix[i] = -1
		prefix[i] = false
	}
	for i := 0; i < m-1; i++ {
		j := i
		l := 0
		for j >= 0 && b[j] == b[m-1-l] {
			j--
			l++
			suffix[l] = j + 1
		}
		if j == -1 {
			prefix[l] = true
		}
	}
	return
}

func moveByPre(j, m int, suffix []int, prefix []bool) int {
	//后缀匹配长度
	l := m - 1 - j
	if suffix[l] != -1 {
		return j - suffix[l] + 1
	}
	num := m - 1
	for i := j + 2; i < num; i++ {
		if prefix[m-i] == true {
			return i
		}
	}
	return m
}
