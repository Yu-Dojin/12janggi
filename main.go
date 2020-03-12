package main

import (
	"log"

	"github.com/Yu-Dojin/12janggi/global"
	"github.com/Yu-Dojin/12janggi/scenemanager"
	"github.com/Yu-Dojin/12janggi/scenes"
	"github.com/hajimehoshi/ebiten"
)

func main() {

	scenemanager.SetScene(&scenes.StartScene{})

	err := ebiten.Run(scenemanager.Update, global.ScreenWidth, global.ScreenHeight, 1.0, "12 janggi")

	if err != nil {
		log.Fatalf("Ebiten run error : %v", err)
	}
}
