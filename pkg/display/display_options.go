package display

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/kishlin/snake/pkg/game"
)

func (display *Display) DrawConfig(config *game.Config, selection int) {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)

	display.writeTitle()
	display.writeSpeedOption(config, selection)
	display.writeWallsAreDeadlyOption(config, selection)
	display.writeStartText(selection)

	rl.EndDrawing()
}

func (display *Display) writeSpeedOption(config *game.Config, selection int) {
	speedText := fmt.Sprintf("Speed: %d/%d", config.Speed, game.MaxSpeed)

	display.writeOption(speedText, 0, selection, int32(rl.GetScreenHeight())/2-fontSize)
}

func (display *Display) writeWallsAreDeadlyOption(config *game.Config, selection int) {
	wallsAreDeadlyText := "Walls are deadly: "
	if config.WallsAreDeadly {
		wallsAreDeadlyText += "ON"
	} else {
		wallsAreDeadlyText += "OFF"
	}

	display.writeOption(wallsAreDeadlyText, 1, selection, int32(rl.GetScreenHeight())/2+fontSize)
}

func (display *Display) writeStartText(selection int) {
	display.writeOption("Press [Enter] to start", 2, selection, int32(rl.GetScreenHeight())-padding*2-fontSize*2)
}

func (display *Display) writeOption(text string, index int, selection int, y int32) {
	color := rl.White
	if index == selection {
		color = rl.Orange
		text = "< " + text + " >"
	}

	display.drawText(
		text,
		textOptions{
			center: true,
			y:      y,
			color:  []rl.Color{color},
		},
	)
}
