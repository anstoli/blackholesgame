package text

import (
	"github.com/anstoli/blackholesgame/internal/blackholes"
	"strconv"
	"strings"
)

func MarshalBoard(b [][]*blackholes.Cell, allOpen bool) string {
	var sb strings.Builder
	sb.WriteString("  ")
	for i := range b[0] {
		sb.WriteString(strconv.Itoa(i))
	}
	sb.WriteString("\n")
	sb.WriteString("  ")
	for _ = range b[0] {
		sb.WriteString("-")
	}
	sb.WriteString("\n")

	for i := range b {
		sb.WriteString(strconv.Itoa(i))
		sb.WriteString("|")
		for j := range b[i] {
			c := b[i][j] // cell
			if c.IsOpen || allOpen {
				sb.WriteString(marshalCellOpen(c))
			} else {
				sb.WriteString(marshalCellClosed(c))
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func marshalCellClosed(_ *blackholes.Cell) string {
	return " "
}

func marshalCellOpen(c *blackholes.Cell) string {
	switch {
	case c.IsBlackHole:
		return "O"
	case c.AdjacentHolesNumber == 0:
		return "-"
	default:
		return strconv.Itoa(c.AdjacentHolesNumber)
	}
}
