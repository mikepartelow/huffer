package huffman_test

import (
	"testing"

	"mp/huffer/pkg/huffman"

	"github.com/stretchr/testify/assert"
)

func TestMakeTable(t *testing.T) {
	testCases := []struct {
		given []rune
		want  huffman.Table
	}{
		{
			// https://en.wikipedia.org/wiki/Huffman_coding#Compression
			given: []rune("A_DEAD_DAD_CEDED_A_BAD_BABE_A_BEADED_ABACA_BED"),
			want: huffman.Table{
				'D': huffman.Code{Len: 2, Value: 0b00},
				'_': huffman.Code{Len: 2, Value: 0b01},
				'A': huffman.Code{Len: 2, Value: 0b10},
				'E': huffman.Code{Len: 3, Value: 0b110},
				'C': huffman.Code{Len: 4, Value: 0b1110},
				'B': huffman.Code{Len: 4, Value: 0b1111},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(string(tC.given), func(t *testing.T) {
			got := huffman.MakeTable(tC.given)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestEncode(t *testing.T) {
	testCases := []struct {
		given []rune
		want  []byte
	}{
		{
			given: []rune("A_DEAD_DAD_CEDED_A_BAD_BABE_A_BEADED_ABACA_BED"),
			want: []byte{
				0x87,
				// and so on
			},
			// want: // 0b1000011101001000110010011101100111001001000111110010011111011111100010001111110100111001001011111011101000111111001
		},
	}
	for _, tC := range testCases {
		t.Run(string(tC.given), func(t *testing.T) {
			got, _ := huffman.Encode(tC.given)
			assert.Equal(t, tC.want, got)
		})
	}
}
