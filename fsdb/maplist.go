package fsdb

import (
	"encoding/json"
	"errors"
	"github.com/qwenode/gogo/ff"
	"github.com/qwenode/gogo/serialize"
	"sync"
)

type MapList[T any] struct {
	data     map[string]T
	lock     *sync.RWMutex
	filename string
}

func (r *MapList[T]) GetData() map[string]T {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.data
}
func (r *MapList[T]) GetDataCopy() map[string]T {
	r.lock.RLock()
	defer r.lock.RUnlock()
	m := make(map[string]T, r.Len())
	for s, t := range r.data {
		m[s] = t
	}
	return m
}
func (r *MapList[T]) Reload() error {
	r.lock.Lock()
	defer r.lock.Unlock()
	if !ff.Exist(r.filename) {
		return errors.New("file not found")
	}
	return json.Unmarshal(ff.GetContentsByte(r.filename), &r.data)
}
func (r *MapList[T]) Add(key string, item T) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.data[key] = item
}
func (r *MapList[T]) Remove(key string) {
	r.lock.Lock()
	defer r.lock.Unlock()
	delete(r.data, key)
}
func (r *MapList[T]) Exist(key string) bool {
	r.lock.RLock()
	defer r.lock.RUnlock()
	if _, ok := r.data[key]; ok {
		return true
	}
	return false
}
func (r *MapList[T]) Get(key string) (T, bool) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	t, ok := r.data[key]
	return t, ok
}
func (r *MapList[T]) Save() error {
	r.lock.Lock()
	defer r.lock.Unlock()
	return ff.PutContents(r.filename, serialize.JsonEncodeByte(r.data))
}
func (r *MapList[T]) SaveTo(path string) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	return ff.PutContents(path, serialize.JsonEncodeByte(r.data))
}
func (r *MapList[T]) Len() int {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return len(r.data)
}
func MapListFromFile[T any](filename string) *MapList[T] {
	m := new(MapList[T])
	m.filename = filename
	m.lock = &sync.RWMutex{}
	m.data = make(map[string]T)
	m.Reload()
	return m
}
