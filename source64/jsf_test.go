package source64_test

import (
	"github.com/jtejido/grand"
	"github.com/jtejido/grand/source64"
	"testing"
)

func TestJSF(t *testing.T) {
	rng := source64.NewJSF(0x012de1babb3c4104)

	r := grand.New(rng)

	expected := []uint64{
		0xb2eb2f629a2818c2, 0xe6c4df3bd8e4a0c8, 0x2b3ab71e4e888b46, 0x12a6088f5960738d,
		0x95715b21fcb1a7d9, 0x7acafc3916723b0f, 0x3a0c5f8c4caff822, 0x9b47b7a1e9784699,
		0x9c399839261a024f, 0x56a2fa6eaa7a62aa, 0xca6995ea5baeb8da, 0x56cad0c4dee9cbb9,
		0xbb5df57850f117a5, 0x147a41dad6a87b7b, 0xf9225f2aa6485812, 0x812b9d2c9b99aaa0,
		0x266ad947cac0acfc, 0x19bcfc1b69831866, 0xc5486e1cfa0eca28, 0x80ca1802e7dd04b7,
		0x003addd1e44ff095, 0xb9eaa245ce7c040b, 0xe607e64b31a6e9b4, 0x1553718b8013007b,
		0x86dcd29120fd807b, 0xeb5b8ec5d73dc39e, 0x3c26147f6b7ff7d7, 0xe0b994497bf55bb5,
		0x24fb3dc33de779c6, 0x022aba70fc48e04a, 0xbcf938e19b81f27f, 0x9022bd08a8ac7511,
		0x79ad91f7404ecef1, 0x291858706a2286db, 0xf395681f493eb602, 0xf85ed536da160b93,
		0x5dd685454dd0d913, 0x150e7b8f99b10f7d, 0xcd1c0b519cc69c05, 0xca92e08bf2676077,
	}

	for i := 0; i < len(expected); i++ {
		rg := r.Uint64()
		if expected[i] != rg {
			t.Errorf("Mismatch. want: %v, got: %v", expected[i], rg)
		}
	}

}
