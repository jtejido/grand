package source64

import (
	"github.com/jtejido/grand"
)

// This is used for seeding any implementations here
var (
	seeder *grand.Rand
)

// all structs should follow this
type source64 interface {
	Uint64() uint64
}

func init() {
	seeder = grand.New(NewSplitMix64(123456789))
}

// This serves as the base struct for holding the starting point for the stream.
// all structs embedding this should handle their own starting point of the stream.
type baseSource64 struct {
	// Golang doesn't override or pickup methods from child (in OOP sense),
	// thus it is required to assign any source64 implementors here to be used by Bool().
	spi source64
	// stream stores the starting point for the stream.
	stream []uint64
	// The idea is to generate uint64 once, and return high and low 32-bits alternating for each call of Uint32
	cachedInt32Source bool
	int32Source       uint64
	// booleanSource caches the most recent uint64
	// booleanBitMask is the bit mask of the boolean source to obtain the boolean bit.
	// This begins at the least significant bit and is gradually shifted upwards until overflow to zero.
	// When zero a new boolean source should be created and the mask set to the least significant bit (i.e. 1).
	booleanSource, booleanBitMask uint64
}

func (bs64 *baseSource64) Uint32() uint32 {
	if bs64.cachedInt32Source {
		bs64.cachedInt32Source = false
		return uint32(bs64.int32Source)
	}

	bs64.cachedInt32Source = true
	bs64.int32Source = bs64.Uint64()

	return uint32(bs64.int32Source >> 32)
}

func (bs64 *baseSource64) Uint64() uint64 { return bs64.spi.Uint64() }

func (bs64 *baseSource64) Bool() bool {
	bs64.booleanBitMask <<= 1
	if bs64.booleanBitMask == 0 {
		bs64.booleanBitMask = 1
		bs64.booleanSource = bs64.Uint64()
	}

	return (bs64.booleanSource & bs64.booleanBitMask) != 0
}

// restores starting values for Bool() and Uint32()
func (bs64 *baseSource64) resetState() {
	bs64.booleanSource = 0
	bs64.booleanBitMask = 0
	bs64.int32Source = 0
	bs64.cachedInt32Source = false
}

// Embeds baseSource64 and store substreams.
type baseJumpableSource64 struct {
	baseSource64
	//  substream stores the starting point of the current substream.
	substream []uint64
}
