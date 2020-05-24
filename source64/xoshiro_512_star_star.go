package source64

// A fast all-purpose generator.
//
// This is a member of the Xor-Shift-Rotate family of generators. Memory footprint is 512 bits
// and the period is 2^512-1. Speed is expected to be slower than XoShiRo256StarStar.
//
// http://xoshiro.di.unimi.it/xoshiro512starstar.c
// http://xoshiro.di.unimi.it/
//
// ========= Summary results of Crush =========
//
// Version:          TestU01 1.2.3
// Generator:        rand.Float64
// Number of statistics:  144
// Total CPU time:   03:34:34.38
//
// All tests were passed
type XoShiRo512StarStar struct {
	baseXoShiRo512
}

func NewXoShiRo512StarStarFromStream(seed []uint64) (*XoShiRo512StarStar, error) {
	err := checkEmptySeed(seed)
	if err != nil {
		return nil, err
	}

	ans := new(XoShiRo512StarStar)

	if len(seed) < xoshiro512_r {
		tmp := make([]uint64, xoshiro512_r)
		fillState(tmp, seed)
		ans.setSeed(tmp)
	} else {
		ans.setSeed(seed)
	}

	return ans, nil
}

// this builds the seed slice from a split_mx64 generator using the seed provided
func NewXoShiRo512StarStar(seed int64) *XoShiRo512StarStar {
	ans := new(XoShiRo512StarStar)
	ans.Seed(seed)
	return ans
}

func (xoshiro *XoShiRo512StarStar) Uint64() uint64 {
	result := rotateLeft(xoshiro.state[1]*5, 7) * 9

	t := xoshiro.state[1] << 11

	xoshiro.state[2] ^= xoshiro.state[0]
	xoshiro.state[5] ^= xoshiro.state[1]
	xoshiro.state[1] ^= xoshiro.state[2]
	xoshiro.state[7] ^= xoshiro.state[3]
	xoshiro.state[3] ^= xoshiro.state[4]
	xoshiro.state[4] ^= xoshiro.state[5]
	xoshiro.state[0] ^= xoshiro.state[6]
	xoshiro.state[6] ^= xoshiro.state[7]

	xoshiro.state[6] ^= t

	xoshiro.state[7] = rotateLeft(xoshiro.state[7], 21)

	return result
}

func (xoshiro *XoShiRo512StarStar) String() string {
	return "XoShiRo512StarStar"
}
