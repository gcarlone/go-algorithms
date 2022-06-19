package quickunion

type components []int

// set id of each object to itself (N array accesses)
func newComponents(n int) components {
	c := make(components, n)
	for i := range c {
		c[i] = i
	}
	return c
}

// chase parent pointers until reach root (depth of i array accesses)
func (c components) getRoot(i int) int {
	for c[i] != i {
		i = c[i]
	}
	return i
}

// change root of p to point to root of q (depth of p and q array accesses)
func (id components) union(p, q int) {
	id[id.getRoot(p)] = id.getRoot(q)
}

// check if p and q have same root (depth of p and q array accesses)
func (c components) connected(p, q int) bool {
	return c.getRoot(p) == c.getRoot(q)
}

func (c components) countConnectedComponents() int {
	counter := 0
	for i := 0; i < len(c); i++ {
		if c[i] == i {
			counter++
		}
	}
	return counter
}
