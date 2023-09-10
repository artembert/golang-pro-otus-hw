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
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	size int
	head *ListItem
	tail *ListItem
}

func NewList() List {
	return new(list)
}

func (list *list) Front() *ListItem {
	//TODO implement me
	panic("implement me")
}

func (list *list) Back() *ListItem {
	//TODO implement me
	panic("implement me")
}

func (list *list) PushFront(value interface{}) *ListItem {
	//TODO implement me
	panic("implement me")
}

func (list *list) PushBack(value interface{}) *ListItem {
	//TODO implement me
	panic("implement me")
}

func (list *list) Remove(i *ListItem) {
	//TODO implement me
	panic("implement me")
}

func (list *list) MoveToFront(i *ListItem) {
	//TODO implement me
	panic("implement me")
}

func NewListItem(value interface{}) *ListItem {
	return &ListItem{Value: value}
}

func (list *list) Len() int {
	return list.size
}
