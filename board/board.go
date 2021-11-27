package board

import (
	"fmt"

	"github.com/elideveloper/tic-tac-toe/pkg/minimax"
)

const (
	boardSize = 8
	numToWin  = 5
	numCells  = 64

	emptyCell = '0'
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

func (b *Board) SetMove(x, y uint, isBot bool) {
	pos := boardSize - 1 - x
	if isBot {
		b.botMoves[y] |= (1 << pos)
	} else {
		b.playerMoves[y] |= (1 << pos)
	}
}

func (b *Board) Print() {
	var rowString string
	for i := 0; i < boardSize; i++ {
		if b.botMoves[i] == 0 && b.playerMoves[i] == 0 {
			rowString = "0 0 0 0 0 0 0 0"
		} else {
			strBot := fmt.Sprintf("%08b", b.botMoves[i])
			strPlayer := fmt.Sprintf("%08b", b.playerMoves[i])
			rowString = ""
			for j := 0; j < boardSize; j++ {
				if strBot[j] != emptyCell {
					rowString += "B "
				} else if strPlayer[j] != emptyCell {
					rowString += "P "
				} else {
					rowString += "0 "
				}
			}
		}
		fmt.Println(rowString)
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

	var pos uint8
	numInColumns := [boardSize]uint8{}
	for i := 0; i < boardSize; i++ {
		for pos = 0; pos < boardSize; pos++ {
			// if is not empty cell
			if isSetBit(needBoard[i], pos) {
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

func (b *Board) winInDiag(isBot bool) bool {
	var needBoard *BoardMatrix
	if isBot {
		needBoard = &b.botMoves
	} else {
		needBoard = &b.playerMoves
	}

	var i, j uint8
	var counterLeft uint
	var counterRight uint
	for j = 0; j < 4; j++ {
		counterLeft = 0
		counterRight = 0
		for i = 0; i < boardSize-j; i++ {
			if isSetBit(needBoard[i], i+j) {
				counterLeft++
				if counterLeft >= numToWin {
					return true
				}
			} else {
				counterLeft = 0
			}

			if isSetBit(needBoard[i], boardSize-1-i-j) {
				counterRight++
				if counterRight >= numToWin {
					return true
				}
			} else {
				counterRight = 0
			}
		}
	}

	for j = 1; j < 4; j++ {
		counterLeft = 0
		counterRight = 0
		for i = boardSize - 1; i >= 0+j; i-- {
			if isSetBit(needBoard[i], i-j) {
				counterLeft++
				if counterLeft >= numToWin {
					return true
				}
			} else {
				counterLeft = 0
			}

			if isSetBit(needBoard[i], j+(boardSize-1-i)) {
				counterRight++
				if counterRight >= numToWin {
					return true
				}
			} else {
				counterRight = 0
			}

		}
	}

	return false
}

func (b *Board) IsWin(isBot bool) bool {
	if b.winInRow(isBot) {
		return true
	}

	if b.winInColumn(isBot) {
		return true
	}

	return b.winInDiag(isBot)
}

func (b *Board) IsEnd() bool {
	countBot := 0
	countPlayer := 0

	for i := 0; i < boardSize; i++ {
		strBot := fmt.Sprintf("%08b", b.botMoves[i])
		strPlayer := fmt.Sprintf("%08b", b.playerMoves[i])

		for j := 0; j < boardSize; j++ {
			if strBot[j] != emptyCell {
				countBot++
			}
			if strPlayer[j] != emptyCell {
				countPlayer++
			}
		}
	}

	// all cells on board were set
	if countBot+countPlayer == numCells {
		return true
	}

	return false
}

func (b Board) Eval() float64 {
	if b.IsWin(true) {
		return 100.0
	}

	if b.IsWin(false) {
		return -100.0
	}

	return 0.0
}

// bot is considered as maximizer
func (b Board) GetChildren(isMaximizer bool) []minimax.State {
	if b.Eval() != 0 {
		// endgame
		return []minimax.State{}
	}

	setMoves := [boardSize][boardSize]bool{}
	var numPossibleChildren uint8 = numCells
	var pos, i uint8
	for i = 0; i < boardSize; i++ {
		for pos = 0; pos < boardSize; pos++ {
			if isSetBit(b.botMoves[i], pos) || isSetBit(b.playerMoves[i], pos) {
				setMoves[i][pos] = true
				numPossibleChildren--
			}
		}
	}

	children := make([]minimax.State, 0, numPossibleChildren)
	for i := 0; i < boardSize; i++ {
		for pos = 0; pos < boardSize; pos++ {
			if !setMoves[i][pos] {
				ch := NewBoard(b.botMoves, b.playerMoves)
				if isMaximizer {
					ch.botMoves[i] |= (1 << pos)
				} else {
					ch.playerMoves[i] |= (1 << pos)
				}
				children = append(children, ch)
			}
		}
	}

	return children
}

func (b *Board) GetPossibleMoves(isBot bool) []Board {
	moves := make([]Board, 0, numCells)
	var pos uint8
	for i := 0; i < boardSize; i++ {
		for pos = 0; pos < boardSize; pos++ {
			if !isSetBit(b.botMoves[i], pos) && !isSetBit(b.playerMoves[i], pos) {
				b := NewBoard(b.botMoves, b.playerMoves)
				if isBot {
					b.botMoves[i] |= (1 << pos)
				} else {
					b.playerMoves[i] |= (1 << pos)
				}
				moves = append(moves, b)
			}
		}
	}

	return moves
}

func isSetBit(row uint8, pos uint8) bool {
	if row&(1<<pos) != 0 {
		return true
	}
	return false
}
