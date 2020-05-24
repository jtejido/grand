package source32

const (
	well1024a_r  int    = 32
	well1024a_m1 uint32 = 3
	well1024a_m2 uint32 = 24
	well1024a_m3 uint32 = 10
)

// This table represents the 1024 coefficients of the following polynomial
// (z^(2^400) mod P(z)) mod 2
// P(z) is the characteristic polynomial of the generator.
// Characteristic Polynomial is defined as:
// P(z) = det(A − zI) = zk − α1zk−1 −...− αk−1z − αk
// where I is the identity matrix and each αj is in F2 (the finite field with two elements, 0 and 1). For each j, the sequences
// {xi, j , i ≥ 0} and { yi, j , i ≥ 0} both obey the linear recurrence
// xi, j = (α1xi−1, j +...+ αk xi−k, j) mod 2
var (
	well1024a_pw = []uint32{
		0xe44294e, 0xef237eff, 0x5e8b6bfb, 0xa724e67a,
		0x59994cfd, 0x6f7c3de1, 0x6735d50d, 0x4bfe199a,
		0x39c28e61, 0xfd075266, 0x96cc6d1f, 0x5dc1a685,
		0xd67fa444, 0xccc01b86, 0x8ff861c, 0xce113725,
		0x66707603, 0x38abb0fd, 0x7681f64, 0x104535c5,
		0xce4ae5f4, 0x50e37105, 0xd0c5f77f, 0x74c1ebf6,
		0x2ccf1505, 0xd1f21b86, 0x9a6c402e, 0xea34a31c,
		0x65e13d13, 0xde8f2f05, 0x89db804f, 0x8dc387f2,
	}
)

// This implements the WELL1024a pseudo-random number generator
// from Panneton, L'Ecuyer and Matsumoto.
//
// This generator is described in a paper by Panneton, L'Ecuyer and Matsumoto
// http://www.iro.umontreal.ca/~lecuyer/myftp/papers/wellrng.pdf
// Improved Long-Period Generators Based on Linear Recurrences Modulo 2
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
//  Total CPU time:   03:18:14.09
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
type WELL1024A struct {
	baseJumpableWell
	table *indexTable
}

func NewWELL1024AFromStream(seed []uint32) (*WELL1024A, error) {
	err := checkEmptySeed(seed)
	if err != nil {
		return nil, err
	}

	ans := new(WELL1024A)
	ans.spi = ans

	if len(seed) < well1024a_r {
		tmp := make([]uint32, well1024a_r)
		fillState(tmp, seed)
		ans.setSeed(tmp)
	} else {
		ans.setSeed(seed)
	}
	ans.table = newIndexTable(well1024a_r, well1024a_m1, well1024a_m2, well1024a_m3)
	return ans, nil
}

func NewWELL1024A(seed int64) *WELL1024A {
	ans := new(WELL1024A)
	ans.spi = ans

	ans.table = newIndexTable(well1024a_r, well1024a_m1, well1024a_m2, well1024a_m3)
	ans.Seed(seed)

	return ans
}

func (w1024 *WELL1024A) Uint32() uint32 {
	indexRm1 := w1024.table.indexPredAt(w1024.state_idx)

	v0 := w1024.state[w1024.state_idx]
	vM1 := w1024.state[w1024.table.indexM1At(w1024.state_idx)]
	vM2 := w1024.state[w1024.table.indexM2At(w1024.state_idx)]
	vM3 := w1024.state[w1024.table.indexM3At(w1024.state_idx)]

	z0 := w1024.state[indexRm1]
	z1 := v0 ^ (vM1 ^ (vM1 >> 8))
	z2 := (vM2 ^ (vM2 << 19)) ^ (vM3 ^ (vM3 << 14))
	z3 := z1 ^ z2
	z4 := (z0 ^ (z0 << 11)) ^ (z1 ^ (z1 << 7)) ^ (z2 ^ (z2 << 13))

	w1024.state[w1024.state_idx] = z3
	w1024.state[indexRm1] = z4
	w1024.state_idx = int(indexRm1)

	return z4
}

func (w1024 *WELL1024A) Seed(seed int64) {
	seeds := make([]uint32, well1024a_r)
	seeder.Seed(seed)
	var i int
	if (well1024a_r & 1) == 1 {
		seeds[i] = uint32(seeder.Uint64() >> 32)
		i++
	}

	for i < well1024a_r {
		v := seeder.Uint64()
		seeds[i] = uint32(v >> 32)
		seeds[i+1] = uint32(v)
		i += 2
	}

	w1024.setSeed(seeds)
}

func (w1024 *WELL1024A) Jump() {
	w1024.advanceSeed(well1024a_pw, well1024a_r)
	w1024.RestartSubstream()
}

func (w1024 *WELL1024A) String() string {
	return "WELL1024A"
}
