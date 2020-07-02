package advancedapplications

type BitMap struct {
	bytes []uint8
	nbits int
}

func NewBitMap(nbits int) *BitMap {
	return &BitMap{
		nbits: nbits,
		bytes: make([]uint8, nbits>>3+1),
	}
}

func (b *BitMap) Set(k int) {
	if k > b.nbits {
		return
	}
	byteIndex := k >> 3
	bitIndex := k % 8
	b.bytes[byteIndex] |= 1 << uint(bitIndex)
}

func (b *BitMap) Get(k int) bool {
	if k > b.nbits {
		return false
	}
	byteIndex := k >> 3
	bitIndex := k % 8
	return b.bytes[byteIndex]&(1<<uint(bitIndex)) != 0
}
