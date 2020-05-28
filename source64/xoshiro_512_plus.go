package source64

// This is slightly faster than the all-purpose generator XoShiRo512StarStar
//
// This is a member of the Xor-Shift-Rotate family of generators. Memory footprint is 512 bits
// and the period is 2^512-1. Speed is expected to be slower than XoShiRo256Plus.
//
// http://xoshiro.di.unimi.it/xoshiro512plus.c
// http://xoshiro.di.unimi.it/
//
// ========= Summary results of Crush =========
//
// Version:          TestU01 1.2.3
// Generator:        rand.Float64
// Number of statistics:  144
// Total CPU time:   03:10:35.87
//
// All tests were passed
type XoShiRo512Plus struct {
	baseXoShiRo512
}

func NewXoShiRo512PlusFromStream(seed []uint64) (*XoShiRo512Plus, error) {
	err := checkEmptySeed(seed)
	if err != nil {
		return nil, err
	}

	ans := new(XoShiRo512Plus)
	ans.spi = ans
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
func NewXoShiRo512Plus(seed int64) *XoShiRo512Plus {
	ans := new(XoShiRo512Plus)
	ans.spi = ans
	ans.Seed(seed)
	return ans
}

func (xoshiro *XoShiRo512Plus) Uint64() uint64 {
	result := xoshiro.state[0] + xoshiro.state[2]
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

func (xoshiro *XoShiRo512Plus) String() string {
	return "XoShiRo512Plus"
}
