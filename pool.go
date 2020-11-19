package merry_go_round

type CreateFn func() interface{}

func NewPool(createFn CreateFn, size int) *pool {
	pool := pool{
		size:  size,
		chSem: make(chan interface{}, size),
	}
	for i := 0; i < size; i++ {
		pool.chSem <- createFn()
	}
	return &pool
}

type pool struct {
	chSem chan interface{}
	size  int
}

func (p *pool) Get() interface{} {
	return <-p.chSem
}

func (p *pool) Put(rs interface{}) {
	p.chSem <- rs
}
