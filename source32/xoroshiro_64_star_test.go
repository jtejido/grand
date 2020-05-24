package source32_test

import (
	"github.com/jtejido/grand"
	"github.com/jtejido/grand/source32"
	"testing"
)

func TestXoRoShiRo64Star(t *testing.T) {
	/*
	 * Data from running the executable compiled from the author's C code:
	 *   http://xoshiro.di.unimi.it/xoroshiro64star.c
	 */
	rng, err := source32.NewXoRoShiRo64StarFromStream([]uint32{0x012de1ba, 0xa5a818b8})

	if err != nil {
		t.Errorf("Error occured.")
	}

	r := grand.New(rng)

	expected := []uint32{
		0xd72accde, 0x29cbd26c, 0xa00fd44a, 0xa4d612c8,
		0xf9c7572b, 0xce94c084, 0x47a3d7ee, 0xb64aa982,
		0x67a9b2a4, 0x0c3d61a8, 0x8f70f7fa, 0xd1edbd63,
		0xac954b3a, 0xd7fe941e, 0xaa38e658, 0x019ecf61,
		0xcded7d7c, 0xd6588891, 0x4414454a, 0xb3c3a124,
		0x4a16fcfe, 0x3fb393c2, 0x4d8d14d6, 0x3a02c906,
		0x0c82f080, 0x174186c4, 0x1199966b, 0x12b83d6a,
		0xe697999e, 0x9df4d2f4, 0x5a5a0879, 0xc44ad6b4,
		0x96a9adc3, 0x4603c20f, 0x3171ca57, 0x66e349c9,
		0xa77dba19, 0xbe4f279d, 0xf5cd3402, 0x1962933d,
	}

	for i := 0; i < len(expected); i++ {
		rg := r.Uint32()
		if expected[i] != rg {
			t.Errorf("Mismatch. want: %v, got: %v", expected[i], rg)
		}
	}

}
