package ecs

import (
	"iter"

	"github.com/xxvzzczxxax/athena/util"
)

type QueryBuilder struct {
	world   *World
	with    []IPool
	without []IPool
}

func new_query_builder(world *World) *QueryBuilder {
	return &QueryBuilder{
		world:   world,
		with:    make([]IPool, 0),
		without: make([]IPool, 0),
	}
}

func (self *QueryBuilder) With(pools ...IPool) *QueryBuilder {
	self.with = append(self.with, pools...)
	return self
}

func (self *QueryBuilder) Without(pools ...IPool) *QueryBuilder {
	self.without = append(self.without, pools...)
	return self
}

func (self *QueryBuilder) Build() *Query {
	if len(self.with) == 0 && len(self.without) == 0 {
		return new_empty_query()
	}

	result := self.world.spawned.Copy()

	for _, pool := range self.with {
		result := util.Intersect(result, pool.BitSet())
		if result.Card() == 0 {
			return new_empty_query()
		}
	}

	for _, pool := range self.without {
		result := util.Diff(result, pool.BitSet())
		if result.Card() == 0 {
			return new_empty_query()
		}
	}

	return &Query{
		entities: result,
	}
}

type Query struct {
	entities *util.BitSet
}

func new_empty_query() *Query {
	return &Query{
		entities: util.NewBitSet(),
	}
}

func (self *Query) Has(entity int) bool {
	return self.entities.In(entity)
}

func (self *Query) Count() int {
	return self.entities.Card()
}

func (self *Query) Iter() iter.Seq[int] {
	return self.entities.Iter()
}

func (self *Query) BitSet() *util.BitSet {
	return self.entities
}
