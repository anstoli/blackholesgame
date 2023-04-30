package blackholes_test

import (
	"github.com/anstoli/blackholesgame/internal/blackholes"
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

		g.Board[0][0].IsBlackHole = false
		g.Board[0][1].IsBlackHole = false
		g.Board[0][2].IsBlackHole = false

		g.Board[1][0].IsBlackHole = false
		g.Board[1][1].IsBlackHole = false
		g.Board[1][2].IsBlackHole = false

		g.Board[2][0].IsBlackHole = false
		g.Board[2][1].IsBlackHole = false
		g.Board[2][2].IsBlackHole = true

		g.Board[0][0].AdjacentHolesNumber = 0
		g.Board[0][1].AdjacentHolesNumber = 0
		g.Board[0][2].AdjacentHolesNumber = 0

		g.Board[1][0].AdjacentHolesNumber = 0
		g.Board[1][1].AdjacentHolesNumber = 1
		g.Board[1][2].AdjacentHolesNumber = 1

		g.Board[2][0].AdjacentHolesNumber = 0
		g.Board[2][1].AdjacentHolesNumber = 1
		g.Board[2][2].AdjacentHolesNumber = 0

		err = g.Open(0, 0)
		if err != nil {
			t.Errorf("Open got error = %v", err)
		}

		if g.State != blackholes.GameStateWon {
			t.Errorf("Open State = %v, want %v", g.State, blackholes.GameStateWon)
		}

		if !g.Board[0][0].IsOpen {
			t.Errorf("Open Board[0][0].IsOpen = %v, want %v", g.Board[0][0].IsOpen, true)
		}

		if !g.Board[0][1].IsOpen {
			t.Errorf("Open Board[0][1].IsOpen = %v, want %v", g.Board[0][1].IsOpen, true)
		}

		if !g.Board[0][2].IsOpen {
			t.Errorf("Open Board[0][1].IsOpen = %v, want %v", g.Board[0][2].IsOpen, true)
		}

		if !g.Board[1][0].IsOpen {
			t.Errorf("Open Board[1][0].IsOpen = %v, want %v", g.Board[1][0].IsOpen, true)
		}

		if !g.Board[1][1].IsOpen {
			t.Errorf("Open Board[1][1].IsOpen = %v, want %v", g.Board[1][1].IsOpen, true)
		}

		if !g.Board[1][2].IsOpen {
			t.Errorf("Open Board[1][0].IsOpen = %v, want %v", g.Board[1][2].IsOpen, true)
		}

		if !g.Board[2][0].IsOpen {
			t.Errorf("Open Board[2][0].IsOpen = %v, want %v", g.Board[2][0].IsOpen, true)
		}

		if !g.Board[2][1].IsOpen {
			t.Errorf("Open Board[2][1].IsOpen = %v, want %v", g.Board[2][1].IsOpen, true)
		}

		if g.Board[2][2].IsOpen {
			t.Errorf("Open Board[2][0].IsOpen = %v, want %v", g.Board[2][2].IsOpen, false)
		}

	})
}
