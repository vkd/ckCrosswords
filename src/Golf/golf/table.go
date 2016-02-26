package golf

import (
	"fmt"
	"strings"

	"bufio"
)

type Table struct {
	Rows, Cols int
	Cells      [][]*Cell

	start_cell  *Cell
	finish_cell *Cell
}

func NewTable(cols, rows int, scanner *bufio.Scanner) (t *Table) {
	t = new(Table)

	t.Rows = rows
	t.Cols = cols

	t.Cells = make([][]*Cell, rows)
	for i := 0; i < rows && scanner.Scan(); i++ {
		t.Cells[i] = make([]*Cell, cols)

		values := strings.Split(scanner.Text(), " ")
		for j := 0; j < cols && j < len(values); j++ {
			t.Cells[i][j] = NewCell(i, j, values[j])
		}
	}

	t.start_cell = t.Cells[0][0]

	t.finish_cell = t.Cells[4][7]
	t.finish_cell.ResultLen = 1

	return t
}

func (t *Table) Strings() []string {
	lines := make([]string, t.Rows)

	for i := 0; i < t.Rows; i++ {
		values := make([]string, t.Cols)
		for j := 0; j < t.Cols; j++ {
			values[j] = t.Cells[i][j].String()
		}
		lines[i] = strings.Join(values, " ")
	}
	return lines
}

func (t *Table) IsExists(p *Point) bool {
	if p.X < 0 || p.X >= t.Rows {
		return false
	}
	if p.Y < 0 || p.Y >= t.Cols {
		return false
	}
	return true
}

func (t *Table) Get(p *Point) *Cell {
	if t.IsExists(p) {
		return t.Cells[p.X][p.Y]
	} else {
		return nil
	}
}

func (t *Table) Solve() {
	basic_directions := &Directions{
		&Direction{1, 0},
		&Direction{-1, 0},
		&Direction{0, 1},
		&Direction{0, -1},
	}

	var next_cells = []*NextCell{
		{
			C: t.finish_cell,
			// Buffer: []*Cell{t.finish_cell},
		},
	}

	for i := 0; i < 30; i++ {
		t.Print()

		var temp_next_cells []*NextCell

		for _, nc := range next_cells {
			if nc.Ds == nil {
				nc.Ds = basic_directions
			}

			for walk_nc := range nc.Ds.Walk(t, nc) {
				temp_next_cells = append(temp_next_cells, walk_nc)
			}
		}

		next_cells = temp_next_cells
		if len(next_cells) == 0 {
			break
		}
	}

	// for j := 0; j < 9; j++ {
	// 	t.Print()

	// 	var ncs []*NextCell
	// 	for _, nc := range next_cells {
	// 		if nc.Ds == nil {
	// 			nc.Ds = basic_directions
	// 		}
	// 		for next := range nc.Ds.Walk(nc) {
	// 			ncs = append(ncs, next)
	// 		}
	// 	}
	// 	next_cells = ncs

	// 	if len(next_cells) == 0 {
	// 		break
	// 	}
	// }
}

func (t *Table) Print() {
	fmt.Println("---------------------------------------")
	fmt.Println(strings.Join(t.Strings(), "\n"))
	fmt.Println("---------------------------------------")
}

func (t *Table) PrintResult() {
	fmt.Println("Result:")
	c := t.start_cell
	fmt.Println(c.ResultLen)
	for i := 0; i < 20 && c != nil; i++ {
		fmt.Println(c.P.String())
		for index, _ := range c.Buffer {
			fmt.Println(c.Buffer[len(c.Buffer)-index-1].P.String(), "b")
		}
		c = c.PrevCell
	}
}
