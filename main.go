package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	b "github.com/elideveloper/tic-tac-toe/board"
	"github.com/elideveloper/tic-tac-toe/pkg/minimax"
)

// TODO need not terminal estimation for movies
// because bot cannot find full path to win with small depth (6-8)

func main() {
	for {
		level := 0
		fmt.Println("Начало игры")
		board := b.Board{}

		for !board.IsEnd() {
			if level == 0 {
				board.SetMove(3, 3, true)
			} else {
				// bot is considered as maximizer player
				board = minimax.FindBestUsingMinimax(board, true).(b.Board)
			}

			fmt.Println("\nМой ход!")
			board.Print()
			if board.IsWin(true) {
				fmt.Println("Я выиграл!!!")
				break
			}
			if board.IsEnd() {
				fmt.Println("Ничья.")
				break
			}
			time.Sleep(time.Second * 2)

			fmt.Println()
			fmt.Println("Возможные ходы: ")
			possibMoves := board.GetPossibleMoves(false)
			for i := range possibMoves {
				fmt.Printf("номер хода '%d'\n", i)
				possibMoves[i].Print()
			}

			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Введите номер выбранного хода: ")

			text, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			fmt.Println("Ваш ход: ", text)

			if text == "q\n" {
				break
			}

			moveIndex, err := strconv.Atoi(text[:len(text)-1])
			if err != nil {
				panic(err)
			}

			board = possibMoves[moveIndex]
			fmt.Println("Доска после вашего хода")
			board.Print()
			if board.IsWin(false) {
				fmt.Println("Вы выиграли.")
				break
			}

			level++
			time.Sleep(time.Second * 1)
		}

		time.Sleep(time.Second * 4)
		fmt.Println()
	}
}
