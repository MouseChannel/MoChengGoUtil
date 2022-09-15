package main

import (
	"fmt"
	container "github.com/MouseChannle/MoChengGoUtil/container"
)

type Test[myType any] struct {
	v myType
	m myType
}

func (test *Test[myType]) testFunc(newa myType) {
	test.m = newa
	fmt.Println(test.m)
}

func main() {

	fmt.Println(1231)

	t := &Test[int]{}
	t.testFunc(12)

	var myQueue container.MoChengQueue[int]

	for n := 0; n < 10; n++ {
		myQueue.Enqueue(n)
	}
	for n := 0; n < 10; n++ {
		myQueue.Dequeue()
	}
	_, b := myQueue.Dequeue()
	if b != nil {
		fmt.Println(b)
	}
	fmt.Print(123)
	myQueue.Dequeue()

	fmt.Print(123)

}
