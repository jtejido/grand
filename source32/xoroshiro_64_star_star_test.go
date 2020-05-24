package source32_test

import (
	"github.com/jtejido/grand"
	"github.com/jtejido/grand/source32"
	"testing"
)

func TestXoRoShiRo64StarStar(t *testing.T) {
	/*
	 * Data from running the executable compiled from the author's C code:
	 *   http://xoshiro.di.unimi.it/xoroshiro64starstar.c
	 */
	rng, err := source32.NewXoRoShiRo64StarStarFromStream([]uint32{0x012de1ba, 0xa5a818b8})

	if err != nil {
		t.Errorf("Error occured.")
	}

	r := grand.New(rng)

	expected := []uint32{
		0x7ac00b42, 0x1f638399, 0x09e4aea4, 0x05cbbd64,
		0x1c967b7b, 0x1cf852fd, 0xc666f4e8, 0xeea9f1ae,
		0xca0fa6bc, 0xa65d0905, 0xa69afc95, 0x34965e62,
		0xdd4f04a9, 0xff1c9342, 0x638ff769, 0x03419ca0,
		0xb46e6dfd, 0xf7555b22, 0x8cab4e68, 0x5a44b6ee,
		0x4e5e1eed, 0xd03c5963, 0x782d05ed, 0x41bda3e3,
		0xd1d65005, 0x88f43a8a, 0xfffe02ea, 0xb326624a,
		0x1ec0034c, 0xb903d8df, 0x78454bd7, 0xaec630f8,
		0x2a0c9a3a, 0xc2594988, 0xe71e767e, 0x4e0e1ddc,
		0xae945004, 0xf178c293, 0xa04081d6, 0xdd9c062f,
	}

	for i := 0; i < len(expected); i++ {
		rg := r.Uint32()
		if expected[i] != rg {
			t.Errorf("Mismatch. want: %v, got: %v", expected[i], rg)
		}
	}

}
