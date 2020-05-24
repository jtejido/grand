package source32_test

import (
	"github.com/jtejido/grand"
	"github.com/jtejido/grand/source32"
	"testing"
)

func TestJSF(t *testing.T) {
	rng := source32.NewJSF(0xb5ad4ece)

	r := grand.New(rng)

	expected := []uint32{
		0x3b05df0d, 0xc1b222b1, 0xdc38504a, 0x5a929fee,
		0x695f52ee, 0x49246926, 0xeaca3aaa, 0xb7ea1598,
		0x6f946a66, 0xf4eddf53, 0x4235b7bf, 0x4b1eb3c6,
		0xffa13fa2, 0x095ab9fc, 0x64dc8c5c, 0x3ad18ba8,
		0xb5f8354d, 0x744ef6de, 0xff9d2943, 0xb3d54756,
		0x096e9c74, 0x142a29c5, 0xcf090298, 0x71823d63,
		0x587052d2, 0xb843e5ed, 0x670e0279, 0xc5bb26d5,
		0xc28d61e0, 0xd31aedaf, 0x52fe2b77, 0x65f50ec7,
		0x522a44c5, 0x25f4baf8, 0x9fd1d806, 0x3a24f3bc,
		0x78f2aac1, 0xce496e14, 0x74d186b8, 0x34ff8809,
	}

	for i := 0; i < len(expected); i++ {
		rg := r.Uint32()
		if expected[i] != rg {
			t.Errorf("Mismatch. want: %v, got: %v", expected[i], rg)
		}
	}

}
