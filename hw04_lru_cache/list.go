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
	Key   Key // needed for cache
}

type list struct {
	length int       // zero value == 0
	first  *ListItem // first item
	last   *ListItem // last item
}

func (l *list) Len() int { // add method to list
	return l.length // zero value == 0
}

func (l *list) Front() *ListItem { // add method to list
	return l.first
}

func (l *list) Back() *ListItem { // add method to list
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem { // add method to list
	data := &ListItem{} // create new instance == list of items
	data.Value = v      // pass empty interface
	if l.first != nil { // new elem become first
		data.Next = l.first // swap current and new as first
		l.first.Prev = data
		if l.last == nil { // no last item means first is last itself
			l.last = l.first
		}
	} else if l.last != nil {
		l.last.Prev = data // old elem.prev point to new elem
		data.Next = l.last // new elem.next point to old elem
	}
	l.first = data // new item is first in list
	l.length++
	return data
}

func (l *list) PushBack(v interface{}) *ListItem { // add method to list
	data := &ListItem{}
	data.Value = v
	// same logic as in PushFront
	if l.last != nil {
		data.Prev = l.last
		l.last.Next = data
		if l.first == nil {
			l.first = l.last
		}
	} else if l.first != nil {
		l.first.Next = data
		data.Prev = l.first
	}
	l.last = data
	l.length++
	return data
}

func (l *list) Remove(i *ListItem) { // add method to list
	// l should not be nil
	// i should not be nil
	if l == nil || i == nil {
		return
	}

	if i.Prev == nil { // removing first item
		l.first = i.Next
	} else { // removing middle item - next
		i.Prev.Next = i.Next
	}
	if i.Next == nil { // removing last item
		l.last = i.Prev
	} else { // removing middle item - prev
		i.Next.Prev = i.Prev
	}
	l.length--
	// set i as removed and return
	i = nil
}

func (l *list) MoveToFront(i *ListItem) { // add method to list
	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	return new(list)
}
