# Golang Object Pool

## use StdPool
```go
import "github.com/dyxushuai/objpool"

type Message struct {
	Content string
}

func (m *Message) Reset() {
	m.Content = ""
}

func NewMessage() objpool.Object {
	return new(Message)
}

p := objpool.NewStdPool(NewMessage)
obj := p.Get()
defer p.Put(obj)
```

## use FixedPool
```go
import "github.com/dyxushuai/objpool"

type Message struct {
	Content string
}

func (m *Message) Reset() {
	m.Content = ""
}

func NewMessage() objpool.Object {
	return new(Message)
}

p := objpool.NewFixedPool(1000, NewMessage)
obj := p.Get()
defer p.Put(obj)
```