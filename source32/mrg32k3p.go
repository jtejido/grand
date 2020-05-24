package source32

const (
	mrg32k3p_m1     uint32 = 2147483647 //2^31 - 1
	mrg32k3p_m2     uint32 = 2147462579 //2^31 - 21069
	mrg32k3p_mask12 uint32 = 511        //2^9 - 1
	mrg32k3p_mask13 uint32 = 16777215   //2^24 - 1
	mrg32k3p_mask2  uint32 = 65535      //2^16 - 1
	mrg32k3p_mult2  uint32 = 21069
	mrg32k3p_r             = 6
)

var (
	a1p72 = [][]uint32{
		{1516919229, 758510237, 499121365},
		{1884998244, 1516919229, 335398200},
		{601897748, 1884998244, 358115744},
	}
	a2p72 = [][]uint32{
		{1228857673, 1496414766, 954677935},
		{1133297478, 1407477216, 1496414766},
		{2002613992, 1639496704, 1407477216},
	}
)

// This implements the MRG32k3P pseudo-random number generator
// from Pierre L'Ecuyer.
//
// This generator is described in a paper by L'Ecuyer.
// https://www.informs-sim.org/wsc00papers/090.PDF
// Fast Combined Multiple Recursive Generators with Multipliers of the Form a=±2q±2r.
// Proceedings of the 2000 Winter Simulation Conference, Dec. 2000, 683--689
// https://github.com/clMathLibraries/clRNG
// TODO: unit test
type MRG32k3P struct {
	baseJumpableSource32
	s [2][3]uint32
}

func NewMRG32k3PFromStream(seed []uint32) (ans *MRG32k3P, err error) {
	err = checkEmptySeed(seed)
	if err != nil {
		return
	}
	ans = new(MRG32k3P)
	ans.spi = ans
	if len(seed) < mrg32k3p_r {
		tmp := make([]uint32, mrg32k3p_r)
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
func NewMRG32k3P(seed int64) *MRG32k3P {
	ans := new(MRG32k3P)
	ans.spi = ans
	ans.Seed(seed)
	return ans
}

func (mrg *MRG32k3P) setSeed(seed []uint32) (err error) {

	err = checkMRGSeed(seed, mrg32k3p_r, mrg32k3p_m1, mrg32k3p_m2)
	if err != nil {
		return
	}

	mrg.stream = append([]uint32{}, seed...)

	mrg.Restart()
	return nil
}

func (mrg *MRG32k3P) Restart() {
	mrg.substream = append([]uint32{}, mrg.stream...)
	mrg.RestartSubstream()

}

func (mrg *MRG32k3P) RestartSubstream() {
	mrg.s[0][0] = mrg.substream[0]
	mrg.s[0][1] = mrg.substream[1]
	mrg.s[0][2] = mrg.substream[2]
	mrg.s[1][0] = mrg.substream[3]
	mrg.s[1][1] = mrg.substream[4]
	mrg.s[1][2] = mrg.substream[5]
	mrg.resetState()
}

func (mrg *MRG32k3P) Jump() {
	multMatVect(mrg.substream, a1p72, mrg32k3p_m1, a1p72, mrg32k3p_m2)
	mrg.RestartSubstream()
}

func (mrg *MRG32k3P) Uint32() uint32 {

	//first component
	y1 := ((mrg.s[0][1] & mrg32k3p_mask12) << 22) + (mrg.s[0][1] >> 9) + ((mrg.s[0][2] & mrg32k3p_mask13) << 7) + (mrg.s[0][2] >> 24)
	if y1&0x80000000 == 1 || y1 >= mrg32k3p_m1 {
		y1 -= mrg32k3p_m1
	}
	y1 += mrg.s[0][2]
	if y1&0x80000000 == 1 || y1 >= mrg32k3p_m1 {
		y1 -= mrg32k3p_m1
	}

	mrg.s[0][2] = mrg.s[0][1]
	mrg.s[0][1] = mrg.s[0][0]
	mrg.s[0][0] = y1

	//second component
	y1 = ((mrg.s[1][0] & mrg32k3p_mask2) << 15) + (mrg32k3p_mult2 * (mrg.s[1][0] >> 16))
	if y1&0x80000000 == 1 || y1 >= mrg32k3p_m2 {
		y1 -= mrg32k3p_m2
	}
	y2 := ((mrg.s[1][2] & mrg32k3p_mask2) << 15) + (mrg32k3p_mult2 * (mrg.s[1][2] >> 16))
	if y2&0x80000000 == 1 || y2 >= mrg32k3p_m2 {
		y2 -= mrg32k3p_m2
	}
	y2 += mrg.s[1][2]
	if y2&0x80000000 == 1 || y2 >= mrg32k3p_m2 {
		y2 -= mrg32k3p_m2
	}
	y2 += y1
	if y2&0x80000000 == 1 || y2 >= mrg32k3p_m2 {
		y2 -= mrg32k3p_m2
	}

	mrg.s[1][2] = mrg.s[1][1]
	mrg.s[1][1] = mrg.s[1][0]
	mrg.s[1][0] = y2

	//Must never return either 0 or 1
	r := mrg.s[0][0] - mrg.s[1][0]

	if mrg.s[0][0] > mrg.s[1][0] {
		return r
	}

	return r + mrg32k3p_m1
}

func (mrg *MRG32k3P) Seed(seed int64) {
	seeds := make([]uint32, mrg32k3p_r)
	seeder.Seed(seed)
	for j := 0; j < 3; j++ {
	again0:
		f := seeder.Intn(int(mrg32k3p_m1))
		if f == 0 {
			goto again0
		}
		seeds[j] = uint32(f)
	}
	for j := 3; j < 6; j++ {
	again1:
		f := seeder.Intn(int(mrg32k3p_m2))
		if f == 0 {
			goto again1
		}
		seeds[j] = uint32(f)
	}

	// Initialize the pool content.
	mrg.setSeed(seeds)
}

func (mrg *MRG32k3P) String() string {
	return "MRG32k3P"
}
