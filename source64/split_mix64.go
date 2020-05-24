package source64

const (
	golden_gamma uint64 = 0x9e3779b97f4a7c15
)

// A fast RNG, with 64 bits of state, that can be used to initialize the
// state of other generators. This was created by Steele, G.L., et al.
//
// http://xorshift.di.unimi.it/splitmix64.c
// http://gee.cs.oswego.edu/dl/papers/oopsla14.pdf
//
// ========= Summary results of Crush =========
//
//  Version:          TestU01 1.2.3
//  Generator:        Generator.Float64
//  Number of statistics:  144
//  Total CPU time:   02:58:50.76
//
// All tests were passed
//
// ========= Summary results of Crush =========
//
//  Version:          TestU01 1.2.3
//  Generator:        rand.Uint32
//  Number of statistics:  144
//  Total CPU time:   02:51:02.27
//  The following tests gave p-values outside [0.001, 0.9990]:
//  (eps  means a value < 1.0e-300):
//  (eps1 means a value < 1.0e-15):
//
//        Test                          p-value
//  ----------------------------------------------
//   8  CollisionOver, t = 8            7.9e-5
//  ----------------------------------------------
//  All other tests were passed
// TO-DO.. Jump() defined as state += (gamma * n) where n is the jump length (this can define JumpBack() when n is negative)
// The maximum length that would wrap the Weyl state would be 2^64.
type SplitMix64 struct {
	baseSource64
	state, substate uint64
}

func NewSplitMix64(seed uint64) *SplitMix64 {
	ans := new(SplitMix64)
	ans.baseSource64.spi = ans
	ans.substate = seed
	ans.state = ans.substate
	return ans
}

func (sm64 *SplitMix64) Uint64() uint64 {
	sm64.state += golden_gamma
	z := sm64.state
	z = (z ^ (z >> 30)) * 0xbf58476d1ce4e5b9
	z = (z ^ (z >> 27)) * 0x94d049bb133111eb
	return z ^ (z >> 31)
}

func (sm64 *SplitMix64) Seed(seed int64) {
	sm64.substate = uint64(seed)
	sm64.Restart()
}

func (sm64 *SplitMix64) Restart() {
	sm64.state = sm64.substate
	sm64.resetState()
}
