package display

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (display *Display) DrawIntro(selection int) {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)

	display.writeTitle()
	display.writeInstructions()
	display.writeContinueText(selection)
	display.writeLeaderboardText(selection)

	rl.EndDrawing()
}

func (display *Display) writeInstructions() {
	instructions := []string{
		"Use the arrow keys to move",
		"Press [Space] to pause",
		"Press [Enter] to restart",
		"Press [Backspace] to go back",
		"Press [Escape] to go exit",
	}

	xPos := int32(rl.GetScreenWidth())/2 - rl.MeasureText(instructions[3], fontSize)/2
	yOffset := int32(rl.GetScreenHeight())/2 - fontSize*2 - fontSize/2

	for i, instruction := range instructions {
		rl.DrawText(instruction, xPos, yOffset+int32(i)*fontSize, fontSize, rl.White)
	}
}

func (display *Display) writeContinueText(selection int) {
	display.writeSelectable(
		"Continue",
		0,
		selection,
		int32(rl.GetScreenHeight())-padding*2-fontSize*2,
	)
}

func (display *Display) writeLeaderboardText(selection int) {
	display.writeSelectable(
		"Leaderboard",
		1,
		selection,
		int32(rl.GetScreenHeight())-padding-fontSize,
	)
}
