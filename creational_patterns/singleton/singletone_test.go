package singleton

import "testing"

func TestNewInstance(t *testing.T) {
	instance1 := NewInstance()
	instance2 := NewInstance()

	if instance1 != instance2 {
		t.Error("It should be same instance")
	}

	instance1.Increment()

	if instance1.counter != instance2.counter {
		t.Error("Counter value should be same for the both instance")
	}
}
