package components

import (
	"fmt"

	"github.com/solarlune/resolv"
)

type HurtBox struct {
	*resolv.Object
	hits []*resolv.Object
}

func NewHurtBox(x, y, w, h float64) *HurtBox {
	var hits []*resolv.Object
	obj := HurtBox{
		resolv.NewObject(x, y, w, h, "hurt"),
		hits,
	}
	return &obj
}

func (hb *HurtBox) AddHits(objs ...*resolv.Object) {
	for _, o := range objs {
		hb.hits = append(hb.hits, o)
		hb.AddToIgnoreList(o)
	}
}
func (hb *HurtBox) Enable() {
	if !hb.HasTags("hurt") {
		hb.AddTags("hurt")
	}
}

func (hb *HurtBox) Disable() {
	hb.RemoveTags("hurt")
	hb.Update()
	fmt.Printf("%v", hb.Tags())
	for _, o := range hb.hits {
		hb.RemoveFromIgnoreList(o)
	}
	hb.hits = nil
}

// returns dx and dy after checking for collisions,
// as well as a boolean indicating if collision occured with
// a hurtbox if c has tag hitbox
