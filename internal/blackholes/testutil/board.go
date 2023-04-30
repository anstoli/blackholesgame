package testutil

import "github.com/anstoli/blackholesgame/internal/blackholes"

func BoardOpenStateAsBools(b [][]*blackholes.Cell) [][]bool {
	o := make([][]bool, len(b)) // open
	for i := range o {
		o[i] = make([]bool, len(b[i]))
	}

	for i := range b {
		for j := range b[i] {
			o[i][j] = b[i][j].IsOpen
		}
	}

	return o
}

func BoardSetBlackHolesFromBools(b [][]*blackholes.Cell, isBlackHoles [][]bool) {
	for i := range b {
		for j := range b[i] {
			b[i][j].IsBlackHole = isBlackHoles[i][j]
		}
	}
}

func BoardSetAdjacentHolesFromInts(b [][]*blackholes.Cell, adjacentHoles [][]int) {
	for i := range b {
		for j := range b[i] {
			b[i][j].AdjacentHolesNumber = adjacentHoles[i][j]
		}
	}
}
