package source32

const (
	jsf32_r = 3
)

// Implement Bob Jenkins's small fast (JSF) 32-bit generator.
// The state size is 128-bits; the shortest period is expected to be about 2^94
// and it expected that about one seed will run into another seed within 2^64 values.
// https://burtleburtle.net/bob/rand/smallprng.html
type JSF struct {
	baseSource32
	state [4]uint32
}

func NewJSF(seed uint32) *JSF {
	ans := new(JSF)
	ans.spi = ans
	ans.Seed(int64(seed)<<32 | int64(seed)&0xffffffff)
	return ans
}

func (jsf *JSF) setSeed(seed []uint32) {
	jsf.stream = append([]uint32{}, seed...)
	jsf.Restart()
}

func (jsf *JSF) Restart() {
	for i := 0; i < jsf32_r+1; i++ {
		jsf.state[i] = jsf.stream[i]
	}

	for i := 0; i < 20; i++ {
		jsf.Uint32()
	}

	jsf.resetState()
}

func (jsf *JSF) Uint32() uint32 {
	e := jsf.state[0] - rotateLeft(jsf.state[1], 27)
	jsf.state[0] = jsf.state[1] ^ rotateLeft(jsf.state[2], 17)
	jsf.state[1] = jsf.state[2] + jsf.state[3]
	jsf.state[2] = jsf.state[3] + e
	jsf.state[3] = e + jsf.state[0]
	return jsf.state[3]
}

func (jsf *JSF) Seed(seed int64) {
	seeds := make([]uint32, jsf32_r+1)
	var i int
	seeds[0] = 0xf1ea5eed
	for i < jsf32_r {
		seeds[i+1] = uint32(seed)
		i++
	}

	jsf.setSeed(seeds)
}

func (jsf *JSF) String() string {
	return "JSF"
}
