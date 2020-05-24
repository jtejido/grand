package source64

const (
	xoshiro256_r = 4
)

var (
	xoshiro256_pw = [...]uint64{
		0x180ec6d33cfd0aba, 0xd5a61266f0c9392c, 0xa9582618e03fc9aa, 0x39abdc4529b1661c,
	}
)

// This is a base for algorithms from the Xor-Shift-Rotate family of 64-bit
// generators with 256-bits of state.
// http://xoshiro.di.unimi.it/
type baseXoShiRo256 struct {
	baseJumpableSource64
	state [4]uint64
}

func (bx *baseXoShiRo256) setSeed(seed []uint64) {
	bx.stream = append([]uint64{}, seed...)
	bx.Restart()
}

func (bx *baseXoShiRo256) Seed(seed int64) {
	seeds := make([]uint64, xoshiro256_r)
	seeder.Seed(seed)
	var i int
	// Fill the remaining pairs
	for i < xoshiro256_r {
		v := seeder.Uint64()
		seeds[i] = v
		i++
	}

	// Initialize the pool content.
	bx.setSeed(seeds)
}

func (bx *baseXoShiRo256) Restart() {
	bx.substream = append([]uint64{}, bx.stream...)
	bx.RestartSubstream()
}

func (bx *baseXoShiRo256) RestartSubstream() {
	for i := 0; i < xoshiro256_r; i++ {
		bx.state[i] = bx.substream[i]
	}

	bx.resetState()
}

func (bx *baseXoShiRo256) Jump() {
	s := make([]uint64, xoshiro256_r)

	for i := 0; i < len(xoshiro256_pw); i++ {
		for b := 0; b < 64; b++ {
			if (xoshiro256_pw[i] & (1 << uint64(b))) != 0 {
				for j := 0; j < len(bx.state); j++ {
					s[j] ^= bx.state[j]
				}
			}
			bx.Uint64()
		}
	}

	bx.substream = append([]uint64{}, s...)

	bx.Restart()
}
