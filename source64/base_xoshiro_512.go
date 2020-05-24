package source64

const (
	xoshiro512_r = 8
)

var (
	xoshiro512_pw = [...]uint64{
		0x33ed89b6e7a353f9, 0x760083d7955323be, 0x2837f2fbb5f22fae, 0x4b8c5674d309511c,
		0xb11ac47a7ba28c25, 0xf1be7667092bcc1c, 0x53851efdb6df0aaf, 0x1ebbc8b23eaf25db,
	}
)

// This is a base for algorithms from the Xor-Shift-Rotate family of 64-bit
// generators with 256-bits of state.
// http://xoshiro.di.unimi.it/
type baseXoShiRo512 struct {
	baseJumpableSource64
	state [8]uint64
}

func (bx *baseXoShiRo512) setSeed(seed []uint64) {
	bx.stream = append([]uint64{}, seed...)
	bx.Restart()
}

func (bx *baseXoShiRo512) Seed(seed int64) {
	seeds := make([]uint64, xoshiro512_r)
	seeder.Seed(seed)
	var i int
	// Fill the remaining pairs
	for i < xoshiro512_r {
		v := seeder.Uint64()
		seeds[i] = v
		i++
	}

	// Initialize the pool content.
	bx.setSeed(seeds)
}

func (bx *baseXoShiRo512) Restart() {
	bx.substream = append([]uint64{}, bx.stream...)
	bx.RestartSubstream()
}

func (bx *baseXoShiRo512) RestartSubstream() {
	for i := 0; i < xoshiro512_r; i++ {
		bx.state[i] = bx.substream[i]
	}

	bx.resetState()
}

func (bx *baseXoShiRo512) Jump() {
	s := make([]uint64, xoshiro512_r)

	for i := 0; i < len(xoshiro512_pw); i++ {
		for b := 0; b < 64; b++ {
			if (xoshiro512_pw[i] & (1 << uint64(b))) != 0 {
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
