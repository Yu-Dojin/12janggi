package main

import (
	"log"

	"github.com/Yu-Dojin/12janggi/scenemanager"
	"github.com/Yu-Dojin/12janggi/scenes"
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

const (
	ScreenWidth  = 480
	ScreenHeight = 362
	GimulStartX  = 20
	GimulStartY  = 23
	GridWidth    = 116
	GridHeight   = 116
	BoardWidth   = 4
	BoardHeight  = 3
)

var (
	board       [BoardWidth][BoardHeight]GimulType
	bgimg       *ebiten.Image
	gimulImgs   [GimulTypeMax]*ebiten.Image
	selectedImg *ebiten.Image
	selected    bool
	selectedX   int
	selectedY   int
	currentTeam TeamType = TeamGreen
	gameover    bool
)

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

func update(screen *ebiten.Image) error {

	screen.DrawImage(bgimg, nil)
	if gameover {
		return nil
	}

	// Input handling
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		i, j := x/GridWidth, y/GridHeight

		if selected {
			if i == selectedX && j == selectedY {
				selected = false
			} else {
				//Move
				move(selectedX, selectedY, i, j)
			}
		} else {
			if board[i][j] != GimulTypeNone && currentTeam == GetTeamType(board[i][j]) {
				selected = true
				selectedX, selectedY = i, j
			}

		}
	}

	// Draw gimuls
	for i := 0; i < BoardWidth; i++ {
		for j := 0; j < BoardHeight; j++ {

			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(GimulStartX+GridWidth*i), float64(GimulStartY+GridHeight*j))

			switch board[i][j] {
			case GimulTypeGreenWang:
				screen.DrawImage(gimulImgs[GimulTypeGreenWang], opts)
			case GimulTypeGreenJa:
				screen.DrawImage(gimulImgs[GimulTypeGreenJa], opts)
			case GimulTypeGreenJang:
				screen.DrawImage(gimulImgs[GimulTypeGreenJang], opts)
			case GimulTypeGreenSang:
				screen.DrawImage(gimulImgs[GimulTypeGreenSang], opts)
			case GimulTypeRedWang:
				screen.DrawImage(gimulImgs[GimulTypeRedWang], opts)
			case GimulTypeRedJa:
				screen.DrawImage(gimulImgs[GimulTypeRedJa], opts)
			case GimulTypeRedJang:
				screen.DrawImage(gimulImgs[GimulTypeRedJang], opts)
			case GimulTypeRedSang:
				screen.DrawImage(gimulImgs[GimulTypeRedSang], opts)
			}
		}
	}

	if selected {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(GimulStartX+GridWidth*selectedX), float64(GimulStartY+GridHeight*selectedY))
		screen.DrawImage(selectedImg, opts)
	}
	return nil
}

func move(prevX, prevY, tarX, tarY int) {
	if isMovable(prevX, prevY, tarX, tarY) {
		OnDie(board[tarX][tarY])
		board[prevX][prevY], board[tarX][tarY] = GimulTypeNone, board[prevX][prevY]
		selected = false
		if currentTeam == TeamGreen {
			currentTeam = TeamRed
		} else {
			currentTeam = TeamGreen
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func isMovable(prevX, prevY, tarX, tarY int) bool {
	if GetTeamType(board[prevX][prevY]) == GetTeamType(board[tarX][tarY]) {
		return false
	}

	if tarX < 0 || tarY < 0 {
		return false
	}
	if tarX >= BoardWidth || tarY >= BoardHeight {
		return false
	}

	switch board[prevX][prevY] {
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
func OnDie(gimulType GimulType) {
	if gimulType == GimulTypeGreenWang ||
		gimulType == GimulTypeRedWang {
		gameover = true
	}
}

func main() {
	var err error
	bgimg, _, err = ebitenutil.NewImageFromFile("./images/bgimg.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}
	gimulImgs[GimulTypeGreenWang], _, err = ebitenutil.NewImageFromFile("./images/green_wang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}
	gimulImgs[GimulTypeGreenJa], _, err = ebitenutil.NewImageFromFile("./images/green_ja.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}
	gimulImgs[GimulTypeGreenJang], _, err = ebitenutil.NewImageFromFile("./images/green_jang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}
	gimulImgs[GimulTypeGreenSang], _, err = ebitenutil.NewImageFromFile("./images/green_sang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}
	gimulImgs[GimulTypeRedWang], _, err = ebitenutil.NewImageFromFile("./images/red_wang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}
	gimulImgs[GimulTypeRedJa], _, err = ebitenutil.NewImageFromFile("./images/red_ja.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}
	gimulImgs[GimulTypeRedJang], _, err = ebitenutil.NewImageFromFile("./images/red_jang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}
	gimulImgs[GimulTypeRedSang], _, err = ebitenutil.NewImageFromFile("./images/red_sang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}

	selectedImg, _, err = ebitenutil.NewImageFromFile("./images/selected.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error : %v", err)
	}

	//Initialize board
	for i := 0; i < BoardWidth; i++ {
		for j := 0; j < BoardHeight; j++ {
			board[i][j] = GimulTypeNone
		}
	}

	board[0][0] = GimulTypeGreenSang
	board[0][1] = GimulTypeGreenWang
	board[0][2] = GimulTypeGreenJang
	board[1][1] = GimulTypeGreenJa

	board[3][0] = GimulTypeRedSang
	board[3][1] = GimulTypeRedWang
	board[3][2] = GimulTypeRedJang
	board[2][1] = GimulTypeRedJa

	scenemanager.SetScene(&scenes.StartScene{})

	err = ebiten.Run(scenemanager.Update, ScreenWidth, ScreenHeight, 1.0, "12 janggi")

	if err != nil {
		log.Fatalf("Ebiten run error : %v", err)
	}
}
