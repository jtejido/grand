package source32

const (
	sfc32_r = 3
)

// Implement the Small, Fast, Chaotic (SFC) 32-bit generator of Chris Doty-Humphrey.
//
// The state size is 128-bits; the period is a minimum of 2^32 and an
// average of approximately 2^127.
//
// http://pracrand.sourceforge.net
//
// ========= Summary results of Crush =========
//
// Version:          TestU01 1.2.3
// Generator:        rand.Float64
// Number of statistics:  144
// Total CPU time:   02:39:04.45
//
// All tests were passed
type SFC struct {
	baseSource32
	state   [3]uint32
	counter uint32
}

func NewSFCFromStream(seed []uint32) (*SFC, error) {
	err := checkEmptySeed(seed)
	if err != nil {
		return nil, err
	}

	ans := new(SFC)
	ans.spi = ans
	if len(seed) < sfc32_r {
		tmp := make([]uint32, sfc32_r)
		fillState(tmp, seed)
		ans.setSeed(tmp)
	} else {
		ans.setSeed(seed)
	}
	return ans, nil
}

// this builds the seed slice from a split_mx64 generator using the seed provided
func NewSFC(seed int64) *SFC {
	ans := new(SFC)
	ans.spi = ans
	ans.Seed(seed)
	return ans
}

func (sfc *SFC) setSeed(seed []uint32) {
	sfc.stream = append([]uint32{}, seed...)
	sfc.Restart()
}

func (sfc *SFC) Restart() {
	for i := 0; i < sfc32_r; i++ {
		sfc.state[i] = sfc.stream[i]
	}

	sfc.counter = 1
	for i := 0; i < 15; i++ {
		sfc.Uint32()
	}

	sfc.resetState()
}

func (sfc *SFC) Uint32() uint32 {
	tmp := sfc.state[0] + sfc.state[1] + sfc.counter
	sfc.counter++
	sfc.state[0] = sfc.state[1] ^ (sfc.state[1] >> 9)
	sfc.state[1] = sfc.state[2] + (sfc.state[2] << 3)
	sfc.state[2] = rotateLeft(sfc.state[2], 21) + tmp
	return tmp
}

func (sfc *SFC) Seed(seed int64) {
	seeds := make([]uint32, sfc32_r)
	seeder.Seed(seed)
	var i int
	// Fill the remaining pairs
	for i < sfc32_r {
		v := seeder.Uint32()
		seeds[i] = v
		i++
	}

	// Initialize the pool content.
	sfc.setSeed(seeds)
}

func (sfc *SFC) String() string {
	return "SFC"
}
