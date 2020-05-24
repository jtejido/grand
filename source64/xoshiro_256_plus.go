package source64

// This is slightly faster than the all-purpose generator XoShiRo256StarStar
//
// This is a member of the Xor-Shift-Rotate family of generators. Memory footprint is 256 bits
// and the period is 2^256-1.
//
// http://xoshiro.di.unimi.it/xoshiro256plus.c
// http://xoshiro.di.unimi.it/
//
// ========= Summary results of Crush =========
//
// Version:          TestU01 1.2.3
// Generator:        rand.Float64
// Number of statistics:  144
// Total CPU time:   02:59:16.46
//
// All tests were passed
type XoShiRo256Plus struct {
	baseXoShiRo256
}

func NewXoShiRo256PlusFromStream(seed []uint64) (*XoShiRo256Plus, error) {
	err := checkEmptySeed(seed)
	if err != nil {
		return nil, err
	}

	ans := new(XoShiRo256Plus)

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
func NewXoShiRo256Plus(seed int64) *XoShiRo256Plus {
	ans := new(XoShiRo256Plus)
	ans.Seed(seed)
	return ans
}

func (xoshiro *XoShiRo256Plus) Uint64() uint64 {

	result := xoshiro.state[0] + xoshiro.state[3]

	t := xoshiro.state[1] << 17

	xoshiro.state[2] ^= xoshiro.state[0]
	xoshiro.state[3] ^= xoshiro.state[1]
	xoshiro.state[1] ^= xoshiro.state[2]
	xoshiro.state[0] ^= xoshiro.state[3]

	xoshiro.state[2] ^= t

	xoshiro.state[3] = rotateLeft(xoshiro.state[3], 45)

	return result
}

func (xoshiro *XoShiRo256Plus) String() string {
	return "XoShiRo256Plus"
}
