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
	len   int
	front *ListItem
	back  *ListItem
}

func NewList() List {
	return new(list)
}

func (l list) Len() int {
	return l.len
}

func (l list) Front() *ListItem {
	return l.front
}

func (l list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := ListItem{Value: v, Next: l.Front()}
	if newItem.Next != nil {
		newItem.Next.Prev = &newItem
	} else {
		l.back = &newItem
	}

	l.front = &newItem
	l.len++
	return l.Front()
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := ListItem{Value: v, Prev: l.Back()}
	if newItem.Prev != nil {
		newItem.Prev.Next = &newItem
	} else {
		l.front = &newItem
	}

	l.back = &newItem
	l.len++
	return l.Back()
}

func (l *list) Remove(i *ListItem) {
	if i != l.Front() {
		i.Prev.Next = i.Next
	} else {
		l.front = i.Next
	}
	if i != l.Back() {
		i.Next.Prev = i.Prev
	} else {
		l.back = i.Prev
	}
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}
