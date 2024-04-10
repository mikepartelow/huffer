package bitstream_test

import (
	"strconv"
	"testing"

	"mp/huffer/pkg/bitstream"

	"github.com/stretchr/testify/assert"
)

func TestBitstream(t *testing.T) {
	testCases := []struct {
		given []int
		want  []byte
	}{
		{
			given: []int{0b0},
			want:  []byte{0},
		},
		{
			given: []int{0b1},
			want:  []byte{0b10000000},
		},
		{
			given: []int{0b1, 0b0},
			want:  []byte{0b10000000},
		},
		{
			given: []int{0b1, 0b0, 0b1, 0b0, 0b1, 0b0, 0b1, 0b0},
			want:  []byte{0b10101010},
		},
		{
			given: []int{0b1, 0b0, 0b1, 0b0, 0b1, 0b0, 0b1, 0b0, 0b1},
			want:  []byte{0b10101010, 0b10000000},
		},
	}
	for i, tC := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var bs bitstream.Buffer
			for _, b := range tC.given {
				err := bs.Write(b)
				assert.NoError(t, err)
			}
			got := bs.Bytes()
			assert.Equal(t, tC.want, got)
		})
	}
}
