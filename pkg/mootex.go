package mootex

import (
	"sync"
)

var (
	Mutex    = make(map[string]*sync.Mutex)
	mapMutex = &sync.Mutex{}
)

func Lock(key string) {

	mapMutex.Lock()
	defer mapMutex.Unlock()

	if Mutex[key] == nil {
		Mutex[key] = &sync.Mutex{}
	}

	Mutex[key].Lock()
}

func Unlock(key string) {

	mapMutex.Lock()
	defer mapMutex.Unlock()

	if Mutex[key] == nil {
		return
	}

	Mutex[key].Unlock()
}
