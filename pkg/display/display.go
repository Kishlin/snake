package display

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/kishlin/snake/pkg/game"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Display struct {
	gridWidth, gridHeight int32

	renderer *sdl.Renderer
	logger   *log.Logger
	font     *ttf.Font
}

const (
	winTitle = "Snake"
	padding  = 25

	fontPath = "./assets/font.ttf"
	fontSize = 32
)

func (display *Display) Draw(snakeGame *game.Game) error {
	renderer := display.renderer

	display.Must(renderer.SetDrawColor(0, 0, 0, 255))
	display.Must(renderer.Clear())

	display.Must(renderer.SetDrawColor(255, 255, 255, 255))
	display.Must(renderer.DrawRect(&sdl.Rect{
		X: padding - 1,
		Y: padding - 1,
		W: display.gridWidth*20 + 2,
		H: display.gridHeight*20 + 2,
	}))

	snakeHead := snakeGame.Snake[0]
	display.Must(renderer.SetDrawColor(0, 255, 0, 255))
	display.Must(renderer.FillRect(&sdl.Rect{
		X: snakeHead.X*20 + padding,
		Y: snakeHead.Y*20 + padding,
		W: 20,
		H: 20,
	}))

	for _, snakePart := range snakeGame.Snake[1:] {
		display.Must(renderer.SetDrawColor(255, 255, 0, 255))
		display.Must(renderer.FillRect(&sdl.Rect{
			X: snakePart.X*20 + padding,
			Y: snakePart.Y*20 + padding,
			W: 20,
			H: 20,
		}))
	}

	display.Must(renderer.SetDrawColor(255, 0, 0, 255))
	display.Must(renderer.FillRect(&sdl.Rect{
		X: snakeGame.Food.X*20 + padding,
		Y: snakeGame.Food.Y*20 + padding,
		W: 20,
		H: 20,
	}))

	scoreString := fmt.Sprintf("Score: %d", snakeGame.Score)
	scoreText, err := display.font.RenderUTF8Blended(scoreString, sdl.Color{R: 255, G: 255, B: 255, A: 255})
	if err != nil {
		display.logger.Fatalf("Failed to render text: %s", err)
		return err
	}
	defer scoreText.Free()

	texture, err := renderer.CreateTextureFromSurface(scoreText)
	if err != nil {
		display.logger.Fatalf("Failed to create texture: %s", err)
		return err
	}
	defer texture.Destroy()

	scoreRect := &sdl.Rect{
		X: padding - 1,
		Y: padding*2 + 20*display.gridHeight + 2,
		W: int32(len(scoreString) * fontSize / 2),
		H: fontSize,
	}
	err = renderer.Copy(texture, nil, scoreRect)
	if err != nil {
		display.logger.Fatalf("Failed to copy texture: %s", err)
		return err
	}

	if snakeGame.IsGameOver == false {
		renderer.Present()
		return nil
	}

	gameOverString := "Game Over"
	gameOverText, err := display.font.RenderUTF8Blended(gameOverString, sdl.Color{R: 255, G: 0, B: 0, A: 255})
	if err != nil {
		display.logger.Fatalf("Failed to render text: %s", err)
		return err
	}
	defer gameOverText.Free()

	texture, err = renderer.CreateTextureFromSurface(gameOverText)
	if err != nil {
		display.logger.Fatalf("Failed to create texture: %s", err)
		return err
	}
	defer texture.Destroy()

	gameOverRect := &sdl.Rect{
		X: padding + display.gridWidth*10 - int32(len(gameOverString)*fontSize/2),
		Y: padding*2 + 20*display.gridHeight + 2,
		W: int32(len(gameOverString) * fontSize),
		H: fontSize,
	}

	err = renderer.Copy(texture, nil, gameOverRect)
	if err != nil {
		display.logger.Fatalf("Failed to copy texture: %s", err)
		return err
	}

	renderer.Present()
	return nil
}

func (display *Display) Must(err error) {
	if err != nil {
		display.logger.Fatalf("Unexpected error in unsafe func: %s", err)
	}
}

func (display *Display) Init(gridWidth, gridHeight int32) (func(), error) {
	display.logger = log.New(os.Stderr, "snake", log.LstdFlags|log.LUTC)
	display.gridWidth, display.gridHeight = gridWidth, gridHeight

	err := ttf.Init()
	if err != nil {
		display.logger.Fatalf("Failed to initialize TTF: %s", err)
		return nil, err
	}

	absFontPath, err := filepath.Abs(fontPath)
	if err != nil {
		display.logger.Fatalf("Failed to get absolute font path: %s", err)
		return nil, err
	}

	var font *ttf.Font
	font, err = ttf.OpenFont(absFontPath, fontSize)
	if err != nil {
		display.logger.Fatalf("Failed to open font: %s", err)
		return nil, err
	}
	display.font = font

	var window *sdl.Window
	window, err = sdl.CreateWindow(
		winTitle,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		gridWidth*20+2*padding,
		gridHeight*20+3*padding+fontSize,
		sdl.WINDOW_SHOWN,
	)

	if err != nil {
		display.logger.Fatalf("Failed to create window: %s", err)
		return nil, err
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		display.logger.Fatalf("Failed to create renderer: %s", err)
		return nil, err
	}
	display.renderer = renderer

	exitFunc := func() {
		defer ttf.Quit()
		defer font.Close()
		defer display.Must(window.Destroy())
		defer display.Must(renderer.Destroy())
	}

	return exitFunc, nil
}
