package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Prev  *ListItem
	Next  *ListItem
}

type list struct {
	size int
	head *ListItem
	tail *ListItem
}

func (list *list) Front() *ListItem {
	return list.head
}

func (list *list) Back() *ListItem {
	return list.tail
}

func (list *list) PushFront(value interface{}) *ListItem {
	if list.head == nil {
		newNode := &ListItem{
			Value: value,
			Prev:  nil,
			Next:  nil,
		}
		list.tail = newNode
		list.head = newNode
		list.size++
		return newNode
	}
	newNode := &ListItem{
		Value: value,
		Prev:  nil,
		Next:  list.head,
	}
	list.head.Prev = newNode
	list.head = newNode
	list.size++
	return newNode
}

func (list *list) PushBack(value interface{}) *ListItem {
	if list.head == nil {
		newNode := &ListItem{
			Value: value,
			Prev:  nil,
			Next:  nil,
		}
		list.head = newNode
		list.tail = newNode
		list.size++
		return newNode
	}
	newNode := &ListItem{
		Value: value,
		Prev:  list.tail,
		Next:  nil,
	}
	list.tail.Next = newNode
	list.tail = newNode
	list.size++
	return newNode
}

func (list *list) Remove(itemToRemove *ListItem) {
	switch {
	case list.head == list.tail:
		if itemToRemove == list.head {
			list.head = nil
			list.tail = nil
		}
	case itemToRemove == list.head:
		list.head = list.head.Next
		list.head.Prev = nil
	case itemToRemove == list.tail:
		list.tail = list.tail.Prev
		list.tail.Next = nil
	default:
		itemToRemove.Prev.Next = itemToRemove.Next
		itemToRemove.Next.Prev = itemToRemove.Prev
	}
	list.size--
}

func (list *list) MoveToFront(itemToMove *ListItem) {
	if itemToMove == list.head {
		return
	}
	list.Remove(itemToMove)
	list.size++
	list.head.Prev = itemToMove
	itemToMove.Next = list.head
	list.head = itemToMove
}

func (list *list) Len() int {
	return list.size
}

func NewList() List {
	return new(list)
}
