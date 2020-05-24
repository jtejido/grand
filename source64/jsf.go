package source64

const (
	jsf64_r = 3
)

// Implement Bob Jenkins's small fast (JSF) 32-bit generator.
// The state size is 128-bits; the shortest period is expected to be about 2^94
// and it expected that about one seed will run into another seed within 2^64 values.
// https://burtleburtle.net/bob/rand/smallprng.html
type JSF struct {
	baseSource64
	state [4]uint64
}

func NewJSF(seed uint64) *JSF {
	ans := new(JSF)
	ans.spi = ans
	ans.Seed(int64(seed))
	return ans
}

func (jsf *JSF) setSeed(seed []uint64) {
	jsf.stream = append([]uint64{}, seed...)
	jsf.Restart()
}

func (jsf *JSF) Restart() {
	for i := 0; i < jsf64_r+1; i++ {
		jsf.state[i] = jsf.stream[i]
	}

	for i := 0; i < 20; i++ {
		jsf.Uint64()
	}

	jsf.resetState()
}

func (jsf *JSF) Uint64() uint64 {
	e := jsf.state[0] - rotateLeft(jsf.state[1], 7)
	jsf.state[0] = jsf.state[1] ^ rotateLeft(jsf.state[2], 13)
	jsf.state[1] = jsf.state[2] + rotateLeft(jsf.state[3], 37)
	jsf.state[2] = jsf.state[3] + e
	jsf.state[3] = e + jsf.state[0]
	return jsf.state[3]
}

func (jsf *JSF) Seed(seed int64) {
	seeds := make([]uint64, jsf64_r+1)
	var i int
	seeds[0] = uint64(0xf1ea5eed)
	for i < jsf64_r {
		seeds[i+1] = uint64(seed)
		i++
	}

	jsf.setSeed(seeds)
}

func (jsf *JSF) String() string {
	return "JSF"
}
