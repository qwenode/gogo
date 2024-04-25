package fsdb

import (
	"encoding/json"
	"github.com/qwenode/gogo/ff"
	"github.com/qwenode/gogo/serialize"
	"sync"
)

type Common[T any] struct {
	data     T
	lock     *sync.RWMutex
	filename string
}

func (r *Common[T]) GetData() T {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.data
}
func (r *Common[T]) SetData(data T) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.data = data
}

func (r *Common[T]) Save() error {
	r.lock.Lock()
	defer r.lock.Unlock()
	return ff.PutContents(r.filename, serialize.JsonEncodeByte(r.data))
}
func (r *Common[T]) Reload() error {
	r.lock.Lock()
	defer r.lock.Unlock()
	if !ff.Exist(r.filename) {
		return nil
	}
	return json.Unmarshal(ff.GetContentsByte(r.filename), &r.data)
}
func (r *Common[T]) SaveTo(path string) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	return ff.PutContents(path, serialize.JsonEncodeByte(r.data))
}
func CommonFromFile[T any](filename string) *Common[T] {
	s := new(Common[T])
	s.filename = filename
	s.lock = &sync.RWMutex{}
	s.Reload()
	return s
}
