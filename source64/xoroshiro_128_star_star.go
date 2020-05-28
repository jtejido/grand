package source64

// A fast all-purpose 64-bit generator.
//
// This is a member of the Xor-Shift-Rotate family of generators. Memory footprint is 128 bits
// and the period is 2^128-1. Speed is expected to be similar to XoShiRo256StarStar
//
// http://xoshiro.di.unimi.it/xoroshiro128statstar.c
// http://xoshiro.di.unimi.it/
//
// ========= Summary results of Crush =========

// Version:          TestU01 1.2.3
// Generator:        rand.Float64
// Number of statistics:  144
// Total CPU time:   03:08:10.65

// All tests were passed
type XoRoShiRo128StarStar struct {
	baseXoRoShiRo128
}

func NewXoRoShiRo128StarStarFromStream(seed []uint64) (*XoRoShiRo128StarStar, error) {
	err := checkEmptySeed(seed)
	if err != nil {
		return nil, err
	}

	ans := new(XoRoShiRo128StarStar)
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
func NewXoRoShiRo128StarStar(seed int64) *XoRoShiRo128StarStar {
	ans := new(XoRoShiRo128StarStar)
	ans.spi = ans
	ans.Seed(seed)
	return ans
}

func (xoshiro *XoRoShiRo128StarStar) Uint64() uint64 {
	s0 := xoshiro.state[0]
	s1 := xoshiro.state[1]
	result := rotateLeft(s0*5, 7) * 9

	s1 ^= s0
	xoshiro.state[0] = rotateLeft(s0, 24) ^ s1 ^ (s1 << 16)
	xoshiro.state[1] = rotateLeft(s1, 37)

	return result
}

func (xoshiro *XoRoShiRo128StarStar) String() string {
	return "XoRoShiRo128StarStar"
}
