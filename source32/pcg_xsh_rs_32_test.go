package source32_test

import (
	"github.com/jtejido/grand"
	"github.com/jtejido/grand/source32"
	"testing"
)

func TestPcgXshRs32(t *testing.T) {
	rng, err := source32.NewPcgXshRs32FromStream([]uint64{0x012de1babb3c4104, 0xc8161b4202294965})

	if err != nil {
		t.Errorf("Error occured.")
	}

	r := grand.New(rng)

	expected := []uint32{
		0xba4138b8, 0xd329a393, 0x75d68d3f, 0xbb7572ca,
		0x7a48d2f2, 0xcb3c1e37, 0xc1374a97, 0x7c2c5bfa,
		0x8a1c8695, 0x30db4fea, 0x95f9a901, 0x72ebfa48,
		0x6a284dbf, 0x0ef11286, 0x37330e11, 0xfeb53893,
		0x77e3adda, 0x64dc86bd, 0xc8d762d7, 0xbf3fb80c,
		0x732dfd12, 0x6088e86d, 0xbc4e79e5, 0x56ece5b1,
		0xe706ac72, 0xee798018, 0xef73de74, 0x3de1f966,
		0x7a36db53, 0x1e921eb2, 0x55e35484, 0x2577c6f2,
		0x0a006e21, 0x8cb811b7, 0x5f26c916, 0x3990837f,
		0x15f2983d, 0x546ccb4a, 0x4eda8716, 0xb8666a25,
	}

	for i := 0; i < len(expected); i++ {
		rg := r.Uint32()
		if expected[i] != rg {
			t.Errorf("Mismatch. want: %v, got: %v", expected[i], rg)
		}
	}

}
