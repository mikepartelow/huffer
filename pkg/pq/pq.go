package pq

import (
	"container/heap"
)

type Item[T any] struct {
	Value    T
	Priority int
}

type PriorityQueue[T any] struct {
	pq pQ[T]
}

func New[T any]() *PriorityQueue[T] {
	return &PriorityQueue[T]{}
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.pq.Len()
}

func (pq *PriorityQueue[T]) Push(i Item[T]) {
	heap.Push(&pq.pq, i)
}

func (pq *PriorityQueue[T]) Pop() *Item[T] {
	if len(pq.pq) == 0 {
		return nil
	}

	i := heap.Pop(&pq.pq)
	ii := i.(Item[T])
	return &ii
}

type pQ[T any] []Item[T]

func (pq pQ[T]) Len() int { return len(pq) }

func (pq pQ[T]) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq pQ[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *pQ[T]) Push(x any) {
	item := x.(Item[T])
	*pq = append(*pq, item)
}

func (pq *pQ[T]) Pop() any {
	old := *pq
	n := len(old)

	item := old[n-1]
	*pq = old[0 : n-1]

	return item
}
