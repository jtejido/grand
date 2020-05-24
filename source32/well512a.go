package source32

const (
	well512a_r  int    = 16
	well512a_m1 uint32 = 13
	well512a_m2 uint32 = 9
	well512a_m3 uint32 = 5
)

// This table represents the 512 coefficients of the following polynomial
// (z^(2^200) mod P(z)) mod 2
// P(z) is the characteristic polynomial of the generator.
// Characteristic Polynomial is defined as:
// P(z) = det(A − zI) = zk − α1zk−1 −...− αk−1z − αk
// where I is the identity matrix and each αj is in F2 (the finite field with two elements, 0 and 1). For each j, the sequences
// {xi, j , i ≥ 0} and { yi, j , i ≥ 0} both obey the linear recurrence
// xi, j = (α1xi−1, j +...+ αk xi−k, j) mod 2
var (
	well512a_pw = []uint32{
		0x280009a9, 0x31e221d0, 0xa00c0296, 0x763d492b,
		0x63875b75, 0xef2acc3a, 0x1400839f, 0x5e0c8526,
		0x514e11b, 0x56b398e4, 0x9436c8b9, 0xa6d8130b,
		0xc0a48a78, 0x26ad57d0, 0xa3a0c62a, 0x3ff16c9b,
	}
)

// This implements the WELL512a pseudo-random number generator
// from Panneton, L'Ecuyer and Matsumoto.
//
// This generator is described in a paper by Panneton, L'Ecuyer and Matsumoto
// http://www.iro.umontreal.ca/~lecuyer/myftp/papers/wellrng.pdf
// Improved Long-Period Generators Based on Linear Recurrences Modulo 2</a>
// ACM Transactions on Mathematical Software, 32, 1 (2006).
// The errata for the paper are in:
// http://www.iro.umontreal.ca/~lecuyer/myftp/papers/wellrng-errata.txt
// http://www.iro.umontreal.ca/~panneton/WELLRNG.html
//
// ========= Summary results of Crush =========
//
//  Version:          TestU01 1.2.3
//  Generator:        Generator.Float64
//  Number of statistics:  144
//  Total CPU time:   02:58:06.79
//  The following tests gave p-values outside [0.001, 0.9990]:
//  (eps  means a value < 1.0e-300):
//  (eps1 means a value < 1.0e-15):
//
//        Test                          p-value
//  ----------------------------------------------
//  60  MatrixRank, 1200 x 1200          eps
//  61  MatrixRank, 1200 x 1200          eps
//  71  LinearComp, r = 0              1 - eps1
//  72  LinearComp, r = 29             1 - eps1
//  ----------------------------------------------
//  All other tests were passed
type WELL512A struct {
	baseJumpableWell
	table *indexTable
}

func NewWELL512AFromStream(seed []uint32) (*WELL512A, error) {
	err := checkEmptySeed(seed)
	if err != nil {
		return nil, err
	}

	ans := new(WELL512A)
	ans.spi = ans

	if len(seed) < well512a_r {
		tmp := make([]uint32, well512a_r)
		fillState(tmp, seed)
		ans.setSeed(tmp)
	} else {
		ans.setSeed(seed)
	}

	ans.table = newIndexTable(well512a_r, well512a_m1, well512a_m2, well512a_m3)
	return ans, nil
}

// this builds the seed slice from a split_mx64 generator using the seed provided
func NewWELL512A(seed int64) *WELL512A {
	ans := new(WELL512A)
	ans.spi = ans

	ans.table = newIndexTable(well512a_r, well512a_m1, well512a_m2, well512a_m3)
	ans.Seed(seed)

	return ans
}

func (w512 *WELL512A) Uint32() uint32 {
	indexRm1 := w512.table.indexPredAt(w512.state_idx)

	vi := w512.state[w512.state_idx]
	vi1 := w512.state[w512.table.indexM1At(w512.state_idx)]
	vi2 := w512.state[w512.table.indexM2At(w512.state_idx)]
	z0 := w512.state[indexRm1]

	// the values below include the errata of the original article
	z1 := (vi ^ (vi << 16)) ^ (vi1 ^ (vi1 << 15))
	z2 := vi2 ^ (vi2 >> 11)
	z3 := z1 ^ z2
	z4 := (z0 ^ (z0 << 2)) ^ (z1 ^ (z1 << 18)) ^ (z2 << 28) ^ (z3 ^ ((z3 << 5) & 0xda442d24))

	w512.state[w512.state_idx] = z3
	w512.state[indexRm1] = z4
	w512.state_idx = int(indexRm1)

	return z4
}

func (w512 *WELL512A) Seed(seed int64) {
	seeds := make([]uint32, well512a_r)
	seeder.Seed(seed)
	var i int

	if (well512a_r & 1) == 1 {
		seeds[i] = uint32(seeder.Uint64() >> 32)
		i++
	}

	for i < well512a_r {
		v := seeder.Uint64()
		seeds[i] = uint32(v >> 32)
		seeds[i+1] = uint32(v)
		i += 2
	}

	w512.setSeed(seeds)
}

func (w512 *WELL512A) Jump() {
	w512.advanceSeed(well512a_pw, well512a_r)
	w512.RestartSubstream()
}

func (w512 *WELL512A) String() string {
	return "WELL512A"
}
