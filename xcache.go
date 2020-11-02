package xcache

import (
	"context"
	"errors"
)

type xCache struct {
	opt     *option
	cache   *lru
	monitor *monitor
}

func NewXCache(cap int, of ...optionF) *xCache {
	if cap <= 0 {
		return nil
	}
	o := &option{}
	for _, f := range of {
		f(o)
	}
	if o.monitorGap <= 0 {
		o.monitorGap = defaultGap
	}
	x := &xCache{
		opt:     o,
		cache:   newLRU(cap),
		monitor: newMonitor(o.monitorGap),
	}
	if o.isOpenMonitor {
		go x.startMonitor()
	}
	return x
}

func (x *xCache) Add(key string, data []byte) {
	x.cache.add(key, data)
}

func (x *xCache) Get(ctx context.Context, key string, f GetF) ([]byte, error) {
	reply, ok := x.cache.get(key)
	if !ok {
		reply, err := f(ctx)
		if err != nil {
			return nil, err
		}
		x.Add(key, reply)
		return reply, nil
	}
	return reply, nil
}

func (x *xCache) Del(key string) {
	x.cache.del(key)
}

func (x *xCache) AddPrisoner(gs ...getter) {
	x.monitor.add(gs...)
}

func (x *xCache) RemovePrisoner(gs ...getter) {
	x.monitor.del(gs...)
}

func (x *xCache) ResetMonitor(gap int) error {
	if gap <= 0 {
		return errors.New("invalid gap")
	}
	x.monitor.resetGap(gap)
	return nil
}

func (x *xCache) startMonitor() {
	defer x.monitor.timer.Stop()
	for range x.monitor.timer.C {
		key, data := x.monitor.run(context.Background())
		if data == nil {
			x.Del(key)
		}
	}
}
