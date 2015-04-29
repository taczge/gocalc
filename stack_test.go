package main

import "testing"

func Test_IsEmpty_ReturnTrueForIninialStack(t *testing.T) {
	stack := NewStack()

	if !stack.IsEmpty() {
		t.Errorf("Initial stack is not empty.")
	}
}

func Test_IsEmpty_ReturnFalseForNonEmptyStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)

	if stack.IsEmpty() {
		t.Errorf("Stack.IsEmpty() for non-empty stack return false.")
	}
}

func Test_Pop_OccurErrorWhenStackIsEmpty(t *testing.T) {
	stack := NewStack()

	if _, err := stack.Pop(); err == nil {
		t.Errorf("Stack.Pop() of empty stack does not return nil.")
	}
}

func Test_Push(t *testing.T) {
	stack := NewStack()
	stack.Push(1)

	if stack.IsEmpty() {
		t.Errorf("Stack.IsEmpty() returns true for stack that pushed element.")
	}
}

func Test_IsEmpty_ReturnTrueForNonEmpyStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Pop()

	if !stack.IsEmpty() {
		t.Errorf("Stack.IsEmpty() does not return true for non-empty stack.")
	}
}

func Test_Pop_ReturnPushedNumber(t *testing.T) {
	stack := NewStack()
	stack.Push(1)

	n, err := stack.Pop()
	if err != nil {
		t.Errorf("An error occurs for non-empty stack.")
	}

	if n != 1 {
		t.Errorf("Stack.Pop() does not return pushed number.")
	}

	if !stack.IsEmpty() {
		t.Errorf("Stack.IsEmpty() for non-empty stack does not return true.")
	}
}

func Test_Push_ForStackThatHasMultipleElement(t *testing.T) {
	stack := NewStack()

	element_max := 5
	for i := 1; i <= element_max; i++ {
		stack.Push(i)
	}

	for i := element_max; i > 0; i-- {
		n, err := stack.Pop()
		if err != nil {
			t.Errorf("An error occurs for non-empty stack.")
		}

		if n != i {
			t.Errorf("Stack.Pop() does not return pushed number.")
		}
	}

	if !stack.IsEmpty() {
		t.Errorf("Stack.IsEmpty() for non-empty stack does not return true.")
	}
}

func Test_Size(t *testing.T) {
	stack := NewStack()
	checkStackSize(t, stack, 0)

	stack.Push(1)
	checkStackSize(t, stack, 1)

	stack.Push(1)
	checkStackSize(t, stack, 2)

	stack.Push(1)
	checkStackSize(t, stack, 3)

	stack.Pop()
	checkStackSize(t, stack, 2)

	stack.Pop()
	checkStackSize(t, stack, 1)

	stack.Pop()
	checkStackSize(t, stack, 0)
}

func checkStackSize(t *testing.T, stack Stack, size int) {
	if stack.Size() != size {
		t.Errorf("wrong stack size.")
	}
}
