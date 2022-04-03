
### About
Talos ECS allows to design your game flow in Entity Component System pattern. Made specially for Go language. 
You can use it with any render or game framework, which can work with Go.

Talos is pretty simple, I'm trying to use minimum amount of abstractions and boiler plate to make it easier to understand.

### Setup guide
To install it in your go project, run this in the terminal:

```go get -u github.com/OlegDzhuraev/talosecs```

### Usage and Examples

#### Basic Talos setup example
You need to run a ECS from some place in your code. There an example.
```go
package main

import ecs "github.com/OlegDzhuraev/talosecs"

func main() {
  ecs.AddSystem(&SystemA{}) // Adding systems first, remember that order is important
  ecs.AddSystem(&SystemB{})
  // ecs.AddSystem(&SystemC{})
  
  // Any other initialization can be placed there
 
  ecs.Init()  // Now we're running all systems initialization
  
  for { // Some update loop, break it when app should be closed
    ecs.Update() // Updating all systems
  }
}
```


#### Components
Component is a simple struct, you can store any data in component.
```go
type Attack struct {
	Damage     float32
	ReloadTime float32
}
```

#### Entities
Making a new entity with some components:

```go
ent := talosecs.NewEntity()
ent.Add(&Move{Speed: 10})
ent.Add(&Attack{Damage: 5, Distance: 100})
```

#### Systems
System handle all registered components, doing some game logic.
```go
type YourSystem struct {
// System can store its own data in its struct
}

func (ys *YourSystem) Init() {
// Your initialization logic
}

func (ys *YourSystem) Update() {
// Your update logic
}
```
Dont forget to add it to the game loop.
```go
ecs.AddSystem(&YourSystem{})
```

#### Filters
Filters used by systems to get all components of specific type.
You can filter from Update in actual version. To optimize it, you can make some reactive system, to prevent re-filter every frame.

Filtering by a 1 component:
```go
guns := talosecs.FilterWith[*Gun]()
```

Filtering by a 2 components:
```go
guns, reloads := talosecs.FilterWith[*Gun, *Reload]()
```

Filtering by a 2 components, but 2nd slice will not be used:
```go
guns, _ := talosecs.FilterWith[*Gun, *Reload]()
```

Filtering by a 1 component, excluding all entities, which have EnemyTag component:
```go
playerCharacters := talosecs.FilterW1Excl1[*Character, *EnemyTag]()
```

Filtering by a 2 components, excluding all entities, which have EnemyTag component:
```go
playerCharacters, playerHealths := talosecs.FilterW2Excl1[*Character, *Health, *EnemyTag]()
```

#### Signals
Work in progress.

### Projects examples
Maybe them will be added in future :)

### License
MIT License
