package linkedlist

import "testing"

func TestPushBack(t *testing.T) {

	list := &List{}

	if list.Len() != 0 {
		t.Fatal("List is not empty")
	}

	list.PushBack(5)

	if list.First() != list.Last() {
		t.Fatal("First item is not equal Last")
	}

	list.PushBack(4)
	list.PushBack(3)

	if list.Len() != 3 {
		t.Fatal("List length is not right")
	}

	if list.First().Value() != 5 {
		t.Fatal("First item is not right")
	}

	if list.Last().Value() != 3 {
		t.Fatal("Last item is not right")
	}

	if list.First().Prev() != nil {
		t.Fatal("Item before the first item is not nil")
	}

	item := list.Last()
	if item.Prev().Value() != 4 {
		t.Fatal("Middle item is not right")
	}
	if item.Next() != nil {
		t.Fatal("Item after the last item is not nil")
	}
	if item.Prev().Next().Value() != 3 {
		t.Fatal("Item after Middle item is not right")
	}
	if item.Prev().Prev().Value() != 5 {
		t.Fatal("Item before Middle item is not right")
	}
}

func TestPushFront(t *testing.T) {

	list := &List{}

	list.PushBack(5)
	list.PushBack(4)
	list.PushBack(3)

	list.PushFront(10)
	list.PushFront(9)
	list.PushFront(8)

	if list.Len() != 6 {
		t.Fatal("List length is not right")
	}

	if list.First().Value() != 8 {
		t.Fatal("First item is not right")
	}

	if list.First().Next().Value() != 9 {
		t.Fatal("Item after First is not right")
	}

}

func TestRemoveItemFirst(t *testing.T) {

	list := &List{}

	list.PushBack(5)
	list.PushBack(4)
	list.PushBack(3)
	list.PushFront(10)
	list.PushFront(9)
	list.PushFront(8)

	item := list.First()
	item.Remove()

	if list.Len() != 5 {
		t.Fatal("List length is not right")
	}

	if list.First().Value() != 9 {
		t.Fatal("First item is not right")
	}
	if list.First().Prev() != nil {
		t.Fatal("First item is not right")
	}
	if list.First().Next().Value() != 10 {
		t.Fatal("First item is not right")
	}

	list.Last().Remove()
	list.Last().Remove()
	list.Last().Remove()
	list.Last().Remove()

	if list.Len() != 1 {
		t.Fatal("List length is not right")
	}

	if list.First() != list.Last() {
		t.Fatal("First item is not equal Last")
	}

	if list.First().Value() != 9 {
		t.Fatal("First item is not right")
	}

	list.Last().Remove()

	if list.Len() != 0 {
		t.Fatal("List length is not right")
	}

	if list.Last() != nil {
		t.Fatal("First item is not equal Last")
	}

	if list.First() != nil {
		t.Fatal("First item is not equal Last")
	}

}

func TestRemoveItemLast(t *testing.T) {

	list := &List{}

	list.PushFront(10)
	list.PushFront(9)
	list.PushFront(8)
	list.PushBack(5)
	list.PushBack(4)
	list.PushBack(3)

	item := list.Last()
	item.Remove()

	if list.Len() != 5 {
		t.Fatal("List length is not right")
	}

	if list.Last().Value() != 4 {
		t.Fatal("First item is not right")
	}
	if list.First().Prev() != nil {
		t.Fatal("First item is not right")
	}
	if list.Last().Prev().Value() != 5 {
		t.Fatal("First item is not right")
	}

	list.First().Remove()
	list.First().Remove()
	list.First().Remove()
	list.First().Remove()

	if list.Len() != 1 {
		t.Fatal("List length is not right")
	}

	if list.First() != list.Last() {
		t.Fatal("First item is not equal Last")
	}

	if list.First().Value() != 4 {
		t.Fatal("First item is not right")
	}

	list.First().Remove()

	if list.Len() != 0 {
		t.Fatal("List length is not right")
	}

	if list.Last() != nil {
		t.Fatal("First item is not equal Last")
	}

	if list.First() != nil {
		t.Fatal("First item is not equal Last")
	}

}

func TestRemoveItemMiddle(t *testing.T) {

	list := &List{}

	list.PushFront(10)
	list.PushFront(9)
	list.PushFront(8)
	item := list.Last()
	list.PushBack(5)
	list.PushBack(4)
	list.PushBack(3)

	item2 := item
	item.Remove()

	if list.Len() != 5 {
		t.Fatal("List length is not right")
	}

	for { // bad
		if item2.Next() == nil {
			break
		}
		item2 = item2.Next()
		item2.Remove()
	}

	if list.Len() != 2 {
		t.Fatal("List length is not right")
	}

}

// all
func TestLinkedList(t *testing.T) {

	list := &List{}

	if list.Len() != 0 {
		t.Fatal("List is not empty")
	}

	list.PushBack(5)

	if list.First() != list.Last() {
		t.Fatal("First item is not equal Last")
	}

	list.PushBack(4)
	list.PushBack(3)

	if list.Len() != 3 {
		t.Fatal("List length is not right")
	}

	if list.First().Value() != 5 {
		t.Fatal("First item is not right")
	}

	if list.Last().Value() != 3 {
		t.Fatal("Last item is not right")
	}

	if list.First().Prev() != nil {
		t.Fatal("Item before the first item is not nil")
	}

	item := list.Last()
	if item.Prev().Value() != 4 {
		t.Fatal("Middle item is not right")
	}
	if item.Next() != nil {
		t.Fatal("Item after the last item is not nil")
	}
	if item.Prev().Next().Value() != 3 {
		t.Fatal("Item after Middle item is not right")
	}
	if item.Prev().Prev().Value() != 5 {
		t.Fatal("Item before Middle item is not right")
	}

	list.PushFront(10)
	list.PushFront(9)
	list.PushFront(8)

	if list.Len() != 6 {
		t.Fatal("List length is not right")
	}

	if list.First().Value() != 8 {
		t.Fatal("First item is not right")
	}

	item = list.First()
	item.Remove()

	if list.Len() != 5 {
		t.Fatal("List length is not right")
	}

	if list.First().Value() != 9 {
		t.Fatal("First item is not right")
	}
	if list.First().Prev() != nil {
		t.Fatal("First item is not right")
	}
	if list.First().Next().Value() != 10 {
		t.Fatal("First item is not right")
	}

	list.Last().Remove()
	list.Last().Remove()
	list.Last().Remove()
	list.Last().Remove()

	if list.Len() != 1 {
		t.Fatal("List length is not right")
	}

	if list.First() != list.Last() {
		t.Fatal("First item is not equal Last")
	}

	if list.First().Value() != 9 {
		t.Fatal("First item is not right")
	}

	list.Last().Remove()

	if list.Len() != 0 {
		t.Fatal("List length is not right")
	}

	if list.Last() != nil {
		t.Fatal("First item is not equal Last")
	}

	if list.First() != nil {
		t.Fatal("First item is not equal Last")
	}

}
