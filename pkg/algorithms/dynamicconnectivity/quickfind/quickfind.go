package dynamicconnectivity

type quickfind struct {
	id                  []int // connected components id
	connectedComponents int
}

// set id of each object to itself (N array accesses)
func newQuickfind(n int) quickfind {
	qf := quickfind{id: make([]int, n), connectedComponents: n}
	for i := range qf.id {
		qf.id[i] = i
	}
	return qf
}

// change all entries with id[p] to id[q] (at most 2N + 2 array accesses)
func (qf *quickfind) union(p, q int) {
	if !qf.connected(p, q) {
		qf.connectedComponents--

		pId := qf.id[p]
		for i := range qf.id {
			if qf.id[i] == pId {
				qf.id[i] = qf.id[q]
			}
		}
	}
}

// check whether p and q are in the same component (2 array accesses)
func (qf quickfind) connected(p, q int) bool {
	return qf.id[p] == qf.id[q]
}
