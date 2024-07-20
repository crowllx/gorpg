package components

import (
	"fmt"

	"github.com/solarlune/resolv"
)

type Collider resolv.Object
type HurtBox struct {
	*resolv.Object
}

func NewCollider(x, y, w, h float64, args ...string) *Collider {
	obj := Collider(*resolv.NewObject(x, y, w, h, args...))
	return &obj
}
func NewHurtBox(x, y, w, h float64) *HurtBox {
	obj := HurtBox{
		resolv.NewObject(x, y, w, h, "hurt"),
	}
	return &obj
}

// returns dx and dy after checking for collisions,
// as well as a boolean indicating if collision occured with
// a hurtbox if c has tag hitbox
func (c *Collider) Check(dx, dy float64) (float64, float64, bool) {
	obj := (*resolv.Object)(c)
	var hit bool
	if obj.HasTags("hit") {
		if col := obj.Check(dx, dy, "hurt"); col != nil {
			fmt.Printf(" I've been hit! \n %+v\n", col.Cells[0])
			hit = true
			fmt.Println(hit)
		}
	} else if obj.HasTags("hurt") {
		fmt.Println("hurt found")
		if col := obj.Check(dx, dy, "hit"); col != nil {
			fmt.Println("landed a hit")
		}
	}
	if check := obj.Check(0, dy, "solid"); check != nil {
		fmt.Println("collision")
		dy = 0
	}
	if check := obj.Check(dx, 0, "solid"); check != nil {
		fmt.Println("collision")
		dx = 0
	}
	return dx, dy, hit
}
