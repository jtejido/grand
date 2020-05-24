package source32

const (
	well19937a_r  int    = 624
	well19937a_m1 uint32 = 70
	well19937a_m2 uint32 = 179
	well19937a_m3 uint32 = 449
)

// This implements the WELL19937a pseudo-random number generator
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
//  Generator:        rand.Float64
//  Number of statistics:  144
//  Total CPU time:   02:46:13.65
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
type WELL19937A struct {
	baseWellNonJumpable
	table *indexTable
}

func NewWELL19937AFromStream(seed []uint32) (*WELL19937A, error) {
	err := checkEmptySeed(seed)
	if err != nil {
		return nil, err
	}

	ans := new(WELL19937A)
	ans.spi = ans

	if len(seed) < well19937a_r {
		tmp := make([]uint32, well19937a_r)
		fillState(tmp, seed)
		ans.setSeed(tmp)
	} else {
		ans.setSeed(seed)
	}
	ans.table = newIndexTable(well19937a_r, well19937a_m1, well19937a_m2, well19937a_m3)
	return ans, nil
}

func NewWELL19937A(seed int64) *WELL19937A {
	ans := new(WELL19937A)
	ans.spi = ans

	ans.table = newIndexTable(well19937a_r, well19937a_m1, well19937a_m2, well19937a_m3)
	ans.Seed(seed)

	return ans
}

func (w19937 *WELL19937A) Uint32() uint32 {
	indexRm1 := w19937.table.indexPredAt(w19937.state_idx)
	indexRm2 := w19937.table.indexPred2At(w19937.state_idx)

	v0 := w19937.state[w19937.state_idx]
	vM1 := w19937.state[w19937.table.indexM1At(w19937.state_idx)]
	vM2 := w19937.state[w19937.table.indexM2At(w19937.state_idx)]
	vM3 := w19937.state[w19937.table.indexM3At(w19937.state_idx)]

	z0 := (0x80000000 & w19937.state[indexRm1]) ^ (0x7FFFFFFF & w19937.state[indexRm2])
	z1 := (v0 ^ (v0 << 25)) ^ (vM1 ^ (vM1 >> 27))
	z2 := (vM2 >> 9) ^ (vM3 ^ (vM3 >> 1))
	z3 := z1 ^ z2
	z4 := z0 ^ (z1 ^ (z1 << 9)) ^ (z2 ^ (z2 << 21)) ^ (z3 ^ (z3 >> 21))

	w19937.state[w19937.state_idx] = z3
	w19937.state[indexRm1] = z4
	w19937.state[indexRm2] &= 0x80000000
	w19937.state_idx = int(indexRm1)

	return z4
}

func (w19937 *WELL19937A) Seed(seed int64) {
	seeds := make([]uint32, well19937a_r)
	seeder.Seed(seed)
	var i int

	if (well19937a_r & 1) == 1 {
		seeds[i] = uint32(seeder.Uint64() >> 32)
		i++
	}

	for i < well19937a_r {
		v := seeder.Uint64()
		seeds[i] = uint32(v >> 32)
		seeds[i+1] = uint32(v)
		i += 2
	}

	w19937.setSeed(seeds)
}

func (w19937 *WELL19937A) String() string {
	return "WELL19937A"
}
