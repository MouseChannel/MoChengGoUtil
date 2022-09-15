package MoChengGoUtil

import (
	"errors"
)

type MoChengQueue[T any] struct {
	queue []T
}

func (q *MoChengQueue[T]) Enqueue(item T) {
	q.queue = append(q.queue, item)
}

func (q *MoChengQueue[T]) Peak() (res T, err error) {

	if len(q.queue) == 0 {

		return res, errors.New("empty queue")

	}
	return q.queue[0], nil
}
func (q *MoChengQueue[T]) Dequeue() (res T, err error) {

	if len(q.queue) == 0 {

		return res, errors.New("empty queue")

	}
	defer func() {
		q.queue = q.queue[1:]
	}()
	return q.queue[0], nil
}
 

func (q *MoChengQueue[T]) Size() int {
	return len(q.queue)
}
func (q *MoChengQueue[T]) IsEmpty() bool {
	return q.Size() == 0
}

func (q *MoChengQueue[T]) Clear() {
	q.queue = []T{}
}

func (q *MoChengQueue[T]) ToSlice() []T {

	return q.queue
}
