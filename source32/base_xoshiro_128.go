package source32

const (
	xoshiro128_r = 4
)

var (
	xoshiro128_pw = [...]uint32{
		0x8764000b, 0xf542d2d3, 0x6fa035c3, 0x77f2db5b,
	}
)

// This implements 32-bit generators with 128-bits of state.
// http://xoshiro.di.unimi.it/
type baseXoShiRo128 struct {
	baseJumpableSource32
	state [4]uint32
}

func (bx *baseXoShiRo128) setSeed(seed []uint32) {
	bx.stream = append([]uint32{}, seed...)
	bx.Restart()
}

func (bx *baseXoShiRo128) Seed(seed int64) {
	seeds := make([]uint32, xoshiro128_r)
	seeder.Seed(seed)
	var i int
	// Fill the remaining pairs
	for i < xoshiro128_r {
		v := seeder.Uint32()
		seeds[i] = v
		i++
	}

	// Initialize the pool content.
	bx.setSeed(seeds)
}

func (bx *baseXoShiRo128) Restart() {
	bx.substream = append([]uint32{}, bx.stream...)
	bx.RestartSubstream()
}

func (bx *baseXoShiRo128) RestartSubstream() {
	for i := 0; i < xoshiro128_r; i++ {
		bx.state[i] = bx.substream[i]
	}
	bx.resetState()
}

func (bx *baseXoShiRo128) Jump() {
	s := make([]uint32, xoshiro128_r)
	for i := 0; i < len(xoshiro128_pw); i++ {
		var b uint32
		for ; b < 32; b++ {
			if (xoshiro128_pw[i] & (1 << b)) != 0 {
				for j := 0; j < len(bx.state); j++ {
					s[j] ^= bx.state[j]
				}
			}
			bx.Uint32()
		}
	}

	bx.substream = append([]uint32{}, s...)

	bx.RestartSubstream()
}
