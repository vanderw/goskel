package utils

import "testing"

func TestSyncMap(t *testing.T) {
	sm := NewSyncMap[int, any]()
	sm.Store(1, 1)
	sm.Store(2, 2)
	v1, _ := sm.Load(1)
	t.Log(v1)
	sm.Delete(1)
	v2, ok := sm.Load(1)
	if !ok {
		t.Log("not exists")
	}
	t.Log(v2)
}
