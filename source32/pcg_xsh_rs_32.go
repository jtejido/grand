package source32

/**
 * A Permuted Congruential Generator (PCG) that is composed of a 64-bit Linear Congruential
 * Generator (LCG) combined with the XSH-RS (xorshift; random shift) output
 * transformation to create 32-bit output.
 *
 * State size is 128 bits and the period is 2^64
 *
 *  PCG, A Family of Better Random Number Generators
 */
// ========= Summary results of Crush =========

// Version:          TestU01 1.2.3
// Generator:        rand.Float64
// Number of statistics:  144
// Total CPU time:   02:59:44.31
// The following tests gave p-values outside [0.001, 0.9990]:
// (eps  means a value < 1.0e-300):
// (eps1 means a value < 1.0e-15):

//       Test                          p-value
// ----------------------------------------------
// 76  LongestHeadRun, r = 0          1 -  9.0e-5
// ----------------------------------------------
// All other tests were passed
type PcgXshRs32 struct {
	basePCG6432
}

func NewPcgXshRs32FromStream(seed []uint64) (*PcgXshRs32, error) {
	err := checkEmptySeed64(seed)
	if err != nil {
		return nil, err
	}

	ans := new(PcgXshRs32)
	ans.spi = ans
	if len(seed) < PCG6432_SEED_SIZE {
		tmp := make([]uint64, PCG6432_SEED_SIZE)
		fillState64(tmp, seed)
		ans.setSeed(tmp)
	} else {
		ans.setSeed(seed)
	}

	return ans, nil
}

// this builds the seed slice from a split_mx64 generator using the seed provided
func NewPcgXshRs32(seed int64) *PcgXshRs32 {
	ans := new(PcgXshRs32)
	ans.spi = ans
	ans.Seed(seed)
	return ans
}

func (rr32 *PcgXshRs32) Uint32() uint32 {
	x := rr32.state
	rr32.state = rr32.state*PCG6432_MULT + rr32.increment

	count := uint32(x >> 61)
	return uint32((x ^ (x >> 22)) >> (22 + count))
}

func (rr32 *PcgXshRs32) String() string {
	return "PcgXshRs32"
}
