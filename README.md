# mecs
this is my golang Entity Component System that i make for fun enjoy it break it and please tell me if (when) you find a bug
and you invited to make forks of this ecs with more things i enjoy to see it
i try to made that similar to bevy in the way you code in it
the code will look in the style of this

```go
package main

import (
	"fmt"

	"github.com/marstone123/mecs"
)

type Position struct{
    x float32
    y float32
}
type Velocity struct{
    x float32
    y float32
}

type Name struct{string}

func updatePositionsWithVelocity(w *mecs.World){
    //i ask for all the position i can change
    positions := w.GetAllComponents("*Position")

    //run for on all the positions
    for i := range positions.Components{

        //the entity id of the correct component
        entityID := positions.GetEntityByIndex(i)

        //i chack if the entity of the correct 
        if w.HasComponent(entityID,"*Velocity"){
            //get a pointer to the velocity of this entity
            velocity_pointer := w.GetComponent(entityID,"*Velocity").(*Velocity)

            //get a pointer to the correct position as *Position
            position_pointer := positions.Components[i].(*Position)

            //add the velocity to the position
            position_pointer.x += velocity_pointer.x
            position_pointer.y += velocity_pointer.y
        }
    }
}

func sayHello(w *mecs.World){
    peoples := w.GetAllComponents("Name")
    
    for i := range peoples.Components{
        fmt.Print("Hi my name is ",peoples.Components[i].(Name).string," ")

        entityID := peoples.GetEntityByIndex(i)

        if w.HasComponent(entityID,"*Velocity"){
            fmt.Print("and my velocity is ",*(w.GetComponent(entityID,"*Velocity").(*Velocity))," ")
        }
        if w.HasComponent(entityID,"*Position"){
            fmt.Print("and im in ",*(w.GetComponent(entityID,"*Position").(*Position)))
        }

        fmt.Println()
    }
}

func main(){
	
    world := mecs.NewWorld()

    world.AddEntityNow(&Position{0,0},&Velocity{0.2,0.2},Name{"Cool guy"})

    world.AddSystem(mecs.EventUpdate,updatePositionsWithVelocity)
    world.AddSystem(mecs.EventDraw,sayHello)

    world.Start()
    for{
        world.Update()
        world.Draw()
    }
}
```
