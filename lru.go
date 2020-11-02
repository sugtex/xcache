package xcache

type lru struct {
	maxCap int
	length int
	mp     map[string]*linkNode
	head   *linkNode
	tail   *linkNode
}

func newLRU(cap int) *lru {
	h, t := initHeadTail()
	return &lru{
		maxCap: cap,
		length: 0,
		mp:     make(map[string]*linkNode),
		head:   h,
		tail:   t,
	}
}

func (l *lru) len() int {
	return l.length
}

func (l *lru) add(key string, value []byte) {
	if n, ok := l.mp[key]; ok {
		n.value = value
		l.update(n)
	} else {
		newNode := &linkNode{key: key, value: value}
		if l.length < l.maxCap {
			l.headInsert(newNode)
			l.length++
		} else {
			l.replace(l.tail.pre, newNode)
		}
		l.mp[key] = newNode
	}
}

func (l *lru) get(key string) ([]byte, bool) {
	if n, ok := l.mp[key]; ok {
		l.update(n)
		return n.value, ok
	}
	return nil, false
}

func (l *lru) del(key string) {
	if n, ok := l.mp[key]; ok {
		n.exit()
		delete(l.mp, key)
	}
}

func initHeadTail() (*linkNode, *linkNode) {
	head, tail := &linkNode{}, &linkNode{}
	head.next = tail
	tail.pre = head
	return head, tail
}

func (l *lru) update(n *linkNode) {
	n.exit()
	l.headInsert(n)
}

func (l *lru) replace(dead, active *linkNode) {
	dead.exit()
	delete(l.mp, dead.key)
	l.headInsert(active)
}

func (l *lru) headInsert(n *linkNode) {
	first := l.head.next
	n.bridging(l.head, first)
	first.pre = n
	l.head.next = n
}
