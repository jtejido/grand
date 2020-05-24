package source32_test

import (
	"github.com/jtejido/grand"
	"github.com/jtejido/grand/source32"
	"testing"
)

func TestPcgMcgXshRs32(t *testing.T) {
	rng := source32.NewPcgMcgXshRs32(0x012de1babb3c4104)

	r := grand.New(rng)

	expected := []uint32{
		0xb786f832, 0x6920834f, 0x5b88b399, 0x6b811447,
		0x91230c70, 0x163c83b5, 0x8dd8bba9, 0xb8bcd10a,
		0xe1964b6e, 0x40b9adc8, 0x75fbee87, 0xed3d1e5c,
		0x82cb437b, 0xea94cea8, 0x76b1726a, 0x9275544a,
		0xed015249, 0x9d46c1cc, 0xe6fddd59, 0x487a0912,
		0xa709c922, 0xd15ac2a2, 0xba36e687, 0x3e40b099,
		0x62ae602c, 0xec0ebb27, 0x94246eda, 0xa40c2daa,
		0xd7e0abb5, 0xf8061587, 0x97f2132a, 0x861cfa5e,
		0xc5b2280b, 0x5fc8ec4e, 0xa9e552ed, 0xbf2ee34f,
		0x0a945eb3, 0x9e578662, 0x292cf72c, 0xc7e04668,
	}

	for i := 0; i < len(expected); i++ {
		rg := r.Uint32()
		if expected[i] != rg {
			t.Errorf("Mismatch. want: %v, got: %v", expected[i], rg)
		}
	}

}
