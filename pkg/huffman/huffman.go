package huffman

import (
	"mp/huffer/pkg/counter"
	"mp/huffer/pkg/pq"
	"slices"
)

func Encode(s []rune) []byte {
	return nil
}

type Code struct {
	Len   int
	Value int
}

type Table map[rune]Code

func MakeTable(s []rune) Table {
	counts := counter.Counts[rune](s)

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

func traverse(n node, value, len int, table Table) {
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
