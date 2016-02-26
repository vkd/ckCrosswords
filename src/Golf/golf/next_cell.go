package golf

type NextCell struct {
	C  *Cell
	Ds *Directions

	Buffer []*Cell

	PrevPositiveCell *Cell
}
