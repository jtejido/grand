package source64

// A fast all-purpose 64-bit generator.
//
// This is a member of the Xor-Shift-Rotate family of generators. Memory footprint is 256 bits
// and the period is 2^256-1
//
// http://xoshiro.di.unimi.it/xoshiro256starstar.c
// http://xoshiro.di.unimi.it/
//
// ========= Summary results of Crush =========
//
// Version:          TestU01 1.2.3
// Generator:        rand.Float64
// Number of statistics:  144
// Total CPU time:   03:10:37.87
//
// All tests were passed
type XoShiRo256StarStar struct {
	baseXoShiRo256
}

func NewXoShiRo256StarStarFromStream(seed []uint64) (*XoShiRo256StarStar, error) {
	err := checkEmptySeed(seed)
	if err != nil {
		return nil, err
	}

	ans := new(XoShiRo256StarStar)

	if len(seed) < xoshiro256_r {
		tmp := make([]uint64, xoshiro256_r)
		fillState(tmp, seed)
		ans.setSeed(tmp)
	} else {
		ans.setSeed(seed)
	}

	return ans, nil
}

// this builds the seed slice from a split_mx64 generator using the seed provided
func NewXoShiRo256StarStar(seed int64) *XoShiRo256StarStar {
	ans := new(XoShiRo256StarStar)
	ans.Seed(seed)
	return ans
}

func (xoshiro *XoShiRo256StarStar) Uint64() uint64 {
	result := rotateLeft(xoshiro.state[1]*5, 7) * 9

	t := xoshiro.state[1] << 17

	xoshiro.state[2] ^= xoshiro.state[0]
	xoshiro.state[3] ^= xoshiro.state[1]
	xoshiro.state[1] ^= xoshiro.state[2]
	xoshiro.state[0] ^= xoshiro.state[3]

	xoshiro.state[2] ^= t

	xoshiro.state[3] = rotateLeft(xoshiro.state[3], 45)

	return result
}

func (xoshiro *XoShiRo256StarStar) String() string {
	return "XoShiRo256StarStar"
}
