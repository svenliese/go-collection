package net

import (
	"container/list"
	"math"
)

type Connection struct {
	source *Neuron
	// 0.0 <= factor <= 1.0
	factor float64
}

func NewConnection(source *Neuron, factor float64) *Connection {
	return new(Connection).init(source, factor)
}

func (c *Connection) init(source *Neuron, factor float64) *Connection {
	c.source = source
	c.factor = factor
	return c
}

func (c *Connection) getValue() float64 {
	return c.factor * c.source.output
}

type Neuron struct {
	id          int
	connections *list.List
	input       float64
	threshold   float64
	output      float64
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
