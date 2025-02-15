package display

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/kishlin/snake/pkg/game"
)

func (display *Display) DrawGame(snakeGame *game.Game) {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)

	display.drawGameArea()

	display.drawSnake(snakeGame)

	display.drawFood(snakeGame)

	display.writeScoreText(snakeGame)

	display.writeFPS()

	if snakeGame.IsGameOver {
		display.writeGameStatusText("GAME OVER", rl.Red)
	} else if snakeGame.IsPaused {
		display.writeGameStatusText("PAUSED", rl.Orange)
	}

	rl.EndDrawing()
}

func (display *Display) drawGameArea() {
	rl.DrawRectangleLines(
		padding-1,
		padding-1,
		display.gridWidth*20+2,
		display.gridHeight*20+2,
		rl.White,
	)
}

func (display *Display) drawSnake(snakeGame *game.Game) {
	snakeHead := snakeGame.Snake[0]
	rl.DrawRectangle(
		snakeHead.X*20+padding,
		snakeHead.Y*20+padding,
		20,
		20,
		rl.Green,
	)

	for _, snakePart := range snakeGame.Snake[1:] {
		rl.DrawRectangle(
			snakePart.X*20+padding,
			snakePart.Y*20+padding,
			20,
			20,
			rl.Yellow,
		)
	}
}

func (display *Display) drawFood(snakeGame *game.Game) {
	rl.DrawRectangle(
		snakeGame.Food.X*20+padding,
		snakeGame.Food.Y*20+padding,
		20,
		20,
		rl.Red,
	)
}

func (display *Display) writeScoreText(snakeGame *game.Game) {
	scoreString := fmt.Sprintf("Score: %d", snakeGame.Score)
	rl.DrawText(
		scoreString,
		padding,
		padding*2+20*display.gridHeight+2,
		fontSize,
		rl.White,
	)
}

func (display *Display) writeFPS() {
	fps := rl.GetFPS()
	fpsString := fmt.Sprintf("FPS: %d", fps)
	rl.DrawText(
		fpsString,
		int32(rl.GetScreenWidth())-rl.MeasureText(fpsString, fontSize)-padding,
		padding*2+20*display.gridHeight+2,
		fontSize,
		rl.White,
	)
}

func (display *Display) writeGameStatusText(status string, color rl.Color) {
	display.drawText(
		status,
		textOptions{
			center: true,
			y:      padding*2 + 20*display.gridHeight + 2,
			color:  []rl.Color{color},
		},
	)
}
