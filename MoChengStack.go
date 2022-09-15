package MoChengGoUtil

import (
	"errors"
)

type MoChengStack[T any] struct {
	stack []T
}

func (q *MoChengStack[T]) Push(item T) {
	q.stack = append(q.stack, item)
}

func (q *MoChengStack[T]) Peak() (res T, err error) {

	if len(q.stack) == 0 {

		return res, errors.New("empty stack")

	}
	return q.stack[len(q.stack)-1], nil
}
func (q *MoChengStack[T]) Pop() (res T, err error) {

	if len(q.stack) == 0 {

		return res, errors.New("empty stack")

	}
	defer func() {
		q.stack = q.stack[:len(q.stack)-1]
	}()
	return q.stack[len(q.stack)-1], nil
}
 

func (q *MoChengStack[T]) Size() int {
	return len(q.stack)
}
func (q *MoChengStack[T]) IsEmpty() bool {
	return q.Size() == 0
}

func (q *MoChengStack[T]) Clear() {
	q.stack = []T{}
}

func (q *MoChengStack[T]) ToSlice() []T {

	return q.stack
}
