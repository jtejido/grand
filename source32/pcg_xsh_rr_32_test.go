package source32_test

import (
	"github.com/jtejido/grand"
	"github.com/jtejido/grand/source32"
	"testing"
)

func TestPcgXshRr32(t *testing.T) {
	rng, err := source32.NewPcgXshRr32FromStream([]uint64{0x012de1babb3c4104, 0xc8161b4202294965})

	if err != nil {
		t.Errorf("Error occured.")
	}

	r := grand.New(rng)

	expected := []uint32{
		0xe860dd24, 0x15d339c0, 0xd9f75c46, 0x00efabb7,
		0xa625e97f, 0xcdeae599, 0x6304e667, 0xbc81be11,
		0x2b8ea285, 0x8e186699, 0xac552be9, 0xd1ae72e5,
		0x5b953ad4, 0xa061dc1b, 0x526006e7, 0xf5a6c623,
		0xfcefea93, 0x3a1964d2, 0xd6f03237, 0xf3e493f7,
		0x0c733750, 0x34a73582, 0xc4f8807b, 0x92b741ca,
		0x0d38bf9c, 0xc39ee6ad, 0xdc24857b, 0x7ba8f7d8,
		0x377a2618, 0x92d83d3f, 0xd22a957a, 0xb6724af4,
		0xe116141a, 0xf465fe45, 0xa95f35bb, 0xf0398d4d,
		0xe880af3e, 0xc2951dfd, 0x984ec575, 0x8679addb,
	}

	for i := 0; i < len(expected); i++ {
		rg := r.Uint32()
		if expected[i] != rg {
			t.Errorf("Mismatch. want: %v, got: %v", expected[i], rg)
		}
	}

}
