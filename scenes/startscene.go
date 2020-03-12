package scenes

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type StartScene struct {
	startImg *ebiten.Image
}

func (s *StartScene) Startup() {
	s.startImg, _, err = ebitenutil.NewImageFromFile("./images/start.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}
}

func (s *StartScene) Update(screen *ebiten.Image) error {
	screen.DrawImage(s.startImg, nil)
	return nil
}
