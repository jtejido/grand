package source32

// A fast 32-bit generator suitable for float generation. This is slightly faster than the all-purpose
// Xoshiro64StarStar generator.
//
// This is a member of the Xor-Shift-Rotate family of generators. Memory footprint is 64 bits.
//
// http://xoshiro.di.unimi.it/xoroshiro64star.c
// http://xoshiro.di.unimi.it/
//
// ========= Summary results of Crush =========
//
// Version:          TestU01 1.2.3
// Generator:        rand.Float64
// Number of statistics:  144
// Total CPU time:   02:42:10.04
// The following tests gave p-values outside [0.001, 0.9990]:
// (eps  means a value < 1.0e-300):
// (eps1 means a value < 1.0e-15):
//
//       Test                          p-value
// ----------------------------------------------
// 11  BirthdaySpacings, t = 2         3.1e-4
// 76  LongestHeadRun, r = 0           5.7e-4
// ----------------------------------------------
// All other tests were passed
type XoRoShiRo64Star struct {
	baseXoRoShiRo64
}

func NewXoRoShiRo64StarFromStream(seed []uint32) (*XoRoShiRo64Star, error) {
	err := checkEmptySeed(seed)
	if err != nil {
		return nil, err
	}

	ans := new(XoRoShiRo64Star)
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
func NewXoRoShiRo64Star(seed int64) *XoRoShiRo64Star {
	ans := new(XoRoShiRo64Star)
	ans.spi = ans
	ans.Seed(seed)
	return ans
}

func (xoshiro *XoRoShiRo64Star) Uint32() uint32 {
	s0 := xoshiro.state[0]
	s1 := xoshiro.state[1]
	result := s0 * 0x9E3779BB

	s1 ^= s0
	xoshiro.state[0] = rotateLeft(s0, 26) ^ s1 ^ (s1 << 9) // a, b
	xoshiro.state[1] = rotateLeft(s1, 13)                  // c

	return result
}

func (xoshiro *XoRoShiRo64Star) String() string {
	return "XoRoShiRo64Star"
}
