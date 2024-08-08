package scenes

import (
	"fmt"
	"gorpg/components"
	"gorpg/enemies"
	"gorpg/player"
	_ "gorpg/player"
	"gorpg/utils"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jakecoffman/cp/v2"
	"github.com/solarlune/ldtkgo"
)

func LoadProject(path string) *ldtkgo.Project {
	proj, err := ldtkgo.Open(path)
	if err != nil {
		panic(err)
	}
	return proj
}

func LoadLevel(level *ldtkgo.Level) (*Scene, *player.Player) {
	entities := level.LayerByIdentifier("Entities")
	intGrid := level.LayerByIdentifier("IntGrid")
	space := cp.NewSpace()
	var es []enemies.Enemy
	var p *player.Player

	for _, t := range intGrid.AllTiles() {
		enums := intGrid.Tileset.EnumsForTile(t.ID)
		if enums.Contains("Block") {
			x, y := float64(t.Position[0]), float64(t.Position[1])
			verts := []cp.Vector{
				{x, y},
				{x + 16, y},
				{x + 16, y + 16},
				{x, y + 16},
			}
			shape := cp.NewPolyShapeRaw(cp.NewStaticBody(), 4, verts, 0)
			filter := cp.NewShapeFilter(
				0,
				components.ENVIRONMENT_LAYER,
				components.ENEMY_LAYER|components.PLAYER_LAYER,
			)
			shape.SetFilter(filter)
			shape.SetCollisionType(components.ENVIRONMENT_TYPE)
			shape.UserData = utils.Collidable{}
			space.AddShape(shape)

			// filter := cp.NewShapeFilter(0, components.ENVIRONMENT_LAYER, components.PLAYER_LAYER|components.ENEMY_LAYER)
			// for _, w := range walls {
			// 	w.SetFilter(filter)
			// 	w.SetCollisionType(components.ENVIRONMENT_TYPE)
			// 	w.UserData = utils.Collidable{}
			// 	s.space.AddShape(w)
		}

	}
	tileMap, _, err := ebitenutil.NewImageFromFile(intGrid.Tileset.Path)

	if err != nil {
		panic(err)
	}
	// entities
	for _, e := range entities.Entities {
		switch e.Identifier {
		case "Enemy":
			slime := enemies.NewSlime(cp.Vector{float64(e.Position[0]), float64(e.Position[1])}, space)
			es = append(es, slime)
			slime.AddToSpace(space)
		case "Player":
			fmt.Printf("entity id: %v\n", e.Identifier)
			if p == nil {
				x, y := float64(e.Position[0]), float64(e.Position[1])
				p = player.New(x, y)
				p.AddSpace(space)
			}
		default:
		}
	}

	SetupCollisionHandlers(space)
	s := &Scene{
		player:  p,
		space:   space,
		tileSet: tileMap,
		tiles:   intGrid.AllTiles(),
		enemies: es,
		debug:   true,
	}
	fmt.Printf("player:%T\n%v\n", p, p)
	return s, p
}
