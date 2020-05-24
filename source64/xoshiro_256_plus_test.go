package source64_test

import (
	"github.com/jtejido/grand"
	"github.com/jtejido/grand/source64"
	"testing"
)

func TestXoShiRo256Plus(t *testing.T) {
	/*
	 * Data from running the executable compiled from the author's C code:
	 *  http://xoshiro.di.unimi.it/xoshiro256plus.c
	 */
	rng, err := source64.NewXoShiRo256PlusFromStream([]uint64{0x012de1babb3c4104, 0xa5a818b8fc5aa503, 0xb124ea2b701f4993, 0x18e0374933d8c782})

	if err != nil {
		t.Errorf("Error occured.")
	}

	r := grand.New(rng)

	expected := []uint64{
		0x1a0e1903ef150886, 0x08b605f47abc5d75, 0xd82176096ac9be31, 0x8fbf2af9b4fa5405,
		0x9ab074b448171964, 0xfd68cc83ab4360aa, 0xf431f7c0c8dc6f2b, 0xc04430be08212638,
		0xc1ad670648f1da03, 0x3eb70d38796ba24a, 0x0e474d0598251ed2, 0xf9b6b3b56482566b,
		0x3d11e529ae07a7c8, 0x3b195f84f4db17e7, 0x09d62e817b8223e2, 0x89dc4db9cd625509,
		0x52e04793fe977846, 0xc052428d6d7d17cd, 0x6fd6f8da306b10ef, 0x64a7996ba5cc80aa,
		0x03abf59b95a1ef20, 0xc5a82fc3cfb50234, 0x0d401229eabb2d39, 0xb537b249f70bd18a,
		0x1af1b703753fcf4d, 0xb84648c1945d9ccb, 0x1d321bea673e1f66, 0x93d4445b268f305f,
		0xc046cfa36d89a312, 0x8cc2d55bbf778790, 0x1d668b0a3d329cc7, 0x81b6d533dfcf82de,
		0x9ca1c49a18537b16, 0x68e55c4054e0cb72, 0x06ed1956cb69afc6, 0x4871e696449da910,
		0xcfbd7a145066d46e, 0x10131cb15004b62d, 0x489c91a322bca3b6, 0x8ec95fa9bef73e66,
	}

	for i := 0; i < len(expected); i++ {
		rg := r.Uint64()
		if expected[i] != rg {
			t.Errorf("Mismatch. want: %v, got: %v", expected[i], rg)
		}
	}

}
