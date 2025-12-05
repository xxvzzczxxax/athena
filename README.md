# What is Athena?
Athena is a small library for game development in the Go programming language.

# Status
The project is currently under development. At present, the [ECS](https://en.wikipedia.org/wiki/Entity_component_system) module is implemented, and work is underway to integrate 2D and 3D rendering.

# ECS 
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

# Feedback
If you have any questions or suggestions, I will be happy to discuss them with you in my [Telegram](https://t.me/xxvzzczxxax).