package main

type Storage interface {
	Get(key string) interface{}
	Set(key string, item interface{})
	Remove(key string)
}

type storageImpl struct {
	storage map[string]interface{}
}

func GetStorage(initialCapacity int) Storage {
	return &storageImpl{
		storage: make(map[string]interface{}, initialCapacity),
	}
}

func (s *storageImpl) Get(key string) interface{} {
	return s.storage[key]
}

func (s *storageImpl) Set(key string, item interface{}) {
	s.storage[key] = item
}

func (s *storageImpl) Remove(key string) {
	delete(s.storage, key)
}
