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
	log.Println("INFO")
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
	possibleMoves := map[string]bool{
		"up":    true,
		"down":  true,
		"left":  true,
		"right": true,
	}

	// Step 0: Don't let your Battlesnake move back in on it's own neck
	myHead := state.You.Body[0] // Coordinates of your head
	myNeck := state.You.Body[1] // Coordinates of body piece directly behind your head (your "neck")

	if myNeck.X < myHead.X {
		possibleMoves["left"] = false
	} else if myNeck.X > myHead.X {
		possibleMoves["right"] = false
	} else if myNeck.Y < myHead.Y {
		possibleMoves["down"] = false
	} else if myNeck.Y > myHead.Y {
		possibleMoves["up"] = false
	}

	// TODO: Step 1 - Don't hit walls.
	boardWidth := state.Board.Width
	boardHeight := state.Board.Height

	// avoid left wall
	if myHead.X == 0 {
		possibleMoves["left"] = false
	}

	// avoid right wall
	if myHead.X == boardWidth-1 {
		possibleMoves["right"] = false
	}

	// avoid bottom wall
	if myHead.Y == 0 {
		possibleMoves["down"] = false
	}

	// avoid top wall
	if myHead.Y == boardHeight-1 {
		possibleMoves["up"] = false
	}

	// TODO: Step 2 - Don't hit yourself.
	mybody := state.You.Body

	for index, bodySegment := range mybody {
		// ignore head
		if index != len(mybody)-1 {
			// avoid body right
			if myHead.X+1 == bodySegment.X {
				possibleMoves["right"] = false
			}

			// avoid body left
			if myHead.X-1 == bodySegment.X {
				possibleMoves["left"] = false
			}

			// avoid body up
			if myHead.Y+1 == bodySegment.Y {
				possibleMoves["up"] = false
			}

			// avoid body down
			if myHead.Y-1 == bodySegment.Y {
				possibleMoves["down"] = false
			}
		}
	}

	// TODO: Step 3 - Don't collide with others.
	// snakes := state.Board.Snakes

	// for _, snake := range snakes {
	// 	snake.ID =
	// 	for _, snakeSegment := range snake.Body {
	// 		// avoid snake left
	// 		if myHead.X == snakeSegment.X+1 {
	// 			possibleMoves["right"] = false
	// 		}

	// 		// avoid snake right
	// 		if myHead.X == snakeSegment.X-1 {
	// 			possibleMoves["left"] = false
	// 		}

	// 		// avoid snake below
	// 		if myHead.Y == snakeSegment.Y+1 {
	// 			possibleMoves["up"] = false
	// 		}

	// 		// avoid snake above
	// 		if myHead.Y == snakeSegment.Y-1 {
	// 			possibleMoves["down"] = false
	// 		}
	// 	}
	// }

	// TODO: Step 4 - Find food.
	// Use information in GameState to seek out and find food.

	// Finally, choose a move from the available safe moves.
	// TODO: Step 5 - Select a move to make based on strategy, rather than random.
	var nextMove string

	safeMoves := []string{}
	for move, isSafe := range possibleMoves {
		if isSafe {
			safeMoves = append(safeMoves, move)
		}
	}

	if len(safeMoves) == 0 {
		nextMove = "down"
		log.Printf("%s MOVE %d: No safe moves detected! Moving %s\n", state.Game.ID, state.Turn, nextMove)
	} else {
		nextMove = safeMoves[rand.Intn(len(safeMoves))]
		log.Printf("%s MOVE %d: %s\n", state.Game.ID, state.Turn, nextMove)
	}

	return BattlesnakeMoveResponse{
		Move: nextMove,
	}
}

func isNeighbour(origin Coord, point Coord) bool {
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if origin.Y+dy == point.Y && origin.X+dx == point.X {
				return true
			}
		}
	}

	return false
}
