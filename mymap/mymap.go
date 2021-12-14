package mymap

import "sync"

type MyMap struct {
	m map[interface{}]interface{}
	mutex *sync.RWMutex
}

func NewMyMap(size int) *MyMap{
	return &MyMap{
		m: make(map[interface{}]interface{}, size),
		mutex : &sync.RWMutex{},
	}
}

func (mm *MyMap) Load(key interface{})(value interface{}, ok bool){
	mm.mutex.RLock()
	value, ok = mm.m[key]
	mm.mutex.RUnlock()
	return
}

func (mm *MyMap) Store(key, value interface{}){
	mm.mutex.Lock()
	mm.m[key] = value
	mm.mutex.Unlock()
}

func (mm *MyMap) Delete(key interface{}){
	mm.mutex.Lock()
	delete(mm.m, key)
	mm.mutex.Unlock()
}

// 如果 key 对应的值存在，删除这个key,并返回其对应的值，loaded返回 true
func (mm *MyMap) LoadAndDelete(key interface{})(value interface{}, loaded bool){
	mm.mutex.Lock()
	if value, loaded = mm.m[key]; loaded {
		delete(mm.m, key)
		mm.mutex.Unlock()
		return
	}
	mm.mutex.Unlock()
	return
}

// 如果 key 对应的值是存在的，则返回存在的值，loaded 返回 false
// 如果 key 不存在，则返回要写入的值，loaded 返回 true
func (mm *MyMap) LoadOrStore(key, value interface{})(actual interface{}, loaded bool){
	mm.mutex.Lock()

	if actual, loaded = mm.m[key]; loaded {
		mm.mutex.Unlock()
		return
	}
	mm.m[key] = value
	mm.mutex.Unlock()
	actual = value
	return
}






