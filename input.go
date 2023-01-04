package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// Input represents the input state.
// Input结构体表示输入状态
type Input struct {
	msg            string
	lastBulletTime time.Time
}

func NewInput(msg string) *Input {
	input := &Input{
		msg: msg,
	}
	// fmt.Println(msg)
	return input
}

func (i *Input) Update(g *Game) {
	// 移动飞船
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		// fmt.Println("←←")
		i.msg = "left pressed"
		g.ship.x -= g.cfg.ShipSpeedFactor
		// 防止飞机出界, 最多出去一半
		if g.ship.x < -float64(g.ship.width)/2 {
			g.ship.x = -float64(g.ship.width) / 2
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		// fmt.Println("→→")
		g.ship.x += g.cfg.ShipSpeedFactor
		i.msg = "right pressed"
		// 防止飞机出界
		if g.ship.x > float64(g.cfg.ScreenWidth)-float64(g.ship.width)/2 {
			g.ship.x = float64(g.cfg.ScreenWidth) - float64(g.ship.width)/2
		}
	}
	// 发射子弹
	if ebiten.IsKeyPressed(ebiten.KeySpace) && time.Since(i.lastBulletTime).Milliseconds() > g.cfg.BulletInterval {
		// fmt.Println("--")
		i.msg = "space pressed"
		if len(g.bullets) < g.cfg.MaxBulletNum {
			bullet := NewBullet(g.cfg, g.ship)
			g.addBullet(bullet)
			i.lastBulletTime = time.Now()
		}
	}
}

func (i *Input) IsKeyStartPressed() bool {
	if ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		return true
	}
	return false
}
