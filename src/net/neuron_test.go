package net

import "testing"

func compareInt(t *testing.T, expected int, value int) {
	if expected != value {
		t.Error("expect", expected, "but got", value)
	}
}

func compareDouble(t *testing.T, expected float64, value float64) {
	if expected != value {
		t.Error("expect", expected, "but got", value)
	}
}

func TestNeuron(t *testing.T) {
	neuron1 := NewNeuron(1, 1.0)
	compareInt(t, 1, neuron1.id)
	compareDouble(t, 1.0, neuron1.threshold)
	compareDouble(t, 0.0, neuron1.input)
	compareDouble(t, 0.0, neuron1.output)

	neuron1.calculateInput()
	compareDouble(t, 0.0, neuron1.input)

	neuron2 := NewSimpleNeuron(2)
	neuron2.SetOutput(4.0)
	connection := NewConnection(neuron2, 0.5)

	neuron1.addConnection(connection)
	neuron1.calculateInput()
	neuron1.calculateOutput()
	compareDouble(t, 1.0, neuron1.output)
}
