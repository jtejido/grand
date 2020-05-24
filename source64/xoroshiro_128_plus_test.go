package source64_test

import (
	"github.com/jtejido/grand"
	"github.com/jtejido/grand/source64"
	"testing"
)

func TestXoRoShiRo128Plus(t *testing.T) {
	/*
	 * Data from running the executable compiled from the author's C code:
	 *   http://xoshiro.di.unimi.it/xoroshiro128plus.c
	 */
	rng, err := source64.NewXoRoShiRo128PlusFromStream([]uint64{0x012de1babb3c4104, 0xa5a818b8fc5aa503})

	if err != nil {
		t.Errorf("Error occured.")
	}

	r := grand.New(rng)

	expected := []uint64{
		0xa6d5fa73b796e607, 0xd419031a381fea2e, 0x28938b88b4972f52, 0x032793a0d51c1a27,
		0x50001cd69cc5b006, 0x44bbf571167cb7f0, 0x172f6a2f093b2bef, 0xe642c831f1e4f7bf,
		0xcec4e4b5d448032a, 0xc0164992807cbd59, 0xb96ff06c68515410, 0x5288e0312a0aae72,
		0x79a891c387d3be2e, 0x6c52f6f710db553b, 0x2ce6f6b1946862b3, 0x87eb1e1b24b47f11,
		0x9f7c3511c5f23bcf, 0x3254897533dcd1ab, 0x89d56ad217fbd1ad, 0x70f6b269f815f6e6,
		0xe8ee60efadfdb8c4, 0x09286db69fdd232b, 0xf440882651fc19e8, 0x6356fea018cc26cd,
		0xf692282b43fcb0c2, 0xef3f084929119bab, 0x355efbf5bedeb114, 0x6cf5089c2acc96dd,
		0x819c19e480f0bfd1, 0x414d12ff4082e261, 0xc9a33a52545dd374, 0x4675247e6fe89b3c,
		0x069f2e55cea155ba, 0x1e8d1dcf349746b8, 0xdf32e487bdd74523, 0xa544710cae2ad7cd,
		0xf5ac505e74fe049d, 0xf039e289da4cdf7e, 0x0a6fbebe9122529c, 0x880c51e0915031a3,
	}

	for i := 0; i < len(expected); i++ {
		rg := r.Uint64()
		if expected[i] != rg {
			t.Errorf("Mismatch. want: %v, got: %v", expected[i], rg)
		}
	}

}
