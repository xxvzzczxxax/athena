package ecs

import (
	"iter"

	"github.com/xxvzzczxxax/athena/util"
)

type World struct {
	spawned   *util.BitSet
	despawned []int
	pools     []IPool
}

func NewWorld(pools ...IPool) *World {
	return &World{
		spawned:   util.NewBitSet(),
		despawned: make([]int, 0),
		pools:     pools,
	}
}

func (self *World) Spawn() int {
	if len(self.despawned) > 0 {
		entity := self.despawned[len(self.despawned)-1]
		self.despawned = self.despawned[:len(self.despawned)-1]
		self.spawned.Incl(entity)
		return entity
	} else {
		entity := self.spawned.Card()
		self.spawned.Incl(entity)
		return entity
	}
}

func (self *World) Despawn(entity int) {
	self.spawned.Excl(entity)
	self.despawned = append(self.despawned, entity)
	for x := range self.pools {
		self.pools[x].Detach(entity)
	}
}

func (self *World) Has(entity int) bool {
	return self.spawned.In(entity)
}

func (self *World) Count() int {
	return self.spawned.Card()
}

func (self *World) Iter() iter.Seq[int] {
	return self.spawned.Iter()
}

func (self *World) With(pools ...IPool) *QueryBuilder {
	return new_query_builder(self).With(pools...)
}

func (self *World) Without(pools ...IPool) *QueryBuilder {
	return new_query_builder(self).Without(pools...)
}

func (self *World) BitSet() *util.BitSet {
	return self.spawned
}
