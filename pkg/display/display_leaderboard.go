package display

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/kishlin/snake/v2/pkg/game"
	"time"
)

func (display *Display) DrawLeaderboard(entries []game.LeaderboardEntry) {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)

	display.writeLeaderboardTitle()
	display.writeLeaderboard(entries)
	display.writeBackText()

	rl.EndDrawing()
}

func (display *Display) writeLeaderboardTitle() {
	display.drawText(
		"Leaderboard",
		textOptions{
			center: true,
			y:      padding*2 + fontSize,
			color:  []rl.Color{rl.White},
		},
	)
}

func (display *Display) writeLeaderboard(entries []game.LeaderboardEntry) {
	yOffset := padding*3 + fontSize*2

	if len(entries) == 0 {
		display.drawText(
			"No entries yet",
			textOptions{
				center: true,
				y:      yOffset,
				color:  []rl.Color{rl.White},
			},
		)

		return
	}

	for i, entry := range entries {
		display.drawText(
			display.stringify(entry),
			textOptions{
				center: true,
				y:      yOffset + int32(i)*fontSize,
				color:  []rl.Color{rl.White},
			},
		)
	}
}

func (display *Display) writeBackText() {
	display.writeSelectable(
		"Back",
		0,
		0,
		int32(rl.GetScreenHeight())-padding-fontSize,
	)
}

func (display *Display) stringify(entry game.LeaderboardEntry) string {
	datetime := time.Unix(entry.Timestamp, 0).Format("2006-01-02")

	return fmt.Sprintf("Score: %d (%s, Speed: %d, Walls: %t)", entry.Score, datetime, entry.SpeedConfig, entry.WallsAreDeadly)
}
