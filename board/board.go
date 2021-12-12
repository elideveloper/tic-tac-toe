package board

import (
	"fmt"

	"github.com/elideveloper/tic-tac-toe/pkg/minimax"
)

const (
	boardSize         uint8 = 8
	numToWin          uint8 = 5
	numCells          uint8 = 64
	numRowWinPatterns uint8 = 10

	winVal float64 = 1000.0

	emptyCell = '0'
)

const (
	winRow_1 uint8 = 31  // 00011111
	winRow_2 uint8 = 62  // 00111110
	winRow_3 uint8 = 124 // 01111100
	winRow_4 uint8 = 248 // 11111000

	winRow_5 uint8 = 126 // 01111110
	winRow_6 uint8 = 252 // 11111100
	winRow_7 uint8 = 63  // 00111111

	winRow_8 uint8 = 127 // 01111111
	winRow_9 uint8 = 254 // 11111110

	winRow_10 uint8 = 255 // 11111111
)

var allRowWinPatterns = [numRowWinPatterns]uint8{
	winRow_1,
	winRow_2,
	winRow_3,
	winRow_4,
	winRow_5,
	winRow_6,
	winRow_7,
	winRow_8,
	winRow_9,
	winRow_10,
}

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

func (b *Board) SetMove(x, y uint8, isBot bool) {
	pos := boardSize - 1 - x
	if isBot {
		b.botMoves[y] |= (1 << pos)
	} else {
		b.playerMoves[y] |= (1 << pos)
	}
}

func (b *Board) Print() {
	var rowString string
	var i, j uint8
	for i = 0; i < boardSize; i++ {
		if b.botMoves[i] == 0 && b.playerMoves[i] == 0 {
			rowString = "0 0 0 0 0 0 0 0"
		} else {
			strBot := fmt.Sprintf("%08b", b.botMoves[i])
			strPlayer := fmt.Sprintf("%08b", b.playerMoves[i])
			rowString = ""
			for j = 0; j < boardSize; j++ {
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
	var countBot, countPlayer, i, j uint8

	for i = 0; i < boardSize; i++ {
		strBot := fmt.Sprintf("%08b", b.botMoves[i])
		strPlayer := fmt.Sprintf("%08b", b.playerMoves[i])

		for j = 0; j < boardSize; j++ {
			if strBot[j] != emptyCell {
				countBot++
			}
			if strPlayer[j] != emptyCell {
				countPlayer++
			}
		}
	}

	// all cells on board were set
	return countBot+countPlayer == numCells
}

func (b Board) Eval() float64 {
	if b.IsWin(true) {
		return winVal
	}

	if b.IsWin(false) {
		return -winVal
	}

	// // attempt to evaluate
	// // TODO does not work right for now
	// var i uint8
	// var valSum float64
	// for i = 0; i < boardSize; i++ {
	// 	valSum += b.evalRow(i, true)
	// }
	// if valSum > 10.0 {
	// 	return valSum
	// }

	return 0.0
}

func (b Board) EvalNotTerminal() float64 {
	// TODO need apply this evaluation to reached states
	var i uint8
	var valSum float64
	var maxVal float64
	for i = 0; i < boardSize; i++ {
		valSum = b.evalRow(i, true)
		if valSum > maxVal {
			maxVal = valSum
		}
	}
	return maxVal
}

// bot is considered as maximizer
func (b Board) GetChildren(isMaximizer bool) []minimax.State {
	val := b.Eval()
	if val == winVal || val == -winVal {
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
	for i = 0; i < boardSize; i++ {
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
	var pos, i uint8
	for i = 0; i < boardSize; i++ {
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

func (b *Board) winInRow(isBot bool) bool {
	needBoard := b.getNeedBoard(isBot)

	var i uint8
	for i = 0; i < boardSize; i++ {
		if needBoard[i] == winRow_1 ||
			needBoard[i] == winRow_2 ||
			needBoard[i] == winRow_3 ||
			needBoard[i] == winRow_4 ||
			needBoard[i] == winRow_5 ||
			needBoard[i] == winRow_6 ||
			needBoard[i] == winRow_7 ||
			needBoard[i] == winRow_8 ||
			needBoard[i] == winRow_9 ||
			needBoard[i] == winRow_10 {
			return true
		}
	}
	return false
}

func (b *Board) winInColumn(isBot bool) bool {
	needBoard := b.getNeedBoard(isBot)

	var pos, i uint8
	numInColumns := [boardSize]uint8{}
	for i = 0; i < boardSize; i++ {
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
	needBoard := b.getNeedBoard(isBot)

	var i, j uint8
	var counterLeft, counterRight uint8
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
		for i = boardSize - 1; i >= j; i-- {
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

func (b *Board) evalRow(rowIndex uint8, isBot bool) float64 {
	evalRow := b.getNeedBoard(isBot)[rowIndex]
	oppRow := b.getNeedBoard(!isBot)[rowIndex]

	possibleWinPatterns := make([]uint8, 0, numRowWinPatterns)
	var i uint8
	for i = 0; i < numRowWinPatterns; i++ {
		if oppRow&allRowWinPatterns[i] == 0 {
			possibleWinPatterns = append(possibleWinPatterns, allRowWinPatterns[i])
		}
	}

	var val, pos uint8
	for i = 0; i < uint8(len(possibleWinPatterns)); i++ {
		var row uint8 = evalRow & possibleWinPatterns[i]
		if row != 0 {
			for pos = 0; pos < boardSize; pos++ {
				// counts number of set position in winning pattern
				if isSetBit(row, pos) {
					val++
				}
			}
		}
	}

	return float64(val)
}

func (b *Board) getNeedBoard(isBot bool) *BoardMatrix {
	if isBot {
		return &b.botMoves
	}

	return &b.playerMoves
}

func isSetBit(row uint8, pos uint8) bool {
	return row&(1<<pos) != 0
}
