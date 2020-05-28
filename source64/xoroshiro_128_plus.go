package source64

// A fast 64-bit generator suitable for float generation. This is slightly faster than the all-purpose
// XoShiRo128StarStar generator.
//
// This is a member of the Xor-Shift-Rotate family of generators. Memory footprint is 128 bits
// and the period is 2^128-1. Speed is expected to be similar to XoShiRo256Plus.
//
// http://xoshiro.di.unimi.it/xoroshiro128plus.c
// http://xoshiro.di.unimi.it/
//
// ========= Summary results of Crush =========

// Version:          TestU01 1.2.3
// Generator:        rand.Float64
// Number of statistics:  144
// Total CPU time:   03:12:41.29
// The following tests gave p-values outside [0.001, 0.9990]:
// (eps  means a value < 1.0e-300):
// (eps1 means a value < 1.0e-15):

//       Test                          p-value
// ----------------------------------------------
//  7  CollisionOver, t = 8            0.9996
// 85  HammingIndep, L = 30            2.0e-4
// ----------------------------------------------
// All other tests were passed
type XoRoShiRo128Plus struct {
	baseXoRoShiRo128
}

func NewXoRoShiRo128PlusFromStream(seed []uint64) (*XoRoShiRo128Plus, error) {
	err := checkEmptySeed(seed)
	if err != nil {
		return nil, err
	}

	ans := new(XoRoShiRo128Plus)
	ans.spi = ans
	if len(seed) < xoroshiro128_r {
		tmp := make([]uint64, xoroshiro128_r)
		fillState(tmp, seed)
		ans.setSeed(tmp)
	} else {
		ans.setSeed(seed)
	}

	return ans, nil
}

// this builds the seed slice from a split_mx64 generator using the seed provided
func NewXoRoShiRo128Plus(seed int64) *XoRoShiRo128Plus {
	ans := new(XoRoShiRo128Plus)
	ans.spi = ans
	ans.Seed(seed)
	return ans
}

func (xoshiro *XoRoShiRo128Plus) Uint64() uint64 {

	s0 := xoshiro.state[0]
	s1 := xoshiro.state[1]
	result := s0 + s1

	s1 ^= s0
	xoshiro.state[0] = rotateLeft(s0, 24) ^ s1 ^ (s1 << 16) // a, b
	xoshiro.state[1] = rotateLeft(s1, 37)                   // c

	return result
}

func (xoshiro *XoRoShiRo128Plus) String() string {
	return "XoRoShiRo128Plus"
}
