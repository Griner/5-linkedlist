package linkedlist

type Item struct {
	value interface{}
	next  *Item
	prev  *Item
	list  *List
}

type List struct {
	length uint64
	first  *Item
	last   *Item
}

func (i *Item) Value() interface{} {
	return i.value
}

func (i *Item) Next() *Item {
	return i.next
}

func (i *Item) Prev() *Item {
	return i.prev
}

func (i *Item) Remove() {

	if i.next != nil && i.prev != nil {
		i.prev.next = i.next
		i.next.prev = i.prev
	} else if i.next != nil {
		i.next.prev = nil
		i.list.first = i.next
	} else if i.prev != nil {
		i.prev.next = nil
		i.list.last = i.prev
	} else {
		i.list.first = nil
		i.list.last = nil
	}

	i.list.length--

	i.next = nil
	i.prev = nil
	i.list = nil
}

func (l *List) Len() uint64 {
	return l.length
}

func (l *List) First() *Item {
	return l.first
}

func (l *List) Last() *Item {
	return l.last
}

func (l *List) PushFront(v interface{}) {

	i := &Item{v, nil, nil, l}

	if l.first != nil {
		i.next = l.first
		l.first.prev = i

	}

	if l.last == nil {
		l.last = i
	}

	l.first = i
	l.length++
}

func (l *List) PushBack(v interface{}) {

	i := &Item{v, nil, nil, l}

	if l.last != nil {
		i.prev = l.last
		l.last.next = i
	}

	if l.first == nil {
		l.first = i
	}

	l.last = i
	l.length++
}
