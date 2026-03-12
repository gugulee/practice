package bitmap

// BitMap ...
type BitMap struct {
	data      []uint64
	bitsCount int
}

// New return bitmap which contains n bits
func New(n int) *BitMap {
	byteIndex := n >> 6

	return &BitMap{
		data:      make([]uint64, byteIndex+1),
		bitsCount: n,
	}
}

// Set set bit
func (b *BitMap) Set(data int) {
	byteIndex := data >> 6
	var bitIndex uint64 = uint64(data) & 63

	b.data[byteIndex] |= (1 << bitIndex)
}

// Get get bit
func (b *BitMap) Get(data int) bool {
	byteIndex := data >> 6
	var bitIndex uint64 = uint64(data) & 63

	return (b.data[byteIndex] & (1 << bitIndex)) != 0
}

// Clean clean bit
func (b *BitMap) Clean(data int) {
	byteIndex := data >> 6
	var bitIndex uint64 = uint64(data) & 63

	b.data[byteIndex] ^= (1 << bitIndex)
}
