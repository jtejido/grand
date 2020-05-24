package source64

const (
	xorshift_r = 16
)

var (
	xorshift_pw = [...]uint64{
		0x84242f96eca9c41d, 0xa3c65b8776f96855, 0x5b34a39f070b5837, 0x4489affce4f31a1e,
		0x2ffeeb0a48316f40, 0xdc2d9891fe68c022, 0x3659132bb12fea70, 0xaac17d8efa43cab8,
		0xc4cb815590989b13, 0x5ee975283d71c93b, 0x691548c86c1bd540, 0x7910c41d10a1e6a5,
		0x0b5fc64563b3e2a8, 0x047f7684e9fc949d, 0xb99181f2d8f685ca, 0x284600e3f30e38c3,
	}
)

// A fast RNG implementing the XorShift1024* algorithm.
//
// Note: This has been superseded by XorShift1024*_Phi (see test for the constant).
// The sequences emitted by both generators are correlated.
//
// http://xorshift.di.unimi.it/xorshift1024star.c
// https://en.wikipedia.org/wiki/Xorshift
// ========= Summary results of Crush =========
//
//  Version:          TestU01 1.2.3
//  Generator:        rand.Float64
//  Number of statistics:  144
//  Total CPU time:   02:53:45.51
//  The following tests gave p-values outside [0.001, 0.9990]:
//  (eps  means a value < 1.0e-300):
//  (eps1 means a value < 1.0e-15):
//
//        Test                          p-value
//  ----------------------------------------------
//  89  HammingIndep, L = 1200          5.5e-4
//  ----------------------------------------------
//  All other tests were passed
type XorShift1024Star struct {
	baseJumpableSource64
	state             [16]uint64
	multiplier, index uint64
}

func NewXorShift1024StarFromStream(seed []uint64) (*XorShift1024Star, error) {
	return newXorShiftStreamWithMultiplier(seed, 1181783497276652981)
}

func NewXorShift1024StarPhiFromStream(seed []uint64) (*XorShift1024Star, error) {
	return newXorShiftStreamWithMultiplier(seed, 0x9e3779b97f4a7c13)
}

// This builds the seed slice from a split_mx64 generator using the seed provided.
func NewXorShift1024Star(seed int64) *XorShift1024Star {
	ans := new(XorShift1024Star)
	ans.multiplier = 1181783497276652981
	ans.Seed(seed)
	return ans
}

// Th xorshift1024starphi uses a different multiplier
func NewXorShift1024StarPhi(seed int64) *XorShift1024Star {
	ans := new(XorShift1024Star)
	ans.multiplier = 0x9e3779b97f4a7c13
	ans.Seed(seed)
	return ans
}

func newXorShiftStreamWithMultiplier(seed []uint64, multiplier uint64) (*XorShift1024Star, error) {
	err := checkEmptySeed(seed)
	if err != nil {
		return nil, err
	}

	ans := new(XorShift1024Star)
	ans.multiplier = multiplier
	if len(seed) < xorshift_r {
		tmp := make([]uint64, xorshift_r)
		fillState(tmp, seed)
		ans.setSeed(tmp)
	} else {
		ans.setSeed(seed)
	}

	return ans, nil
}

func (xs *XorShift1024Star) setSeed(seed []uint64) {
	xs.stream = append([]uint64{}, seed...)
	xs.Restart()
}

func (xs *XorShift1024Star) Seed(seed int64) {
	seeds := make([]uint64, xorshift_r)
	seeder.Seed(seed)
	var i int
	// Fill the remaining pairs
	for i < xorshift_r {
		v := seeder.Uint64()
		seeds[i] = v
		i++
	}

	// Initialize the pool content.
	xs.setSeed(seeds)
}

func (xs *XorShift1024Star) Restart() {
	xs.substream = append([]uint64{}, xs.stream...)
	xs.RestartSubstream()
}

func (xs *XorShift1024Star) RestartSubstream() {
	for i := 0; i < xorshift_r; i++ {
		xs.state[i] = xs.substream[i]
	}

	xs.index = 0
	xs.resetState()
}

// Perform the jump to advance the generator state.
func (xs *XorShift1024Star) Jump() {
	s := make([]uint64, xorshift_r)
	xs.substream = make([]uint64, xorshift_r)

	for i := 0; i < len(xorshift_pw); i++ {
		for b := 0; b < 64; b++ {
			if (xorshift_pw[i] & (1 << uint64(b))) != 0 {
				for j := 0; j < len(xs.state); j++ {
					s[j] ^= xs.state[(uint64(i)+xs.index)&15]
				}
			}
			xs.Uint64()
		}
	}

	for i := 0; i < xorshift_r; i++ {
		xs.substream[(uint64(i)+xs.index)&15] = s[i]
	}

	xs.Restart()
}

func (xs *XorShift1024Star) Uint64() uint64 {
	s0 := xs.state[xs.index]
	xs.index = (xs.index + 1) & 15
	s1 := xs.state[xs.index]
	s1 ^= s1 << 31                                         // a
	xs.state[xs.index] = s1 ^ s0 ^ (s1 >> 11) ^ (s0 >> 30) // b,c
	return xs.state[xs.index] * xs.multiplier
}

func (xs *XorShift1024Star) String() string {
	return "XorShift1024Star"
}
