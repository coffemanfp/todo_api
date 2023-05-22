package main

import "fmt"

type node struct {
	data int
	next *node
	prev *node
}

type linkedList struct {
	length int
	head   *node
}

func (ll *linkedList) prepend(data int) {
	n := &node{
		data: data,
		next: ll.head,
	}

	ll.head = n

	ll.length++
}

func (ll *linkedList) append(data int) {
	n := &node{
		data: data,
	}

	if ll.head == nil {
		ll.head = n
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}

		current.next = n
	}

	ll.length++
}

func (ll *linkedList) insert(index int, data int) {
	n := &node{
		data: data,
	}

	if index == 0 {
		ll.prepend(data)
	} else {
		curr := ll.head
		for i := 0; i < ll.length; i++ {
			if i == index-1 {
				n.next = curr.next
				curr.next = n
				break
			}
			curr = curr.next
		}
	}
	ll.length++
}

func (ll *linkedList) remove(index int) {

	if index == 0 {
		ll.head = ll.head.next
	} else {

	}

	curr := ll.head
	for i := 0; i < ll.length; i++ {
		if i == index-1 {
			curr.next = curr.next.next
			break
		}
		curr = curr.next
	}
	ll.length--
}

func (ll *linkedList) display() {
	curr := ll.head
	fmt.Println("length", ll.length)
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
}
