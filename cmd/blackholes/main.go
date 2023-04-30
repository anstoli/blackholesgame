package main

import (
	"fmt"
	"github.com/anstoli/blackholesgame/internal/blackholes"
	"github.com/anstoli/blackholesgame/internal/text"
	"strconv"
)

func main() {
	var g *blackholes.Game
	for g == nil {
		n := ReadInt("Enter size of the board (N)")
		k := ReadInt("Enter number of black holes (K)")
		var err error
		g, err = blackholes.NewGame(n, k)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
	}
	for g.State == blackholes.GameStateActive {
		fmt.Println(text.MarshalBoard(g.Board, false))
		r := ReadInt("Enter row of cell to open")
		c := ReadInt("Enter column of cell to open")
		err := g.Open(r, c)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
	}
	fmt.Println(text.MarshalBoard(g.Board, true))
	fmt.Println(g.State)
}

func ReadInt(hint string) int {
	for true {
		fmt.Print(hint + ": ")
		var inputRaw string
		_, err := fmt.Scan(&inputRaw)
		if err != nil {
			fmt.Printf("error: %v", err)
		}
		input, err := strconv.Atoi(inputRaw)
		if err != nil {
			fmt.Printf("error: %v", err)
		}
		return input
	}
	return 0
}
