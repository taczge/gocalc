package main

import (
	"errors"
	"fmt"
)

type Element struct {
	data int
	next *Element
}

func NewElement(data int, next *Element) *Element {
	return &Element{
		data: data,
		next: next,
	}
}

func (this *Element) String() string {
	if this.next == nil {
		return fmt.Sprintf("%d -> nil", this.data)
	} else {
		return fmt.Sprintf("%d -> %s", this.data, this.next)
	}
}

type Stack struct {
	top *Element
	size int
}

func NewStack() Stack {
	return Stack{}
}

func (this *Stack) String() string {
	if this.IsEmpty() {
		return "nil"
	}

	return this.top.String()
}

func (this *Stack) IsEmpty() bool {
	return this.top == nil
}

func (this *Stack) Push(n int) {
	this.size++

	if this.top == nil {
		this.top = NewElement(n, nil)
	} else {
		this.top = NewElement(n, this.top)
	}
}

func (this *Stack) Pop() (int, error) {
	if this.IsEmpty() {
		return 0, errors.New("Stack.Pop() is called for empty stack.")
	}

	data := this.top.data
	this.top = this.top.next
	this.size--

	return data, nil
}

func (this *Stack) Size() int {
	return this.size
}
