package mootex

import (
	"sync"
)

var (
	Mutex    = make(map[string]*sync.Mutex)
	mapMutex = &sync.Mutex{}
)

func Lock(key string) {

	if Mutex[key] == nil {
		Mutex[key] = &sync.Mutex{}
	}

	mapMutex.Lock()
	Mutex[key].Lock()
}

func Unlock(key string) {

	if Mutex[key] == nil {
		return
	}

	Mutex[key].Unlock()
	mapMutex.Unlock()
}
