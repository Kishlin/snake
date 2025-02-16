package display

import "C"
import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Display struct {
	gridWidth, gridHeight int32

	haltTextDrawn bool
}

const (
	padding = 25

	fontSize int32 = 30
)

func (display *Display) Init(gridWidth, gridHeight int32) {
	display.gridWidth, display.gridHeight, display.haltTextDrawn = gridWidth, gridHeight, false

	rl.InitWindow(gridWidth*20+2*padding, gridHeight*20+3*padding+fontSize, "Snake")

	rl.SetTargetFPS(60)
}

func (display *Display) ShouldClose() bool {
	return rl.WindowShouldClose()
}

func (display *Display) Close() {
	rl.CloseWindow()
}

func (display *Display) writeTitle() {
	display.drawText(
		"SNAKE",
		textOptions{
			center: true,
			y:      padding * 2,
			color:  []rl.Color{rl.White},
		},
	)
}

type textOptions struct {
	x, y   int32
	center bool
	color  []rl.Color
	ratio  float32
}

func (options *textOptions) validate() {
	if options.center && options.x != 0 {
		panic("center and x options cannot be used together")
	}

	if options.center == false && options.x == 0 {
		panic("center or x option must be used")
	}
}

func (display *Display) writeSelectable(text string, index int, selection int, y int32) {
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

func (display *Display) drawText(text string, options textOptions) {
	options.validate()

	if options.center {
		options.x = int32(rl.GetScreenWidth())/2 - rl.MeasureText(text, fontSize)/2
	}

	if options.ratio == 0 {
		options.ratio = 1
	}

	if len(options.color) == 0 {
		options.color = []rl.Color{rl.White}
	}

	color := options.color[0]
	if len(options.color) > 1 {
		timeInTenths := int(rl.GetTime() * 10)
		color = options.color[int(timeInTenths)/3%len(options.color)]
	}

	rl.DrawText(
		text,
		options.x,
		options.y,
		int32(float32(fontSize)*options.ratio),
		color,
	)
}
