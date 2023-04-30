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
	R, C                int // Raw, Column of the cell is used for debugging purposes in graph view
}

type Game struct {
	Board       [][]*Cell // Board represents game state with first index is a row and second is a column of a cell.
	N           int
	K           int // K is number of black holes
	State       GameState
	CellsToOpen int
}

type GameState string

const (
	GameStateActive GameState = "active"
	GameStateWon    GameState = "won"
	GameStateLost   GameState = "lost"
)

var (
	ErrInvalidBoardSize        = errors.New("invalid board size (N)")
	ErrInvalidBlackHolesNumber = errors.New("invalid black holes number (K)")
)

// NewGame creates and initializes new ready to play *Game with NxN board and K black holes.
func NewGame(n int, k int) (*Game, error) {
	if n <= 0 {
		return nil, ErrInvalidBoardSize
	}
	if k <= 0 || k > n*n {
		return nil, ErrInvalidBlackHolesNumber
	}
	return &Game{
		Board:       newBoard(n, k),
		N:           n,
		K:           k,
		State:       GameStateActive,
		CellsToOpen: n * n,
	}, nil
}

func newBoard(n int, k int) [][]*Cell {
	board := newBoardOfCells(n)
	placeBlackHoles(n, k, board)
	calcAdjacentBlackHoles(n, board)
	return board
}

func newBoardOfCells(n int) [][]*Cell {
	b := make([][]*Cell, n) // board
	for i := range b {
		b[i] = make([]*Cell, n)
	}
	for i := range b {
		for j := range b[i] {
			b[i][j] = &Cell{
				R: i,
				C: j,
			}
		}
	}
	return b
}

func placeBlackHoles(n int, k int, board [][]*Cell) {
	for k > 0 {
		r := rand.Intn(n) // row
		c := rand.Intn(n) // column
		if board[r][c].IsBlackHole {
			continue
		}
		board[r][c].IsBlackHole = true
		k--
	}
}

// adjacentCells are indexes shift of adjacent cells starting from right one in clockwise order
var adjacentCells = [][]int{
	{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1},
}

func calcAdjacentBlackHoles(n int, board [][]*Cell) {
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
				board[i][j].AdjacentCells = append(board[i][j].AdjacentCells, board[r][c])
				if board[r][c].IsBlackHole {
					adjacentHoles++
				}
			}
			board[i][j].AdjacentHolesNumber = adjacentHoles
		}
	}
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

	g.Board[r][c].IsOpen = true
	g.CellsToOpen--

	// if current cell has 0 adjacent black holes, open neighbour cells
	var toVisit []*Cell
	toVisit = append(toVisit, g.Board[r][c])
	visited := make(map[*Cell]struct{})
	for len(toVisit) > 0 {
		adjCell := toVisit[len(toVisit)-1]
		toVisit = toVisit[0 : len(toVisit)-1]
		if _, ok := visited[adjCell]; ok {
			continue
		}
		visited[adjCell] = struct{}{}
		if !adjCell.IsOpen {
			adjCell.IsOpen = true
			g.CellsToOpen--
		}
		if adjCell.AdjacentHolesNumber != 0 {
			continue
		}
		toVisit = append(toVisit, adjCell.AdjacentCells...)
	}

	if g.CellsToOpen <= g.K {
		g.State = GameStateWon
	}
	return nil
}
