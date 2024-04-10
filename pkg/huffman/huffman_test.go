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
				// this encoding is different from the one on wikipedia because huffman codes are not unique, particularly in the case of a "tie"
				// in frequency counts, which this example has.
				0b10010011, 0b01000010, 0b01000011, 0b11011000, 0b11000011, 0b00111111, 0b00001111, 0b11011111, 0b10011001, 0b11111101, 0b00011000, 0b01101111, 0b10111010, 0b01111111, 0b00000000,
			},
		},
	}

	for _, tC := range testCases {
		t.Run(string(tC.given), func(t *testing.T) {
			got, _ := huffman.Encode(tC.given)
			assert.Equal(t, tC.want, got)
		})
	}
}
