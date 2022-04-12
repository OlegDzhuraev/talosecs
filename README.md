
### About
Talos ECS allows to design your game flow in Entity Component System pattern. Made specially for Go language. 
You can use it with any render or game framework, which can work with Go.

Talos is pretty simple, I'm trying to use minimum amount of abstractions and boiler plate to make it easier to understand.

**Warning: Talos still work in progress**, so there will be a lot of api changes, it is not production-ready.

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
    mainLayer := ecs.NewLayer() // adding main game loop layer. You can have several layers, each one can group systems by same feature for example.

    mainLayer.Add(&SystemA{}) // Adding systems first, remember that order is important
    mainLayer.Add(&SystemB{})
    // mainLayer.Add(&SystemC{})
    
    ecs.AddLayer(mainLayer) // you can add layer to the ECS for autorun like this. Or, you can run it directly by Init and Update methods.
  
    // Any other initialization can be placed there
    
    ecs.Init()  // Now we're running all layers with systems initialization
  
    for { // Some update loop, break it when app should be closed
        ecs.Update() // Updating all layers with systems
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
Entity is a simple number id to connect components with one object. Making a new entity with some components:

```go
ent := talosecs.NewEntity()
ent.Add(&Move{Speed: 10})
ent.Add(&Attack{Damage: 5, Distance: 100})
```
#### Layers
Layer unions several systems into one group. You can use it to separate features. For example, one layer for gameplay loop, another - for render, third - for UI render, etc.
But you also can use one layer for all game systems. :)

Make new layer example:
```go
    mainLayer := ecs.NewLayer() 
    ecs.AddLayer(mainLayer) // not necessary step, you can also call mainLayer.Init() and mainLayer.Update() directly.
```

All layers have access to all game entities and components.

You also can disable Layer update loop (for example, if you need to stop render UI):
```go
mainLayer.Active = false
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
layer.AddSystem(&YourSystem{})
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

Full usage example:
```go
func (ms *MiningSystem) Update() {
    mines := talosecs.FilterWith[*MineBuilding]()
	
    for _, mineBuilding := range mines {
    // Proceed some logic with mineBuilding 
    }
}
```

### One-frame components
Sometimes you need a component, which is live only one frame, something like event, applied to the entity. All one-frame components removed on the end of frame.

Example how to mark your component one-frame:
```go
talosecs.AddOneFrame(entity, &YourComponent{})
// or
entity.OneFrame(&YourComponent{})
```

#### Signals
You can use signals to send a global event. It can be useful when you don't want a specific entity for adding component to it, so you just register this component as Signal.

Signal will be able in all systems, which is ordered below system register.
Example:
```go
// Signal structure, same to usual component
type BuildSignal struct {
    Position Vector2
}

// Registering a new signal, Try will return false if same signal was already registered.
talosecs.TryAddSignal(&signals.BuildSignal{Position: Vector2(X: 100, Y: 200)})

// Reading the signal:
if signal, ok := talosecs.GetSignal[*BuildSignal](); ok {
    // Do something
}
```

### Projects examples
Maybe them will be added in future :)

### License
MIT License
