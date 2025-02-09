package loop

import (
	"github.com/kishlin/snake/pkg/display"
	"github.com/kishlin/snake/pkg/game"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
	"time"
)

const gameSpeed = 900
const gridWidth, gridHeight = 40, 21

func Run() int {
	ui := display.Display{}
	exitFunc, err := ui.Init(gridWidth, gridHeight)
	if err != nil {
		return 1
	}
	defer exitFunc()

	snakeGame := game.Game{}
	snakeGame.Init(gridWidth, gridHeight)

	rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		tickStart := sdl.GetTicks64()
		directionChange := -1

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return 0
			case *sdl.KeyboardEvent:
				keyCode := event.(*sdl.KeyboardEvent).Keysym.Sym
				if event.(*sdl.KeyboardEvent).Type == sdl.KEYUP {
					continue
				}
				switch keyCode {
				case sdl.K_UP:
					directionChange = game.DirectionUp
				case sdl.K_DOWN:
					directionChange = game.DirectionDown
				case sdl.K_LEFT:
					directionChange = game.DirectionLeft
				case sdl.K_RIGHT:
					directionChange = game.DirectionRight
				case sdl.K_KP_ENTER, sdl.K_RETURN:
					if snakeGame.IsGameOver {
						snakeGame.NewGame()
					}
				}
			}
		}

		if directionChange != -1 {
			snakeGame.ChangeSnakeDirection(directionChange)
		}

		snakeGame.MoveSnake()

		err = ui.Draw(&snakeGame)
		if err != nil {
			return 2
		}

		tickEnd := sdl.GetTicks64()

		sdl.Delay(1000 - gameSpeed - (uint32(tickEnd) - uint32(tickStart)))
	}
}
