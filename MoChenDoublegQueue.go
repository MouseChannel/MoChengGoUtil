package MoChengGoUtil

import (
	"errors"
)

type MoChenDoublegQueue[T any] struct {
	queue []T
}

func (q *MoChenDoublegQueue[T]) EnqueueFront(item T) {

	q.queue = append([]T{item}, q.queue...)
}
func (q *MoChenDoublegQueue[T]) EnqueueBack(item T) {
	q.queue = append(q.queue, item)
}

func (q *MoChenDoublegQueue[T]) First() (res T, err error) {

	if len(q.queue) == 0 {

		return res, errors.New("empty queue")

	}
	return q.queue[0], nil
}
func (q *MoChenDoublegQueue[T]) Last() (res T, err error) {

	if len(q.queue) == 0 {

		return res, errors.New("empty queue")

	}
	return q.queue[len(q.queue)-1], nil
}
func (q *MoChenDoublegQueue[T]) DequeueFront() (res T, err error) {

	if len(q.queue) == 0 {

		return res, errors.New("empty queue")

	}
	defer func() {
		q.queue = q.queue[1:]
	}()
	return q.queue[0], nil
}
func (q *MoChenDoublegQueue[T]) DequeueBack() (res T, err error) {

	if len(q.queue) == 0 {

		return res, errors.New("empty queue")

	}
	defer func() {
		q.queue = q.queue[:len(q.queue)-1]
	}()
	return q.queue[0], nil
}

func (q *MoChenDoublegQueue[T]) Size() int {
	return len(q.queue)
}
func (q *MoChenDoublegQueue[T]) IsEmpty() bool {
	return q.Size() == 0
}

func (q *MoChenDoublegQueue[T]) Clear() {
	q.queue = []T{}
}

func (q *MoChenDoublegQueue[T]) ToSlice() []T {

	return q.queue
}
