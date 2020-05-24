package source32

const (
	mwc_r                = 256
	mwc_seed_size        = mwc_r + 1
	mwc_a         uint32 = 809430660
)

// Port from Marsaglia's "Multiply-With-Carry" Algorithm
// https://en.wikipedia.org/wiki/Multiply-with-carry
//
// Implementation is based on the C code reproduced on:
// http://school.anhb.uwa.edu.au/personalpages/kwessen/shared/Marsaglia03.html
//
type MultiplyWithCarry256 struct {
	baseSource32
	state        [256]uint32
	index, carry uint32
}

func NewMultiplyWithCarry256FromStream(seeds []uint32) (*MultiplyWithCarry256, error) {
	err := checkEmptySeed(seeds)
	if err != nil {
		return nil, err
	}

	ans := new(MultiplyWithCarry256)
	ans.spi = ans
	ans.setSeed(seeds)
	return ans, nil
}

// this builds the seed slice from a split_mx64 generator using the seed provided
func NewMultiplyWithCarry256(seed int64) *MultiplyWithCarry256 {
	ans := new(MultiplyWithCarry256)
	ans.spi = ans
	ans.Seed(seed)
	return ans
}

func (mwc256 *MultiplyWithCarry256) setSeed(seed []uint32) {
	seeds := make([]uint32, mwc_seed_size)
	fillState(seeds, seed)
	c := seeds[0]
	mwc256.carry = c % mwc_a
	mwc256.stream = append([]uint32{}, seeds[1:1+mwc_r]...)

	mwc256.Restart()
}

func (mwc256 *MultiplyWithCarry256) Restart() {
	for i := 0; i < mwc_r; i++ {
		mwc256.state[i] = mwc256.stream[i]
	}
	mwc256.index = mwc_r
	mwc256.resetState()
}

func (mwc256 *MultiplyWithCarry256) Seed(seed int64) {
	seeds := make([]uint32, mwc_seed_size)
	seeder.Seed(seed)
	var i int
	for i < mwc_seed_size {
		v := seeder.Uint32()
		seeds[i] = v
		i++
	}

	mwc256.setSeed(seeds)
}

func (mwc256 *MultiplyWithCarry256) Uint32() uint32 {
	mwc256.index &= 0xff
	t := uint64(mwc_a)*(uint64(mwc256.state[mwc256.index])&0xffffffff) + uint64(mwc256.carry)
	mwc256.carry = uint32(t >> 32)
	mwc256.state[mwc256.index] = uint32(t)
	ret := uint32(t)
	mwc256.index++
	return ret
}

func (mwc256 *MultiplyWithCarry256) String() string {
	return "MultiplyWithCarry256"
}
