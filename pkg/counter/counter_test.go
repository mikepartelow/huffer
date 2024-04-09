package counter_test

import (
	"mp/huffer/pkg/counter"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	testCases := []struct {
		s    []rune
		want map[rune]int
	}{
		{
			s: []rune("foo bar"),
			want: map[rune]int{
				'a': 1,
				'b': 1,
				'f': 1,
				'o': 2,
				'r': 1,
				' ': 1,
			},
		},
	}
	for idx, tC := range testCases {
		t.Run(strconv.Itoa(idx), func(t *testing.T) {
			got := counter.Counts(tC.s)
			assert.Equal(t, tC.want, got)
		})
	}
}
