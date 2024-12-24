package server

type Entity struct {
	Id int
	x_pos int
	y_pos int
}

func (e *Entity) Move(x int, y int) {
	e.x_pos = x
	e.y_pos = y
}

func (e *Entity) GetPosition() (int, int) {
	return e.x_pos, e.y_pos
}

func NewEntity(id int, x int, y int) *Entity {
	return &Entity{id, x, y}
}

