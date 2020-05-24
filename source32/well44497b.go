package source32

// This implements the WELL44497B pseudo-random number generator
// from Panneton, L'Ecuyer and Matsumoto.
//
// This generator is described in a paper by Panneton, L'Ecuyer and Matsumoto
// http://www.iro.umontreal.ca/~lecuyer/myftp/papers/wellrng.pdf
// Improved Long-Period Generators Based on Linear Recurrences Modulo 2</a>
// ACM Transactions on Mathematical Software, 32, 1 (2006).
// The errata for the paper are in:
// http://www.iro.umontreal.ca/~lecuyer/myftp/papers/wellrng-errata.txt
// http://www.iro.umontreal.ca/~panneton/WELLRNG.html
// ========= Summary results of Crush =========
//
//  Version:          TestU01 1.2.3
//  Generator:        rand.Float64
//  Number of statistics:  144
//  Total CPU time:   03:17:48.42
//  The following tests gave p-values outside [0.001, 0.9990]:
//  (eps  means a value < 1.0e-300):
//  (eps1 means a value < 1.0e-15):
//
//        Test                          p-value
//  ----------------------------------------------
//  71  LinearComp, r = 0              1 - eps1
//  72  LinearComp, r = 29             1 - eps1
//  ----------------------------------------------
//  All other tests were passed
// TO-DO. Coefficients for Jump()
type WELL44497B struct {
	*WELL44497A
}

func NewWELL44497BFromStream(seed []uint32) (*WELL44497B, error) {
	ans := new(WELL44497B)
	w, err := NewWELL44497AFromStream(seed)
	if err != nil {
		return nil, err
	}
	ans.WELL44497A = w
	ans.spi = ans
	return ans, nil
}

// this builds the seed slice from a split_mx64 generator using the seed provided
func NewWELL44497B(seed int64) *WELL44497B {
	ans := new(WELL44497B)

	ans.WELL44497A = NewWELL44497A(seed)
	ans.spi = ans
	return ans
}

func (w44497 *WELL44497B) Uint32() uint32 {
	z4 := w44497.WELL44497A.Uint32()

	// Matsumoto-Kurita tempering to get a maximally equidistributed generator.
	z4 ^= (z4 << 7) & 0x93dd1400
	z4 ^= (z4 << 15) & 0xfa118000

	return z4
}

func (w44497 *WELL44497B) String() string {
	return "WELL44497B"
}
