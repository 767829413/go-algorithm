package greedy

import (
	"testing"
)

func TestNewHuffManTree(t *testing.T) {
	str := []rune(`对输入的英文大写字母进行统计概率 然后构建哈夫曼树,输出是按照概率降序排序输出Huffman编码`)
	h := NewHuffManTree(str)
	h.Print()
}
