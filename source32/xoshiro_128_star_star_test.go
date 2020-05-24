package source32_test

import (
	"github.com/jtejido/grand"
	"github.com/jtejido/grand/source32"
	"testing"
)

func TestXoShiRo128StarStar(t *testing.T) {
	/*
	 * Data from running the executable compiled from the author's C code:
	 *   http://xoshiro.di.unimi.it/xoshiro128starstar.c
	 */
	rng, err := source32.NewXoShiRo128StarStarFromStream([]uint32{0x012de1ba, 0xa5a818b8, 0xb124ea2b, 0x18e03749})

	if err != nil {
		t.Errorf("Error occured.")
	}

	r := grand.New(rng)

	expected := []uint32{
		0x8856d912, 0xf2a19a86, 0x7693f66d, 0x23516f86,
		0x4895054e, 0xf4503fe6, 0x40e04672, 0x99244e34,
		0xb971815c, 0x3008b82c, 0x0ee73b58, 0x88aad2c6,
		0x7923f2e9, 0xfde55485, 0x7aed95f5, 0xeb8abb59,
		0xca78183a, 0x80ecdd68, 0xfd404b06, 0x248b9c9e,
		0xa2c69c6f, 0x1723b375, 0x879f37b0, 0xe98fd208,
		0x75de84a9, 0x717d6df8, 0x92cd7bc7, 0x46380167,
		0x7f08600b, 0x58566f2b, 0x7f781475, 0xe34ec04d,
		0x6d5ef889, 0xb76ff6d8, 0x501f5df6, 0x4cf70ccb,
		0xd7375b26, 0x457ea1ab, 0x7439e565, 0x355855af,
	}

	for i := 0; i < len(expected); i++ {
		rg := r.Uint32()
		if expected[i] != rg {
			t.Errorf("Mismatch. want: %v, got: %v", expected[i], rg)
		}
	}

}
