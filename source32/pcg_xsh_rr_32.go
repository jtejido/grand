package source32

/**
 * A Permuted Congruential Generator (PCG) that is composed of a 64-bit Linear Congruential
 * Generator (LCG) combined with the XSH-RR (xorshift; random rotate) output
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
// Total CPU time:   03:41:16.85
// The following tests gave p-values outside [0.001, 0.9990]:
// (eps  means a value < 1.0e-300):
// (eps1 means a value < 1.0e-15):

//       Test                          p-value
// ----------------------------------------------
// 19  ClosePairs mNP2, t = 3          9.4e-4
// ----------------------------------------------
// All other tests were passed
type PcgXshRr32 struct {
	basePCG6432
}

func NewPcgXshRr32FromStream(seed []uint64) (*PcgXshRr32, error) {
	err := checkEmptySeed64(seed)
	if err != nil {
		return nil, err
	}

	ans := new(PcgXshRr32)
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
func NewPcgXshRr32(seed int64) *PcgXshRr32 {
	ans := new(PcgXshRr32)
	ans.spi = ans
	ans.Seed(seed)
	return ans
}

func (rr32 *PcgXshRr32) Uint32() uint32 {
	oldstate := rr32.state
	rr32.state = oldstate*PCG6432_MULT + rr32.increment
	xorshifted := uint32(((oldstate >> 18) ^ oldstate) >> 27)
	rot := uint32(oldstate >> 59)
	return (xorshifted >> rot) | (xorshifted << ((-rot) & 31))
}

func (rr32 *PcgXshRr32) String() string {
	return "PcgXshRr32"
}
