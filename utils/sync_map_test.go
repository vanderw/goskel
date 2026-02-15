package utils

import "testing"

func TestSyncMap(t *testing.T) {
	sm := NewSyncMap[int, any]()
	sm.Put(1, 1)
	sm.Put(2, 2)
	v1, _ := sm.Get(1)
	t.Log(v1)
	sm.Del(1)
	v2, ok := sm.Get(1)
	if !ok {
		t.Log("not exists")
	}
	t.Log(v2)
}
