package scenes

import (
	"fmt"
	"gorpg/enemies"
	"gorpg/player"
	_ "gorpg/player"

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
	tileMap, _, err := ebitenutil.NewImageFromFile(intGrid.Tileset.Path)
	if err != nil {
		panic(err)
	}
	var es []enemies.Enemy
	var p *player.Player
	// entities
	space := cp.NewSpace()
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
