package ecs

import "github.com/xxvzzczxxax/athena/util"

type IPool interface {
	BitSet() *util.BitSet
	Detach(int)
}

type Pool[T any] struct {
	entities   *util.BitSet
	components map[int]T
}

func NewPool[T any]() *Pool[T] {
	return &Pool[T]{
		entities:   util.NewBitSet(),
		components: make(map[int]T),
	}
}

func (self *Pool[T]) Attach(entity int, component T) {
	self.entities.Incl(entity)
	self.components[entity] = component
}

func (self *Pool[T]) Get(entity int) (component T, ok bool) {
	component, ok = self.components[entity]
	return component, ok
}

func (self *Pool[T]) Detach(entity int) {
	self.entities.Excl(entity)
	delete(self.components, entity)
}

func (self *Pool[T]) BitSet() *util.BitSet {
	return self.entities
}
