package golf

type Direction struct {
	Dx, Dy int
}

func (d *Direction) GetPoint(start_point *Point, step int) *Point {
	return &Point{start_point.X + d.Dx*step, start_point.Y + d.Dy*step}
}

func (d *Direction) Negative() *Direction {
	return &Direction{d.Dx * -1, d.Dy * -1}
}

// type Tabler interface {
// 	GetCell(*Point) (*Cell, bool)
// }

type Directions []*Direction

type Tabler interface {
	Get(p *Point) *Cell
}

func (d *Directions) Walk(t Tabler, start_nc *NextCell) <-chan *NextCell {
	walk_cells := make(chan *NextCell)

	go func() {
		for _, direction := range *d {
			for i := 0; ; i++ {
				p := direction.GetPoint(start_nc.C.P, i)
				cell := t.Get(p)

				if cell == nil {
					break
				}

				if cell.AbsEq(i) {
					var nc = &NextCell{
						C: cell,
						// PrevPositiveCell: start_nc.C,
					}

					prev_cell := start_nc.PrevPositiveCell
					if prev_cell == nil {
						prev_cell = start_nc.C
					}

					if cell.IsNegative() {
						nc.Ds = &Directions{
							direction.Negative(),
						}
						nc.Buffer = append(start_nc.Buffer, cell)
						nc.PrevPositiveCell = prev_cell
					} else {
						if cell.IsCalculated() {
							// cell.Update(start_nc.C)
							continue
						}
						cell.Update(prev_cell, start_nc.Buffer)
					}
					walk_cells <- nc
				}
			}
		}
		close(walk_cells)
	}()

	return walk_cells
}

// func (d *Directions) Walk(prev_cell *NextCell) <-chan *NextCell {
// 	cells_chan := make(chan *NextCell)

// 	go func() {
// 		start_cell := prev_cell.C

// 		for _, direction := range d.Ds {
// 			for i := 0; ; i++ {
// 				p := direction.GetPoint(start_cell.P, i)
// 				c, is_ok := d.T.GetCell(p)
// 				if !is_ok {
// 					break
// 				}

// 				if !c.AbsEq(i) {
// 					continue
// 				}

// 				if c.IsCalculated && c.IsPositive() {
// 					if len(start_cell.PrevCells)+len(prev_cell.BufferCells) < len(c.PrevCells) {
// 						c.Update(start_cell, prev_cell.Len)
// 					}
// 					continue
// 				}
// 				if c.IsPositive() {
// 					c.Update(start_cell, prev_cell.Len)
// 				}

// 				next_cell := &NextCell{C: c}

// 				if !c.IsPositive() {
// 					next_cell.Ds = &Directions{
// 						T:  d.T,
// 						Ds: []*Direction{direction.Negative()},
// 					}
// 					next_cell.Len = prev_cell.Len + 1
// 					next_cell.BufferCells = append(next_cell.BufferCells, c)
// 				}

// 				cells_chan <- next_cell
// 			}
// 		}
// 		close(cells_chan)
// 	}()

// 	return cells_chan
// }
