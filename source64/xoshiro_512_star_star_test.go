package source64_test

import (
	"github.com/jtejido/grand"
	"github.com/jtejido/grand/source64"
	"testing"
)

func TestXoShiRo512StarStar(t *testing.T) {
	/*
	 * Data from running the executable compiled from the author's C code:
	 * http://xoshiro.di.unimi.it/xoshiro512starstar.c
	 */
	rng, err := source64.NewXoShiRo512StarStarFromStream([]uint64{
		0x012de1babb3c4104, 0xa5a818b8fc5aa503, 0xb124ea2b701f4993, 0x18e0374933d8c782,
		0x2af8df668d68ad55, 0x76e56f59daa06243, 0xf58c016f0f01e30f, 0x8eeafa41683dbbf4,
	})

	if err != nil {
		t.Errorf("Error occured.")
	}

	r := grand.New(rng)

	expected := []uint64{
		0x462c422df780c48e, 0xa82f1f6031c183e6, 0x60559add0e1e369a, 0xf956a2b900083a8d,
		0x0e5c039df1576573, 0x2f35cef71b14aa24, 0x5809ea8aa1d5a045, 0x3e695e3189ccf9bd,
		0x1eb940ee4bcb1a08, 0x78b72a0927bd9257, 0xe1a8e8d6dc64600b, 0x3993bff6e1378a4b,
		0x439161ee3b5d5cc8, 0xac6ca2359fe7f321, 0xc4238c5785d320e2, 0x75cf64526530aed5,
		0x679241ffc120e2b1, 0xded30a8f20b24c73, 0xff8ac62cff0deb9b, 0xe63a25973df23c45,
		0x74742f9096c56401, 0xc573afa2368288ac, 0x9b1048cf2daf9f9d, 0xe7d9720c2f51ca5f,
		0x38a21e1f7a441ced, 0x78835d75a9bd17a6, 0xeb64167a723de35f, 0x9455dd663e40620c,
		0x88693a769f203ed1, 0xea5f0997a281cffc, 0x2662b83f835f3273, 0x5e90efde2150ed04,
		0xd481b14551c8f8d9, 0xf2e4d714a0ab22d7, 0xdfb1a8f0637a2013, 0x8cd8d8c353640028,
		0xb4ce3b66785e0cc6, 0xa51386e09b6ab734, 0xfeac4151ac4a3f8d, 0x0e5679853ab5180b,
	}

	for i := 0; i < len(expected); i++ {
		rg := r.Uint64()
		if expected[i] != rg {
			t.Errorf("Mismatch. want: %v, got: %v", expected[i], rg)
		}
	}

}
