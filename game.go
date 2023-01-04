package main

import (
	"fmt"
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// Game represents the game.
// Game结构体
type Game struct {
	mode    Mode
	i       uint8
	input   *Input
	cfg     *Config
	ship    *Ship
	bullets map[*Bullet]struct{}
	aliens  map[*Alien]struct{} // Game结构中的map用来存储外星人对象
}

var (
	titleArcadeFont font.Face
	arcadeFont      font.Face
	smallArcadeFont font.Face
)

func NewGame() *Game {
	cfg := loadConfig()
	// 窗口大小
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	// 设置窗口标题
	ebiten.SetWindowTitle(cfg.Title)

	g := &Game{
		input:   NewInput("Hello, Airship War!"),
		cfg:     cfg,
		ship:    NewShip(cfg.ScreenWidth, cfg.ScreenHeight),
		bullets: make(map[*Bullet]struct{}),
		aliens:  make(map[*Alien]struct{}),
	}
	g.init()
	return g
}

func (g *Game) init() {
	// 调用 CreateAliens 创建一组外星人
	g.CreateAliens()
	g.CreateFonts()
}

// 帧， 每个tick都会被调用。tick是引擎更新的一个时间单位，默认为1/60s
func (g *Game) Update() error {
	switch g.mode {
	case ModeTitle:
		if g.input.IsKeyPressed() {
			g.mode = ModeGame
		}
	case ModeGame:
		// 更新子弹
		for bullet := range g.bullets {
			bullet.y -= bullet.speedFactor
			if bullet.outOfScreen() {
				delete(g.bullets, bullet)
			}
		}
		// 更新外星人
		for alien := range g.aliens {
			alien.y += alien.speedFactor
		}
		g.input.Update(g)
		g.CheckCollision()
	case ModeOver:
		if g.input.IsKeyPressed() {
			g.init()
			g.mode = ModeTitle
		}
	}
	return nil
}

// 每帧（frame）调用。帧是渲染使用的一个时间单位，依赖显示器的刷新
func (g *Game) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, "Hello, World!") //在屏幕上输出
	g.i++
	if g.i < 255 {
		screen.Fill(Hex2RGB("#0dceda", g.i))
	} else {
		g.i = 0
	}
	// screen.Fill(g.cfg.BgColor)

	var titleTexts []string
	var texts []string
	switch g.mode {
	case ModeTitle:
		titleTexts = []string{"ALIEN INVASION"}
		texts = []string{"", "", "", "", "", "", "", "PRESS SPACE KEY", "", "OR LEFT MOUSE"}
	case ModeGame:
		// 绘制飞船
		g.ship.Draw(screen)
		// 绘制子弹
		for bullet := range g.bullets {
			bullet.Draw(screen)
		}

		// 绘制外星人
		for alien := range g.aliens {
			alien.Draw(screen)
		}

		// 输出帧率和tps
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Hello, Airship War!\nTPS: %0.2f\nFPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))

	case ModeOver:
		texts = []string{"", "GAME OVER!"}
	}

	for i, l := range titleTexts {
		x := (g.cfg.ScreenWidth - len(l)*g.cfg.TitleFontSize) / 2
		text.Draw(screen, l, titleArcadeFont, x, (i+4)*g.cfg.TitleFontSize, color.White)
	}
	for i, l := range texts {
		x := (g.cfg.ScreenWidth - len(l)*g.cfg.FontSize) / 2
		text.Draw(screen, l, arcadeFont, x, (i+4)*g.cfg.FontSize, color.White)
	}
}

// 该方法接收游戏窗口的尺寸作为参数，返回游戏的逻辑屏幕大小
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	//窗口分辨率
	return g.cfg.ScreenWidth, g.cfg.ScreenHeight
}

// 添加子弹
func (g *Game) addBullet(bullet *Bullet) {
	g.bullets[bullet] = struct{}{}
}

// 创建外星人
func (g *Game) CreateAliens() {
	alien := NewAlien(g.cfg)

	availableSpaceX := g.cfg.ScreenWidth - 2*alien.width
	numAliens := availableSpaceX / (2 * alien.width)

	for row := 0; row < 2; row++ {
		for i := 0; i < numAliens; i++ {
			alien = NewAlien(g.cfg)
			alien.x = float64(alien.width + 2*alien.width*i)
			alien.y = float64(alien.height*row) * 1.5
			g.addAlien(alien)
		}
	}
}

// 添加外星人
func (g *Game) addAlien(alien *Alien) {
	g.aliens[alien] = struct{}{}
}

// 检查碰撞
func (g *Game) CheckCollision() {
	for alien := range g.aliens {
		for bullet := range g.bullets {
			if CheckCollision(bullet, alien) {
				delete(g.aliens, alien)
				delete(g.bullets, bullet)
			}
		}
	}
}

// 创建字体
func (g *Game) CreateFonts() {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	titleArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.cfg.TitleFontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	arcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.cfg.FontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	smallArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(g.cfg.SmallFontSize),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func Hex2RGB(color16 string, alpha uint8) color.RGBA {
	r, _ := strconv.ParseInt(color16[:2], 16, 10)
	g, _ := strconv.ParseInt(color16[2:4], 16, 18)
	b, _ := strconv.ParseInt(color16[4:], 16, 10)
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: alpha}
}
