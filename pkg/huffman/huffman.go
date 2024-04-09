package huffman

import (
	"mp/huffer/pkg/counter"
	"mp/huffer/pkg/pq"
)

func Encode[T comparable](s []T) []byte {
	return nil
}

type Code struct {
	Len   int
	Value int
}

type Table[T comparable] map[T]Code

func MakeTable[T comparable](s []T) Table[T] {
	counts := counter.Counts[T](s)

	q := pq.New[node[T]]()

	for r, count := range counts {
		q.Push(pq.Item[node[T]]{
			Value:    node[T]{R: r},
			Priority: count,
		})
	}

	for q.Len() > 1 {
		a, b := q.Pop(), q.Pop()
		n := node[T]{
			Left:  &a.Value,
			Right: &b.Value,
		}
		q.Push(pq.Item[node[T]]{
			Value:    n,
			Priority: a.Priority + b.Priority,
		})
	}

	root := q.Pop()
	table := make(Table[T])

	traverse(root.Value, 0, 0, table)

	return table
}

type node[T comparable] struct {
	R     T
	Left  *node[T]
	Right *node[T]
}

func traverse[T comparable](n node[T], value, len int, table Table[T]) {
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
