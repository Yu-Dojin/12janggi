package scenes

import (
	"log"

	"github.com/Yu-Dojin/12janggi/global"
	"github.com/Yu-Dojin/12janggi/scenemanager"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type GimulType int

const ( //각 칸에 어떤 상태인지 표시할 변수.
	GimulTypeNone GimulType = -1 + iota
	GimulTypeGreenWang
	GimulTypeGreenJa
	GimulTypeGreenJang
	GimulTypeGreenSang
	GimulTypeRedWang
	GimulTypeRedJa
	GimulTypeRedJang
	GimulTypeRedSang
	GimulTypeMax
)

type TeamType int

const (
	TeamNone TeamType = iota
	TeamGreen
	TeamRed
)

type GameScene struct {
	bgimg       *ebiten.Image
	board       [global.BoardWidth][global.BoardHeight]GimulType
	gimulImgs   [GimulTypeMax]*ebiten.Image
	selectedImg *ebiten.Image
	selected    bool
	selectedX   int
	selectedY   int
	currentTeam TeamType
	gameover    bool
}

func GetTeamType(gimulType GimulType) TeamType {
	if gimulType == GimulTypeGreenJa ||
		gimulType == GimulTypeGreenJang ||
		gimulType == GimulTypeGreenSang ||
		gimulType == GimulTypeGreenWang {
		return TeamGreen
	}
	if gimulType == GimulTypeRedJa ||
		gimulType == GimulTypeRedJang ||
		gimulType == GimulTypeRedSang ||
		gimulType == GimulTypeRedWang {
		return TeamRed
	}
	return TeamNone
}

func (g *GameScene) Startup() {
	g.gameover = false
	g.currentTeam = TeamGreen

	var err error
	g.bgimg, _, err = ebitenutil.NewImageFromFile("./images/bgimg.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}
	g.gimulImgs[GimulTypeGreenWang], _, err = ebitenutil.NewImageFromFile("./images/green_wang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}
	g.gimulImgs[GimulTypeGreenJa], _, err = ebitenutil.NewImageFromFile("./images/green_ja.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}
	g.gimulImgs[GimulTypeGreenJang], _, err = ebitenutil.NewImageFromFile("./images/green_jang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}
	g.gimulImgs[GimulTypeGreenSang], _, err = ebitenutil.NewImageFromFile("./images/green_sang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}
	g.gimulImgs[GimulTypeRedWang], _, err = ebitenutil.NewImageFromFile("./images/red_wang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}
	g.gimulImgs[GimulTypeRedJa], _, err = ebitenutil.NewImageFromFile("./images/red_ja.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}
	g.gimulImgs[GimulTypeRedJang], _, err = ebitenutil.NewImageFromFile("./images/red_jang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}
	g.gimulImgs[GimulTypeRedSang], _, err = ebitenutil.NewImageFromFile("./images/red_sang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}

	g.selectedImg, _, err = ebitenutil.NewImageFromFile("./images/selected.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}

	//Initialize board
	for i := 0; i < global.BoardWidth; i++ {
		for j := 0; j < global.BoardHeight; j++ {
			g.board[i][j] = GimulTypeNone
		}
	}

	g.board[0][0] = GimulTypeGreenSang
	g.board[0][1] = GimulTypeGreenWang
	g.board[0][2] = GimulTypeGreenJang
	g.board[1][1] = GimulTypeGreenJa

	g.board[3][0] = GimulTypeRedSang
	g.board[3][1] = GimulTypeRedWang
	g.board[3][2] = GimulTypeRedJang
	g.board[2][1] = GimulTypeRedJa
}

func (g *GameScene) Update(screen *ebiten.Image) error {

	screen.DrawImage(g.bgimg, nil)
	if g.gameover {
		return nil
	}

	// Input handling
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		i, j := x/global.GridWidth, y/global.GridHeight

		if g.selected {
			if i == g.selectedX && j == g.selectedY {
				g.selected = false
			} else {
				//Move
				g.move(g.selectedX, g.selectedY, i, j)
			}
		} else {
			if g.board[i][j] != GimulTypeNone && g.currentTeam == GetTeamType(g.board[i][j]) {
				g.selected = true
				g.selectedX, g.selectedY = i, j
			}

		}
	}

	// Draw gimuls
	for i := 0; i < global.BoardWidth; i++ {
		for j := 0; j < global.BoardHeight; j++ {

			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(global.GimulStartX+global.GridWidth*i), float64(global.GimulStartY+global.GridHeight*j))

			switch g.board[i][j] {
			case GimulTypeGreenWang:
				screen.DrawImage(g.gimulImgs[GimulTypeGreenWang], opts)
			case GimulTypeGreenJa:
				screen.DrawImage(g.gimulImgs[GimulTypeGreenJa], opts)
			case GimulTypeGreenJang:
				screen.DrawImage(g.gimulImgs[GimulTypeGreenJang], opts)
			case GimulTypeGreenSang:
				screen.DrawImage(g.gimulImgs[GimulTypeGreenSang], opts)
			case GimulTypeRedWang:
				screen.DrawImage(g.gimulImgs[GimulTypeRedWang], opts)
			case GimulTypeRedJa:
				screen.DrawImage(g.gimulImgs[GimulTypeRedJa], opts)
			case GimulTypeRedJang:
				screen.DrawImage(g.gimulImgs[GimulTypeRedJang], opts)
			case GimulTypeRedSang:
				screen.DrawImage(g.gimulImgs[GimulTypeRedSang], opts)
			}
		}
	}

	if g.selected {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(global.GimulStartX+global.GridWidth*g.selectedX), float64(global.GimulStartY+global.GridHeight*g.selectedY))
		screen.DrawImage(g.selectedImg, opts)
	}
	return nil
}

func (g *GameScene) move(prevX, prevY, tarX, tarY int) {
	if g.isMovable(prevX, prevY, tarX, tarY) {
		g.OnDie(g.board[tarX][tarY])
		g.board[prevX][prevY], g.board[tarX][tarY] = GimulTypeNone, g.board[prevX][prevY]
		g.selected = false
		if g.currentTeam == TeamGreen {
			g.currentTeam = TeamRed
		} else {
			g.currentTeam = TeamGreen
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (g *GameScene) isMovable(prevX, prevY, tarX, tarY int) bool {
	if GetTeamType(g.board[prevX][prevY]) == GetTeamType(g.board[tarX][tarY]) {
		return false
	}

	if tarX < 0 || tarY < 0 {
		return false
	}
	if tarX >= global.BoardWidth || tarY >= global.BoardHeight {
		return false
	}

	switch g.board[prevX][prevY] {
	case GimulTypeGreenJa:
		return prevX+1 == tarX && prevY == tarY
	case GimulTypeRedJa:
		return prevX-1 == tarX && prevY == tarY
	case GimulTypeGreenSang, GimulTypeRedSang:
		return abs(prevX-tarX) == 1 && abs(prevY-tarY) == 1
	case GimulTypeGreenJang, GimulTypeRedJang:
		return abs(prevX-tarX)+abs(prevY-tarY) == 1
	case GimulTypeGreenWang, GimulTypeRedWang:
		return abs(prevX-tarX) <= 1 && abs(prevY-tarY) <= 1
	}
	return false
}

// OnDie calls when gimul is died
func (g *GameScene) OnDie(gimulType GimulType) {
	if gimulType == GimulTypeGreenWang ||
		gimulType == GimulTypeRedWang {
		scenemanager.SetScene(&GameoverScene{})
	}
}
