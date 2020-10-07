package ch29

import "sync"

type Node struct {
	key  int
	next *Node
}

type NodeL struct {
	key  int
	next *NodeL
	mu   sync.Mutex
}

type List struct {
	head *Node
	mu   sync.Mutex
}

type ListL struct {
	head *NodeL
	mu   sync.Mutex
}

// Insert inserts key into the linked list.
func (l *List) Insert(key int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	new := &Node{key: key, next: l.head}
	l.head = new
}

// Lookup return the node whose key matches the given key.
func (l *List) Lookup(key int) *Node {
	l.mu.Lock()
	defer l.mu.Unlock()
	current := l.head
	for current != nil {
		if current.key == key {
			return current
		}
		current = current.next
	}
	return nil
}

// Insert inserts key into the linked list.
func (l *ListL) InsertL(key int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	new := &NodeL{key: key, next: l.head}
	l.head = new
}

// Lookup return the node whose key matches the given key.
func (l *ListL) LookupL(key int) *NodeL {
	l.mu.Lock()
	current := l.head
	l.mu.Unlock()
	for current != nil {
		current.mu.Lock()
		lKey := current.key
		lNext := current.next
		current.mu.Unlock()
		if lKey == key {
			return current
		}
		current = lNext
	}
	return nil
}
