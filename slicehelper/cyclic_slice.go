package slicehelper

type CyclicSlice[T any] struct {
	data  []T
	size  int
	index int
	full  bool
}

func NewCyclicSlice[T any](size int) *CyclicSlice[T] {
	return &CyclicSlice[T]{
		data:  make([]T, size),
		size:  size,
		index: 0,
		full:  false,
	}
}

func (cs *CyclicSlice[T]) GetData() []T {
	d := make([]T, 0, cs.GetCurrentSize())
	if cs.IsFull() {
		d = append(d, cs.data[cs.index:cs.size]...)
	}
	d = append(d, cs.data[0:cs.index]...)
	return d
}

func (cs *CyclicSlice[T]) Add(value T) {
	cs.data[cs.index] = value
	cs.index = (cs.index + 1) % cs.size
	if cs.index == 0 {
		cs.full = true
	}
}

func (cs *CyclicSlice[T]) SetAt(index int, value T) {
	cs.data[cs.getIndex(index)] = value
}

func (cs *CyclicSlice[T]) GetAt(index int) T {
	return cs.data[cs.getIndex(index)]
}

func (cs *CyclicSlice[T]) getIndex(index int) int {
	if index < 0 || index >= cs.GetCurrentSize() {
		panic("index out of bounds")
	}
	if cs.full {
		return (cs.index + index) % cs.size
	} else {
		return index
	}
}

func (cs *CyclicSlice[T]) IsFull() bool {
	return cs.full
}

func (cs *CyclicSlice[T]) GetCurrentSize() int {
	if cs.full {
		return cs.size
	}
	return cs.index
}

func (cs *CyclicSlice[T]) GetLast() T {
	if cs.index == 0 && !cs.full {
		panic("cyclic slice is empty")
	}
	return cs.GetAt(cs.GetCurrentSize() - 1)
}

func (cs *CyclicSlice[T]) GetPrevious() T {
	if cs.index == 0 && !cs.full {
		panic("cyclic slice is empty")
	}
	return cs.GetAt(cs.GetCurrentSize() - 2)
}

func (cs *CyclicSlice[T]) Reset() {
	cs.full = false
	cs.index = 0
	//cs.data = cs.data[:0]
}
