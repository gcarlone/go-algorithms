package dynamicconnectivity

type components []int

// set id of each object to itself (N array accesses)
func newComponents(n int) components {
	c := make(components, n)
	for i := range c {
		c[i] = i
	}
	return c
}

// change all entries with id[p] to id[q] (at most 2N + 2 array accesses)
func (id components) union(p, q int) {
	pId := id[p]
	for i := range id {
		if id[i] == pId {
			id[i] = id[q]
		}
	}
}

// check whether p and q are in the same component (2 array accesses)
func (c components) connected(p, q int) bool {
	return c[p] == c[q]
}
