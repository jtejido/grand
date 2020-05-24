package source64_test

import (
	"github.com/jtejido/grand"
	"github.com/jtejido/grand/source64"
	"testing"
)

func TestSFC(t *testing.T) {
	rng, err := source64.NewSFCFromStream([]uint64{0x012de1babb3c4104, 0xc8161b4202294965, 0xb5ad4eceda1ce2a9})

	if err != nil {
		t.Errorf("Error occured.")
	}

	r := grand.New(rng)

	expected := []uint64{
		0x383be11f844db7f4, 0x563e7e24056ad886, 0x959e56afde1c3f72, 0x7924b83a8ac40b01,
		0xe3096acc85876ae6, 0x9932c32968faf17e, 0x5df8e164496c717b, 0x443e63b0f0636d11,
		0xa0c1255bd56fb4ce, 0xe9b12d67fbae4394, 0x87a6b8f68968124b, 0xe7a29a2c9eb466b6,
		0xcfbcb67cac7ffb22, 0x9a77f8d8860be8e5, 0x51c287e3d450bf11, 0x9518d0a2cd3f16a3,
		0x36fdfd2044cbbb67, 0x94d6e5b7e50ed797, 0x01c80459dcc9ba5e, 0x913aa13874b1da2a,
		0x136f9eb31f816b8d, 0xbb68f2aba658e9f5, 0x455f38462bb2e598, 0x216693ead3d4036d,
		0x2e697d6093522eee, 0x8aa3e5e922c68cec, 0x55f38b99e6e9fadc, 0xc3b18937baf48d2f,
		0xd3a84a0f0781ef03, 0x0374b8766ea7b9a7, 0x354736eb92044fc2, 0x7e78cca53d9bb584,
		0x6b44e298f16ca140, 0xf1c7b84c51d8b1d8, 0x0bee55dd0ea4439d, 0xd9a26515c0a88471,
		0xda4c3174cafc57f8, 0x6193f4b96362eb4b, 0x207e9a94b58041af, 0x5451bd65c481d8fc,
	}

	for i := 0; i < len(expected); i++ {
		rg := r.Uint64()
		if expected[i] != rg {
			t.Errorf("Mismatch. want: %v, got: %v", expected[i], rg)
		}
	}

}
