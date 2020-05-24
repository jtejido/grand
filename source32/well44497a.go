package source32

const (
	well44497a_r  int    = 1391
	well44497a_m1 uint32 = 23
	well44497a_m2 uint32 = 481
	well44497a_m3 uint32 = 229
)

// This implements the WELL44497A pseudo-random number generator
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
//  Total CPU time:   04:21:23.22
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
type WELL44497A struct {
	baseWellNonJumpable
	table *indexTable
}

func NewWELL44497AFromStream(seed []uint32) (*WELL44497A, error) {
	err := checkEmptySeed(seed)
	if err != nil {
		return nil, err
	}

	ans := new(WELL44497A)
	ans.spi = ans

	if len(seed) < well44497a_r {
		tmp := make([]uint32, well44497a_r)
		fillState(tmp, seed)
		ans.setSeed(tmp)
	} else {
		ans.setSeed(seed)
	}

	ans.table = newIndexTable(well44497a_r, well44497a_m1, well44497a_m2, well44497a_m3)
	return ans, nil
}

// this builds the seed slice from a split_mx64 generator using the seed provided
func NewWELL44497A(seed int64) *WELL44497A {
	ans := new(WELL44497A)
	ans.spi = ans

	ans.table = newIndexTable(well44497a_r, well44497a_m1, well44497a_m2, well44497a_m3)
	ans.Seed(seed)

	return ans
}

func (w44497 *WELL44497A) Uint32() uint32 {
	var z2Second uint32
	indexRm1 := w44497.table.indexPredAt(w44497.state_idx)
	indexRm2 := w44497.table.indexPred2At(w44497.state_idx)

	v0 := w44497.state[w44497.state_idx]
	vM1 := w44497.state[w44497.table.indexM1At(w44497.state_idx)]
	vM2 := w44497.state[w44497.table.indexM2At(w44497.state_idx)]
	vM3 := w44497.state[w44497.table.indexM3At(w44497.state_idx)]

	z0 := (0xFFFF8000 & w44497.state[indexRm1]) ^ (0x00007FFF & w44497.state[indexRm2])
	z1 := (v0 ^ (v0 << 24)) ^ (vM1 ^ (vM1 >> 30))
	z2 := (vM2 ^ (vM2 << 10)) ^ (vM3 << 26)
	z3 := z1 ^ z2
	z2Prime := ((z2 << 9) ^ (z2 >> 23)) & 0xfbffffff
	if (z2 & 0x00020000) == 0 {
		z2Second = z2Prime
	} else {
		z2Second = (z2Prime ^ 0xb729fcec)
	}

	z4 := z0 ^ (z1 ^ (z1 >> 20)) ^ z2Second ^ z3

	w44497.state[w44497.state_idx] = z3
	w44497.state[indexRm1] = z4
	w44497.state[indexRm2] &= 0xFFFF8000
	w44497.state_idx = int(indexRm1)

	return z4
}

func (w44497 *WELL44497A) Seed(seed int64) {
	seeds := make([]uint32, well44497a_r)
	seeder.Seed(seed)
	var i int

	if (well44497a_r & 1) == 1 {
		seeds[i] = uint32(seeder.Uint64() >> 32)
		i++
	}

	for i < well44497a_r {
		v := seeder.Uint64()
		seeds[i] = uint32(v >> 32)
		seeds[i+1] = uint32(v)
		i += 2
	}

	w44497.setSeed(seeds)
}

func (w44497 *WELL44497A) String() string {
	return "WELL44497A"
}
