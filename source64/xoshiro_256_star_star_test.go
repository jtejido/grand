package source64_test

import (
	"github.com/jtejido/grand"
	"github.com/jtejido/grand/source64"
	"testing"
)

func TestXoShiRo256StarStar(t *testing.T) {
	/*
	 * Data from running the executable compiled from the author's C code:
	 * http://xoshiro.di.unimi.it/xoshiro256starstar.c
	 */
	rng, err := source64.NewXoShiRo256StarStarFromStream([]uint64{0x012de1babb3c4104, 0xa5a818b8fc5aa503, 0xb124ea2b701f4993, 0x18e0374933d8c782})

	if err != nil {
		t.Errorf("Error occured.")
	}

	r := grand.New(rng)

	expected := []uint64{
		0x462c422df780c48e, 0xa82f1f6031c183e6, 0x8a113820e8d2ca8d, 0x1ac7023a26534958,
		0xac8e41d0101e109c, 0x46e34bc13edd63c4, 0x3a26776adcd665c3, 0x9ac6c9bea8fc518c,
		0x1cef0aa07cc738c4, 0x5136a5f070244b1d, 0x12e2e12edee691ff, 0x28942b20799b71b4,
		0xbe2d5c4267af2469, 0x9dbec53728b2b9b7, 0x893cf86611b14a96, 0x712c226c79f066d6,
		0x1a8a11ef81d2ac60, 0x28171739ef8f2f46, 0x073baa93525f8b1d, 0xa73c7f3cb93df678,
		0xae5633ab977a3531, 0x25314041ba2d047e, 0x31e6819dea142672, 0x9479fa694f4c2965,
		0xde5b771a968472b7, 0xf0501965d9eeb4a3, 0xef25a2a8ec90b911, 0x1f58f71a75392659,
		0x32d9547188781f3c, 0x2d13b036ccf65bc0, 0x289f9cc038dd952f, 0x6ae2d5231e50824a,
		0x75651acfb42ab170, 0x7369aeb4f10056cf, 0x0297ed632a97cf75, 0x19f534c778015b72,
		0x5d1d111c5ff182a8, 0x861cdfe8e8014b96, 0x07c6071e08112c83, 0x15601582dcf4e4fe,
	}

	for i := 0; i < len(expected); i++ {
		rg := r.Uint64()
		if expected[i] != rg {
			t.Errorf("Mismatch. want: %v, got: %v", expected[i], rg)
		}
	}

}
