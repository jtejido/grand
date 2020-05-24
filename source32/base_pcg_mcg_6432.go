package source32

const (
	PCGMCG6432_MULT = 6364136223846793005
)

type basePCGMCG6432 struct {
	baseSource32
	state, temp_state uint64
}

func (bpcgmcg *basePCGMCG6432) setSeed(seed uint64) {
	bpcgmcg.temp_state = seed | 3
	bpcgmcg.Restart()
}

func (bpcgmcg *basePCGMCG6432) Seed(seed int64) {
	bpcgmcg.setSeed(uint64(seed))
}

func (bpcgmcg *basePCGMCG6432) Restart() {
	bpcgmcg.state = bpcgmcg.temp_state
	bpcgmcg.resetState()
}
