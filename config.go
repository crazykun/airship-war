package main

import (
	"encoding/json"
	"image/color"
	"log"
	"os"
)

type Config struct {
	ScreenWidth       int        `json:"screenWidth"`       // 屏幕宽度
	ScreenHeight      int        `json:"screenHeight"`      //
	Title             string     `json:"title"`             // 窗口标题
	BgColor           color.RGBA `json:"bgColor"`           // 背景颜色
	ShipSpeedFactor   float64    `json:"shipSpeedFactor"`   // 飞船移动速度
	BulletWidth       int        `json:"bulletWidth"`       // 子弹宽度
	BulletHeight      int        `json:"bulletHeight"`      // 子弹高度
	BulletSpeedFactor float64    `json:"bulletSpeedFactor"` // 子弹移动速度
	BulletColor       color.RGBA `json:"bulletColor"`       // 子弹颜色
	MaxBulletNum      int        `json:"maxBulletNum"`      // 最大子弹数量
	BulletInterval    int64      `json:"bulletInterval"`    // 子弹发射间隔
}

func loadConfig() *Config {
	f, err := os.Open("./config.json")
	if err != nil {
		log.Fatalf("os.Open failed: %v\n", err)
	}

	var cfg Config
	err = json.NewDecoder(f).Decode(&cfg)
	if err != nil {
		log.Fatalf("json.Decode failed: %v\n", err)
	}

	return &cfg
}
