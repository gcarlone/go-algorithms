package dynamicconnectivity

import (
	"fmt"
	"testing"
)

type tdco struct {
	p         int
	q         int
	connected bool
}

func TestUnion(t *testing.T) {
	n := 10
	values := []tdco{
		{4, 3, false},
		{3, 8, false},
		{6, 5, false},
		{6, 5, true},
		{9, 4, false},
		{2, 1, false},
		{8, 9, true},
		{5, 0, false},
		{7, 2, false},
		{6, 1, false},
		{1, 0, true},
		{6, 7, true},
	}

	components := newComponents(n)

	fmt.Println(components)

	for _, o := range values {
		c := components.connected(o.p, o.q)
		if c != o.connected {
			t.Errorf("Expected %v - %v connected was %v before union but got %v", o.p, o.q, o.connected, c)
		}

		components.union(o.p, o.q)

		c = components.connected(o.p, o.q)
		if c != true {
			t.Errorf("Expected %v - %v was connected after union but got %v", o.p, o.q, o.connected)
		}

		fmt.Println(components)
	}

	connectedComponents := components.countConnectedComponents()
	if connectedComponents != 2 {
		t.Errorf("Expected %v connected compoenents but got %v", 2, connectedComponents)
	}
}
