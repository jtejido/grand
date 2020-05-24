package source32

const (
	xoroshiro_r = 2
)

// This implements 32-bit generators with 64-bits of state.
// http://xoshiro.di.unimi.it/
type baseXoRoShiRo64 struct {
	baseSource32
	state [2]uint32
}

func (bx *baseXoRoShiRo64) setSeed(seed []uint32) {
	bx.stream = append([]uint32{}, seed...)
	bx.Restart()
}

func (bx *baseXoRoShiRo64) Seed(seed int64) {
	seeds := make([]uint32, xoroshiro_r)
	seeder.Seed(seed)
	var i int
	// Fill the remaining pairs
	for i < xoroshiro_r {
		v := seeder.Uint32()
		seeds[i] = v
		i++
	}

	// Initialize the pool content.
	bx.setSeed(seeds)
}

func (bx *baseXoRoShiRo64) Restart() {
	for i := 0; i < xoroshiro_r; i++ {
		bx.state[i] = bx.stream[i]
	}
	bx.resetState()
}
