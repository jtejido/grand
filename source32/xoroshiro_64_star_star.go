package source32

// A fast all-purpose 32-bit generator. For faster generation of float values try the
// Xoshiro64Star generator.
//
// This is a member of the Xor-Shift-Rotate family of generators. Memory footprint is 64 bits.
//
// http://xoshiro.di.unimi.it/xoroshiro64starstar.c
// http://xoshiro.di.unimi.it/
//
// ========= Summary results of Crush =========
//
// Version:          TestU01 1.2.3
// Generator:        rand.Float64
// Number of statistics:  144
// Total CPU time:   03:38:58.87
//
// All tests were passed
type XoRoShiRo64StarStar struct {
	baseXoRoShiRo64
}

func NewXoRoShiRo64StarStarFromStream(seed []uint32) (*XoRoShiRo64StarStar, error) {
	err := checkEmptySeed(seed)
	if err != nil {
		return nil, err
	}

	ans := new(XoRoShiRo64StarStar)
	ans.spi = ans
	if len(seed) < xoroshiro_r {
		tmp := make([]uint32, xoroshiro_r)
		fillState(tmp, seed)
		ans.setSeed(tmp)
	} else {
		ans.setSeed(seed)
	}

	return ans, nil
}

// this builds the seed slice from a split_mx64 generator using the seed provided
func NewXoRoShiRo64StarStar(seed int64) *XoRoShiRo64StarStar {
	ans := new(XoRoShiRo64StarStar)
	ans.spi = ans
	ans.Seed(seed)
	return ans
}

func (xoshiro *XoRoShiRo64StarStar) Uint32() uint32 {
	s0 := xoshiro.state[0]
	s1 := xoshiro.state[1]
	result := rotateLeft(s0*0x9E3779BB, 5) * 5

	s1 ^= s0
	xoshiro.state[0] = rotateLeft(s0, 26) ^ s1 ^ (s1 << 9) // a, b
	xoshiro.state[1] = rotateLeft(s1, 13)                  // c

	return result
}

func (xoshiro *XoRoShiRo64StarStar) String() string {
	return "XoRoShiRo64StarStar"
}
