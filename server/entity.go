package server

type Entity struct {
	Id int `json:"id"`
	X_pos int `json:"x_pos"`
	Y_pos int `json:"y_pos"`
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

