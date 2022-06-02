package main

import (
	"log"
	"math/rand"
)

// This function is called when you register your Battlesnake on play.battlesnake.com
// See https://docs.battlesnake.com/guides/getting-started#step-4-register-your-battlesnake
// It controls your Battlesnake appearance and author permissions.
// For customization options, see https://docs.battlesnake.com/references/personalization
// TIP: If you open your Battlesnake URL in browser you should see this data.
func info() BattlesnakeInfoResponse {
	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "Jon Hammond",
		Color:      "#00c2d4",
		Head:       "all-seeing",
		Tail:       "hook",
	}
}

// This function is called everytime your Battlesnake is entered into a game.
// The provided GameState contains information about the game that's about to be played.
// It's purely for informational purposes, you don't have to make any decisions here.
func start(state GameState) {
	log.Printf("%s START\n", state.Game.ID)
}

// This function is called when a game your Battlesnake was in has ended.
// It's purely for informational purposes, you don't have to make any decisions here.
func end(state GameState) {
	log.Printf("%s END\n\n", state.Game.ID)
}

// This function is called on every turn of a game. Use the provided GameState to decide
// where to move -- valid moves are "up", "down", "left", or "right".
// We've provided some code and comments to get you started.
func move(state GameState) BattlesnakeMoveResponse {
	myHead := state.You.Body[0]

	possibleMoves := map[string]PossibleMove{
		"up": {
			Safe: checkForCollision(myHead, state),
			Coord: Coord{
				X: myHead.X + 1,
				Y: myHead.Y,
			},
		},
		"down": {
			Safe: checkForCollision(myHead, state),
			Coord: Coord{
				X: myHead.X - 1,
				Y: myHead.Y,
			},
		},
		"left": {
			Safe: checkForCollision(myHead, state),
			Coord: Coord{
				X: myHead.X,
				Y: myHead.Y - 1,
			},
		},
		"right": {
			Safe: checkForCollision(myHead, state),
			Coord: Coord{
				X: myHead.X,
				Y: myHead.Y + 1,
			},
		},
	}

	// choose a move from the available safe moves.
	var nextMove string

	safeMoves := []string{}
	for move, possibleMove := range possibleMoves {
		if possibleMove.Safe {
			safeMoves = append(safeMoves, move)
		}
	}

	if len(safeMoves) == 0 {
		nextMove = "down"
	} else {
		nextMove = safeMoves[rand.Intn(len(safeMoves))]
	}

	var battlesnakeMoveResponse BattlesnakeMoveResponse

	battlesnakeMoveResponse.Move = nextMove

	return battlesnakeMoveResponse
}

func checkForCollision(coord Coord, state GameState) bool {
	myBody := state.You.Body
	boardMaxWidthIndex := state.Board.Width - 1
	boardMaxHeightIndex := state.Board.Height - 1

	// avoid left wall
	if coord.X == -1 || coord.X > boardMaxWidthIndex || coord.Y == -1 || coord.Y > boardMaxHeightIndex {
		return false
	}

	// avoid body
	for _, bodySegment := range myBody {
		if coord.X == bodySegment.X && coord.Y == bodySegment.Y {
			return false
		}
	}

	// avoid snakes
	for _, snake := range state.Board.Snakes {
		for _, snakeSegment := range snake.Body {
			if coord.X == snakeSegment.X && coord.Y == snakeSegment.Y {
				return false
			}
		}
	}

	return true
}
