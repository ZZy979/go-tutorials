package main

import (
	"errors"
	"fmt"
)

type Node[T any] struct {
	next *Node[T]
	val  T
}

// List represents a singly-linked list that holds values of any type.
type List[T any] struct {
	head *Node[T]
	size int
}

// MakeList creates a new List
func MakeList[T any]() *List[T] {
	return &List[T]{
		head: &Node[T]{}, // virtual head node
		size: 0,
	}
}

func (l *List[T]) Front() (T, error) {
	if l.size == 0 {
		var t T
		return t, errors.New("list is empty")
	}
	return l.head.next.val, nil
}

// Add inserts val to the beginning of l
func (l *List[T]) Add(val T) {
	node := &Node[T]{l.head.next, val}
	l.head.next = node
	l.size++
}

// Remove removes the first element of l
func (l *List[T]) Remove() error {
	if l.size == 0 {
		return errors.New("list is empty")
	}
	l.head.next = l.head.next.next
	l.size--
	return nil
}

// Len returns the number of elements in l
func (l *List[T]) Len() int {
	return l.size
}

// ForEach applies f to every element of l
func (l *List[T]) ForEach(f func(val T)) {
	for cur := l.head.next; cur != nil; cur = cur.next {
		f(cur.val)
	}
}

func main() {
	list := MakeList[int]()
	for i := 1; i <= 10; i++ {
		list.Add(i)
	}
	printList(list)

	for i := 1; i <= 5; i++ {
		if err := list.Remove(); err != nil {
			fmt.Println(err)
		}
	}
	printList(list)

	for i := 1; i <= 8; i++ {
		if err := list.Remove(); err != nil {
			fmt.Println(err)
		}
	}
	printList(list)
}

func printList[T any](list *List[T]) {
	fmt.Println("list.Len():", list.Len())
	if v, err := list.Front(); err == nil {
		fmt.Println("list.Front():", v)
	} else {
		fmt.Println(err)
	}
	f := func(val T) { fmt.Printf("%v ", val) }
	list.ForEach(f)
	fmt.Println()
}
