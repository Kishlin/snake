package loop

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/kishlin/snake/pkg/display"
	"github.com/kishlin/snake/pkg/game"
)

const gridWidth, gridHeight = 40, 21

const (
	StepExit = iota
	StepTitleScreen
	StepConfig
	StepGame
)

type Loop struct {
	step int

	ui *display.Display
}

func (loop *Loop) Run() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	rl.SetExitKey(rl.KeyNull)

	loop.step = StepTitleScreen

	ui := display.Display{}
	ui.Init(gridWidth, gridHeight)

	snakeConfig := game.Config{}
	snakeConfig.Init()

	for !ui.ShouldClose() && loop.step != StepExit {
		rl.PollInputEvents()

		var nextStep int
		switch loop.step {
		case StepTitleScreen:
			nextStep = loop.TitleScreen(&ui)
		case StepConfig:
			nextStep = loop.ConfigScreen(&ui, &snakeConfig)
		case StepGame:
			nextStep = loop.GameScreen(&ui, &snakeConfig)
		default:
			break
		}

		loop.step = nextStep
	}

	ui.Close()
}

func (loop *Loop) TitleScreen(ui *display.Display) int {
	for {
		if ui.ShouldClose() {
			return StepExit
		} else if rl.IsKeyPressed(rl.KeyBackspace) {
			return StepExit
		} else if rl.IsKeyPressed(rl.KeyEnter) || rl.IsKeyPressed(rl.KeyKpEnter) {
			return StepConfig
		}

		ui.DrawIntro()
	}
}

func (loop *Loop) ConfigScreen(ui *display.Display, snakeConfig *game.Config) int {
	selected := 0

	for {
		if ui.ShouldClose() {
			return StepExit
		} else if rl.IsKeyPressed(rl.KeyBackspace) {
			return StepTitleScreen
		} else if rl.IsKeyPressed(rl.KeyUp) {
			selected = (selected + 3 - 1) % 3
		} else if rl.IsKeyPressed(rl.KeyDown) {
			selected = (selected + 1) % 3
		} else if rl.IsKeyPressed(rl.KeyEnter) || rl.IsKeyPressed(rl.KeyKpEnter) {
			return StepGame
		} else if rl.IsKeyPressed(rl.KeyLeft) {
			if selected == 0 {
				snakeConfig.DecreaseSpeed()
			} else if selected == 1 {
				snakeConfig.ToggleWallsAreDeadly()
			}
		} else if rl.IsKeyPressed(rl.KeyRight) {
			if selected == 0 {
				snakeConfig.IncreaseSpeed()
			} else if selected == 1 {
				snakeConfig.ToggleWallsAreDeadly()
			}
		}

		ui.DrawConfig(snakeConfig, selected)
	}
}

func (loop *Loop) GameScreen(ui *display.Display, snakeConfig *game.Config) int {
	snakeGame := game.Game{}
	snakeGame.Init(snakeConfig, gridWidth, gridHeight)

	framesCount := 0

	for {
		if ui.ShouldClose() {
			return StepExit
		} else if rl.IsKeyPressed(rl.KeyBackspace) {
			return StepConfig
		} else if rl.IsKeyPressed(rl.KeyUp) {
			snakeGame.RecordDirectionChange(game.DirectionUp)
		} else if rl.IsKeyPressed(rl.KeyDown) {
			snakeGame.RecordDirectionChange(game.DirectionDown)
		} else if rl.IsKeyPressed(rl.KeyLeft) {
			snakeGame.RecordDirectionChange(game.DirectionLeft)
		} else if rl.IsKeyPressed(rl.KeyRight) {
			snakeGame.RecordDirectionChange(game.DirectionRight)
		} else if rl.IsKeyPressed(rl.KeyEnter) || rl.IsKeyPressed(rl.KeyKpEnter) {
			if snakeGame.IsGameOver {
				snakeGame.NewGame()
			}
		} else if rl.IsKeyPressed(rl.KeySpace) {
			snakeGame.TogglePause()
		}

		framesCount++
		if framesCount > (10 - snakeConfig.Speed) {
			snakeGame.MoveSnake()
			framesCount = 0
		}

		ui.DrawGame(&snakeGame)
	}
}
