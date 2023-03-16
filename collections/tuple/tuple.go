package tuple

type Tuple interface {
	Size() int
	SetElements(elems []any)
	GetElements() []interface{}
	At(index int) interface{}
	Set(index int, x interface{}) interface{}
}

type tuple struct {
	size int
	data []interface{}
}

func New(size int) Tuple {
	t := &tuple{
		size: size,
	}
	return t
}

func NewFrom(elems []any) Tuple {
	size := len(elems)
	t := New(size)
	t.SetElements(elems)
	return t
}

func (t *tuple) Size() int {
	return t.size
}

func (t *tuple) SetElements(elems []any) {
	if t.data == nil || len(t.data) == 0 {
		t.data = make([]interface{}, len(elems))
	}
	for i := range t.data {
		t.data[i] = elems[i]
	}
}

func (t *tuple) GetElements() []interface{} {
	return t.data
}

func (t *tuple) At(index int) interface{} {
	if index < 0 || index >= t.Size() {
		return nil
	}
	return t.data[index]
}

func (t *tuple) Set(index int, x interface{}) interface{} {
	if index < 0 || index >= t.Size() {
		return nil
	}
	old := t.data[index]
	t.data[index] = x
	return old
}
