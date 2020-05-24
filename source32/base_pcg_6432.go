package source32

const (
	PCG6432_SEED_SIZE = 2
	PCG6432_MULT      = 6364136223846793005
)

type basePCG6432 struct {
	baseSource32
	stream64         []uint64
	state, increment uint64
}

func (bpcg *basePCG6432) setSeed(seed []uint64) {
	bpcg.increment = (seed[1] << 1) | 1
	bpcg.state = (seed[0]+bpcg.increment)*PCG6432_MULT + bpcg.increment
	bpcg.stream64 = append([]uint64{}, seed...)
	bpcg.Restart()
}

func (bpcg *basePCG6432) Seed(seed int64) {
	seeds := make([]uint64, PCG6432_SEED_SIZE)
	seeder.Seed(seed)
	var i int
	// Fill the remaining pairs
	for i < PCG6432_SEED_SIZE {
		v := seeder.Uint64()
		seeds[i] = v
		i++
	}

	// Initialize the pool content.
	bpcg.setSeed(seeds)
}

func (bpcg *basePCG6432) Restart() {
	bpcg.state = (bpcg.stream64[0]+bpcg.increment)*PCG6432_MULT + bpcg.increment
	bpcg.increment = (bpcg.stream64[1] << 1) | 1
	bpcg.resetState()
}
