package objpool

import "sync"

// Object the pool can stored
type Object interface {
	// Reset the object to the initial state
	Reset()
}

// Pool implementations with these methods
type Pool interface {
	// Get a Object from Pool
	Get() Object
	// Giveback Object to Poll
	Put(Object)
}

// New a Object factory method
type New func() Object

// StdPool with golang basic libary sync.Pool
type StdPool struct {
	p *sync.Pool
}

// NewStdPool create a StdPool instance
func NewStdPool(fac New) *StdPool {
	return &StdPool{
		p: &sync.Pool{
			New: func() interface{} {
				return fac()
			},
		},
	}
}

func (p *StdPool) Get() Object {
	obj := p.p.Get().(Object)
	obj.Reset()
	return obj
}

func (p *StdPool) Put(obj Object) {
	p.p.Put(obj)
}

// FixedPool with a given capacity pool
type FixedPool struct {
	objs chan Object
	fac  New
}

// NewFixedPool create a FixedPool instance with a given capacity
func NewFixedPool(capacity int, fac New) *FixedPool {
	return &FixedPool{
		objs: make(chan Object, capacity),
		fac:  fac,
	}
}

func (p *FixedPool) Get() Object {
	var obj Object
	select {
	case obj = <-p.objs:
	default:
		obj = p.fac()
	}
	obj.Reset()
	return obj
}

func (p *FixedPool) Put(obj Object) {
	select {
	case p.objs <- obj:
	default:
	}
}
