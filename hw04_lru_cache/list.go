package hw04lrucache

import (
	"fmt"
	"reflect"
)

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
	//List // Remove me after realization.
	// Place your code here.
	length int // zero value == 0
	next   *ListItem
	prev   *ListItem
}

func (l *list) Len() int { // add method to list
	fmt.Println("list.Len")
	fmt.Println(l.length)
	return l.length // zero value == 0
}

func (l *list) Front() *ListItem { // add method to list
	fmt.Println("list.Front")
	fmt.Println(reflect.TypeOf(l.next))
	fmt.Println(l.next)
	return l.next
}

func (l *list) Back() *ListItem { // add method to list
	fmt.Println("list.Back")
	fmt.Println(l.prev)
	return l.prev
}

func (l *list) PushFront(v interface{}) *ListItem { // add method to list
	fmt.Println("list.PushFront")
	return nil
}

func (l *list) PushBack(v interface{}) *ListItem { // add method to list
	fmt.Println("list.PushBack")
	return nil
}

func (l *list) Remove(i *ListItem) { // add method to list
	fmt.Println("list.Remove")
}

func (l *list) MoveToFront(i *ListItem) { // add method to list
	fmt.Println("list.MoveToFront")
}

func (l *list) Test() *ListItem { // add method to list
	fmt.Println("list.TEST")
	return nil
}

func NewList() List {
	fmt.Println("list.NewList")
	ll := new(list)
	fmt.Println(reflect.TypeOf(ll))
	fmt.Println(ll)
	fmt.Println(ll.Len())
	fmt.Println(ll.Front())
	fmt.Println(ll.Back())
	fmt.Println(ll.Test())
	var i interface{}
	fmt.Println(ll.PushFront(i))
	return ll
}
