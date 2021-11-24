package board

import (
	"fmt"
)

const (
	boardSize = 8
	numToWin  = 5
)

type BoardMatrix [boardSize]uint8

type Board struct {
	botMoves    BoardMatrix
	playerMoves BoardMatrix
}

func NewBoard(botMoves, playerMoves BoardMatrix) Board {
	return Board{
		botMoves:    botMoves,
		playerMoves: playerMoves,
	}
}

func (b *Board) Print() {
	for i := 0; i < boardSize; i++ {
		if b.botMoves[i] == 0 && b.playerMoves[i] == 0 {
			fmt.Println("0 0 0 0 0 0 0 0")
		} else {
			strBot := fmt.Sprintf("%08b", b.botMoves[i])
			strPlayer := fmt.Sprintf("%08b", b.playerMoves[i])
			finalString := ""
			for j := 0; j < boardSize; j++ {
				if strBot[j] != '0' {
					finalString += "B "
				} else if strPlayer[j] != '0' {
					finalString += "P "
				} else {
					finalString += "0 "
				}
			}
			fmt.Println(finalString)
		}
	}
}

func (b *Board) winInRow(isBot bool) bool {
	var needBoard *BoardMatrix
	if isBot {
		needBoard = &b.botMoves
	} else {
		needBoard = &b.playerMoves
	}

	// TODO move to named constants
	for i := 0; i < boardSize; i++ {
		if needBoard[i] == 31 ||
			needBoard[i] == 62 ||
			needBoard[i] == 124 ||
			needBoard[i] == 248 ||
			needBoard[i] == 126 ||
			needBoard[i] == 252 ||
			needBoard[i] == 63 ||
			needBoard[i] == 127 ||
			needBoard[i] == 254 ||
			needBoard[i] == 255 {
			return true
		}
	}
	return false
}

func (b *Board) winInColumn(isBot bool) bool {
	var needBoard *BoardMatrix
	if isBot {
		needBoard = &b.botMoves
	} else {
		needBoard = &b.playerMoves
	}

	numInColumns := [boardSize]uint8{}
	for i := 0; i < boardSize; i++ {
		for pos := 0; pos < boardSize; pos++ {
			if (needBoard[i])&(1<<pos) != 0 {
				numInColumns[pos]++
				if numInColumns[pos] >= numToWin {
					return true
				}
			} else {
				numInColumns[pos] = 0
			}
		}
	}
	return false
}

func (b *Board) IsWin(isBot bool) bool {
	if b.winInRow(isBot) {
		return true
	}

	return b.winInColumn(isBot)
}

// TODO implement
func (b *Board) IsEnd() bool { return false }

// type Move struct {
// 	X uint
// 	Y uint
// }

// func (b *Board) SetMove(m Move, playerID int) *Board {
// 	// TODO possible to add validation for emptiness

// 	nb := *b
// 	if nb.matrix[m.X][m.Y] != 0 {
// 		panic("cannot move!")
// 	}
// 	nb.matrix[m.X][m.Y] = playerID

// 	return &nb
// }

// func (b *Board) GetAllPossibleMoves() []Move {
// 	possMoves := []Move{}
// 	var i, j uint
// 	for i = 0; i < b.size; i++ {
// 		for j = 0; j < b.size; j++ {
// 			if b.matrix[i][j] == 0 {
// 				possMoves = append(possMoves, Move{
// 					X: i,
// 					Y: j,
// 				})
// 			}
// 		}
// 	}

// 	return possMoves
// }
