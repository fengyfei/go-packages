// Package list provides a generic doulbe-linked list implementation like it in Linux kernel.
package list

// Element provides the general operations to handle a link.
type Element interface {
	Next() Element    // fetch next element
	Prev() Element    // fetch previous element
	LinkNext(Element) // link a next element
	LinkPrev(Element) // link a previous element
}

// DoubleLinkedList is a general purpose double-linked list.
type DoubleLinkedList struct {
	length int
	head   Element
	tail   Element
}

// Length returns the number of elements linked in the list.
func (l *DoubleLinkedList) Length() int {
	return l.length
}

// Reset drops all elements linked in the list.
func (l *DoubleLinkedList) Reset() {
	l.length = 0
	l.head = nil
	l.tail = nil
}

// Front returns the first element of the list.
func (l *DoubleLinkedList) Front() Element {
	return l.head
}

// Back returns the last element of the list.
func (l *DoubleLinkedList) Back() Element {
	return l.tail
}

// PushFront links a element at the front of the list.
func (l *DoubleLinkedList) PushFront(e Element) {
	e.LinkNext(l.head)
	e.LinkPrev(nil)

	if l.head != nil {
		l.head.LinkPrev(e)
	} else {
		l.tail = e
	}
	l.head = e

	l.length++
}

// PushBack links a element at the end of the list.
func (l *DoubleLinkedList) PushBack(e Element) {
	e.LinkPrev(l.tail)
	e.LinkNext(nil)

	if l.tail != nil {
		l.tail.LinkNext(e)
	} else {
		l.head = e
	}
	l.tail = e

	l.length++
}

// InsertAfter inserts a element after element p.
func (l *DoubleLinkedList) InsertAfter(p, e Element) {
	next := p.Next()
	e.LinkNext(next)
	e.LinkPrev(p)
	p.LinkNext(e)

	// insert after the last elements
	if p == l.tail {
		l.tail = e
	} else {
		next.LinkPrev(e)
	}

	l.length++
}

// InsertBefore inserts a element before element p.
func (l *DoubleLinkedList) InsertBefore(p, e Element) {
	prev := p.Prev()
	e.LinkNext(p)
	e.LinkPrev(prev)
	p.LinkPrev(e)

	if p == l.head {
		l.head = e
	} else {
		prev.LinkNext(e)
	}

	l.length++
}

// Remove a element at position e.
func (l *DoubleLinkedList) Remove(e Element) {
	next := e.Next()
	prev := e.Prev()

	if next != nil {
		next.LinkPrev(prev)
	} else {
		l.tail = prev
	}

	if prev != nil {
		prev.LinkNext(next)
	} else {
		l.head = next
	}

	l.length--
}

// Item is a struct link Linux List, contains only two fields next, prev.
type Item struct {
	next Element
	prev Element
}

// Next implements the Next method of Element interface.
func (d *Item) Next() Element {
	return d.next
}

// Prev implements the Prev method of Element interface.
func (d *Item) Prev() Element {
	return d.prev
}

// LinkNext implements the LinkNext method of Element interface.
func (d *Item) LinkNext(e Element) {
	d.next = e
}

// LinkPrev implements the LinkPrev method of Element interface.
func (d *Item) LinkPrev(e Element) {
	d.prev = e
}
