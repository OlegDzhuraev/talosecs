
### About
Talos ECS allows to design your game flow in Entity Component System pattern. Made specially for Go language. 
You can use it with any render or game framework, which can work with Go.

Talos is pretty simple, I'm trying to use minimum amount of abstractions and boiler plate to make it easier to understand.

### Setup guide
To install it in your go project, run this in the terminal:

```go get -u github.com/OlegDzhuraev/talosecs```

### Usage and Examples
Info for this partition will be added soon. Meanwhile, you can use documentation comments in code.

Basic Talos setup example
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

### Projects examples
Maybe them will be added in future :)

### License
MIT License
