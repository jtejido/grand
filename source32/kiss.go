package source32

const (
	kiss_r = 4
)

// Port from Marsaglia's "KISS" algorithm
// This version contains the correction referred to
// https://programmingpraxis.com/2010/10/05/george-marsaglias-random-number-generators
//
// https://en.wikipedia.org/wiki/KISS_(algorithm)
// ========= Summary results of Crush =========
//
//  Version:          TestU01 1.2.3
//  Generator:        rand.Float64
//  Number of statistics:  144
//  Total CPU time:   02:48:43.60
//
//  All tests were passed
type KISS struct {
	baseSource32
	state [4]uint32
}

func NewKISSFromStream(seed []uint32) (*KISS, error) {
	err := checkEmptySeed(seed)
	if err != nil {
		return nil, err
	}

	ans := new(KISS)
	ans.spi = ans
	if len(seed) < kiss_r {
		tmp := make([]uint32, kiss_r)
		fillState(tmp, seed)
		ans.setSeed(tmp)
	} else {
		ans.setSeed(seed)
	}
	return ans, nil
}

func NewKISS(seed int64) *KISS {
	ans := new(KISS)
	ans.spi = ans
	ans.Seed(seed)
	return ans
}

func (kiss *KISS) setSeed(seed []uint32) {
	kiss.stream = append([]uint32{}, seed...)
	kiss.Restart()
}

func (kiss *KISS) Restart() {
	for i := 0; i < kiss_r; i++ {
		kiss.state[i] = kiss.stream[i]
	}

	kiss.resetState()
}

func (kiss *KISS) Seed(seed int64) {
	seeds := make([]uint32, kiss_r)
	seeder.Seed(seed)
	var i int
	// Fill the remaining pairs
	for i < kiss_r {
		v := seeder.Uint32()
		seeds[i] = v
		i++
	}

	// Initialize the pool content.
	kiss.setSeed(seeds)
}

func (kiss *KISS) Uint32() uint32 {
	kiss.state[0] = kiss.computeNew(36969, kiss.state[0])
	kiss.state[1] = kiss.computeNew(18000, kiss.state[1])
	mwc := (kiss.state[0] << 16) + kiss.state[1]

	// Cf. correction mentioned in the reply to the original post:
	//   https://programmingpraxis.com/2010/10/05/george-marsaglias-random-number-generators/
	kiss.state[2] ^= kiss.state[2] << 13
	kiss.state[2] ^= kiss.state[2] >> 17
	kiss.state[2] ^= kiss.state[2] << 5

	kiss.state[3] = 69069*kiss.state[3] + 1234567

	return (mwc ^ kiss.state[3]) + kiss.state[2]
}

func (kiss *KISS) computeNew(mult, previous uint32) uint32 {
	return mult*(previous&65535) + (previous >> 16)
}

func (kiss *KISS) String() string {
	return "KISS"
}
