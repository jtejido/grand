package source32

const (
	lfsr113_r = 4
)

// This implements the LFSR113 pseudo-random number generator
// from Pierre L'Ecuyer.
//
// This generator is described in a paper by L'Ecuyer.
// http://www-labs.iro.umontreal.ca/~lecuyer/myftp/papers/tausme.ps
// Maximally equidistributed combined Tausworthe generators.
// Mathematics of Computation, 65, 213 (1996), 203--213.
//
// ========= Summary results of Crush =========

//  Version:          TestU01 1.2.3
//  Generator:        rand.Float64
//  Number of statistics:  144
//  Total CPU time:   03:11:53.53
//  The following tests gave p-values outside [0.001, 0.9990]:
//  (eps  means a value < 1.0e-300):
//  (eps1 means a value < 1.0e-15):

//        Test                          p-value
//  ----------------------------------------------
//  18  ClosePairs mNP1, t = 2          0.9991
//  38  Permutation, r = 15             0.9998
//  44  MaxOft AD, t = 30              1 -  2.9e-6
//  58  MatrixRank, 300 x 300            eps
//  59  MatrixRank, 300 x 300            eps
//  60  MatrixRank, 1200 x 1200          eps
//  61  MatrixRank, 1200 x 1200          eps
//  71  LinearComp, r = 0              1 - eps1
//  72  LinearComp, r = 29             1 - eps1
//  ----------------------------------------------
//  All other tests were passed
type LFSR113 struct {
	baseJumpableSource32
	state [4]uint32
}

func NewLFSR113FromStream(seed []uint32) (*LFSR113, error) {
	err := checkEmptySeed(seed)
	if err != nil {
		return nil, err
	}

	ans := new(LFSR113)
	ans.spi = ans
	if len(seed) < lfsr113_r {
		tmp := make([]uint32, lfsr113_r)
		fillState(tmp, seed)
		ans.setSeed(tmp)
	} else {
		ans.setSeed(seed)
	}

	return ans, nil
}

// this builds the seed slice from a split_mx64 generator using the seed provided
func NewLFSR113(seed int64) *LFSR113 {
	ans := new(LFSR113)
	ans.spi = ans
	ans.Seed(seed)
	return ans
}

func (lfsr *LFSR113) Uint32() uint32 {

	b := ((lfsr.state[0] << 6) ^ lfsr.state[0]) >> 13
	lfsr.state[0] = ((lfsr.state[0] & 4294967294) << 18) ^ b
	b = ((lfsr.state[1] << 2) ^ lfsr.state[1]) >> 27
	lfsr.state[1] = ((lfsr.state[1] & 4294967288) << 2) ^ b
	b = ((lfsr.state[2] << 13) ^ lfsr.state[2]) >> 21
	lfsr.state[2] = ((lfsr.state[2] & 4294967280) << 7) ^ b
	b = ((lfsr.state[3] << 3) ^ lfsr.state[3]) >> 12
	lfsr.state[3] = ((lfsr.state[3] & 4294967168) << 13) ^ b
	return (lfsr.state[0] ^ lfsr.state[1] ^ lfsr.state[2] ^ lfsr.state[3])
}

func (lfsr *LFSR113) setSeed(seed []uint32) {
	lfsr.checkSeed(seed)
	lfsr.stream = append([]uint32{}, seed...)
	lfsr.Restart()
}

func (lfsr *LFSR113) Restart() {
	lfsr.substream = append([]uint32{}, lfsr.stream...)
	lfsr.RestartSubstream()
}

func (lfsr *LFSR113) RestartSubstream() {
	for j := 0; j < len(lfsr.substream); j++ {
		lfsr.state[j] = lfsr.substream[j]
	}

	lfsr.resetState()
}

func (lfsr *LFSR113) Jump() {

	z := lfsr.substream[0] & 4294967294
	b := (z << 6) ^ z
	z = (z) ^ (z << 3) ^ (z << 4) ^ (z << 6) ^ (z << 7) ^
		(z << 8) ^ (z << 10) ^ (z << 11) ^ (z << 13) ^ (z << 14) ^
		(z << 16) ^ (z << 17) ^ (z << 18) ^ (z << 22) ^
		(z << 24) ^ (z << 25) ^ (z << 26) ^ (z << 28) ^ (z << 30)
	z ^= (b >> 1) ^ (b >> 3) ^ (b >> 5) ^ (b >> 6) ^
		(b >> 7) ^ (b >> 9) ^ (b >> 13) ^ (b >> 14) ^
		(b >> 15) ^ (b >> 17) ^ (b >> 18) ^ (b >> 20) ^
		(b >> 21) ^ (b >> 23) ^ (b >> 24) ^ (b >> 25) ^
		(b >> 26) ^ (b >> 27) ^ (b >> 30)
	lfsr.substream[0] = z

	z = lfsr.substream[1] & 4294967288
	b = z ^ (z << 1)
	b ^= (b << 2)
	b ^= (b << 4)
	b ^= (b << 8)

	b <<= 8
	b ^= (z << 22) ^ (z << 25) ^ (z << 27)
	if (z & 0x80000000) == 1 {
		b ^= 0xABFFF000
	}
	if (z & 0x40000000) == 1 {
		b ^= 0x55FFF800
	}
	z = b ^ (z >> 7) ^ (z >> 20) ^ (z >> 21)
	lfsr.substream[1] = z

	z = lfsr.substream[2] & 4294967280
	b = (z << 13) ^ z
	z = (b >> 3) ^ (b >> 17) ^ (z << 10) ^ (z << 11) ^ (z << 25)
	lfsr.substream[2] = z

	z = lfsr.substream[3] & 4294967168
	b = (z << 3) ^ z
	z = (z << 14) ^ (z << 16) ^ (z << 20) ^ (b >> 5) ^ (b >> 9) ^ (b >> 11)
	lfsr.substream[3] = z
	lfsr.RestartSubstream()
}

func (lfsr *LFSR113) checkSeed(seed []uint32) {
	if (seed[0] >= 0 && seed[0] < 2) || (seed[1] >= 0 && seed[1] < 8) || (seed[2] >= 0 && seed[2] < 16) || (seed[3] >= 0 && seed[3] < 128) {
		panic("The seed elements must be greater than 1, 7, 15 and 127 respectively")
	}
}

func (lfsr *LFSR113) Seed(seed int64) {
	seeds := make([]uint32, lfsr113_r)
	seeder.Seed(seed)

	for j := 0; j < lfsr113_r; j++ {
	again:
		f := seeder.Uint32()

		if (j == 0 && f <= 1) || (j == 1 && f <= 7) || (j == 2 && f <= 15) || (j == 3 && f <= 127) {
			goto again
		}

		seeds[j] = f

	}

	// Initialize the pool content.
	lfsr.setSeed(seeds)
}

func (lfsr *LFSR113) String() string {
	return "LFSR113"
}
