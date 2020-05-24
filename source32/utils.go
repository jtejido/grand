package source32

import (
	"errors"
	"fmt"
	"math"
)

const (
	two17 int = 131072
	two53 int = 9007199254740992
)

func checkEmptySeed(seed []uint32) error {
	if seed == nil || len(seed) == 0 {
		return errors.New("stream cannot be empty")
	}

	return nil
}

func checkEmptySeed64(seed []uint64) error {
	if seed == nil || len(seed) == 0 {
		return errors.New("stream cannot be empty")
	}

	return nil
}

/**
 * Simple filling procedure.
 * state will be pre-filled
 * seed cannot be nil
 */
func fillState64(state, seed []uint64) {
	stateSize := len(state)
	seedSize := len(seed)
	l := int(math.Min(float64(seedSize), float64(stateSize)))
	copy(state[:l], seed[:l])

	if seedSize < stateSize {
		for i := seedSize; i < stateSize; i++ {
			state[i] = scrambleWell(state[i-len(seed)], uint(i))
		}
	}
}

/**
 * Simple filling procedure.
 * state will be pre-filled
 * seed cannot be nil
 */
func fillState(state, seed []uint32) {
	stateSize := len(state)
	seedSize := len(seed)
	l := int(math.Min(float64(seedSize), float64(stateSize)))
	copy(state[:l], seed[:l])

	if seedSize < stateSize {
		for i := seedSize; i < stateSize; i++ {
			state[i] = uint32(scrambleWell(uint64(state[i-len(seed)]), uint(i)) & 0xffffffff)
		}
	}
}

func scramble(n, mult uint64, shift, add uint) uint64 {
	return mult*(n^(n>>shift)) + uint64(add)
}

func scrambleWell(n uint64, add uint) uint64 {
	return scramble(n, 1812433253, 30, add)
}

func rotateLeft(x uint32, k int) uint32 {
	n := uint(32)
	s := uint(k) & (n - 1)
	return x<<s | x>>(n-s)
}

func checkMRGSeed(seed []uint32, r int, m1, m2 uint32) error {
	for i := 0; i < r; i++ {
		if seed[i] == 0 {
			return fmt.Errorf("seed values must not be 0.")
		}

		if (i == 0 || i == 1 || i == 2) && seed[i] >= m1 {
			return fmt.Errorf("The first 3 seed values must be less than %d", m1)
		} else if seed[i] >= m2 {
			return fmt.Errorf("The last 3 seed values must be less than %d", m2)
		}
	}

	return nil
}

//multiply the first half of substream by A with a modulo of m1 and the second half by B with a modulo of m2
func multMatVect(substream []uint32, A [][]uint32, m1 uint32, B [][]uint32, m2 uint32) {
	vv := make([]uint32, 3)
	for i := 0; i < 3; i++ {
		vv[i] = substream[i]
	}
	matVecModM(A, vv, vv, m1)
	for i := 0; i < 3; i++ {
		substream[i] = vv[i]
	}

	for i := 0; i < 3; i++ {
		vv[i] = substream[i+3]
	}
	matVecModM(B, vv, vv, m2)
	for i := 0; i < 3; i++ {
		substream[i+3] = vv[i]
	}
}

func matVecModM(A [][]uint32, s, v []uint32, m uint32) {

	x := make([]uint32, len(v))
	for i := 0; i < len(v); i++ {
		x[i] = 0
		for j := 0; j < len(s); j++ {
			x[i] = multModM(int(A[i][j]), int(s[j]), int(x[i]), int(m))
		}
	}

	for i := 0; i < len(v); i++ {
		v[i] = x[i]
	}
}

func multModM(a, s, c, m int) uint32 {
	var a1 int
	v := a*s + c

	if v >= two53 || v <= -two53 {
		a1 = (a / two17)
		a -= a1 * two17
		v = a1 * s
		a1 = (v / m)
		v -= a1 * m
		v = v*two17 + a*s + c
	}

	a1 = (v / m)
	v -= a1 * m

	if v < 0 {
		v += m
		return uint32(v)
	}

	return uint32(v)
}
