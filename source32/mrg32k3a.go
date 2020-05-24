package source32

const (
	mrg32k3a_m1   uint32 = 4294967087
	mrg32k3a_m2   uint32 = 4294944443
	mrg32k3a_a12  uint32 = 1403580
	mrg32k3a_a13n uint32 = 810728
	mrg32k3a_a21  uint32 = 527612
	mrg32k3a_a23n uint32 = 1370589
	mrg32k3a_r           = 6
)

var (
	a1p76 = [][]uint32{
		{82758667, 1871391091, 4127413238},
		{3672831523, 69195019, 1871391091},
		{3672091415, 3528743235, 69195019},
	}
	a2p76 = [][]uint32{
		{1511326704, 3759209742, 1610795712},
		{4292754251, 1511326704, 3889917532},
		{3859662829, 4292754251, 3708466080},
	}
)

// This implements the MRG32k3A pseudo-random number generator
// from Pierre L'Ecuyer.
//
// This generator is described in a paper by L'Ecuyer.
// https://www.iro.umontreal.ca/~lecuyer/myftp/papers/opres-combmrg2-1999.pdf
// Good Parameter Sets for Combined Multiple Recursive Random Number Generators.
// Operations Research, 1999, 47-1, 159--164
// ========= Summary results of Crush =========
//
//  Version:          TestU01 1.2.3
//  Generator:        rand.Uint32
//  Number of statistics:  144
//  Total CPU time:   03:07:09.75
//  The following tests gave p-values outside [0.001, 0.9990]:
//  (eps  means a value < 1.0e-300):
//  (eps1 means a value < 1.0e-15):
//
//        Test                          p-value
//  ----------------------------------------------
//   8  CollisionOver, t = 8           1.4e-12
//  ----------------------------------------------
//  All other tests were passed
type MRG32k3A struct {
	baseJumpableSource32
	s [2][3]uint32
}

func NewMRG32k3AFromStream(seed []uint32) (ans *MRG32k3A, err error) {
	err = checkEmptySeed(seed)
	if err != nil {
		return
	}

	ans = new(MRG32k3A)
	ans.spi = ans
	if len(seed) < mrg32k3a_r {
		tmp := make([]uint32, mrg32k3a_r)
		fillState(tmp, seed)
		err = ans.setSeed(tmp)
		if err != nil {
			return
		}
	} else {
		err = ans.setSeed(seed)
		if err != nil {
			return
		}
	}

	return
}

// this builds the seed slice from a split_mx64 generator using the seed provided
func NewMRG32k3A(seed int64) *MRG32k3A {
	ans := new(MRG32k3A)
	ans.spi = ans
	ans.Seed(seed)
	return ans
}

func (mrg *MRG32k3A) setSeed(seed []uint32) (err error) {

	err = checkMRGSeed(seed, mrg32k3a_r, mrg32k3a_m1, mrg32k3a_m2)
	if err != nil {
		return
	}

	mrg.stream = append([]uint32{}, seed...)

	mrg.Restart()
	return nil
}

func (mrg *MRG32k3A) Restart() {
	mrg.substream = append([]uint32{}, mrg.stream...)
	mrg.RestartSubstream()
}

func (mrg *MRG32k3A) RestartSubstream() {
	mrg.s[0][0] = mrg.substream[0]
	mrg.s[0][1] = mrg.substream[1]
	mrg.s[0][2] = mrg.substream[2]
	mrg.s[1][0] = mrg.substream[3]
	mrg.s[1][1] = mrg.substream[4]
	mrg.s[1][2] = mrg.substream[5]

	mrg.resetState()
}

func (mrg *MRG32k3A) Jump() {
	multMatVect(mrg.substream, a1p76, mrg32k3a_m1, a2p76, mrg32k3a_m2)
	mrg.RestartSubstream()
}

func (mrg *MRG32k3A) Uint32() uint32 {

	/* Component 1 */
	p1 := (mrg32k3a_a12*mrg.s[0][1] - mrg32k3a_a13n*mrg.s[0][0]) % mrg32k3a_m1

	if p1&0x80000000 == 1 {
		p1 += mrg32k3a_m1
	}
	mrg.s[0][0] = mrg.s[0][1]
	mrg.s[0][1] = mrg.s[0][2]
	mrg.s[0][2] = p1

	/* Component 2 */
	p2 := (mrg32k3a_a21*mrg.s[1][2] - mrg32k3a_a23n*mrg.s[1][0]) % mrg32k3a_m2
	if p2&0x80000000 == 1 {
		p2 += mrg32k3a_m2
	}
	mrg.s[1][0] = mrg.s[1][1]
	mrg.s[1][1] = mrg.s[1][2]
	mrg.s[1][2] = p2

	/* Combination */
	r := p1 - p2

	if p1 > p2 {
		return r
	}

	return r + mrg32k3a_m1
}

func (mrg *MRG32k3A) Seed(seed int64) {
	seeds := make([]uint32, mrg32k3a_r)
	seeder.Seed(seed)
	for j := 0; j < 3; j++ {
	again0:
		f := seeder.Intn(int(mrg32k3a_m1))
		if f == 0 {
			goto again0
		}
		seeds[j] = uint32(f)
	}
	for j := 3; j < 6; j++ {
	again1:
		f := seeder.Intn(int(mrg32k3a_m2))
		if f == 0 {
			goto again1
		}
		seeds[j] = uint32(f)
	}

	// Initialize the pool content.
	mrg.setSeed(seeds)
}

func (mrg *MRG32k3A) String() string {
	return "MRG32k3A"
}
