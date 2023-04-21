package mmap

import "sync"

type duplicateCheck[T int | string | float64 | float32] struct {
    data     map[T]T
    lock     *sync.RWMutex
    needLock bool
}

// NewDuplicateChecker 重复检查器
func NewDuplicateChecker[T string | int | float64 | float32](needLock bool) *duplicateCheck[T] {
    ob := &duplicateCheck[T]{}
    ob.data = map[T]T{}
    ob.needLock = needLock
    if ob.needLock {
        ob.lock = new(sync.RWMutex)
    }
    return ob
}

// Exist 检查数据是否已存在
func (r *duplicateCheck[T]) Exist(v T) bool {
    if r.needLock {
        r.lock.RLock()
        defer r.lock.RUnlock()
    }
    _, ok := r.data[v]
    return ok
}

// Add 添加一条数据
func (r *duplicateCheck[T]) Add(v T) {
    if r.needLock {
        r.lock.Lock()
        defer r.lock.Unlock()
    }
    r.data[v] = v
}

// GetRWMutex 获取锁
func (r *duplicateCheck[T]) GetRWMutex() *sync.RWMutex {
    return r.lock
}