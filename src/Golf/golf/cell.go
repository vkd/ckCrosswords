package golf

import (
	"fmt"
	"strconv"
)

type Cell struct {
	P     *Point
	Value int

	Buffer    []*Cell
	PrevCell  *Cell
	ResultLen int
}

func NewCell(x, y int, value string) (c *Cell) {
	c = new(Cell)
	c.P = &Point{x, y}

	i, err := strconv.Atoi(value)
	if err != nil {
		i = 0
	}
	c.Value = i

	// if c.Value == 0 {
	// 	c.IsFinish = true
	// }
	return c
}

func (c *Cell) String() string {
	prefix := " "
	postfix := " "
	if c.IsCalculated() {
		prefix = "["
		postfix = "]"
	}

	if c.Value > 0 {
		return prefix + fmt.Sprintf("+%d", c.Value) + postfix
	} else if c.Value < 0 {
		return prefix + fmt.Sprintf("%d", c.Value) + postfix
	} else {
		return prefix + "00" + postfix
	}
}

func (c *Cell) AbsEq(i int) bool {
	if c.Value < 0 {
		return (c.Value * -1) == i
	} else {
		return c.Value == i
	}
}

func (c *Cell) IsPositive() bool {
	return c.Value > 0
}

func (c *Cell) IsNegative() bool {
	return c.Value < 0
}

func (c *Cell) IsCalculated() bool {
	return c.ResultLen > 0
}

func (c *Cell) Update(prev_cell *Cell, buffer []*Cell) {
	if c.IsNegative() {
		return
	}

	nc_len := prev_cell.ResultLen + len(buffer) + 1

	if c.ResultLen == 0 || nc_len < c.ResultLen {
		c.Buffer = buffer
		c.PrevCell = prev_cell
		c.ResultLen = nc_len
	}
}
