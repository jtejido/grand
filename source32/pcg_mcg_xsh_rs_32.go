package source32

/**
 * A Permuted Congruential Generator (PCG) that is composed of a 64-bit Multiplicative Congruential
 * Generator (MCG) combined with the XSH-RR (xorshift; random rotate) output
 * transformation to create 32-bit output.
 *
 * State size is 64 bits and the period is 2^62
 *
 *  PCG, A Family of Better Random Number Generators
 */
// ========= Summary results of Crush =========

// Version:          TestU01 1.2.3
// Generator:        rand.Float64
// Number of statistics:  144
// Total CPU time:   02:52:34.69
// The following tests gave p-values outside [0.001, 0.9990]:
// (eps  means a value < 1.0e-300):
// (eps1 means a value < 1.0e-15):

//       Test                          p-value
// ----------------------------------------------
// 19  ClosePairs mNP, t = 3           0.9992
// ----------------------------------------------
// All other tests were passed
type PcgMcgXshRs32 struct {
	basePCGMCG6432
}

// this builds the seed slice from a split_mx64 generator using the seed provided
func NewPcgMcgXshRs32(seed uint64) *PcgMcgXshRs32 {
	ans := new(PcgMcgXshRs32)
	ans.spi = ans
	ans.setSeed(seed)
	return ans
}

func (rs32 *PcgMcgXshRs32) Uint32() uint32 {
	x := rs32.state
	rs32.state = rs32.state * PCGMCG6432_MULT

	count := uint32(x >> 61)
	return uint32((x ^ (x >> 22)) >> (22 + count))
}

func (rs32 *PcgMcgXshRs32) String() string {
	return "PcgMcgXshRs32"
}
