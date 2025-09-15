package util

import (
	"iter"
	"math/bits"
)

func div_word_size(x int) int {
	return x >> 0x6
}

func mod_word_size(x int) int {
	return x & 0x3F
}

func mul_word_size(x int) int {
	return x << 0x6
}

type BitSet struct {
	buf []uint64
}

func NewBitSet() *BitSet {
	return &BitSet{
		buf: make([]uint64, 0),
	}
}

func NewBitSetWith(len int) *BitSet {
	return &BitSet{
		buf: make([]uint64, len),
	}
}

func (self *BitSet) Incl(x int) {
	w, b := div_word_size(x), mod_word_size(x)
	if w >= len(self.buf) {
		n := w + 1 - len(self.buf)
		s := make([]uint64, n)
		self.buf = append(self.buf, s...)
	}
	self.buf[w] |= uint64(1) << b
}

func (self *BitSet) In(x int) bool {
	w, b := div_word_size(x), mod_word_size(x)
	if w >= len(self.buf) {
		return false
	}
	return (self.buf[w] & (uint64(1) << b)) != 0
}

func (self *BitSet) Excl(x int) {
	w, b := div_word_size(x), mod_word_size(x)
	if w >= len(self.buf) {
		return
	}
	self.buf[w] &^= uint64(1) << b
}

func (self *BitSet) Iter() iter.Seq[int] {
	return func(yield func(int) bool) {
		for i, w := range self.buf {
			for w != 0 {
				k := bits.TrailingZeros64(w)
				j := mul_word_size(i) + k
				if !yield(j) {
					return
				}
				w &^= uint64(1) << k
			}
		}
	}
}

func (self *BitSet) Card() int {
	sum := 0
	for i := 0; i < len(self.buf); i++ {
		sum += bits.OnesCount64(self.buf[i])
	}
	return sum
}

func (self *BitSet) Copy() *BitSet {
	result := NewBitSetWith(len(self.buf))
	copy(result.buf, self.buf)
	return result
}

func Intersect(x, y *BitSet) *BitSet {
	if len(x.buf) > len(y.buf) {
		x.buf = x.buf[:len(y.buf)]
	}
	for i := range x.buf {
		x.buf[i] &= y.buf[i]
	}
	return x
}

func Union(x, y *BitSet) *BitSet {
	if len(y.buf) > len(x.buf) {
		n := len(y.buf) - len(x.buf)
		s := make([]uint64, n)
		x.buf = append(x.buf, s...)
	}
	for i := 0; i < len(y.buf); i++ {
		x.buf[i] |= y.buf[i]
	}
	return x
}

func SymDiff(x, y *BitSet) *BitSet {
	if len(y.buf) > len(x.buf) {
		n := len(y.buf) - len(x.buf)
		s := make([]uint64, n)
		x.buf = append(x.buf, s...)
	}
	for i := 0; i < len(y.buf); i++ {
		x.buf[i] ^= y.buf[i]
	}
	return x
}

func Diff(x, y *BitSet) *BitSet {
	if len(y.buf) < len(x.buf) {
		for i := 0; i < len(y.buf); i++ {
			x.buf[i] &^= y.buf[i]
		}
	} else {
		for i := 0; i < len(x.buf); i++ {
			x.buf[i] &^= y.buf[i]
		}
	}
	return x
}
