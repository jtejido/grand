package source32

// This implements the WELL19937C pseudo-random number generator
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
//  Total CPU time:   02:59:07.80
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
type WELL19937C struct {
	*WELL19937A
}

func NewWELL19937CFromStream(seed []uint32) (*WELL19937C, error) {

	ans := new(WELL19937C)

	w, err := NewWELL19937AFromStream(seed)

	if err != nil {
		return nil, err
	}

	ans.WELL19937A = w
	ans.spi = ans
	return ans, nil
}

// this builds the seed slice from a split_mx64 generator using the seed provided
func NewWELL19937C(seed int64) *WELL19937C {
	ans := new(WELL19937C)

	ans.WELL19937A = NewWELL19937A(seed)
	ans.spi = ans
	return ans
}

func (w19937 *WELL19937C) Uint32() uint32 {
	z4 := w19937.WELL19937A.Uint32()
	// Matsumoto-Kurita tempering to get a maximally equidistributed generator.
	z4 ^= (z4 << 7) & 0xe46e1700
	z4 ^= (z4 << 15) & 0x9b868000

	return z4
}

func (w19937 *WELL19937C) String() string {
	return "WELL19937C"
}
