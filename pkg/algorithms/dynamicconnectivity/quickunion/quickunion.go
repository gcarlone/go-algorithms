package quickunion

type quickunion struct {
	parent              []int
	connectedComponents int
}

// set id of each object to itself (N array accesses)
func newQuickunion(n int) quickunion {
	qu := quickunion{parent: make([]int, n), connectedComponents: n}
	for i := range qu.parent {
		qu.parent[i] = i
	}
	return qu
}

// chase parent pointers until reach root (depth of i array accesses)
func (qu quickunion) root(i int) int {
	for qu.parent[i] != i {
		i = qu.parent[i]
	}
	return i
}

// change root of p to point to root of q (depth of p and q array accesses)
func (qu *quickunion) union(p, q int) {
	pRoot := qu.root(p)
	qRoot := qu.root(q)

	if qu.parent[pRoot] != qRoot {
		qu.connectedComponents--

		qu.parent[pRoot] = qRoot
	}
}

// check if p and q have same root (depth of p and q array accesses)
func (qu quickunion) connected(p, q int) bool {
	return qu.root(p) == qu.root(q)
}
