package utils

import "sync"

// type SyncMapKey interface {
// 	int | int32 | int64 | uint | uint32 | uint64 | float64 | string
// }

type SyncMap[K comparable, V any] struct {
	sync.RWMutex
	m map[K]V
}

func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{
		m: make(map[K]V),
	}
}

func (s *SyncMap[K, V]) Get(key K) (V, bool) {
	s.RLock()
	defer s.RUnlock()
	val, ok := s.m[key]
	return val, ok
}

func (s *SyncMap[K, V]) Put(key K, value V) {
	s.Lock()
	defer s.Unlock()
	s.m[key] = value
}

func (s *SyncMap[K, V]) Del(key K) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, key)
}
