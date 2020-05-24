package source32_test

import (
	"github.com/jtejido/grand"
	"github.com/jtejido/grand/source32"
	"testing"
)

func TestLFSR88(t *testing.T) {
	rng, err := source32.NewLFSR88FromStream([]uint32{12345, 12345, 12345})

	if err != nil {
		t.Errorf("Error occured.")
	}

	r := grand.New(rng)

	expected := []uint32{
		1667269494,
		944790115,
		468047577,
		2424864938,
		995604853,
		2397442114,
		1827202028,
		964928120,
		246799526,
		2882122420,
		2730418434,
		4088964130,
		1803479424,
		617134916,
		2877139558,
		2290944594,
		132514910,
		3554423975,
		1833490227,
		3508282062,
		3719726999,
		173904632,
		1980325773,
		2206809856,
		2134273071,
		2485031560,
		2214057493,
		192612306,
		626565564,
		4232675098,
		3554341213,
		4211454268,
		1528935362,
		4212278447,
		55809696,
		2560487979,
		1096261590,
		1149151038,
		1678650882,
		320874650,
		234739846,
		2499729879,
		1316960977,
		603666633,
		2273116174,
		3886976459,
		3517795460,
		24959057,
		1550208826,
		2281000253,
		2617187324,
		678653465,
		4289035384,
		49735687,
		1120862736,
		2506506554,
		1418896048,
		796069851,
		3004281257,
		3876988827,
		3762857204,
		2201263032,
		2508818887,
		341729822,
		3615933095,
		1260059056,
		2972259213,
		2935606581,
		2074931220,
		1088637115,
		3452714511,
		1637828060,
		3293990344,
		3935430410,
		3877234248,
		1605252034,
		3993644804,
		436412609,
		3778847398,
		982943544,
		1875264518,
		2649872789,
		1878021489,
		1129525513,
		3936243390,
		1906853951,
		982886443,
		2377124342,
		2693019144,
		2044484340,
		3279856917,
		1471641653,
		3593190793,
		1845935623,
		1366588450,
		2599010224,
		2040029398,
		1844265475,
		3914785705,
		1633969602,
	}

	for i := 0; i < len(expected); i++ {
		rg := r.Uint32()
		if expected[i] != rg {
			t.Errorf("Mismatch. want: %v, got: %v", expected[i], rg)
		}
	}

}
