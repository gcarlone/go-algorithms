package weigthedquickunion

type weigthedQU struct {
	parent              []int
	size                []int
	connectedComponents int
}

// set parent id of each object to itself (N array accesses)
// set initial tree size to 1 (the object itself)
func newWeigthedQuickUnion(n int) weigthedQU {
	c := weigthedQU{parent: make([]int, n), size: make([]int, n), connectedComponents: n}

	for i := range c.parent {
		c.parent[i] = i
		c.size[i] = 1
	}

	return c
}

// chase parent pointers until reach root (depth of i array accesses)
func (qu weigthedQU) root(i int) int {
	for qu.parent[i] != i {
		i = qu.parent[i]
	}
	return i
}

// change root of the tree with less objects to point to root of
// tree with the more objects (depth of p and q array accesses)
func (qu *weigthedQU) union(p, q int) {
	if pRoot, qRoot := qu.root(p), qu.root(q); pRoot != qRoot {
		if qu.size[pRoot] < qu.size[qRoot] {
			qu.parent[pRoot] = qRoot
			qu.size[qRoot] += qu.size[pRoot]
		} else {
			qu.parent[qRoot] = pRoot
			qu.size[pRoot] += qu.size[qRoot]
		}
		qu.connectedComponents--
	}
}

// check if p and q have same root (depth of p and q array accesses)
func (qu weigthedQU) connected(p, q int) bool {
	return qu.root(p) == qu.root(q)
}
