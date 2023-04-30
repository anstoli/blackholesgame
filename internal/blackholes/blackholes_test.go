package blackholes_test

import (
	"github.com/anstoli/blackholesgame/internal/blackholes"
	"testing"
)

func TestNewGame(t *testing.T) {
	t.Run("black holes number", func(t *testing.T) {
		const n = 5
		const k = 5
		g := blackholes.NewGame(n, k)

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
