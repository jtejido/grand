package source32

import (
	"math"
)

const (
	mt19937_n                      = 624
	mt19937_m                      = 397
	mt19937_upper_mask      uint32 = 0x80000000
	mt19937_lower_mask      uint32 = 0x7fffffff
	mt19937_upper_mask_long uint64 = 0x80000000
	mt19937_lower_mask_long uint64 = 0x7fffffff
	mt19937_mask_long       uint64 = 0xffffffff
)

var (
	mt19937_mult_matrix_a = [...]uint32{0x0, 0x9908b0df}
)

// Implements a powerful pseudo-random number generator
// developed by Makoto Matsumoto and Takuji Nishimura during
// 1996-1997.
//
// This generator features an extremely long period
// (2^19937-1) and 623-dimensional equidistribution up to
// 32 bits accuracy.
// http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/emt.html
//
// This generator is described in a paper by Makoto Matsumoto and
// Takuji Nishimura in 1998:
// http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/ARTICLES/mt.pdf
// Mersenne Twister: A 623-Dimensionally Equidistributed Uniform Pseudo-Random
// Number Generator,
// ACM Transactions on Modeling and Computer Simulation, Vol. 8, No. 1,
// January 1998, pp 3--30
//
// 2002-01-26 version of the generator written in C by Makoto Matsumoto
// and Takuji Nishimura.
// http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/MT2002/emt19937ar.html
type MT19937 struct {
	baseSource32
	state [624]uint32
	index uint32
}

func NewMT19937FromStream(seeds []uint32) (*MT19937, error) {
	err := checkEmptySeed(seeds)
	if err != nil {
		return nil, err
	}

	ans := new(MT19937)
	ans.spi = ans
	ans.setSeed(seeds)
	return ans, nil
}

func NewMT19937(seed int64) *MT19937 {
	ans := new(MT19937)
	ans.spi = ans
	ans.Seed(seed)
	return ans
}

func (mt *MT19937) setSeed(stream []uint32) {
	seed := stream
	if len(stream) == 0 {
		seed = make([]uint32, 1)
	}

	mt.stream = make([]uint32, mt19937_n)

	mt.initialiseStream()

	nextIndex := mt.mixSeedAndStream(seed)

	mt.mixStream(nextIndex)

	mt.stream[0] = uint32(mt19937_upper_mask_long)

	mt.Restart()
}

func (mt *MT19937) initialiseStream() {
	m := 19650218 & mt19937_mask_long
	mt.stream[0] = uint32(m)
	for i := 1; i < len(mt.stream); i++ {
		m = (1812433253*(m^(m>>30)) + uint64(i)) & mt19937_mask_long
		mt.stream[i] = uint32(m)
	}
}

func (mt *MT19937) mixSeedAndStream(seed []uint32) int {
	stateSize := len(mt.stream)

	i := 1
	var j uint64

	for k := math.Max(float64(stateSize), float64(len(seed))); k > 0; k-- {
		a := uint64(mt.stream[i])
		b := uint64(mt.stream[i-1])
		c := (a ^ ((b ^ (b >> 30)) * 1664525)) + uint64(seed[j]) + j
		mt.stream[i] = uint32(c)
		i++
		j++
		if i >= stateSize {
			mt.stream[0] = mt.stream[stateSize-1]
			i = 1
		}
		if j >= uint64(len(seed)) {
			j = 0
		}
	}
	return i
}

func (mt *MT19937) mixStream(startIndex int) {
	stateSize := len(mt.stream)

	i := startIndex
	for k := stateSize - 1; k > 0; k-- {
		a := uint64(mt.stream[i])
		b := uint64(mt.stream[i-1])
		c := (a ^ ((b ^ (b >> 30)) * 1566083941)) - uint64(i)
		mt.stream[i] = uint32(c)
		i++
		if i >= stateSize {
			mt.stream[0] = mt.stream[stateSize-1]
			i = 1
		}
	}
}

func (mt *MT19937) Restart() {
	for i := 0; i < mt19937_n; i++ {
		mt.state[i] = mt.stream[i]
	}
	mt.index = mt19937_n
	mt.resetState()
}

func (mt *MT19937) Seed(seed int64) {
	seeds := make([]uint32, mt19937_n)
	seeder.Seed(seed)
	var i int
	// Fill the remaining pairs
	for i < mt19937_n {
		v := seeder.Uint32()
		seeds[i] = v
		i++
	}

	// Initialize the pool content.
	mt.setSeed(seeds)
}

func (mt *MT19937) Uint32() uint32 {
	var y uint32
	if mt.index >= mt19937_n {
		var kk int

		for kk < mt19937_n-mt19937_m {
			y = (mt.state[kk] & mt19937_upper_mask) | (mt.state[kk+1] & mt19937_lower_mask)
			mt.state[kk] = mt.state[kk+mt19937_m] ^ (y >> 1) ^ mt19937_mult_matrix_a[y&0x1]
			kk++
		}
		for kk < mt19937_n-1 {
			y = (mt.state[kk] & mt19937_upper_mask) | (mt.state[kk+1] & mt19937_lower_mask)
			mt.state[kk] = mt.state[kk+(mt19937_m-mt19937_n)] ^ (y >> 1) ^ mt19937_mult_matrix_a[y&0x1]
			kk++
		}
		y = (mt.state[mt19937_n-1] & mt19937_upper_mask) | (mt.state[0] & mt19937_lower_mask)
		mt.state[mt19937_n-1] = mt.state[mt19937_m-1] ^ (y >> 1) ^ mt19937_mult_matrix_a[y&0x1]
		mt.index = 0
	}

	y = mt.state[mt.index]

	// Tempering.
	y ^= y >> 11
	y ^= (y << 7) & 0x9d2c5680
	y ^= (y << 15) & 0xefc60000
	y ^= y >> 18
	mt.index++
	return y
}

func (mt *MT19937) String() string {
	return "MT19937"
}
