package main

import (
	"errors"
	"sync"
)

var rwStore = struct {
	sync.RWMutex
	m map[string]string
}{m: make(map[string]string)}

func Put(k, v string) error {
	rwStore.Lock()
	rwStore.m[k] = v
	rwStore.Unlock()

	return nil
}

var ErrorKeyNotFound = errors.New("key not found")

func Get(k string) (string, error) {
	rwStore.RLock()
	v, ok := rwStore.m[k]
	rwStore.RUnlock()

	if !ok {
		return "", ErrorKeyNotFound
	}

	return v, nil
}

func Del(k string) error {
	rwStore.Lock()
	delete(rwStore.m, k)
	rwStore.Unlock()

	return nil
}
