package mathbeta

type BloomFilter struct {
	bits   []byte
	hashes []func(string) int
}

func NewBloomFilter(bits []byte, hashes []func(string) int) *BloomFilter {
	return &BloomFilter{bits: bits, hashes: hashes}
}

func (bf *BloomFilter) Add(els ...string) {
	if len(bf.hashes) > 0 {
		for _, el := range els {
			for i := 0; i < len(bf.hashes); i++ {
				index := bf.hashes[i](el)
				j := index / 8
				k := uint(index % 8)
				n := byte(1 << k)
				bf.bits[j] = bf.bits[j] | n
			}
		}
	}
}

func (bf *BloomFilter) Contains(el string) bool {
	if el != "" && len(bf.hashes) > 0 {
		for i := 0; i < len(bf.hashes); i++ {
			index := bf.hashes[i](el)
			j := index / 8
			k := uint(index % 8)
			n := byte(1 << k)
			m := bf.bits[j] & n
			if m == 0 {
				return false
			}
		}
	}
	return true
}

func SimpleHash(cap, seed int) func(string) int {
	return func(el string) int {
		result := 0
		length := len(el)
		for i := 0; i < length; i++ {
			result = seed*result + int(el[i])
		}
		return (cap - 1) & result
	}
}
