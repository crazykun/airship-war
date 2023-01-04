package main

import (
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Ship represents a ship.
// Ship结构体表示飞船
type Ship struct {
	GameObject
	image *ebiten.Image
}

func NewShip(screenWidth, screenHeight int) *Ship {
	img, _, err := ebitenutil.NewImageFromFile("./images/ship.png")
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()
	ship := &Ship{
		image: img,
		GameObject: GameObject{
			width:  width,
			height: height,
			x:      float64(screenWidth-width) / 2,
			y:      float64(screenHeight - height),
		},
	}

	return ship
}

func (ship *Ship) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	// 移动像素点
	op.GeoM.Translate(ship.x, ship.y)
	// 等比缩放
	// op.GeoM.Scale(1.5, 1)
	screen.DrawImage(ship.image, op)
}
