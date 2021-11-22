package main

type node[T comparable] struct {
	data T
	prev *node[T]
	next *node[T]
}

type list[T comparable] struct {
	head, tail *node[T]
}

func (l *list[T]) isEmpty() bool {
	return l.head == nil && l.tail == nil
}

func (l *list[T]) add(data T) {
	n := &node[T] {
		data: data,
		prev: l.tail,
		next: nil,
	}
	if l.isEmpty() {
		l.head = n
		l.tail = n
	}
	l.head.prev = n
	l.head = n
}

func (l *list[T]) push(data T) {

}

func (l *list[T]) del(data T) {

}

func (l *list[T]) print() {

}
