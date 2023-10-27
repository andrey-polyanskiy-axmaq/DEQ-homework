package main

import (
	"fmt"
)

//Just double-ended queue data structure
//needed funcs: MakeClearDEQ, PopLeft, PopRight, AddLeft, AddRight, OutputDEQ.
//needed structs : DEQ, Node.

type DEQ struct {
	PntLeft      *Node
	PntRight     *Node
	NodeQuantity int
}

type Node struct {
	RightNode *Node
	LeftNode  *Node
	Value     int
}

func MakeClearDeq() *DEQ {
	Queue := DEQ{nil, nil, 0}
	return &Queue
}

func (Queue *DEQ) AddLeft(Value int) {
	if Queue.NodeQuantity > 0 {
		NewNode := Node{Queue.PntLeft, nil, Value}
		Queue.PntLeft.LeftNode = &NewNode
		Queue.PntLeft = &NewNode
		Queue.NodeQuantity += 1
	} else {
		NewNode := Node{nil, nil, Value}
		Queue.PntLeft = &NewNode
		Queue.PntRight = &NewNode
		Queue.NodeQuantity += 1
	}
}

func (Queue *DEQ) AddRight(Value int) {
	if Queue.NodeQuantity > 0 {
		NewNode := Node{nil, Queue.PntRight, Value}
		Queue.PntRight.RightNode = &NewNode
		Queue.PntRight = &NewNode
		Queue.NodeQuantity += 1
	} else {
		NewNode := Node{nil, nil, Value}
		Queue.PntLeft = &NewNode
		Queue.PntRight = &NewNode
		Queue.NodeQuantity += 1
	}
}

func (Queue *DEQ) PopLeft() {
	Queue.PntLeft = Queue.PntLeft.RightNode
	Queue.PntLeft.LeftNode = nil
	Queue.NodeQuantity -= 1
}

func (Queue *DEQ) PopRight() {
	Queue.PntRight = Queue.PntRight.LeftNode
	Queue.PntRight.RightNode = nil
	Queue.NodeQuantity -= 1
}

func OutputDEQ(Queue *DEQ) {
	OutputPnt := Queue.PntLeft
	for i := 0; i < Queue.NodeQuantity; i++ {
		fmt.Print(OutputPnt.Value, " ")
		OutputPnt = OutputPnt.RightNode
	}
}

func Как_же_легко_EZ() {
	fmt.Printf("\n Я ПРОСТОЙ ИВАН ГОРОД ТВЕРЬ 我了解编程的一切 \n")
}

func main() {
	a := MakeClearDeq()
	a.AddRight(228)
	a.AddLeft(228)
	a.AddRight(228)
	a.AddLeft(228)
	a.PopRight()
	a.PopLeft()
	OutputDEQ(a)
	Как_же_легко_EZ()
}
