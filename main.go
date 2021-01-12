// Package main implements a collection of functions that allow a user to play
// tic tac toe on the command line again the computer.
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// The gameState type holds an array describing the 9 squares of the board, a bool
// descibing whether or not a player has won, and the name of the winning player as
// a string.
type gameState struct {
	board      [9]string
	isWinner   bool
	isDraw     bool
	winnerName string
}

var game gameState = gameState{}

// userMove prompts the user for input and then updates the current game accordingly.
func (g gameState) userMove() {
	var move string
	legal := false

	for !legal {
		fmt.Println("Please enter a digit 1-9:")
		fmt.Scanln(&move)
		switch move {
		case "1":
			if game.board[0] == " " {
				game.board[0] = "X"
				legal = true
			}
		case "2":
			if game.board[1] == " " {
				game.board[1] = "X"
				legal = true
			}
		case "3":
			if game.board[2] == " " {
				game.board[2] = "X"
				legal = true
			}
		case "4":
			if game.board[3] == " " {
				game.board[3] = "X"
				legal = true
			}
		case "5":
			if game.board[4] == " " {
				game.board[4] = "X"
				legal = true
			}
		case "6":
			if game.board[5] == " " {
				game.board[5] = "X"
				legal = true
			}
		case "7":
			if game.board[6] == " " {
				game.board[6] = "X"
				legal = true
			}
		case "8":
			if game.board[7] == " " {
				game.board[7] = "X"
				legal = true
			}
		case "9":
			if game.board[8] == " " {
				game.board[8] = "X"
				legal = true
			}
		}
	}
}

// checkWinnerName assumes that a winner has been found, and looks up a given
// square on the board, one of the winner's three-in-a-row, to infer who has won.
func (g gameState) checkWinnerName(square string) {
	if square == "X" {
		game.winnerName = "user"
	}
}

// checkIsWinner looks at the current game and updates the game struct if either player has won.
func (g gameState) checkIsWinner() {
	if game.board[0] == game.board[1] && game.board[0] == game.board[2] && game.board[0] != " " {
		game.isWinner = true
		game.checkWinnerName(game.board[0])
	}
	if game.board[3] == game.board[4] && game.board[3] == game.board[5] && game.board[2] != " " {
		game.isWinner = true
		game.checkWinnerName(game.board[3])
	}
	if game.board[6] == game.board[7] && game.board[6] == game.board[8] && game.board[6] != " " {
		game.isWinner = true
		game.checkWinnerName(game.board[6])
	}
	if game.board[0] == game.board[4] && game.board[0] == game.board[8] && game.board[0] != " " {
		game.isWinner = true
		game.checkWinnerName(game.board[0])
	}
	if game.board[2] == game.board[4] && game.board[2] == game.board[6] && game.board[2] != " " {
		game.isWinner = true
		game.checkWinnerName(game.board[2])
	}
	if game.board[0] == game.board[3] && game.board[0] == game.board[6] && game.board[0] != " " {
		game.isWinner = true
		game.checkWinnerName(game.board[0])
	}
	if game.board[1] == game.board[4] && game.board[1] == game.board[7] && game.board[1] != " " {
		game.isWinner = true
		game.checkWinnerName(game.board[1])
	}
	if game.board[2] == game.board[5] && game.board[2] == game.board[8] && game.board[2] != " " {
		game.isWinner = true
		game.checkWinnerName(game.board[2])
	}

	game.isDraw = true
	for _, i := range game.board {
		if i == " " {
			game.isDraw = false
		}
	}
}

// computerMove randomly selects an empty space for the computer to play its
// move and updates the current game accordingly.
func (g gameState) computerMove() {
	var moveString string

	for i, square := range game.board {
		if square == " " {
			moveString += strconv.Itoa(i)
		}
	}

	possibleMoves := []rune(moveString)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	move := r1.Intn(len(possibleMoves) - 1)

	squarePicked := int(possibleMoves[move]) - '0'

	game.board[squarePicked] = "O"
}

// PrintGame is a helper method that prints the current game state to the console
// in a readable format.
func (g gameState) printGame() {
	fmt.Printf("\n%v | %v | %v \n⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯\n%v | %v | %v\n⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯\n%v | %v | %v\n",
		game.board[0], game.board[1], game.board[2], game.board[3], game.board[4], game.board[5], game.board[6], game.board[7], game.board[8])
}

// main handles the primary logic for the game, allowing the players to take turns,
// ending the game if there's a winner or it's a draw, and prints instructions
// to the user.
func main() {
	// Just to make the board print nicely to the console. Should probably do this
	// on initialisation, but if the game gets any bigger then this way is better.
	for i := range game.board {
		game.board[i] = " "
	}

	fmt.Println("Welcome to Tic Tac Go!")
	fmt.Println("You'll play as Xs, and the computer will play as Os.")
	fmt.Println("Each square is numbered from left to right, top to bottom, like this:")
	fmt.Println("\n1 | 2 | 3 \n⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯\n4 | 5 | 6\n⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯\n7 | 8 | 9")
	fmt.Println(" ")
	fmt.Println("You'll be prompted to enter a number from 1-9 every turn.")
	fmt.Println("You get to go first!")

	for !game.isWinner && !game.isDraw {
		game.printGame()
		game.userMove()
		game.checkIsWinner()
		if game.isWinner || game.isDraw {
			break
		}
		game.computerMove()
		game.checkIsWinner()
	}

	game.printGame()
	if game.winnerName == "user" {
		fmt.Println("Congratulations, you're the winner!")
	} else if game.isDraw {
		fmt.Println("A closely fought contest, it's a draw!")
	} else {
		fmt.Println("Unlucky, the computer beat you this time!")
	}
}
