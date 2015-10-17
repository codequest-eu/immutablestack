package immutablestack

import (
	"errors"
	"testing"
)

func TestTop(t *testing.T) {
	str := "chunky bacon"
	stack := New().Push(str)
	top := stack.Top()
	if top != str {
		t.Errorf("stack.Top() = %v, expected %v", top, str)
	}
}

func TestPop(t *testing.T) {
	start := New()
	nextStack := start.Push("chunky bacon")
	popped := nextStack.Pop()
	if popped != start {
		t.Errorf("nextStack.Pop() = %v, expected %v", popped, start)
	}
}

func TestSize(t *testing.T) {
	stack := New().Push(1).Push(2).Push(3)
	stackSize := stack.Size()
	if stackSize != 3 {
		t.Errorf("stack.Size() = %d, expected 3", stackSize)
	}
}

func TestForEachIterationNotStopped(t *testing.T) {
	stack := New().Push(1).Push(2).Push(3)
	var slice []int
	iterator := func(s ImmutableStack) error {
		slice = append(slice, s.Top().(int))
		return nil
	}
	stack.ForEach(iterator)
	sliceLen := len(slice)
	if sliceLen != 3 {
		t.Errorf("sliceSize = %d, expected 3", sliceLen)
	}
	// Make sure everything is in the right position.
	if slice[0] != 3 {
		t.Errorf("slice[0] = %d, expected 3", slice[0])
	}
}

func TestForEachIterationStopped(t *testing.T) {
	stack := New().Push(1).Push(2).Push(3)
	var slice []int
	iterator := func(s ImmutableStack) error {
		if s.Top().(int) < 2 {
			return errors.New("I'm scared of small numbers")
		}
		slice = append(slice, s.Top().(int))
		return nil
	}
	stack.ForEach(iterator)
	sliceLen := len(slice)
	if sliceLen != 2 {
		t.Errorf("sliceSize = %d, expected 2", sliceLen)
	}
	// Make sure everything is in the right position.
	if slice[0] != 3 {
		t.Errorf("slice[0] = %d, expected 3", slice[0])
	}
}

func TestFMap(t *testing.T) {
	stack := New().Push(1).Push(2).Push(3)
	functor := func(in interface{}) interface{} {
		return in.(int) * 2
	}
	newStack := stack.FMap(functor)
	newSize := newStack.Size()
	if newSize != 3 {
		t.Errorf("newStack.Size() = %d, expected 3", newSize)
	}
	newTop := newStack.Top()
	if newTop != 6 {
		t.Errorf("newStack.Top() = %v, expected 6", newTop)
	}
}
