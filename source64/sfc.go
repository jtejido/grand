package source64

const (
	sfc_r = 3
)

// Implement the Small, Fast, Chaotic (SFC) 32-bit generator of Chris Doty-Humphrey.
//
// The state size is 128-bits; the period is a minimum of 2^32 and an
// average of approximately 2^127.
//
// http://pracrand.sourceforge.net
type SFC struct {
	baseSource64
	state   [3]uint64
	counter uint64
}

func NewSFCFromStream(seed []uint64) (*SFC, error) {
	err := checkEmptySeed(seed)
	if err != nil {
		return nil, err
	}

	ans := new(SFC)
	ans.spi = ans
	if len(seed) < sfc_r {
		tmp := make([]uint64, sfc_r)
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

func (sfc *SFC) setSeed(seed []uint64) {
	sfc.stream = append([]uint64{}, seed...)
	sfc.Restart()
}

func (sfc *SFC) Restart() {
	for i := 0; i < sfc_r; i++ {
		sfc.state[i] = sfc.stream[i]
	}

	sfc.counter = 1
	for i := 0; i < 18; i++ {
		sfc.Uint64()
	}

	sfc.resetState()
}

func (sfc *SFC) Uint64() uint64 {
	tmp := sfc.state[0] + sfc.state[1] + sfc.counter
	sfc.counter++
	sfc.state[0] = sfc.state[1] ^ (sfc.state[1] >> 11)
	sfc.state[1] = sfc.state[2] + (sfc.state[2] << 3)
	sfc.state[2] = rotateLeft(sfc.state[2], 24) + tmp
	return tmp
}

func (sfc *SFC) Seed(seed int64) {
	seeds := make([]uint64, sfc_r)
	seeder.Seed(seed)
	var i int
	// Fill the remaining pairs
	for i < sfc_r {
		v := seeder.Uint64()
		seeds[i] = v
		i++
	}

	// Initialize the pool content.
	sfc.setSeed(seeds)
}

func (sfc *SFC) String() string {
	return "SFC"
}
