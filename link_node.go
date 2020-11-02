package xcache

type linkNode struct {
	key   string
	value []byte
	pre   *linkNode
	next  *linkNode
}

func (ln *linkNode) exit() {
	ln.pre.next = ln.next
	ln.next.pre = ln.pre
}

// 桥接
func (ln *linkNode) bridging(pf, nx *linkNode) {
	ln.pre = pf
	ln.next = nx
}
