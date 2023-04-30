package blackholes_test

import (
	"github.com/anstoli/blackholesgame/internal/blackholes"
	bhtestutil "github.com/anstoli/blackholesgame/internal/blackholes/testutil"
	"reflect"
	"testing"
)

func TestNewGame(t *testing.T) {
	t.Run("black holes number", func(t *testing.T) {
		const n = 5
		const k = 5
		g, err := blackholes.NewGame(n, k)
		if err != nil {
			t.Fatalf("NewGame error %v", err)
		}

		bhN := 0 // black holes number
		for i := range g.Board {
			for j := range g.Board[i] {
				if g.Board[i][j].IsBlackHole {
					bhN++
				}
			}
		}

		if bhN != 5 {
			t.Errorf("NewGame black holes number = %v, want %v", bhN, k)
		}
	})

}

func TestOpen(t *testing.T) {
	t.Run("auto opens only zeros", func(t *testing.T) {
		g, err := blackholes.NewGame(3, 1)
		if err != nil {
			t.Fatalf("NewGame error %v", err)
		}

		bhtestutil.BoardSetBlackHolesFromBools(g.Board, [][]bool{
			{false, false, false},
			{false, false, false},
			{false, false, true},
		})

		bhtestutil.BoardSetAdjacentHolesFromInts(g.Board, [][]int{
			{0, 0, 0},
			{0, 1, 1},
			{0, 1, 0},
		})

		err = g.Open(0, 0)
		if err != nil {
			t.Errorf("Open got error = %v", err)
		}

		if g.State != blackholes.GameStateWon {
			t.Errorf("Open State = %v, want %v", g.State, blackholes.GameStateWon)
		}

		isOpen := bhtestutil.BoardOpenStateAsBools(g.Board)
		isOpenExpected := [][]bool{
			{true, true, true},
			{true, true, true},
			{true, true, false},
		}

		if !reflect.DeepEqual(isOpen, isOpenExpected) {
			t.Errorf("Game.Board.IsOpen =\n%v\n want\n%v", isOpen, isOpenExpected)
		}
	})
}
