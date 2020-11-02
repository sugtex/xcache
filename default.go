package xcache

import "context"

const (
	defaultCap = 1000
	defaultGap = 10
)

var defaultXCache = NewXCache(defaultCap, WithOpenMonitor(defaultGap))

func Add(key string, data []byte) {
	defaultXCache.Add(key, data)
}

func Get(ctx context.Context, key string, f GetF) ([]byte, error) {
	return defaultXCache.Get(ctx, key, f)
}

func Del(key string) {
	defaultXCache.Del(key)
}

func AddPrisoner(gs ...getter) {
	defaultXCache.AddPrisoner(gs...)
}

func RemovePrisoner(gs ...getter) {
	defaultXCache.RemovePrisoner(gs...)
}
