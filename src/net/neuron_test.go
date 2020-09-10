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

	neuron2 := NewNeuron(2, 1.0)
	connection := NewConnection(neuron1, 0.5)
	neuron2.addConnection(connection)
	neuron1.setOutput(4.0)

	neuron2.calculateInput()
	neuron2.calculateOutput()
	compareDouble(t, 1.0, neuron2.output)
}
