package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Bullet represents a bullet fired by the ship.
// Bullet结构体表示飞船发射的子弹
type Bullet struct {
	GameObject
	image       *ebiten.Image
	speedFactor float64
}

func NewBullet(cfg *Config, ship *Ship) *Bullet {
	rect := image.Rect(0, 0, cfg.BulletWidth, cfg.BulletHeight)
	img := ebiten.NewImageWithOptions(rect, nil)
	img.Fill(cfg.BulletColor)

	return &Bullet{
		image: img,
		GameObject: GameObject{
			width:  cfg.BulletWidth,
			height: cfg.BulletHeight,
			x:      ship.x + float64(ship.width-cfg.BulletWidth)/2,
			y:      float64(cfg.ScreenHeight - ship.height - cfg.BulletHeight),
		},
		speedFactor: cfg.BulletSpeedFactor,
	}
}

func (bullet *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(bullet.x, bullet.y)
	screen.DrawImage(bullet.image, op)
}

func (bullet *Bullet) outOfScreen() bool {
	return bullet.y < -float64(bullet.height)
}
