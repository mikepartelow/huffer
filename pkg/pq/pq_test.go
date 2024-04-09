package pq_test

import (
	"strconv"
	"testing"

	"mp/huffer/pkg/pq"

	"github.com/stretchr/testify/assert"
)

func TestPQ(t *testing.T) {
	testCases := []struct {
		given []pq.Item[rune]
		want  []rune
	}{
		{
			given: []pq.Item[rune]{{Value: 'c', Priority: 100}, {Value: 'b', Priority: 50}, {Value: 'a', Priority: 0}},
			want:  []rune{'a', 'b', 'c'},
		},
		{
			given: []pq.Item[rune]{{Value: 'c', Priority: 0}, {Value: 'b', Priority: 50}, {Value: 'a', Priority: 100}},
			want:  []rune{'c', 'b', 'a'},
		},
	}
	for x, tC := range testCases {
		t.Run(strconv.Itoa(x), func(t *testing.T) {
			q := pq.New[rune]()
			for _, i := range tC.given {
				q.Push(i)
			}
			var got []rune
			for i := q.Pop(); i != nil; i = q.Pop() {
				got = append(got, i.Value)
			}
			assert.Equal(t, tC.want, got)
		})
	}
}
