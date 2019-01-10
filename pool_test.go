package objpool_test

import (
	"testing"

	"github.com/dyxushuai/objpool"
	"github.com/stretchr/testify/assert"
)

type Message struct {
	Content string
}

func (m *Message) Reset() {
	m.Content = ""
}

func NewMessage() objpool.Object {
	return new(Message)
}

func TestStdPool(t *testing.T) {
	p := objpool.NewStdPool(NewMessage)
	obj := p.Get()
	obj.(*Message).Content = "hello"
	p.Put(obj)
	assert.Equal(t, "hello", obj.(*Message).Content)
	obj1 := p.Get()
	assert.Equal(t, obj, obj1)
	assert.Equal(t, "", obj1.(*Message).Content)
}

func TestFixedPool(t *testing.T) {
	p := objpool.NewFixedPool(1, NewMessage)
	obj := p.Get()
	obj.(*Message).Content = "hello"
	p.Put(obj)
	assert.Equal(t, "hello", obj.(*Message).Content)
	obj1 := p.Get()
	assert.Equal(t, obj, obj1)
	assert.Equal(t, "", obj1.(*Message).Content)
}

func TestFixedPoolZeroCap(t *testing.T) {
	p := objpool.NewFixedPool(0, NewMessage)
	obj := p.Get()
	obj.(*Message).Content = "hello"
	p.Put(obj)
	assert.Equal(t, "hello", obj.(*Message).Content)
	obj1 := p.Get()
	assert.NotEqual(t, obj, obj1)
	assert.Equal(t, "", obj1.(*Message).Content)
}
