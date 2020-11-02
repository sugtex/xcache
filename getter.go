package xcache

import "context"

type GetF func(context.Context) ([]byte, error)

type getter struct {
	key string
	f   GetF
}

func WithPrisoner(k string,f GetF)getter{
	return getter{
		key: k,
		f: f,
	}
}
