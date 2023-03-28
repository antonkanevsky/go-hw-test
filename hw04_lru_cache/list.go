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
	length int
	front  *ListItem
	back   *ListItem
}

func NewList() List {
	return new(list)
}

func (list *list) Len() int {
	return list.length
}

func (list *list) Front() *ListItem {
	return list.front
}

func (list *list) Back() *ListItem {
	return list.back
}

func (list *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}
	if list.front != nil {
		newItem.Next = list.front
		list.front.Prev = newItem
		list.front = newItem
	} else {
		list.front = newItem
		list.back = newItem
	}

	list.length++

	return newItem
}

func (list *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}
	if list.back != nil {
		newItem.Prev = list.back
		list.back.Next = newItem
		list.back = newItem
	} else {
		list.front = newItem
		list.back = newItem
	}

	list.length++

	return newItem
}

func (list *list) Remove(i *ListItem) {
	defer func() {
		list.length--
	}()

	switch {
	case list.length == 1:
		list.front = nil
		list.back = nil
	case i == list.front:
		i.Next.Prev = nil
		list.front = i.Next
	case i == list.back:
		i.Prev.Next = nil
		list.back = i.Prev
	default:
		i.Prev.Next = i.Next
	}
}

func (list *list) MoveToFront(i *ListItem) {
	switch {
	case i == list.front:
		return
	case i == list.back:
		i.Prev.Next = nil
		list.back = i.Prev
	default:
		i.Prev.Next = i.Next
	}

	i.Next = list.front
	list.front.Prev = i
	list.front = i
}
