package source32_test

import (
	"github.com/jtejido/grand"
	"github.com/jtejido/grand/source32"
	"testing"
)

func TestXoShiRo128Plus(t *testing.T) {
	/*
	 * Data from running the executable compiled from the author's C code:
	 *   http://xoshiro.di.unimi.it/xoshiro128plus.c
	 */
	rng, err := source32.NewXoShiRo128PlusFromStream([]uint32{0x012de1ba, 0xa5a818b8, 0xb124ea2b, 0x18e03749})

	if err != nil {
		t.Errorf("Error occured.")
	}

	r := grand.New(rng)

	expected := []uint32{
		0x1a0e1903, 0xfde55c35, 0xddb16b2e, 0xab949ac5,
		0xb5519fea, 0xc6a97473, 0x1f0403d9, 0x1bb46995,
		0x79c99a12, 0xe447ebce, 0xa8c31d78, 0x54d8bbe3,
		0x4984a039, 0xb411e932, 0x9c1f2c5e, 0x5f53c469,
		0x7f333552, 0xb368c7a1, 0xa57b8e66, 0xb29a9444,
		0x5c389bfa, 0x8e7d3758, 0xfe17a1bb, 0xcd0aad57,
		0xde83c4bb, 0x1402339d, 0xb557a080, 0x4f828bc9,
		0xde14892d, 0xbba8eaed, 0xab62ebbb, 0x4ad959a4,
		0x3c6ee9c7, 0x4f6a6fd3, 0xd5785eed, 0x1a0227d1,
		0x81314acb, 0xfabdfb97, 0x7e1b7e90, 0x57544e23,
	}

	for i := 0; i < len(expected); i++ {
		rg := r.Uint32()
		if expected[i] != rg {
			t.Errorf("Mismatch. want: %v, got: %v", expected[i], rg)
		}
	}

}
