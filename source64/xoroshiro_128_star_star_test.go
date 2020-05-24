package source64_test

import (
	"github.com/jtejido/grand"
	"github.com/jtejido/grand/source64"
	"testing"
)

func TestXoRoShiRo128StarStar(t *testing.T) {
	/*
	 * Data from running the executable compiled from the author's C code:
	 *  http://xorshift.di.unimi.it/xorshift1024star.c
	 */
	rng, err := source64.NewXoRoShiRo128StarStarFromStream([]uint64{0x012de1babb3c4104, 0xa5a818b8fc5aa503})

	if err != nil {
		t.Errorf("Error occured.")
	}

	r := grand.New(rng)

	expected := []uint64{
		0x8856e974cbb6da12, 0xd1704f3601beb952, 0x368a027941bc61e7, 0x882b24dcfa3ada58,
		0xfa8dafb3363143fb, 0x2eb9417a5dcf7654, 0xed8722e0a73e975b, 0x435fff57a631d485,
		0x954f1ad2377632b8, 0x9aa2f4dcba28ab71, 0xaca10369f96ac911, 0x968088e7277d0369,
		0x662e442ae32c42b4, 0xe1cd476f71dd058e, 0xb462a3c2bbb650f8, 0x74749215e8c07d08,
		0x1629f3cb1a671dbb, 0x3636dcc702eadf55, 0x97ae682e61cb3f57, 0xfdf8fc5ea9541f3b,
		0x2dfdb23d99c34acc, 0x68bef4f41a8f4113, 0x5cd03dc43f7af892, 0xdc2184abe0565da1,
		0x1dfaece40d9f96d0, 0x7d7b19285818ab71, 0xedea7fd3a0e47018, 0x23542ee7ed294823,
		0x1719f2b97bfc26c4, 0x2c7b7e288b399818, 0x49fa00786a1f5ad9, 0xd97cdfbe81700be2,
		0x557480baa4d9e5b2, 0x840a0403c0e85d92, 0xb4d5c6b2dc19dab2, 0xdf1b570e3bf1cf1b,
		0x26d1ac9455ccc75f, 0xdcc0e5fe06d1e231, 0x5164b7650568120e, 0x5fa82f6598483607,
	}

	for i := 0; i < len(expected); i++ {
		rg := r.Uint64()
		if expected[i] != rg {
			t.Errorf("Mismatch. want: %v, got: %v", expected[i], rg)
		}
	}

}
