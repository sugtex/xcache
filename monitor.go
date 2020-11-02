package xcache

import (
	"context"
	"sync"
	"time"
)

type monitor struct {
	lc    sync.Mutex
	timer *time.Timer
	index int // 当前索引位置
	// 对应关系
	keys     []string
	getFList []GetF
}

func newMonitor(g int) *monitor {
	return &monitor{
		timer:    time.NewTimer(time.Duration(g) * time.Second),
		index:    0,
		keys:     make([]string, 0),
		getFList: make([]GetF, 0),
	}
}

func (m *monitor) add(gs ...getter) {
	m.lc.Lock()
	defer m.lc.Unlock()
	mp := m.only(gs...)
	for k, v := range mp {
		m.keys = append(m.keys, k)
		m.getFList = append(m.getFList, v)
	}
}

func (m *monitor) del(gs ...getter) {
	m.lc.Lock()
	defer m.lc.Unlock()
	mp := m.only(gs...)
	for i, v := range m.keys {
		if _, ok := mp[v]; ok {
			m.keys = append(m.keys[:i], m.keys[i+1:]...)
			m.getFList = append(m.getFList[:i], m.getFList[i+1:]...)
		}
	}
}

// 去重
func (m *monitor) only(gs ...getter) map[string]GetF {
	res := make(map[string]GetF)
	for _, g := range gs {
		res[g.key] = g.f
	}
	return res
}

func (m *monitor) run(ctx context.Context) (string, []byte) {
	m.lc.Lock()
	defer m.lc.Unlock()
	reply, err := m.getFList[m.index](ctx)
	if err != nil {
		return "", nil
	}
	key := m.keys[m.index]
	m.index = (m.index + 1) % len(m.keys)
	return key, reply
}

func (m *monitor) resetGap(gap int) {
	m.timer.Reset(time.Duration(gap) * time.Second)
}
