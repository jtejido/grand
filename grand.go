// Package rand implements pseudo-random number generators.
//
// This is a modified version of Golang's math.Rand that accepts 32 and 64-bit sources.
//
// The idea is to make it easy for others to extend or create source easily since most, i.e. known, resources in
// Pseudo-Random Number Generation are implemented in either 32 or 64 bit form.
// It will be as easy as translating directly from the authors' papers (or mostly C source codes).
//
// All sources should allow restarting streams from its initial seed by calling Restart().
// Some sources are allowed for jumping and restarting substreams (see JumpableSource).
// Restartable streams and substreams are useful for simulations and debugging (reproducibility).
// We also modified it to publish LockedSource and LockedJumpableSource for concurrency uses.
package grand

import (
	"sync"
)

const (
	// The multiplier to convert the least significant 24-bits of a uint32 to a float32.
	// 1.0 / (1 << 24)
	float32_multiplier = 5.960464477539063e-08
	//T he multiplier to convert the least significant 53-bits of a uint64 to a float64.
	// 1.0 / (1 << 53)
	float64_multiplier = 1.1102230246251565e-16
)

type Source interface {
	Uint32() uint32
	Bool() bool
	Seed(seed int64)
	Restart()
}

type JumpableSource interface {
	Source
	RestartSubstream()
	Jump()
}

type Source64 interface {
	Source
	Uint64() uint64
}

type Rand struct {
	src Source
	s64 Source64
}

func New(src Source) *Rand {
	s64, _ := src.(Source64)
	return &Rand{src: src, s64: s64}
}

// Seed uses the provided seed value to initialize the generator to a deterministic state.
// Seed should not be called concurrently with any other Rand method.
func (r *Rand) Seed(seed int64) { r.src.Seed(seed) }

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (r *Rand) Int63() int64 { return (int64(r.Uint32()) << 31) | int64(r.Uint32()) }

// Uint32 returns a pseudo-random 32-bit value as a uint32.
func (r *Rand) Uint32() uint32 { return r.src.Uint32() }

// Int31 returns a non-negative pseudo-random 31-bit integer as an int32.
func (r *Rand) Int31() int32 { return int32(r.Uint32() << 1 >> 1) }

// Int returns a non-negative pseudo-random int.
func (r *Rand) Int() int {
	u := uint(r.Int63())
	return int(u)
}

// Uint64 returns a pseudo-random 64-bit value as a uint64.
func (r *Rand) Uint64() uint64 {
	if r.s64 != nil {
		return r.s64.Uint64()
	}

	return (uint64(r.Uint32()) << 32) | uint64(r.Uint32())
}

// Int63n returns, as an int64, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func (r *Rand) Int63n(n int64) int64 {
	if n <= 0 {
		panic("invalid argument to Int63n")
	}
	if n&(n-1) == 0 { // n is power of two, can mask
		return r.Int63() & (n - 1)
	}
	max := int64((1 << 63) - 1 - (1<<63)%uint64(n))
	v := r.Int63()
	for v > max {
		v = r.Int63()
	}
	return v % n
}

// Int31n returns, as an int32, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func (r *Rand) Int31n(n int32) int32 {
	if n <= 0 {
		panic("invalid argument to Int31n")
	}
	if n&(n-1) == 0 { // n is power of two, can mask
		return r.Int31() & (n - 1)
	}
	max := int32((1 << 31) - 1 - (1<<31)%uint32(n))
	v := r.Int31()
	for v > max {
		v = r.Int31()
	}
	return v % n
}

// int31n returns, as an int32, a non-negative pseudo-random number in [0,n).
// n must be > 0, but int31n does not check this; the caller must ensure it.
// int31n exists because Int31n is inefficient, but Go 1 compatibility
// requires that the stream of values produced by math/rand remain unchanged.
// int31n can thus only be used internally, by newly introduced APIs.
//
// For implementation details, see:
// https://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction
// https://lemire.me/blog/2016/06/30/fast-random-shuffling
func (r *Rand) int31n(n int32) int32 {
	v := r.Uint32()
	prod := uint64(v) * uint64(n)
	low := uint32(prod)
	if low < uint32(n) {
		thresh := uint32(-n) % uint32(n)
		for low < thresh {
			v = r.Uint32()
			prod = uint64(v) * uint64(n)
			low = uint32(prod)
		}
	}
	return int32(prod >> 32)
}

// Intn returns, as an int, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func (r *Rand) Intn(n int) int {
	if n <= 0 {
		panic("invalid argument to Intn")
	}
	if n <= 1<<31-1 {
		return int(r.Int31n(int32(n)))
	}
	return int(r.Int63n(int64(n)))
}

// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0).
func (r *Rand) Float64() float64 {
	if r.s64 != nil {
		// Require the least significant 53-bits so shift the higher bits across
		return float64(r.Uint64()>>11) * float64_multiplier
	}

	// Require the least significant 53-bits from a long.
	// Join the most significant 26 first uint32 with 27 from second uint32.
	high := (uint64(r.Uint32() >> 6)) << 27 // 26-bits remain
	low := uint64(r.Uint32() >> 5)          // 27-bits remain

	return float64(high|low) * float64_multiplier
}

// Float32 returns, as a float32, a pseudo-random number in [0.0,1.0).
func (r *Rand) Float32() float32 {
	// Require the least significant 24-bits so shift the higher bits across
	return float32(r.Uint32()>>8) * float32_multiplier
}

// Returns the next bool.
func (r *Rand) Bool() bool {
	if r.s64 != nil {
		return r.s64.Bool()
	}

	return r.src.Bool()
}

// Restarts the stream to its initial state.
func (r *Rand) Restart() {
	r.src.Restart()
}

type JumpableRand struct {
	Rand
}

func NewJumpable(src JumpableSource) *JumpableRand {
	s64, _ := src.(Source64)
	ans := new(JumpableRand)
	ans.src = src
	ans.s64 = s64
	return ans
}

// Restarts the stream to the beginning of its current substream.
func (r *JumpableRand) RestartSubstream() {
	if js, ok := r.src.(JumpableSource); ok {
		js.RestartSubstream()
	}
}

// Restarts the stream to the beginning of its next substream.
func (r *JumpableRand) Jump() {
	if js, ok := r.src.(JumpableSource); ok {
		js.Jump()
	}
}

type LockedSource struct {
	lk  sync.Mutex
	src Source64
}

func NewLockedSource(src Source) *LockedSource {
	return &LockedSource{src: src.(Source64)}
}

func (r *LockedSource) Uint32() (n uint32) {
	r.lk.Lock()
	n = r.src.Uint32()
	r.lk.Unlock()
	return
}

func (r *LockedSource) Bool() (ok bool) {
	r.lk.Lock()
	ok = r.src.Bool()
	r.lk.Unlock()
	return
}

func (r *LockedSource) Uint64() (n uint64) {
	r.lk.Lock()
	n = r.src.Uint64()
	r.lk.Unlock()
	return
}

func (r *LockedSource) Seed(seed int64) {
	r.lk.Lock()
	r.src.Seed(seed)
	r.lk.Unlock()
}

func (r *LockedSource) Restart() {
	r.lk.Lock()
	r.src.Restart()
	r.lk.Unlock()
}

type LockedJumpableSource struct {
	LockedSource
}

func (r *LockedJumpableSource) RestartSubstream() {
	r.lk.Lock()
	if js, ok := r.src.(JumpableSource); ok {
		js.RestartSubstream()
	}
	r.lk.Unlock()
}

func (r *LockedJumpableSource) Jump() {
	r.lk.Lock()
	if js, ok := r.src.(JumpableSource); ok {
		js.Jump()
	}
	r.lk.Unlock()
}
