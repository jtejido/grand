package source32_test

import (
	"github.com/jtejido/grand"
	"github.com/jtejido/grand/source32"
	"testing"
)

func TestSFC(t *testing.T) {
	rng, err := source32.NewSFCFromStream([]uint32{0xbb3c4104, 0x02294965, 0xda1ce2a9})

	if err != nil {
		t.Errorf("Error occured.")
	}

	r := grand.New(rng)

	expected := []uint32{
		0x89b5c414, 0x7ee57639, 0xdbe18f7b, 0x94aa0162,
		0xa22bff0a, 0x21c91fb8, 0x2c6fd6fe, 0xcda90d13,
		0x019684ca, 0xe5b400c0, 0x459d8590, 0x3aec0a1e,
		0x254dac77, 0xbe10ae80, 0x9ac27819, 0xd17d10c6,
		0x71a69026, 0x4bb2bdda, 0x70853646, 0xda28f272,
		0x879200d9, 0x8c2f8b5b, 0x8a87cb78, 0x27ffdced,
		0x988a2b7b, 0xf220ef9b, 0x13b8984f, 0x345d1732,
		0x8f5bc6fc, 0x092b09ff, 0x046bd2b0, 0xa5a99fc5,
		0x19400604, 0xb76e7394, 0x037addd5, 0xe916ed79,
		0x94f10dc6, 0xf2ecb45e, 0x69834355, 0xb814aeb2,
	}

	for i := 0; i < len(expected); i++ {
		rg := r.Uint32()
		if expected[i] != rg {
			t.Errorf("Mismatch. want: %v, got: %v", expected[i], rg)
		}
	}

}
