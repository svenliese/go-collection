package net

import (
	"container/list"
	"math"
)

type INeuron interface {
	GetId() int
	GetOutput() float64
	SetOutput(output float64)
}

//
// SimpleNeuron
//

type SimpleNeuron struct {
	id     int
	output float64
}

func NewSimpleNeuron(id int) *SimpleNeuron {
	return new(SimpleNeuron).init(id)
}

func (n *SimpleNeuron) init(id int) *SimpleNeuron {
	n.id = id
	n.output = 0.0
	return n
}

func (n *SimpleNeuron) GetId() int {
	return n.id
}

func (n *SimpleNeuron) GetOutput() float64 {
	return n.output
}

func (n *SimpleNeuron) SetOutput(output float64) {
	n.output = output
}

//
// Connection
//

type Connection struct {
	source INeuron
	// 0.0 <= factor <= 1.0
	factor float64
}

func NewConnection(source INeuron, factor float64) *Connection {
	return new(Connection).init(source, factor)
}

func (c *Connection) init(source INeuron, factor float64) *Connection {
	c.source = source
	c.factor = factor
	return c
}

func (c *Connection) getValue() float64 {
	return c.factor * c.source.GetOutput()
}

//
// Neuron
//

type Neuron struct {
	SimpleNeuron

	connections *list.List
	input       float64
	threshold   float64
}

func NewNeuron(id int, threshold float64) *Neuron {
	return new(Neuron).init(id, threshold)
}

func (n *Neuron) init(id int, threshold float64) *Neuron {
	n.id = id
	n.connections = list.New()
	n.input = 0
	n.threshold = threshold
	n.output = 0
	return n
}

func (n *Neuron) calculateInput() {
	var sum = 0.0
	for listElement := n.connections.Front(); listElement != nil; listElement = listElement.Next() {
		connection := listElement.Value.(*Connection)
		sum += connection.getValue()
	}
	n.input = sum
}

func (n *Neuron) addConnection(con *Connection) {
	n.connections.PushBack(con)
}

func (n *Neuron) setOutput(output float64) {
	n.output = output
}

func (n *Neuron) calculateOutput() {
	n.output = math.Max(0.0, n.input-n.threshold)
}
