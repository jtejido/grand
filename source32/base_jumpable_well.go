package source32

type baseJumpableWell struct {
	baseJumpableSource32
	state_idx int
	state     []uint32
}

func (bw *baseJumpableWell) advanceSeed(p []uint32, seedSize int) {
	var b uint32
	x := make([]uint32, seedSize)

	for i := 0; i < seedSize; i++ {
		bw.state[i] = bw.substream[i]
	}

	bw.state_idx = 0

	for j := 0; j < seedSize; j++ {
		b = p[j]
		for k := 0; k < block_size_well; k++ {
			if (b & 1) == 1 {
				for i := 0; i < seedSize; i++ {
					x[i] ^= bw.state[(bw.state_idx+i)&(seedSize-1)]
				}
			}
			b >>= 1

			bw.Uint32()
		}
	}

	for i := 0; i < seedSize; i++ {
		bw.substream[i] = x[i]
	}
}

func (bw *baseJumpableWell) setSeed(seed []uint32) {
	bw.stream = append([]uint32{}, seed...)
	bw.Restart()
}

func (bw *baseJumpableWell) Restart() {
	bw.substream = append([]uint32{}, bw.stream...)
	bw.RestartSubstream()
}

func (bw *baseJumpableWell) RestartSubstream() {
	bw.state_idx = 0
	bw.state = append([]uint32{}, bw.substream...)
	bw.resetState()
}
