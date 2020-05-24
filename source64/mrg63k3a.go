package source64

import (
	"fmt"
)

const (
	mrg63k3a_m1   uint64 = 9223372036854769163
	mrg63k3a_m2   uint64 = 9223372036854754679
	mrg63k3a_a12  uint64 = 1754669720
	mrg63k3a_q12  uint64 = 5256471877
	mrg63k3a_r12  uint64 = 251304723
	mrg63k3a_a13n uint64 = 3182104042
	mrg63k3a_q13  uint64 = 2898513661
	mrg63k3a_r13  uint64 = 394451401
	mrg63k3a_a21  uint64 = 31387477935
	mrg63k3a_q21  uint64 = 293855150
	mrg63k3a_r21  uint64 = 143639429
	mrg63k3a_a23n uint64 = 6199136374
	mrg63k3a_q23  uint64 = 1487847900
	mrg63k3a_r23  uint64 = 985240079
	mrg63k3a_r           = 6
)

// This implements the MRG63k3A pseudo-random number generator
// from Pierre L'Ecuyer.
//
// This generator is described in a paper by L'Ecuyer.
// https://www.iro.umontreal.ca/~lecuyer/myftp/papers/opres-combmrg2-1999.pdf
// Good Parameter Sets for Combined Multiple Recursive Random Number Generators.
// Operations Research, 1999, 47-1, 159--164
// TO-DO. Jump()
type MRG63k3A struct {
	// State variable [][]s must be 2-vector 64-bit integer.
	// The seeds for s[0][0], s[0][1], s[0][2] must be integers in [0, m1 - 1] and not all 0.
	// The seeds for s[1][0], s[1][1], s[1][2] must be integers in [0, m2 - 1] and not all 0.
	baseSource64
	s [2][3]uint64
}

func NewMRG63k3AFromStream(seed []uint64) (ans *MRG63k3A, err error) {
	err = checkEmptySeed(seed)
	if err != nil {
		return
	}
	ans = new(MRG63k3A)
	ans.spi = ans
	if len(seed) < mrg63k3a_r {
		tmp := make([]uint64, mrg63k3a_r)
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
func NewMRG63k3A(seed int64) *MRG63k3A {
	ans := new(MRG63k3A)
	ans.spi = ans
	ans.Seed(seed)
	return ans
}

func (mrg *MRG63k3A) setSeed(seed []uint64) (err error) {

	err = mrg.checkSeed(seed)

	if err != nil {
		return
	}

	mrg.stream = append([]uint64{}, seed...)

	mrg.Restart()
	return nil
}

func (mrg *MRG63k3A) checkSeed(seed []uint64) error {

	for i := 0; i < mrg63k3a_r; i++ {
		if seed[i] == 0 {
			return fmt.Errorf("seed values must not be 0.")
		}

		if (i == 0 || i == 1 || i == 2) && seed[i] >= mrg63k3a_m1 {
			return fmt.Errorf("The first 3 seed values must be less than %d", mrg63k3a_m1)
		} else if seed[i] >= mrg63k3a_m2 {
			return fmt.Errorf("The last 3 seed values must be less than %d", mrg63k3a_m2)
		}
	}

	return nil
}

func (mrg *MRG63k3A) Restart() {
	mrg.s[0][0] = mrg.stream[0]
	mrg.s[0][1] = mrg.stream[1]
	mrg.s[0][2] = mrg.stream[2]
	mrg.s[1][0] = mrg.stream[3]
	mrg.s[1][1] = mrg.stream[4]
	mrg.s[1][2] = mrg.stream[5]
	mrg.resetState()
}

func (mrg *MRG63k3A) Uint64() uint64 {

	/* Component 1 */
	h := mrg.s[0][0] / mrg63k3a_q13
	p13 := mrg63k3a_a13n*(mrg.s[0][0]-h*mrg63k3a_q13) - h*mrg63k3a_r13
	h = mrg.s[0][1] / mrg63k3a_q12
	p12 := mrg63k3a_a12*(mrg.s[0][1]-h*mrg63k3a_q12) - h*mrg63k3a_r12
	if p13&0x8000000000000000 == 1 {
		p13 += mrg63k3a_m1
	}
	if p12&0x8000000000000000 == 1 {
		p12 += mrg63k3a_m1 - p13
	} else {
		p12 -= p13
	}
	if p12&0x8000000000000000 == 1 {
		p12 += mrg63k3a_m1
	}
	mrg.s[0][0] = mrg.s[0][1]
	mrg.s[0][1] = mrg.s[0][2]
	mrg.s[0][2] = p12

	/* Component 2 */
	h = mrg.s[1][0] / mrg63k3a_q23
	p23 := mrg63k3a_a23n*(mrg.s[1][0]-h*mrg63k3a_q23) - h*mrg63k3a_r23
	h = mrg.s[1][2] / mrg63k3a_q21
	p21 := mrg63k3a_a21*(mrg.s[1][2]-h*mrg63k3a_q21) - h*mrg63k3a_r21
	if p23&0x8000000000000000 == 1 {
		p23 += mrg63k3a_m2
	}
	if p21&0x8000000000000000 == 1 {
		p21 += mrg63k3a_m2 - p23
	} else {
		p21 -= p23
	}
	if p21&0x8000000000000000 == 1 {
		p21 += mrg63k3a_m2
	}
	mrg.s[1][0] = mrg.s[1][1]
	mrg.s[1][1] = mrg.s[1][2]
	mrg.s[1][2] = p21

	/* Combination */
	if p12 > p21 {
		return (p12 - p21)
	}
	return (p12 - p21 + mrg63k3a_m1)
}

func (mrg *MRG63k3A) Seed(seed int64) {
	seeds := make([]uint64, mrg63k3a_r)
	seeder.Seed(seed)
	for j := 0; j < 3; j++ {
	again0:
		f := seeder.Intn(int(mrg63k3a_m1))
		if f == 0 {
			goto again0
		}
		seeds[j] = uint64(f)
	}
	for j := 3; j < 6; j++ {
	again1:
		f := seeder.Intn(int(mrg63k3a_m2))
		if f == 0 {
			goto again1
		}
		seeds[j] = uint64(f)
	}

	// Initialize the pool content.
	mrg.setSeed(seeds)
}

func (mrg *MRG63k3A) String() string {
	return "MRG63k3A"
}
