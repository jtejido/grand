package source64

const (
	xoroshiro128_r = 2
)

var (
	xoroshiro128_pw = [...]uint64{
		0xdf900294d8f554a5, 0x170865df4b3201fc,
	}
)

// This implements 64-bit generators with 128-bits of state.
// http://xoshiro.di.unimi.it/
type baseXoRoShiRo128 struct {
	baseJumpableSource64
	state [2]uint64
}

func (bx *baseXoRoShiRo128) setSeed(seed []uint64) {
	bx.stream = append([]uint64{}, seed...)
	bx.Restart()
}

func (bx *baseXoRoShiRo128) Seed(seed int64) {
	seeds := make([]uint64, xoroshiro128_r)
	seeder.Seed(seed)
	var i int
	// Fill the remaining pairs
	for i < xoroshiro128_r {
		v := seeder.Uint64()
		seeds[i] = v
		i++
	}

	// Initialize the pool content.
	bx.setSeed(seeds)
}

func (bx *baseXoRoShiRo128) Restart() {
	bx.substream = append([]uint64{}, bx.stream...)
	bx.RestartSubstream()
}

func (bx *baseXoRoShiRo128) RestartSubstream() {
	for i := 0; i < xoroshiro128_r; i++ {
		bx.state[i] = bx.substream[i]
	}

	bx.resetState()
}

// The jump size is the equivalent of 2^64 calls to Uint64().
// It can provide up to 2^64 non-overlapping subsequences.
func (bx *baseXoRoShiRo128) Jump() {
	s := make([]uint64, xoroshiro128_r)

	for i := 0; i < len(xoroshiro128_pw); i++ {
		for b := 0; b < 64; b++ {
			if (xoroshiro128_pw[i] & (1 << uint64(b))) != 0 {
				for j := 0; j < len(bx.state); j++ {
					s[j] ^= bx.state[j]
				}
			}
			bx.Uint64()
		}
	}

	bx.substream = append([]uint64{}, s...)

	bx.RestartSubstream()
}
