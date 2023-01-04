package ds

// Pos = Position for storing row,col or x,y
type Pos struct {
	Row int
	Col int
}

// NewPos Pointer to new position struct
func NewPos() *Pos {
	return &Pos{}
}
