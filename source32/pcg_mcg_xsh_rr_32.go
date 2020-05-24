package source32

// A Permuted Congruential Generator (PCG) that is composed of a 64-bit Multiplicative Congruential
// Generator (MCG) combined with the XSH-RR (xorshift; random shift) output
// transformation to create 32-bit output.
//
// State size is 64 bits and the period is 2^62
//
//  PCG, A Family of Better Random Number Generators
//
// ========= Summary results of Crush =========
//
// Version:          TestU01 1.2.3
// Generator:        rand.Float64
// Number of statistics:  144
// Total CPU time:   03:01:14.26
//
// All tests were passed
type PcgMcgXshRr32 struct {
	basePCGMCG6432
}

func NewPcgMcgXshRr32(seed uint64) *PcgMcgXshRr32 {
	ans := new(PcgMcgXshRr32)
	ans.spi = ans
	ans.setSeed(seed)
	return ans
}

func (rr32 *PcgMcgXshRr32) Uint32() uint32 {
	oldstate := rr32.state
	rr32.state = rr32.state * PCGMCG6432_MULT
	xorshifted := uint32(((oldstate >> 18) ^ oldstate) >> 27)
	rot := uint32(oldstate >> 59)
	return (xorshifted >> rot) | (xorshifted << ((-rot) & 31))
}

func (rr32 *PcgMcgXshRr32) String() string {
	return "PcgMcgXshRr32"
}
