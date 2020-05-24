package source64

import (
	"errors"
)

const (
	lfsr258_r            = 5
	lfsr258_norm float64 = 0.5 / 0x7FFFFFFFFFFFFC00
	seed_inv_err string  = "The seed elements must be either negative or greater than 1, 7, 15, 127 and 8388607 respectively"
)

// This implements the LFSR258 pseudo-random number generator
// from Pierre L'Ecuyer.
//
// This generator is described in a paper by L'Ecuyer.
// http://www-labs.iro.umontreal.ca/~lecuyer/myftp/papers/tausme.ps
// Maximally equidistributed combined Tausworthe generators.
// Mathematics of Computation, 65, 213 (1996), 203--213.
//
// ========= Summary results of Crush =========
//
//  Version:          TestU01 1.2.3
//  Generator:        rand.Float64
//  Number of statistics:  144
//  Total CPU time:   04:22:38.89
//  The following tests gave p-values outside [0.001, 0.9990]:
//  (eps  means a value < 1.0e-300):
//  (eps1 means a value < 1.0e-15):
//
//        Test                          p-value
//  ----------------------------------------------
//  58  MatrixRank, 300 x 300            eps
//  59  MatrixRank, 300 x 300            eps
//  60  MatrixRank, 1200 x 1200          eps
//  61  MatrixRank, 1200 x 1200          eps
//  71  LinearComp, r = 0              1 - eps1
//  72  LinearComp, r = 29             1 - eps1
//  ----------------------------------------------
//  All other tests were passed
type LFSR258 struct {
	baseJumpableSource64
	state [5]uint64
}

func NewLFSR258FromStream(seed []uint64) (ans *LFSR258, err error) {
	err = checkEmptySeed(seed)
	if err != nil {
		return
	}

	ans = new(LFSR258)
	ans.spi = ans
	if len(seed) < lfsr258_r {
		tmp := make([]uint64, lfsr258_r)
		fillState(tmp, seed)
		err = ans.setSeed(tmp)
		if err != nil {
			return
		}
	} else {
		err = ans.setSeed(seed)
		if err != nil {
			return
		}
	}

	return
}

// this builds the seed slice from a split_mx64 generator using the seed provided
func NewLFSR258(seed int64) *LFSR258 {
	ans := new(LFSR258)
	ans.spi = ans
	ans.Seed(seed)
	return ans
}

func (lfsr *LFSR258) setSeed(seed []uint64) (err error) {
	err = lfsr.checkSeed(seed)
	if err != nil {
		return
	}

	lfsr.stream = append([]uint64{}, seed...)
	lfsr.Restart()
	return
}

func (lfsr *LFSR258) Restart() {
	lfsr.substream = append([]uint64{}, lfsr.stream...)
	lfsr.RestartSubstream()
}

func (lfsr *LFSR258) RestartSubstream() {
	for j := 0; j < lfsr258_r; j++ {
		lfsr.state[j] = lfsr.substream[j]
	}

	lfsr.resetState()
}

func (lfsr *LFSR258) Jump() {

	z := lfsr.substream[0] & 0xfffffffffffffffe
	b := z ^ (z << 1)
	z = (b >> 61) ^ (b >> 59) ^ (b >> 58) ^ (b >> 57) ^ (b >> 51) ^
		(b >> 47) ^ (b >> 46) ^ (b >> 45) ^ (b >> 43) ^ (b >> 39) ^
		(b >> 30) ^ (b >> 29) ^ (b >> 23) ^ (b >> 15) ^ (z << 2) ^
		(z << 4) ^ (z << 5) ^ (z << 6) ^ (z << 12) ^ (z << 16) ^
		(z << 17) ^ (z << 18) ^ (z << 20) ^ (z << 24) ^ (z << 33) ^
		(z << 34) ^ (z << 40) ^ (z << 48)
	lfsr.substream[0] = z

	z = lfsr.substream[1] & 0xfffffffffffffe00
	b = z ^ (z << 24)
	z = (b >> 52) ^ (b >> 50) ^ (b >> 49) ^ (b >> 46) ^ (b >> 43) ^
		(b >> 40) ^ (b >> 37) ^ (b >> 34) ^ (b >> 30) ^ (b >> 28) ^
		(b >> 26) ^ (b >> 25) ^ (b >> 23) ^ (b >> 21) ^ (b >> 20) ^
		(b >> 19) ^ (b >> 17) ^ (b >> 15) ^ (b >> 13) ^ (b >> 12) ^
		(b >> 10) ^ (b >> 8) ^ (b >> 7) ^ (b >> 6) ^ (b >> 2) ^
		(z << 1) ^ (z << 4) ^ (z << 6) ^ (z << 7) ^ (z << 11) ^ (z << 14) ^
		(z << 15) ^ (z << 16) ^ (z << 17) ^ (z << 21) ^ (z << 22) ^
		(z << 25) ^ (z << 27) ^ (z << 29) ^ (z << 30) ^ (z << 32) ^
		(z << 34) ^ (z << 35) ^ (z << 36) ^ (z << 38) ^ (z << 40) ^
		(z << 42) ^ (z << 43) ^ (z << 45) ^ (z << 47) ^ (z << 48) ^
		(z << 49) ^ (z << 53)
	lfsr.substream[1] = z

	z = lfsr.substream[2] & 0xfffffffffffff000
	b = z ^ (z << 3)
	z = (b >> 49) ^ (b >> 45) ^ (b >> 41) ^ (b >> 40) ^ (b >> 32) ^
		(b >> 27) ^ (b >> 23) ^ (b >> 14) ^ (b >> 1) ^ (z << 2) ^
		(z << 3) ^ (z << 7) ^ (z << 11) ^ (z << 12) ^ (z << 20) ^
		(z << 25) ^ (z << 29) ^ (z << 38) ^ (z << 51)
	lfsr.substream[2] = z

	z = lfsr.substream[3] & 0xfffffffffffe0000
	b = z ^ (z << 5)
	z = (b >> 45) ^ (b >> 32) ^ (b >> 27) ^ (b >> 22) ^ (b >> 17) ^
		(b >> 13) ^ (b >> 12) ^ (b >> 7) ^ (b >> 3) ^ (b >> 2) ^
		(z << 3) ^ (z << 15) ^ (z << 20) ^ (z << 25) ^ (z << 30) ^
		(z << 34) ^ (z << 35) ^ (z << 40) ^ (z << 44) ^ (z << 45)
	lfsr.substream[3] = z

	z = lfsr.substream[4] & 0xffffffffff800000
	b = z ^ (z << 3)
	z = (b >> 40) ^ (b >> 39) ^ (b >> 38) ^ (b >> 37) ^ (b >> 35) ^
		(b >> 34) ^ (b >> 31) ^ (b >> 30) ^ (b >> 29) ^ (b >> 28) ^
		(b >> 27) ^ (b >> 26) ^ (b >> 24) ^ (b >> 23) ^ (b >> 21) ^
		(b >> 20) ^ (b >> 18) ^ (b >> 15) ^ (b >> 12) ^ (b >> 10) ^
		(b >> 9) ^ (b >> 7) ^ (b >> 6) ^ (b >> 5) ^ (b >> 4) ^
		(b >> 3) ^ (z << 1) ^ (z << 2) ^ (z << 3) ^ (z << 4) ^ (z << 6) ^
		(z << 7) ^ (z << 10) ^ (z << 11) ^ (z << 12) ^ (z << 13) ^
		(z << 14) ^ (z << 15) ^ (z << 17) ^ (z << 18) ^ (z << 20) ^
		(z << 21) ^ (z << 23) ^ (z << 26) ^ (z << 29) ^ (z << 31) ^
		(z << 32) ^ (z << 34) ^ (z << 35) ^ (z << 36) ^ (z << 37) ^
		(z << 38)
	lfsr.substream[4] = z
	lfsr.RestartSubstream()
}

func (lfsr *LFSR258) Uint64() uint64 {
	b := (((lfsr.state[0] << 1) ^ lfsr.state[0]) >> 53)
	lfsr.state[0] = (((lfsr.state[0] & 0xFFFFFFFFFFFFFFFE) << 10) ^ b)
	b = (((lfsr.state[1] << 24) ^ lfsr.state[1]) >> 50)
	lfsr.state[1] = (((lfsr.state[1] & 0xFFFFFFFFFFFFFE00) << 5) ^ b)
	b = (((lfsr.state[2] << 3) ^ lfsr.state[2]) >> 23)
	lfsr.state[2] = (((lfsr.state[2] & 0xFFFFFFFFFFFFF000) << 29) ^ b)
	b = (((lfsr.state[3] << 5) ^ lfsr.state[3]) >> 24)
	lfsr.state[3] = (((lfsr.state[3] & 0xFFFFFFFFFFFE0000) << 23) ^ b)
	b = (((lfsr.state[4] << 3) ^ lfsr.state[4]) >> 33)
	lfsr.state[4] = (((lfsr.state[4] & 0xFFFFFFFFFF800000) << 8) ^ b)

	return (lfsr.state[0] ^ lfsr.state[1] ^ lfsr.state[2] ^ lfsr.state[3] ^ lfsr.state[4])
}

func (lfsr *LFSR258) checkSeed(seed []uint64) error {

	if (seed[0] >= 0 && seed[0] < 2) ||
		(seed[1] >= 0 && seed[1] < 8) ||
		(seed[2] >= 0 && seed[2] < 16) ||
		(seed[3] >= 0 && seed[3] < 128) ||
		(seed[4] >= 0 && seed[4] < 8388608) {
		return errors.New(seed_inv_err)
	}

	return nil
}

func (lfsr *LFSR258) Seed(seed int64) {
	seeds := make([]uint64, lfsr258_r)
	seeder.Seed(seed)

	for j := 0; j < lfsr258_r; j++ {
	again:
		f := seeder.Uint64()

		if (j == 0 && f <= 1) || (j == 1 && f <= 7) || (j == 2 && f <= 15) || (j == 3 && f <= 127) || (j == 4 && f <= 8388607) {
			goto again
		}

		seeds[j] = f

	}

	// Initialize the pool content.
	lfsr.setSeed(seeds)
}

func (lfsr *LFSR258) String() string {
	return "LFSR258"
}
