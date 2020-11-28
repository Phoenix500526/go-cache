package lru

import (
	"reflect"
	"testing"
)

type String string

func (d String) Len() int {
	return len(d)
}

func TestAdd(t *testing.T) {
	lru := NewLRU(int64(0), nil)
	lru.Add("key", String("1"))
	lru.Add("key", String("111"))
	if lru.nbytes != int64(len("key")+len("111")) {
		t.Fatal("nbytes should be 6 but ", lru.nbytes)
	}
}

func TestGet(t *testing.T) {
	lru := NewLRU(int64(0), nil)
	lru.Add("key1", String("1234"))
	if v, ok := lru.Get("key1"); !ok || string(v.(String)) != "1234" {
		t.Fatalf("cache hit failed! \"{key1,1234}\"")
	}
	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	}
}

func TestRemoveoldest(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "key3"
	v1, v2, v3 := "value1", "value2", "value3"
	cap := len(k1 + k2 + v1 + v2)
	lru := NewLRU(int64(cap), nil)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))
	if _, ok := lru.Get("key1"); ok || lru.Len() != 2 {
		t.Fatalf("Removeoldest key1 failed")
	}
}

func TestOnEvicted(t *testing.T) {
	keys := make([]string, 0)
	cb := func(key string, value Value) {
		keys = append(keys, key)
	}
	lru := NewLRU(int64(10), cb)
	lru.Add("key1", String("value1"))
	lru.Add("key2", String("value2"))
	lru.Add("key3", String("value3"))
	lru.Add("key4", String("value4"))

	expect := []string{"key1", "key2", "key3"}
	if !reflect.DeepEqual(expect, keys) {
		t.Fatalf("Execute OnEvicted failed, expect keys equal to %s", expect)
	}

}
