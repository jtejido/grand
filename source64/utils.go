package source64

import (
	"errors"
	"math"
)

func checkEmptySeed(seed []uint64) error {
	if seed == nil || len(seed) == 0 {
		return errors.New("stream cannot be empty")
	}

	return nil
}

func fillState(state, seed []uint64) {
	stateSize := len(state)
	seedSize := len(seed)
	l := int(math.Min(float64(seedSize), float64(stateSize)))
	copy(state[0:l], seed[0:l])

	if seedSize < stateSize {
		for i := seedSize; i < stateSize; i++ {
			state[i] = scrambleWell(state[i-len(seed)], uint(i))
		}
	}
}

func rotateLeft(x uint64, k int) uint64 {
	n := uint(64)
	s := uint(k) & (n - 1)
	return x<<s | x>>(n-s)
}

/**
 * Transformation used to scramble the initial state of
 * a generator.
 *
 * n = seed element.
 * add = Offset.
 */
func scrambleWell(n uint64, add uint) uint64 {
	return scramble(n, 1812433253, 30, add)
}

/**
 * Transformation used to scramble the initial state of
 * a generator.
 *
 * n = seed element.
 * mult = Multiplier.
 * shift = Shift.
 * add = Offset.
 */
func scramble(n, mult uint64, shift, add uint) uint64 {
	// Code inspired from "AbstractWell" class.
	return mult*(n^(n>>shift)) + uint64(add)
}
