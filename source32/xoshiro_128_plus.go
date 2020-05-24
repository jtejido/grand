package source32

// A fast 32-bit generator suitable for float generation. This is slightly faster than the all-purpose
// XoShiRo128StarStar generator.
//
// This is a member of the Xor-Shift-Rotate family of generators. Memory footprint is 128
// bits.
//
// http://xoshiro.di.unimi.it/xoshiro128plus.c
// http://xoshiro.di.unimi.it/
//
// ========= Summary results of Crush =========
//
// Version:          TestU01 1.2.3
// Generator:        rand.Float64
// Number of statistics:  144
// Total CPU time:   03:57:26.63
//
// All tests were passed
type XoShiRo128Plus struct {
	baseXoShiRo128
}

func NewXoShiRo128PlusFromStream(seed []uint32) (*XoShiRo128Plus, error) {
	err := checkEmptySeed(seed)
	if err != nil {
		return nil, err
	}
	ans := new(XoShiRo128Plus)
	ans.spi = ans
	if len(seed) < xoshiro128_r {
		tmp := make([]uint32, xoshiro128_r)
		fillState(tmp, seed)
		ans.setSeed(tmp)
	} else {
		ans.setSeed(seed)
	}

	return ans, nil
}

// this builds the seed slice from a split_mx64 generator using the seed provided
func NewXoShiRo128Plus(seed int64) *XoShiRo128Plus {
	ans := new(XoShiRo128Plus)
	ans.spi = ans
	ans.Seed(seed)
	return ans
}

func (xoshiro *XoShiRo128Plus) Uint32() uint32 {
	result := xoshiro.state[0] + xoshiro.state[3]

	t := xoshiro.state[1] << 9

	xoshiro.state[2] ^= xoshiro.state[0]
	xoshiro.state[3] ^= xoshiro.state[1]
	xoshiro.state[1] ^= xoshiro.state[2]
	xoshiro.state[0] ^= xoshiro.state[3]

	xoshiro.state[2] ^= t

	xoshiro.state[3] = rotateLeft(xoshiro.state[3], 11)

	return result
}

func (xoshiro *XoShiRo128Plus) String() string {
	return "XoShiRo128Plus"
}
