package loop

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/kishlin/snake/v2/pkg/display"
	"github.com/kishlin/snake/v2/pkg/game"
	"github.com/kishlin/snake/v2/pkg/storage"
)

const gridWidth, gridHeight = 40, 21

const (
	StepExit = iota
	StepTitleScreen
	StepLeaderboard
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

	var store game.Storage
	store = &storage.Storage{}
	err := store.Init()

	if err != nil {
		rl.TraceLog(rl.LogError, "Error initializing storage: %v", err)
		ui.Close()
		return
	}

	snakeConfig := game.Config{}
	snakeConfig.Init()

	leaderboard := game.Leaderboard{}
	leaderboard.Init(&store)

	for !ui.ShouldClose() && loop.step != StepExit {
		rl.PollInputEvents()

		var nextStep int
		switch loop.step {
		case StepTitleScreen:
			nextStep = loop.TitleScreen(&ui)
		case StepLeaderboard:
			nextStep = loop.LeaderboardScreen(&ui, &leaderboard)
		case StepConfig:
			nextStep = loop.ConfigScreen(&ui, &snakeConfig)
		case StepGame:
			nextStep = loop.GameScreen(&ui, &snakeConfig, &leaderboard)
		default:
			break
		}

		loop.step = nextStep
	}

	ui.Close()
}

func (loop *Loop) TitleScreen(ui *display.Display) int {
	selected := 0

	for {
		if ui.ShouldClose() {
			return StepExit
		} else if rl.IsKeyPressed(rl.KeyBackspace) {
			return StepExit
		} else if rl.IsKeyPressed(rl.KeyUp) {
			selected = (selected + 2 - 1) % 2
		} else if rl.IsKeyPressed(rl.KeyDown) {
			selected = (selected + 1) % 2
		} else if rl.IsKeyPressed(rl.KeyEnter) || rl.IsKeyPressed(rl.KeyKpEnter) {
			if selected == 0 {
				return StepConfig
			} else {
				return StepLeaderboard
			}
		}

		ui.DrawIntro(selected)
	}
}

func (loop *Loop) LeaderboardScreen(ui *display.Display, leaderboard *game.Leaderboard) int {
	top, err := leaderboard.GetTop()
	if err != nil {
		rl.TraceLog(rl.LogError, "Error getting top scores: %v", err)
		return StepTitleScreen
	}

	for {
		if ui.ShouldClose() {
			return StepExit
		} else if rl.IsKeyPressed(rl.KeyBackspace) || rl.IsKeyPressed(rl.KeyEnter) || rl.IsKeyPressed(rl.KeyKpEnter) {
			return StepTitleScreen
		}

		ui.DrawLeaderboard(top)
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

func (loop *Loop) GameScreen(ui *display.Display, snakeConfig *game.Config, leaderboard *game.Leaderboard) int {
	snakeGame := game.Game{}
	snakeGame.Init(snakeConfig, leaderboard, gridWidth, gridHeight)

	framesCount := 0

	for {
		if ui.ShouldClose() {
			loop.saveLeaderboardScoreBeforeExit(snakeGame)
			return StepExit
		} else if rl.IsKeyPressed(rl.KeyBackspace) {
			loop.saveLeaderboardScoreBeforeExit(snakeGame)
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
			err := snakeGame.MoveSnake()
			if err != nil {
				rl.TraceLog(rl.LogError, "Error moving snake: %v", err)
			}
			framesCount = 0
		}

		ui.DrawGame(&snakeGame)
	}
}

func (loop *Loop) saveLeaderboardScoreBeforeExit(snakeGame game.Game) {
	err := snakeGame.GameOver()
	if err != nil {
		rl.TraceLog(rl.LogError, "Error ending game: %v", err)
	}
}
