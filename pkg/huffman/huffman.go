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
		for i := range code.Len {
			bit := (int(code.Value) >> (code.Len - i - 1)) & 1

			err := buf.Write(bit)
			if err != nil {
				panic(err)
			}
		}
	}

	return buf.Bytes(), table
}

func Decode(s []byte, len int, table Table) ([]rune, error) {
	var runes []rune

	rtable := make(map[Encoding]rune)
	for r, c := range table {
		rtable[c.Value] = r
		fmt.Printf(" 0b%b = %c", c.Value, r)
	}
	fmt.Println("")

	for _, b := range s {
		fmt.Println("++", b)

		var code Encoding

		for i := uint32(0); ; i++ {

			if i == 8 {
				return nil, fmt.Errorf("could not decode input")
			}
			bit := (Encoding(b) << Encoding(i)) & (1 << Encoding(i))
			fmt.Printf("0b%b :: 0b%b :: 0b%b\n", Encoding(i), code, bit)

			code <<= 1
			code |= bit
			r, ok := rtable[code]
			if ok {
				runes = append(runes, r)
				fmt.Printf("-> %c\n", r)
				break
			}
		}
	}

	return runes, nil
}

type Encoding uint32

type Code struct {
	Len   int
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
