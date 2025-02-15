package display

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (display *Display) DrawIntro() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)

	display.writeTitle()
	display.writeInstructions()
	display.writeContinueText()

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

func (display *Display) writeContinueText() {
	display.drawText(
		"< Press [Enter] to continue >",
		textOptions{
			center: true,
			y:      int32(rl.GetScreenHeight()) - padding*2 - fontSize*2,
			color:  []rl.Color{rl.Orange, rl.White},
		},
	)
}
