package main

import (
	"fmt"

	"github.com/elideveloper/tic-tac-toe/board"
)

func main() {
	plMoves := board.BoardMatrix{}
	plMoves[2] = 64
	plMoves[1] = 64
	plMoves[3] = 255
	plMoves[4] = 64
	plMoves[5] = 64
	botMoves := board.BoardMatrix{}
	botMoves[6] = 128
	botMoves[5] = 128
	botMoves[4] = 128
	botMoves[3] = 128
	botMoves[2] = 128
	b := board.NewBoard(botMoves, plMoves)
	b.Print()

	fmt.Println(b.IsWin(false))
	fmt.Println(b.IsWin(true))
}
