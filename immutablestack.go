package immutablestack

// Iterator is a function taking a each substack of a given stack.
// It can return an error to break iteration.
type Iterator func(ImmutableStack) error

// Functor is a function capable of modifying stack elements.
type Functor func(interface{}) interface{}

// ImmutableStack is an abstract immutable LIFO data structure.
type ImmutableStack interface {
	// Top returns the top element on the stack.
	Top() interface{}

	// Pop returns the next substack.
	Pop() ImmutableStack

	// Size returns the size of the stack.
	Size() uint64

	// Push pushes a new element onto the stack and returns a new stack.
	Push(interface{}) ImmutableStack

	// ForEach iterates over the stack and applies the iterator function to each
	// substack until it finds the empty one or the iterator returns an error.
	ForEach(Iterator)

	// FMap applies a function to each element on the stack and returns the new
	// stack with results of the function application.
	FMap(f Functor) ImmutableStack
}

// ImmutableStack is an abstract immutable LIFO data structure.
type immutableStackImpl struct {
	top  interface{}
	pop  ImmutableStack
	size uint64
}

// New returns a new instance of ImmutableStack.
func New() ImmutableStack {
	return &immutableStackImpl{
		top:  nil,
		pop:  nil,
		size: 0,
	}
}

func (i *immutableStackImpl) Top() interface{} {
	return i.top
}

func (i *immutableStackImpl) Pop() ImmutableStack {
	return i.pop
}

func (i *immutableStackImpl) Size() uint64 {
	return i.size
}

func (i *immutableStackImpl) Push(element interface{}) ImmutableStack {
	return &immutableStackImpl{
		top:  element,
		pop:  i,
		size: i.size + 1,
	}
}

func (i *immutableStackImpl) ForEach(iterator Iterator) {
	var cursor ImmutableStack = i
	for {
		if err := iterator(cursor); err != nil {
			break
		}
		cursor = cursor.Pop()
		if cursor.Top() == nil {
			break
		}
	}
}

func (i *immutableStackImpl) FMap(f Functor) ImmutableStack {
	if i.top == nil {
		return New()
	}
	return &immutableStackImpl{
		top:  f(i.top),
		pop:  i.Pop().FMap(f),
		size: i.size,
	}
}
