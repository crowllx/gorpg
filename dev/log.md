## 2024-07-21

added hitbox/hurtbox detection, still needs logic added to make things happen.
The way used is by using the PreSolveFunc callback in `cp` to have hitboxes and hurtboxes
collide with eachother but always return false pre solve and so they can overlap.

An alternate approach could be to not have hurtboxes collide at all and instead 
have a function when activated will do a shape query on the space to identify all
shapes in range. I think this has a benefit of not needing to track when a hitbox is enabled/disabled
and only need to track a duration to continuously search. I think this approach will definately be more
suited for implementing `lineofsight`/`targeting`.


## 2024-07-22

today I want to start implementing area/targeting, I did an example of checking for line of sight
using `cp.Space.SegmentFirstQuery`. I think for targeting I will first implement a basic 2d area component
that will create a cp.Shape with a given area and a callback function to activate when players enter, and have an Update()
function that will run callback when ever something appropriate enters it's range. I guess it will have to take a filter argument 
as well in order to generate areas with different collision filters
```go

// possibly area could instead be a interface that requires the implementation of 
// an update method

type Struct Area {
    maxDistance float64
    shape *cp.Shape
    enabled bool
    onEnter func() // does this need to be limited by input params?
}

func CreateArea(distance float64, enabled bool, callback func()) *Area {
   return &Area{
    maxDistance: distance,
    shape: cp.NewCircle(distance....),
    enabled: enabled,
    onEnter: callback,
   }
}
func Update() {
    // space.ShapeQuery for all shapes overlapping
}

```

