package game

import (
	"math/rand/v2"
	"time"
)

const (
	DirectionUp = iota
	DirectionRight
	DirectionDown
	DirectionLeft
)

type Position struct {
	X, Y int32
}

type Game struct {
	gridWidth, gridHeight int32

	snakeDirection int
	nextDirection  int

	IsGameOver bool
	IsPaused   bool

	Snake []Position

	Food Position

	Score int

	Config *Config

	Leaderboard *Leaderboard
}

func (game *Game) Init(config *Config, leaderboard *Leaderboard, gridWidth, gridHeight int32) {
	game.gridWidth, game.gridHeight = gridWidth, gridHeight

	game.Leaderboard = leaderboard
	game.Config = config

	game.NewGame()
}

func (game *Game) NewGame() {
	game.Snake = []Position{
		{X: 10, Y: 10},
		{X: 9, Y: 10},
		{X: 8, Y: 10},
		{X: 7, Y: 10},
		{X: 6, Y: 10},
	}

	game.snakeDirection = DirectionRight
	game.nextDirection = DirectionRight

	game.IsGameOver = false
	game.IsPaused = false

	game.Score = 0

	game.spawnFood()
}

func (game *Game) RecordDirectionChange(direction int) {
	game.nextDirection = direction
}

func (game *Game) MoveSnake() error {
	if game.IsGameOver || game.IsPaused {
		return nil
	}

	if (game.nextDirection == DirectionLeft && game.snakeDirection != DirectionRight) ||
		(game.nextDirection == DirectionRight && game.snakeDirection != DirectionLeft) ||
		(game.nextDirection == DirectionUp && game.snakeDirection != DirectionDown) ||
		(game.nextDirection == DirectionDown && game.snakeDirection != DirectionUp) {
		game.snakeDirection = game.nextDirection
	}

	previous := game.Snake
	var newHead Position

	switch game.snakeDirection {
	case DirectionUp:
		newHead = Position{
			X: previous[0].X,
			Y: previous[0].Y - 1,
		}
	case DirectionRight:
		newHead = Position{
			X: previous[0].X + 1,
			Y: previous[0].Y,
		}
	case DirectionDown:
		newHead = Position{
			X: previous[0].X,
			Y: previous[0].Y + 1,
		}
	case DirectionLeft:
		newHead = Position{
			X: previous[0].X - 1,
			Y: previous[0].Y,
		}
	}

	if game.Config.WallsAreDeadly && (newHead.X < 0 || newHead.X >= game.gridWidth || newHead.Y < 0 || newHead.Y >= game.gridHeight) {
		return game.GameOver()
	}

	newHead.X = (newHead.X + game.gridWidth) % game.gridWidth
	newHead.Y = (newHead.Y + game.gridHeight) % game.gridHeight

	for _, pos := range game.Snake {
		if pos.X == newHead.X && pos.Y == newHead.Y {
			return game.GameOver()
		}
	}

	if newHead == game.Food {
		game.Snake = append(
			[]Position{newHead},
			previous...,
		)

		points := 1 * game.Config.Speed
		if game.Config.WallsAreDeadly {
			points *= 2
		}

		game.Score += points

		game.spawnFood()
	} else {
		game.Snake = append(
			[]Position{newHead},
			previous[:len(previous)-1]...,
		)
	}

	return nil
}

func (game *Game) TogglePause() {
	if !game.IsGameOver {
		game.IsPaused = !game.IsPaused
	}
}

func (game *Game) GameOver() error {
	game.IsGameOver = true

	return game.Leaderboard.Add(
		LeaderboardEntry{
			Score:          game.Score,
			SpeedConfig:    game.Config.Speed,
			WallsAreDeadly: game.Config.WallsAreDeadly,
			Timestamp:      time.Now().Unix(),
		},
	)
}

func (game *Game) spawnFood() {
	snakePositions := make(map[Position]bool)
	for _, pos := range game.Snake {
		snakePositions[pos] = true
	}

	var food Position
	for {
		food = Position{
			X: rand.Int32N(game.gridWidth),
			Y: rand.Int32N(game.gridHeight),
		}
		if !snakePositions[food] {
			break
		}
	}

	game.Food = food
}
