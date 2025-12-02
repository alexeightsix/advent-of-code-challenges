package main

// ***NO AI USED***

import (
	"bufio"
	"os"
	"strconv"

	"github.com/goforj/godump"
)

type Node struct {
	digit int
	next  *Node
	prev  *Node
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (n *Node) Next() *Node {
	return n.next
}

type LinkedList struct {
	head *Node
	tail *Node
	len  int
}

func NewLinkedList() LinkedList {
	return LinkedList{}
}

func (list *LinkedList) Add(node *Node) {
	if list.len == 0 {
		list.head = node
	} else if list.len == 1 {
		list.tail = node
		list.tail.prev = list.head
		list.tail.next = list.head
		list.head.next = list.tail
		list.head.prev = list.tail
	} else if list.len > 1 {
		tmp := list.tail
		list.tail = node
		list.tail.prev = tmp
		list.tail.next = list.head
		list.head.prev = node
		tmp.next = list.tail
	}
	list.len++
}

func (list *LinkedList) FindByDigit(q int) *Node {
	last := list.head
	for range list.len {
		if last.digit == q {
			return last
		}
		last = last.Next()
	}
	return nil
}

type Pos string

const POS_R Pos = "R"
const POS_L Pos = "L"

type Position struct {
	pos   Pos
	digit int
}

func getPositions() []Position {
	file, err := os.OpenFile("./day-1.input", os.O_RDONLY, os.ModePerm)

	if err != nil {
		panic(err)
	}

	x := map[int]string{
		76: "L",
		82: "R",
	}

	buf := bufio.NewScanner(file)

	positions := []Position{}

	for buf.Scan() {
		text := buf.Text()
		pos := text[0]
		digit := text[1:]
		d, err := strconv.Atoi(digit)
		if err != nil {
			panic(err)
		}
		p := Position{
			pos:   Pos(x[int(pos)]),
			digit: d,
		}

		positions = append(positions, p)

	}

	err = file.Close()
	if err != nil {
		panic(err)
	}

	return positions
}

func main() {
	list := NewLinkedList()

	for i := range 100 {
		list.Add(&Node{
			digit: i,
		})
	}

	c := 0

	currNode := list.FindByDigit(50)

	for _, pos := range getPositions() {
		if pos.pos == POS_L {
			for range pos.digit {
				currNode = currNode.Prev()
			}
		}

		if pos.pos == POS_R {
			for range pos.digit {
				currNode = currNode.Next()
			}
		}

		if currNode.digit == 0 {
			c++
		}
	}

	godump.Dump("Day 1 Part 1:", c)

	c = 0

	for _, pos := range getPositions() {
		if pos.pos == POS_L {
			for range pos.digit {
				currNode = currNode.Prev()
				if currNode.digit == 0 {
					c++
				}
			}
		}

		if pos.pos == POS_R {
			for range pos.digit {
				currNode = currNode.Next()
				if currNode.digit == 0 {
					c++
				}
			}
		}
	}

	godump.Dump("Day 1 Part 2:", c)
}
