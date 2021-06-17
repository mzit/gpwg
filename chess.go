package main

import "fmt"

type boardType [8][8]rune

func main() {

	var board boardType
	//黑棋
	black := [...]rune{
		'r', 'n', 'b', 'q', 'k', 'b', 'n', 'r',
	}
	//白棋
	white := [...]rune{
		'R', 'N', 'B', 'Q', 'K', 'B', 'N', 'R',
	}
	board = fill(board, black, 0)
	board = fill(board, white, 7)

	display(board)
}

func fill(board boardType, color [8]rune, line int) boardType {
	for column := range color {
		board[line][column] = color[column]
	}
	return board
}
func display(board boardType) {
	for _, row := range board {
		for _, column := range row {
			if column == 0 {
				fmt.Print("  ")
			} else {
				fmt.Printf("%c ", column)
			}
		}
		fmt.Println()
	}
}
