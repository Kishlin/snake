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
	display.writeHelpText()

	rl.EndDrawing()
}

func (display *Display) writeSpeedOption(config *game.Config, selection int) {
	speedText := fmt.Sprintf("Speed: %d/%d", config.Speed, game.MaxSpeed)

	display.writeSelectable(speedText, 0, selection, int32(rl.GetScreenHeight())/2-fontSize)
}

func (display *Display) writeWallsAreDeadlyOption(config *game.Config, selection int) {
	wallsAreDeadlyText := "Walls are deadly: "
	if config.WallsAreDeadly {
		wallsAreDeadlyText += "ON"
	} else {
		wallsAreDeadlyText += "OFF"
	}

	display.writeSelectable(wallsAreDeadlyText, 1, selection, int32(rl.GetScreenHeight())/2+fontSize)
}

func (display *Display) writeStartText(selection int) {
	display.writeSelectable("Press [Enter] to start", 2, selection, int32(rl.GetScreenHeight())-padding*2-fontSize*2)
}

func (display *Display) writeHelpText() {
	display.drawText(
		"Use left/right arrows to change settings",
		textOptions{
			center: true,
			y:      int32(rl.GetScreenHeight()) - padding - fontSize,
			color:  []rl.Color{rl.LightGray},
		},
	)
}
