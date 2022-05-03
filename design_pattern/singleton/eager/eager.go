package eager

import "sync/atomic"

var instance = &idGenerator{}

type idGenerator struct {
	id uint64
}

func Instance() *idGenerator {
	return instance
}

// func (i *idGenerator) GetId() uint64 {
// 	i.id++
// 	return i.id
// }

func (i *idGenerator) GetId() uint64 {
	return atomic.AddUint64(&i.id, 1)
}

func (i *idGenerator) Clear() {
	atomic.StoreUint64(&i.id, 0)
}



