package source32

// A fast all-purpose 32-bit generator. For faster generation of {@code float} values try the
// xoshiro128Plus generator.
//
// This is a member of the Xor-Shift-Rotate family of generators. Memory footprint is 128
// bits.
//
// http://xoshiro.di.unimi.it/xoshiro128starstar.c
// http://xoshiro.di.unimi.it/
//
// ========= Summary results of Crush =========
//
// Version:          TestU01 1.2.3
// Generator:        rand.Float64
// Number of statistics:  144
// Total CPU time:   02:55:12.54
//
// All tests were passed
type XoShiRo128StarStar struct {
	baseXoShiRo128
}

func NewXoShiRo128StarStarFromStream(seed []uint32) (*XoShiRo128StarStar, error) {
	err := checkEmptySeed(seed)
	if err != nil {
		return nil, err
	}
	ans := new(XoShiRo128StarStar)
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
func NewXoShiRo128StarStar(seed int64) *XoShiRo128StarStar {
	ans := new(XoShiRo128StarStar)
	ans.spi = ans
	ans.Seed(seed)
	return ans
}

func (xoshiro *XoShiRo128StarStar) Uint32() uint32 {
	result := rotateLeft(xoshiro.state[0]*5, 7) * 9

	t := xoshiro.state[1] << 9

	xoshiro.state[2] ^= xoshiro.state[0]
	xoshiro.state[3] ^= xoshiro.state[1]
	xoshiro.state[1] ^= xoshiro.state[2]
	xoshiro.state[0] ^= xoshiro.state[3]

	xoshiro.state[2] ^= t

	xoshiro.state[3] = rotateLeft(xoshiro.state[3], 11)

	return result
}

func (xoshiro *XoShiRo128StarStar) String() string {
	return "XoShiRo128StarStar"
}
