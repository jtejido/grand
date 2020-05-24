package source64_test

import (
	"github.com/jtejido/grand"
	"github.com/jtejido/grand/source64"
	"testing"
)

func TestXoShiRo512Plus(t *testing.T) {
	/*
	 * Data from running the executable compiled from the author's C code:
	 * http://xoshiro.di.unimi.it/xoshiro512plus.c
	 */
	rng, err := source64.NewXoShiRo512PlusFromStream([]uint64{
		0x012de1babb3c4104, 0xa5a818b8fc5aa503, 0xb124ea2b701f4993, 0x18e0374933d8c782,
		0x2af8df668d68ad55, 0x76e56f59daa06243, 0xf58c016f0f01e30f, 0x8eeafa41683dbbf4,
	})

	if err != nil {
		t.Errorf("Error occured.")
	}

	r := grand.New(rng)

	expected := []uint64{
		0xb252cbe62b5b8a97, 0xa4aaec677f60aaa2, 0x1c8bd694b50fd00e, 0x02753e0294233973,
		0xbfec0be86d152e2d, 0x5b9cd7265f320e98, 0xf8ec45eccc703724, 0x83fcbefa0359b3c1,
		0xbd27fcdb7e79265d, 0x88934227d8bf3cf0, 0x99e1e79384f40371, 0xe7e7fd0af2014912,
		0xebdd19cbcd35745d, 0x218994e1747243ee, 0x80628718e5d310da, 0x88ba1395debd989c,
		0x72e025c0928c6f55, 0x51400eaa050bbb0a, 0x72542ad3e7fe29e9, 0x3a3355b9dcb9c8b0,
		0x2f6618f3df6126f4, 0x34307608d886d40f, 0x34f5a22e98fe3375, 0x558f6560d08b9ec3,
		0xae78928bcb041d6c, 0xe7afe32a7caf4587, 0x22dcfb5ca129d4bd, 0x7c5a41864a6f2cf6,
		0xbe1186add0fe46a7, 0xd019fabc10dc96a5, 0xafa642ef6837d342, 0xdc4924811f62cf03,
		0xdeb486ccebccf747, 0xd827b16c9189f637, 0xf1aab3c3c690a71d, 0x6551214a7f04a2a5,
		0x44b8edb239f2a141, 0xb840cb37cfbeab59, 0x0e9558adc0987ca2, 0xc60442d5ff290606,
	}

	for i := 0; i < len(expected); i++ {
		rg := r.Uint64()
		if expected[i] != rg {
			t.Errorf("Mismatch. want: %v, got: %v", expected[i], rg)
		}
	}

}
