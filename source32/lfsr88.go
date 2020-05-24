package source32

const (
	lfsr88_r = 3
)

// This implements the LFSR88 or Taus88 (Tausworthe) pseudo-random number generator
// from Pierre L'Ecuyer.
//
// This generator is described in a paper by L'Ecuyer.
// http://www-labs.iro.umontreal.ca/~lecuyer/myftp/papers/tausme.ps
// Maximally equidistributed combined Tausworthe generators.
// Mathematics of Computation, 65, 213 (1996), 203--213.
// ========= Summary results of Crush =========
//
//  Version:          TestU01 1.2.3
//  Generator:        rand.Float64
//  Number of statistics:  144
//  Total CPU time:   02:43:38.35
//  The following tests gave p-values outside [0.001, 0.9990]:
//  (eps  means a value < 1.0e-300):
//  (eps1 means a value < 1.0e-15):
//
//        Test                          p-value
//  ----------------------------------------------
//  58  MatrixRank, 300 x 300            eps
//  59  MatrixRank, 300 x 300            eps
//  60  MatrixRank, 1200 x 1200          eps
//  61  MatrixRank, 1200 x 1200          eps
//  71  LinearComp, r = 0              1 - eps1
//  72  LinearComp, r = 29             1 - eps1
//  72  LinearComp, r = 29              8.4e-4
//  ----------------------------------------------
//  All other tests were passed
// TO-DO.. Jump()
type LFSR88 struct {
	baseSource32
	state [3]uint32
	b     uint32
}

func NewLFSR88FromStream(seed []uint32) (*LFSR88, error) {
	err := checkEmptySeed(seed)
	if err != nil {
		return nil, err
	}

	ans := new(LFSR88)
	ans.spi = ans
	if len(seed) < lfsr88_r {
		tmp := make([]uint32, lfsr88_r)
		fillState(tmp, seed)
		ans.setSeed(tmp)
	} else {
		ans.setSeed(seed)
	}
	return ans, nil
}

// this builds the seed slice from a split_mx64 generator using the seed provided
func NewLFSR88(seed int64) *LFSR88 {
	ans := new(LFSR88)
	ans.spi = ans
	ans.Seed(seed)
	return ans
}

func (lfsr *LFSR88) setSeed(seed []uint32) {
	lfsr.checkSeed(seed)
	lfsr.stream = append([]uint32{}, seed...)
	lfsr.Restart()
}

func (lfsr *LFSR88) Restart() {
	for j := 0; j < lfsr88_r; j++ {
		lfsr.state[j] = lfsr.stream[j]
	}

	lfsr.b = 0

	lfsr.resetState()
}

func (lfsr *LFSR88) Uint32() uint32 {
	lfsr.b = (((lfsr.state[0] << 13) ^ lfsr.state[0]) >> 19)
	lfsr.state[0] = (((lfsr.state[0] & 4294967294) << 12) ^ lfsr.b)
	lfsr.b = (((lfsr.state[1] << 2) ^ lfsr.state[1]) >> 25)
	lfsr.state[1] = (((lfsr.state[1] & 4294967288) << 4) ^ lfsr.b)
	lfsr.b = (((lfsr.state[2] << 3) ^ lfsr.state[2]) >> 11)
	lfsr.state[2] = (((lfsr.state[2] & 4294967280) << 17) ^ lfsr.b)
	return (lfsr.state[0] ^ lfsr.state[1] ^ lfsr.state[2])
}

func (lfsr *LFSR88) checkSeed(seed []uint32) {
	if (seed[0] >= 0 && seed[0] < 2) || (seed[1] >= 0 && seed[1] < 8) || (seed[2] >= 0 && seed[2] < 16) {
		panic("The seed elements must be greater than 1, 7 and 15 respectively")
	}
}

func (lfsr *LFSR88) Seed(seed int64) {
	seeds := make([]uint32, lfsr88_r)
	seeder.Seed(seed)

	for j := 0; j < lfsr88_r; j++ {
	again:
		f := seeder.Uint32()

		if (j == 0 && f <= 1) || (j == 1 && f <= 7) || (j == 2 && f <= 15) {
			goto again
		}

		seeds[j] = f

	}

	// Initialize the pool content.
	lfsr.setSeed(seeds)
}

func (lfsr *LFSR88) String() string {
	return "LFSR88"
}
