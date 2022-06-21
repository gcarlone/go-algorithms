package weigthedquickunion

type components struct {
	parents []int
	sizes   []int
}

// set parent id of each object to itself (N array accesses)
// set initial tree size to 1 (the object itself)
func newComponents(n int) components {
	c := components{}
	c.parents = make([]int, n)
	c.sizes = make([]int, n)

	for i := range c.parents {
		c.parents[i] = i
		c.sizes[i] = 1
	}

	return c
}

// chase parent pointers until reach root (depth of i array accesses)
func (c components) getRoot(i int) int {
	for c.parents[i] != i {
		i = c.parents[i]
	}
	return i
}

// change root of the tree with less objects to point to root of
// tree with the more objects (depth of p and q array accesses)
func (c components) union(p, q int) {
	pRoot := c.getRoot(p)
	qRoot := c.getRoot(q)

	if pRoot == qRoot {
		return
	}

	if c.sizes[pRoot] < c.sizes[qRoot] {
		c.parents[pRoot] = qRoot
		c.sizes[qRoot] += c.sizes[pRoot]
	} else {
		c.parents[qRoot] = pRoot
		c.sizes[pRoot] += c.sizes[qRoot]
	}
}

// check if p and q have same root (depth of p and q array accesses)
func (c components) connected(p, q int) bool {
	return c.getRoot(p) == c.getRoot(q)
}

func (c components) countConnectedComponents() int {
	counter := 0
	for i := 0; i < len(c.parents); i++ {
		if c.parents[i] == i {
			counter++
		}
	}
	return counter
}
