package server

type Entity struct {
	Id int
	X_pos int
	Y_pos int
}

func (e *Entity) Move(x int, y int) {
	e.X_pos = x
	e.Y_pos = y
}

func (e *Entity) GetPosition() (int, int) {
	return e.X_pos, e.Y_pos
}

func NewEntity(id int, x int, y int) *Entity {
	return &Entity{id, x, y}
}

