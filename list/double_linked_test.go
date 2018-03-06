package list

import (
	"testing"
)

type integer struct {
	value int
	Item
}

func newInteger(v int) *integer {
	return &integer{
		value: v,
	}
}

func valueOf(e Element) int {
	return e.(*integer).value
}

func TestPushFront(t *testing.T) {
	count := 100
	sum := (count + 1) * count / 2

	l := new(DoubleLinkedList)

	for i := 0; i < count; i++ {
		l.PushFront(newInteger(i + 1))
	}

	if l.Front().(*integer).value != count {
		t.Fatalf("Front expect %d, but get %d", count, l.Front().(*integer).value)
	}

	if l.Length() != count {
		t.Fatalf("Length expected %d, but get %d", count, l.Length())
	}

	for it := l.Front(); it != nil; it = it.Next() {
		sum -= it.(*integer).value
	}

	if sum != 0 {
		t.Fatalf("Sum expected to be 0, but get %d", sum)
	}
}

func TestPushBack(t *testing.T) {
	count := 100
	sum := (count + 1) * count / 2

	l := new(DoubleLinkedList)

	for i := 0; i < count; i++ {
		l.PushBack(newInteger(i + 1))
	}

	if l.Front().(*integer).value != 1 {
		t.Fatalf("Front expect 1, but get %d", l.Back().(*integer).value)
	}

	if l.Length() != count {
		t.Fatalf("Length expected %d, but get %d", count, l.Length())
	}

	for it := l.Front(); it != nil; it = it.Next() {
		sum -= it.(*integer).value
	}

	if sum != 0 {
		t.Fatalf("Sum expected to be 0, but get %d", sum)
	}
}

func TestList(t *testing.T) {
	l := new(DoubleLinkedList)

	l.PushFront(newInteger(1))            // 1
	l.InsertBefore(l.head, newInteger(2)) // 2, 1
	l.InsertAfter(l.head, newInteger(3))  // 2, 3, 1
	l.InsertBefore(l.tail, newInteger(4)) // 2, 3, 4, 1

	if valueOf(l.Front()) != 2 {
		t.Fatalf("check: 2 expected but get %d", valueOf(l.Front()))
	}

	if valueOf(l.Front().Next()) != 3 {
		t.Fatalf("check: 3 expected but get %d", valueOf(l.Front().Next()))
	}

	if valueOf(l.Front().Next().Next()) != 4 {
		t.Fatalf("check: 4 expected but get %d", valueOf(l.Front().Next().Next()))
	}

	if valueOf(l.Front().Next().Next().Next()) != 1 {
		t.Fatalf("check: 1 expected but get %d", valueOf(l.Front().Next().Next().Next()))
	}

	l.Remove(l.Front().Next()) // 2, 4, 1
	if valueOf(l.Front().Next()) != 4 {
		t.Fatalf("remove : 4 expected but get %d", valueOf(l.Front().Next()))
	}

	l.Remove(l.Back().Prev()) // 2, 1
	if valueOf(l.Back().Prev()) != 2 {
		t.Fatalf("remove : 2 expected but get %d", valueOf(l.Back().Prev()))
	}

	l.Remove(l.Front()) // 1
	if valueOf(l.Front()) != 1 {
		t.Fatalf("remove : 1 expected but get %d", valueOf(l.Front()))
	}

	l.PushBack(newInteger(5)) // 1, 5
	if valueOf(l.Back()) != 5 {
		t.Fatalf("remove : 5 expected but get %d", valueOf(l.Back()))
	}

	for it := l.Front(); it != nil; it = it.Next() {
		l.Remove(it)
	}

	if l.Length() != 0 || l.Front() != nil || l.Back() != nil {
		t.Fatalf("Length should be 0, but get %d", l.Length())
	}
}
