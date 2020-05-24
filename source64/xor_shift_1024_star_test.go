package source64_test

import (
	"github.com/jtejido/grand"
	"github.com/jtejido/grand/source64"
	"testing"
)

func TestXorShift1024Star(t *testing.T) {
	/*
	 * Data from running the executable compiled from the author's C code:
	 * http://xorshift.di.unimi.it/xorshift1024star.c
	 */
	rng, err := source64.NewXorShift1024StarFromStream([]uint64{
		0x012de1babb3c4104, 0xa5a818b8fc5aa503, 0xb124ea2b701f4993, 0x18e0374933d8c782,
		0x2af8df668d68ad55, 0x76e56f59daa06243, 0xf58c016f0f01e30f, 0x8eeafa41683dbbf4,
		0x7bf121347c06677f, 0x4fd0c88d25db5ccb, 0x99af3be9ebe0a272, 0x94f2b33b74d0bdcb,
		0x24b5d9d7a00a3140, 0x79d983d781a34a3c, 0x582e4a84d595f5ec, 0x7316fe8b0f606d20,
	})

	if err != nil {
		t.Errorf("Error occured.")
	}

	r := grand.New(rng)

	expected := []uint64{
		0xd85e9fc0855614cd, 0xaf4965c9c1ac6a3d, 0x067da398791111d8, 0x2771c41db58d7644,
		0xf71a471e1ac2b03e, 0x953449ae275f7409, 0x8aa570c72de0af5e, 0xae59db2acdae32be,
		0x3d46f316b8f97301, 0x72dc8399b7a70957, 0xf5624d788b3b6f4e, 0xb7a79275f6c0e7b1,
		0xf79354208377d498, 0x0e5d2f2ac2b4f28f, 0x0f8f57edc8aa802f, 0x5e918ea72ece0c36,
		0xeeb8dbdb00ac7a5a, 0xf16f88dfef0d6047, 0x1244c29e0e0d8d2d, 0xaa94f1cc42691eb7,
		0xd06425dd329e5de5, 0x968b1c2e016f159c, 0x6aadff7055065295, 0x3bce2efcb0d00876,
		0xb28d5b69ad8fb719, 0x1e4040c451376920, 0x6b0801a8a00de7d7, 0x891ba2cbe2a4675b,
		0x6355008481852527, 0x7a47bcd9960126f3, 0x07f72fcd4ebe3580, 0x4658b29c126840cc,
		0xdc7b36d3037c7539, 0x9e30aab0410122e8, 0x7215126e0fce932a, 0xda63f12a489fc8de,
		0x769997671b2a0158, 0xfa9cd84e0ffc174d, 0x34df1cd959dca211, 0xccea41a33ec1f763,
	}

	for i := 0; i < len(expected); i++ {
		rg := r.Uint64()
		if expected[i] != rg {
			t.Errorf("Mismatch. want: %v, got: %v", expected[i], rg)
		}
	}
}

func TestXorShift1024StarPhi(t *testing.T) {
	/*
	 * Data from running the executable compiled from the author's C code:
	 * http://xorshift.di.unimi.it/xorshift1024star.c
	 */
	rng, err := source64.NewXorShift1024StarPhiFromStream([]uint64{
		0x012de1babb3c4104, 0xa5a818b8fc5aa503, 0xb124ea2b701f4993, 0x18e0374933d8c782,
		0x2af8df668d68ad55, 0x76e56f59daa06243, 0xf58c016f0f01e30f, 0x8eeafa41683dbbf4,
		0x7bf121347c06677f, 0x4fd0c88d25db5ccb, 0x99af3be9ebe0a272, 0x94f2b33b74d0bdcb,
		0x24b5d9d7a00a3140, 0x79d983d781a34a3c, 0x582e4a84d595f5ec, 0x7316fe8b0f606d20,
	})

	if err != nil {
		t.Errorf("Error occured.")
	}

	r := grand.New(rng)

	expected := []uint64{
		0xc9351be6ae9af4bb, 0x2696a1a51e3040cb, 0xdcbbc38b838b4be8, 0xc989eee03351a25c,
		0xc4ad829b653ada72, 0x1cff4000cc0118df, 0x988f3aaf7bfb2852, 0x3a621d4d5fb27bf2,
		0x0153d81cf33ff4a7, 0x8a1b5adb974750c1, 0x182799e238df6de2, 0x92d9bda951cd6377,
		0x601f077d2a659728, 0x90536cc64ad5bc49, 0x5d99d9e84e3d7fa9, 0xfc66f4610240613a,
		0x0ff67da640cdd6b6, 0x973c7a6afbb41751, 0x5089cb5236ac1b5b, 0x7ed6edc1e4d7e261,
		0x3e37630df0c00b63, 0x49ec234a0d03bcc4, 0x2bcbe2fa4b80fa33, 0xbaafc47b960baefa,
		0x1855fa47be98c84f, 0x8d59cb57e34a73e0, 0x256b15bb001bf641, 0x28ad378895f5615d,
		0x865547335ba2a571, 0xfefe4c356e154585, 0xeb87f7a74e076680, 0x990d2f5d1e60b914,
		0x3bf0f6864688af2f, 0x8c6304df9b945d58, 0x63bc09c335b63666, 0x1038139f53734ad2,
		0xf41b58faf5680868, 0xa50ba830813c163b, 0x7dc1ca503ae39817, 0xea3d0f2f37f5ce95,
	}

	for i := 0; i < len(expected); i++ {
		rg := r.Uint64()
		if expected[i] != rg {
			t.Errorf("Mismatch. want: %v, got: %v", expected[i], rg)
		}
	}

}
