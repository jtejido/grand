package source32

const (
	block_size_well = 32
)

type baseWellNonJumpable struct {
	baseSource32
	state_idx int
	state     []uint32
}

func (bw *baseWellNonJumpable) advanceSeed(p []uint32, seedSize int) {
	var b uint32
	x := make([]uint32, seedSize)

	for i := 0; i < seedSize; i++ {
		bw.state[i] = bw.stream[i]
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
		bw.stream[i] = x[i]
	}
}

func (bw *baseWellNonJumpable) setSeed(seed []uint32) {
	bw.stream = append([]uint32{}, seed...)
	bw.Restart()
}

func (bw *baseWellNonJumpable) Restart() {
	bw.state_idx = 0
	bw.state = append([]uint32{}, bw.stream...)
	bw.resetState()
}

type indexTable struct {
	iRm1, iRm2, i1, i2, i3 []uint32
}

// Internal index table given a specific WELL algo's parameters.
func newIndexTable(k int, m1, m2, m3 uint32) *indexTable {
	ans := new(indexTable)
	ans.iRm1 = make([]uint32, k)
	ans.iRm2 = make([]uint32, k)
	ans.i1 = make([]uint32, k)
	ans.i2 = make([]uint32, k)
	ans.i3 = make([]uint32, k)
	for jj := 0; jj < k; jj++ {
		r := uint32(k)
		j := uint32(jj)
		ans.iRm1[j] = (j + r - 1) % r
		ans.iRm2[j] = (j + r - 2) % r
		ans.i1[j] = (j + m1) % r
		ans.i2[j] = (j + m2) % r
		ans.i3[j] = (j + m3) % r
	}

	return ans
}

// Returns the predecessor of the given index modulo the table size.
func (it *indexTable) indexPredAt(index int) uint32 {
	return it.iRm1[index]
}

// Returns the second predecessor of the given index modulo the table size.
func (it *indexTable) indexPred2At(index int) uint32 {
	return it.iRm2[index]
}

// Returns index + M1 modulo the table size.
func (it *indexTable) indexM1At(index int) uint32 {
	return it.i1[index]
}

// Returns index + M2 modulo the table size.
func (it *indexTable) indexM2At(index int) uint32 {
	return it.i2[index]
}

// Returns index + M3 modulo the table size.
func (it *indexTable) indexM3At(index int) uint32 {
	return it.i3[index]
}
