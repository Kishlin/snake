package game

import (
	"math/rand/v2"
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

	SnakeDirection int
	Snake          []Position

	Food Position

	Score int

	IsGameOver bool
}

func (game *Game) ChangeSnakeDirection(direction int) {
	if (direction == DirectionLeft && game.SnakeDirection != DirectionRight) ||
		(direction == DirectionRight && game.SnakeDirection != DirectionLeft) ||
		(direction == DirectionUp && game.SnakeDirection != DirectionDown) ||
		(direction == DirectionDown && game.SnakeDirection != DirectionUp) {
		game.SnakeDirection = direction
	}
}

func (game *Game) MoveSnake() {
	if game.IsGameOver {
		return
	}

	previous := game.Snake
	var newHead Position

	switch game.SnakeDirection {
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

	if newHead.X >= game.gridWidth {
		newHead.X = 0
	}
	if newHead.Y >= game.gridHeight {
		newHead.Y = 0
	}
	if newHead.X < 0 {
		newHead.X = game.gridWidth - 1
	}
	if newHead.Y < 0 {
		newHead.Y = game.gridHeight - 1
	}

	for _, pos := range game.Snake {
		if pos.X == newHead.X && pos.Y == newHead.Y {
			game.IsGameOver = true
			return
		}
	}

	if newHead == game.Food {
		game.Snake = append(
			[]Position{newHead},
			previous...,
		)

		game.Score += 1

		game.SpawnFood()
	} else {
		game.Snake = append(
			[]Position{newHead},
			previous[:len(previous)-1]...,
		)
	}
}

func (game *Game) SpawnFood() {
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

func (game *Game) NewGame() {
	game.Snake = []Position{
		{X: 10, Y: 10},
		{X: 9, Y: 10},
		{X: 8, Y: 10},
		{X: 7, Y: 10},
		{X: 6, Y: 10},
	}

	game.SnakeDirection = DirectionRight

	game.Score = 0

	game.SpawnFood()

	game.IsGameOver = false
}

func (game *Game) Init(gridWidth, gridHeight int32) {
	game.gridWidth, game.gridHeight = gridWidth, gridHeight

	game.NewGame()
}
