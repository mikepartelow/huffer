package huffman

import (
	"fmt"
	"mp/huffer/pkg/bitstream"
	"mp/huffer/pkg/counter"
	"mp/huffer/pkg/pq"
	"slices"
)

func Encode(s []rune) ([]byte, Table) {
	var buf bitstream.Buffer

	table := MakeTable(s)

	for _, r := range s {
		code := table[r]
		for i := range code.Width {
			bit := (int(code.Value) >> (code.Width - i - 1)) & 1
			err := buf.Write(bit)
			if err != nil {
				panic(err)
			}
		}
	}

	return buf.Bytes(), table
}

func Decode(s []byte, length int, table Table) ([]rune, error) {
	var runes []rune

	rtable := make(map[Code]rune)
	for r, c := range table {
		rtable[c] = r
	}

	buf := bitstream.NewBuffer(s)

	var e Encoding

	for w := 1; len(runes) < length; w++ {
		bit, err := buf.Read()
		if err != nil {
			return nil, fmt.Errorf("error decoding")
		}

		e = (e << 1) | Encoding(bit)

		if r, ok := rtable[Code{Width: w, Value: e}]; ok {
			runes = append(runes, r)
			e, w = 0, 0
		}
	}

	return runes, nil
}

type Encoding uint32

type Code struct {
	Width int
	Value Encoding
}

type Table map[rune]Code

func MakeTable(s []rune) Table {
	counts := counter.Counts[rune](s)

	// sort the counts by rune so that we make the same table each time
	var counted_runes []rune
	for r := range counts {
		counted_runes = append(counted_runes, r)
	}
	slices.Sort(counted_runes)

	q := pq.New[node]()

	for _, r := range counted_runes {
		q.Push(pq.Item[node]{
			Value:    node{R: r},
			Priority: counts[r],
		})
	}

	for q.Len() > 1 {
		a, b := q.Pop(), q.Pop()
		n := node{
			Left:  &a.Value,
			Right: &b.Value,
		}
		q.Push(pq.Item[node]{
			Value:    n,
			Priority: a.Priority + b.Priority,
		})
	}

	root := q.Pop()
	table := make(Table)

	traverse(root.Value, 0, 0, table)

	return table
}

type node struct {
	R     rune
	Left  *node
	Right *node
}

func traverse(n node, value Encoding, len int, table Table) {
	if n.Left == nil && n.Right == nil {
		table[n.R] = Code{Value: value, Width: len}
		return
	}
	if n.Left != nil {
		traverse(*n.Left, value<<1, len+1, table)
	}
	if n.Right != nil {
		traverse(*n.Right, value<<1|0b1, len+1, table)
	}
}
