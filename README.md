# What is Athena?
Athena is a small library for game development in the Go programming language.

# Status
The project is currently under development. At present, the [ECS](https://en.wikipedia.org/wiki/Entity_component_system) module is implemented, and work is underway to integrate 2D and 3D rendering.

# Feedback
If you have any questions or suggestions, I will be happy to discuss them with you in our [Telegram](https://t.me/athena_ecs).

# ECS 
## World
`World` is a container where all entities live.  
- `Spawn()` — creates a new entity  
- `Despawn(entity)` — removes an entity and releases its resources  
- `Has(entity)` — checks if the entity exists  
- `Count()` — returns the number of active entities  
- `Iter()` — returns an iterator over all active entities  
## Pool
`Pool` is a storage for components of a specific type.  
- `Attach(entity, component)` — attach a component to an entity  
- `Detach(entity)` — detach a component from an entity  
- `Get(entity)` — get a component by entity (with an `ok` flag)  
- `Iter()` — returns an iterator over all entities and their components  
## Query
`Query` allows filtering entities by sets of components.  
Queries can be built using `With(...)` and `Without(...)`.  
- `With(pools...)` — select entities that have the given components  
- `Without(pools...)` — exclude entities with these components  
- `Iter()` — returns an iterator over all matching entities  
- `Has(entity)` — check if an entity matches the `Query`  
- `Count()` — get the number of entities in the `Query`
## Example
```go
	type Hero   struct{ name string }
	type Undead struct{}
	type Human  struct{}

	heroes := ecs.NewPool[Hero]()
	humans := ecs.NewPool[Human]()
	undead := ecs.NewPool[Undead]()

	world := ecs.NewWorld(
		heroes,
		humans,
		undead,
	)

	paladin := world.Spawn()
	heroes.Attach(paladin, Hero{"Paladin"})
	humans.Attach(paladin, Human{})

	archmage := world.Spawn()
	heroes.Attach(archmage, Hero{"Archmage"})
	humans.Attach(archmage, Human{})

	lich := world.Spawn()
	heroes.Attach(lich, Hero{"Lich"})
	undead.Attach(lich, Undead{})

	it := world.With(heroes, undead).Without(humans).Build().Iter()
	for entity := range it {
		if hero, ok := heroes.Get(entity); ok {
			fmt.Println(hero.name)
		}
	}
```