package huffman

import (
	"bytes"
	"mp/huffer/pkg/counter"
	"mp/huffer/pkg/pq"
	"slices"
)

func Encode(s []rune) ([]byte, Table) {
	var buf bytes.Buffer

	table := MakeTable(s)

	// var b uint32
	// var l int

	// for _, r := range s {
	// 	code := table[r]

	// 	b = (b << code.Len) | code.Value
	// 	must(buf.WriteByte(byte(b & 0xff)))
	// 	must(buf.WriteByte(byte(b & 0xff << 2)))
	// 	must(buf.WriteByte(byte(b & 0xff << 4)))
	// 	must(buf.WriteByte(byte(b & 0xff << 6)))

	// 	err := buf.WriteByte(b)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	return buf.Bytes(), table
}

type Code struct {
	Len   int
	Value uint32
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

func traverse(n node, value uint32, len int, table Table) {
	if n.Left == nil && n.Right == nil {
		table[n.R] = Code{Value: value, Len: len}
		return
	}
	if n.Left != nil {
		traverse(*n.Left, value<<1, len+1, table)
	}
	if n.Right != nil {
		traverse(*n.Right, value<<1|0b1, len+1, table)
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
