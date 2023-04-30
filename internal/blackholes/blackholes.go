package blackholes

import (
	"errors"
	"math/rand"
)

type Cell struct {
	AdjacentCells       []*Cell
	IsBlackHole         bool
	IsOpen              bool
	AdjacentHolesNumber int
}

type Game struct {
	Board       [][]*Cell // Matrix represents game state with first index is a row and second is a column of a cell.
	N           int
	State       GameState
	CellsToOpen int
}

type GameState string

const (
	GameStateActive GameState = "active"
	GameStateWon    GameState = "won"
	GameStateLost   GameState = "lost"
)

// NewGame creates and initializes new ready to play *Game with NxN board and K black holes.
func NewGame(n int, k int) *Game {
	return &Game{
		Board:       newBoard(n, k),
		N:           n,
		State:       GameStateActive,
		CellsToOpen: n * n,
	}
}

func newBoard(n int, k int) [][]*Cell {
	board := make([][]*Cell, n)
	for i := range board {
		board[i] = make([]*Cell, n)
	}
	for i := range board {
		for j := range board[i] {
			board[i][j] = &Cell{}
		}
	}

	for k > 0 {
		r := rand.Intn(n) // row
		c := rand.Intn(n) // column
		if board[r][c].IsBlackHole {
			continue
		}
		board[r][c].IsBlackHole = true
		k--
	}

	for i := range board {
		for j := range board[i] {
			adjacentHoles := 0
			for _, adj := range adjacentCells {
				r := i + adj[0]
				if r < 0 || r >= n {
					continue
				}
				c := j + adj[1]
				if c < 0 || c >= n {
					continue
				}
				board[r][c].AdjacentCells = append(board[i][j].AdjacentCells, board[r][c])
				if board[r][c].IsBlackHole {
					adjacentHoles++
				}
			}
			board[i][j].AdjacentHolesNumber = adjacentHoles
		}
	}
	return board
}

// adjacentCells are indexes shift of adjacent cells starting from right one in clockwise order
var adjacentCells = [][]int{
	{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1},
}

var (
	ErrInvalidRowIndex    = errors.New("row index is invalid")
	ErrInvalidColumnIndex = errors.New("column index is invalid")
	ErrCellAlreadyOpened  = errors.New("cell is already opened")
)

func (g *Game) Open(r, c int) error {
	if r < 0 || r >= g.N {
		return ErrInvalidRowIndex
	}
	if c < 0 || c >= g.N {
		return ErrInvalidColumnIndex
	}
	if g.Board[r][c].IsOpen {
		return ErrCellAlreadyOpened
	}
	if g.Board[r][c].IsBlackHole {
		g.State = GameStateLost
		return nil
	}

	g.CellsToOpen--
	toVisit := make([]*Cell, len(g.Board[r][c].AdjacentCells))
	visited := make(map[*Cell]struct{})
	copy(toVisit, g.Board[r][c].AdjacentCells)
	for len(toVisit) > 0 {
		adjCell := toVisit[len(toVisit)-1]
		toVisit = toVisit[:len(toVisit)-1]
		if adjCell.IsBlackHole {
			continue
		}
		adjCell.IsOpen = true
		g.CellsToOpen--
		for _, adjAdj := range adjCell.AdjacentCells {
			if _, ok := visited[adjAdj]; ok {
				continue
			}
			toVisit = append(toVisit, adjAdj)
		}
		visited[adjCell] = struct{}{}
	}

	if g.CellsToOpen <= 0 {
		g.State = GameStateWon
	}
	return nil
}
