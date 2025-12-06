package sorter

import (
	"strings"
	"sync"
)

type StringPool struct {
	pool sync.Pool
}

func NewStringPool() *StringPool {
	return &StringPool{
		pool: sync.Pool{
			New: func() interface{} {
				return &strings.Builder{}
			},
		},
	}
}

func (p *StringPool) Get() *strings.Builder {
	return p.pool.Get().(*strings.Builder)
}

func (p *StringPool) Put(b *strings.Builder) {
	b.Reset()
	p.pool.Put(b)
}
