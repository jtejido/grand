package source64

import (
	"math"
)

const (
	mt19937_n                 = 312
	mt19937_m                 = 156
	mt19937_upper_mask uint64 = 0xffffffff80000000
	mt19937_lower_mask uint64 = 0x7fffffff
)

var (
	mt19937_mult_matrix_a = []uint64{0x0, 0xb5026f5aa96619e9}
)

// Implements the 64-bits version of the originally 32-bits
// 2014/2/23 version of the generator written in C by Makoto Matsumoto
// and Takuji Nishimura.
// http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/emt64.html
// ========= Summary results of Crush =========
//
//  Version:          TestU01 1.2.3
//  Generator:        rand.Float64
//  Number of statistics:  144
//  Total CPU time:   03:01:51.54
//  The following tests gave p-values outside [0.001, 0.9990]:
//  (eps  means a value < 1.0e-300):
//  (eps1 means a value < 1.0e-15):
//
//        Test                          p-value
//  ----------------------------------------------
//  71  LinearComp, r = 0              1 - eps1
//  72  LinearComp, r = 29             1 - eps1
//  ----------------------------------------------
//  All other tests were passed
type MT19937 struct {
	baseSource64
	state [312]uint64
	index uint64
}

func NewMT19937FromStream(stream []uint64) (*MT19937, error) {
	err := checkEmptySeed(stream)
	if err != nil {
		return nil, err
	}

	ans := new(MT19937)
	ans.baseSource64.spi = ans
	ans.setSeed(stream)
	return ans, nil
}

// this builds the seed slice from a split_mx64 generator using the seed provided
func NewMT19937(seed int64) *MT19937 {
	ans := new(MT19937)
	ans.baseSource64.spi = ans
	ans.Seed(seed)
	return ans
}

func (mt *MT19937) setSeed(seed []uint64) {

	if len(seed) == 0 {
		seed = make([]uint64, 1)
	}

	mt.stream = make([]uint64, mt19937_n)

	mt.stream[0] = 19650218
	for mt.index = 1; mt.index < mt19937_n; mt.index++ {
		mm1 := mt.stream[mt.index-1]
		mt.stream[mt.index] = 0x5851f42d4c957f2d*(mm1^(mm1>>62)) + mt.index
	}

	i := 1
	j := 0

	for k := int(math.Max(float64(mt19937_n), float64(len(seed)))); k != 0; k-- {
		mm1 := mt.stream[i-1]
		mt.stream[i] = (mt.stream[i] ^ ((mm1 ^ (mm1 >> 62)) * 0x369dea0f31a53f85)) + seed[j] + uint64(j) // non linear
		i++
		j++
		if i >= mt19937_n {
			mt.stream[0] = mt.stream[mt19937_n-1]
			i = 1
		}
		if j >= len(seed) {
			j = 0
		}
	}
	for k := mt19937_n - 1; k != 0; k-- {
		mm1 := mt.stream[i-1]
		mt.stream[i] = (mt.stream[i] ^ ((mm1 ^ (mm1 >> 62)) * 0x27bb2ee687b0b0fd)) - uint64(i) // non linear
		i++
		if i >= mt19937_n {
			mt.stream[0] = mt.stream[mt19937_n-1]
			i = 1
		}
	}

	mt.stream[0] = 0x8000000000000000 // MSB is 1; assuring non-zero initial array

	mt.Restart()
}

func (mt *MT19937) Restart() {
	for i := 0; i < len(mt.stream); i++ {
		mt.state[i] = mt.stream[i]
	}

	mt.resetState()
}

func (mt *MT19937) Seed(seed int64) {
	seeds := make([]uint64, mt19937_n)
	seeder.Seed(seed)
	var i int
	// Fill the remaining pairs
	for i < mt19937_n {
		v := seeder.Uint64()
		seeds[i] = v
		i++
	}

	// Initialize the pool content.
	mt.setSeed(seeds)
}

func (mt *MT19937) Uint64() uint64 {
	var x uint64
	if mt.index >= mt19937_n {
		for i := 0; i < mt19937_n-mt19937_m; i++ {
			x = (mt.state[i] & mt19937_upper_mask) | (mt.state[i+1] & mt19937_lower_mask)
			mt.state[i] = mt.state[i+mt19937_m] ^ (x >> 1) ^ mt19937_mult_matrix_a[int(x&0x1)]
		}
		for i := mt19937_n - mt19937_m; i < mt19937_n-1; i++ {
			x = (mt.state[i] & mt19937_upper_mask) | (mt.state[i+1] & mt19937_lower_mask)
			mt.state[i] = mt.state[i+(mt19937_m-mt19937_n)] ^ (x >> 1) ^ mt19937_mult_matrix_a[int(x&0x1)]
		}

		x = (mt.state[mt19937_n-1] & mt19937_upper_mask) | (mt.state[0] & mt19937_lower_mask)
		mt.state[mt19937_n-1] = mt.state[mt19937_m-1] ^ (x >> 1) ^ mt19937_mult_matrix_a[int(x&0x1)]

		mt.index = 0
	}

	x = mt.state[mt.index]

	x ^= (x >> 29) & 0x5555555555555555
	x ^= (x << 17) & 0x71d67fffeda60000
	x ^= (x << 37) & 0xfff7eee000000000
	x ^= x >> 43
	mt.index++
	return x
}

func (mt *MT19937) String() string {
	return "MT19937"
}
