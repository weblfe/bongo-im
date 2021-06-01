package internal

import (
	"errors"
	"sync"
)

type (
	gDi struct {
		sync.RWMutex
		entry map[string]interface{}
	}
	Constructor func() interface{}
)

var (
	_di        = &gDi{}
	ErrorNilDi = errors.New("no found in Di")
)

func GetDi() *gDi {
	return _di
}

// 安全添加
func Register(key string, entry interface{}) bool {
	var di = GetDi()
	return di.Count() < di.Set(key, entry).Count()
}

// 不安全添加
func UnSafeAdd(key string, entry interface{}) bool {
	GetDi().load(key, entry)
	return true
}

func (di *gDi) Exist(key string) bool {
	di.RLock()
	defer di.RUnlock()
	if _, ok := di.entry[key]; ok {
		return true
	}
	return false
}

func (di *gDi) load(key string, v interface{}) *gDi {
	di.Lock()
	defer di.Unlock()
	di.entry[key] = v
	return di
}

func (di *gDi) Set(key string, v interface{}) *gDi {
	if key == "" {
		return di
	}
	if di.Exist(key) {
		return di
	}
	di.Lock()
	defer di.Unlock()
	di.entry[key] = v
	return di
}

func (di *gDi) Count() int {
	di.RLock()
	defer di.RUnlock()
	return len(di.entry)
}

func (di *gDi) Get(key string) interface{} {
	di.RUnlock()
	defer di.RUnlock()
	v, ok := di.entry[key]
	if !ok {
		return ErrorNilDi
	}
	return v
}

func (di *gDi) Resolver(key string) interface{} {
	var v = di.Get(key)
	di.RUnlock()
	defer di.RUnlock()
	switch v.(type) {
	case Constructor:
		constructor := v.(Constructor)
		return constructor()
	case Newer:
		newer := v.(Newer)
		return newer.New()
	case Creator:
		creator := v.(Creator)
		return creator.Create()
	case error:
		return v
	}
	return v
}

func (di *gDi) Remove(key string) *gDi {
	di.Lock()
	defer di.Unlock()
	delete(di.entry, key)
	return di
}

// 获取
func Get(key string) interface{} {
	var di = GetDi()
	return di.Resolver(key)
}

// 清理
func Clear() int {
	var (
		n  = 0
		di = GetDi()
	)
	di.Unlock()
	defer di.Unlock()
	for key, _ := range di.entry {
		n++
		delete(di.entry, key)
	}
	return n
}

func IsError(v interface{}) bool {
	switch v.(type) {
	case error:
		return true
	case *error:
		return true
	}
	return false
}
