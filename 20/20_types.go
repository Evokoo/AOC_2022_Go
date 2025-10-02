package day20

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
)

// ========================
// LIST NODE
// ========================
type Node struct {
	value int
	next  *Node
	prev  *Node
}

func NewNode(value int) *Node {
	return &Node{value: value}
}

// ========================
// DOUBLE LINKED LIST
// ========================
type List struct {
	head   *Node
	tail   *Node
	length int
}

func NewList() *List {
	return &List{}
}
func (l *List) Add(node *Node) {
	if l.head == nil {
		l.head = node
		l.tail = node
		node.next = l.head
		node.prev = l.tail
	} else {
		node.prev = l.tail
		node.next = l.head
		l.tail.next = node
		l.head.prev = node
		l.tail = node
	}
	l.length++
}
func (l *List) Remove(node *Node) (*Node, *Node) {
	if l.length == 0 || node == nil {
		return nil, nil
	}
	if l.length == 1 {
		l.head = nil
		l.tail = nil
		l.length--
		return node, nil
	}

	node.prev.next = node.next
	node.next.prev = node.prev

	if node == l.head {
		l.head = node.next
	}
	if node == l.tail {
		l.tail = node.prev
	}

	next := node.next

	node.next = nil
	node.prev = nil
	l.length--

	return node, next
}
func (l *List) InsertAfter(existing, added *Node) {
	if existing == nil || added == nil {
		return
	}

	added.prev = existing
	added.next = existing.next

	existing.next.prev = added
	existing.next = added

	if existing == l.tail {
		l.tail = added
	}

	l.length++
}
func (l *List) InsertBefore(existing, added *Node) {
	if existing == nil || added == nil {
		return
	}

	added.next = existing
	added.prev = existing.prev

	existing.prev.next = added
	existing.prev = added

	if existing == l.head {
		l.head = added
	}

	l.length++
}
func (l *List) Step(start *Node, steps int) *Node {
	if start == nil || l.length == 0 || steps == 0 {
		return start
	}

	current := start
	n := l.length
	steps = steps % n

	if steps > 0 {
		for i := 0; i < steps; i++ {
			current = current.next
		}
	} else if steps < 0 {
		for i := 0; i < -steps; i++ {
			current = current.prev
		}
	}

	return current
}
func (l *List) ToArray() []int {
	if l.head == nil {
		return []int{}
	}

	output := make([]int, 0, l.length)
	current := l.head

	for i := 0; i < l.length; i++ {
		output = append(output, current.value)
		current = current.next
	}

	return output
}
func GenerateList(file string, key int) (*List, Reference) {
	data := tools.ReadFile(file)
	list := NewList()
	ref := make(Reference, 0, len(data))

	for line := range strings.SplitSeq(data, "\n") {
		value, _ := strconv.Atoi(line)
		node := NewNode(value * key)
		ref = append(ref, node)
		list.Add(node)
	}

	return list, ref
}

// ========================
// REFERENCE ARRAY
// ========================
type Reference []*Node

// ========================
// HELPERS
// ========================
func IndexOf(arr []int, target int) int {
	for i, n := range arr {
		if n == target {
			return i
		}
	}
	return -1
}
func GetModifiers(part int) (int, int) {
	if part == 2 {
		return 811589153, 10
	}
	return 1, 1
}
func GetGrooveCoordinate(list *List) (sum int) {
	arr := list.ToArray()
	offset := IndexOf(arr, 0)

	for _, nth := range []int{1000, 2000, 3000} {
		index := nth + offset
		sum += arr[index%len(arr)]
	}
	return
}
