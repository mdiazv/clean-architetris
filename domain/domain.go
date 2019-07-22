package domain

type Point struct {
	X, Y int
}

type Piece interface {
	GetPoints() []Point
	Rotate()
	move(dx, dy int)
}

type World interface {
	Cleared() int
	Dimensions() (int, int)
	Drop() bool
	Spawn() bool
	IsFree(int, int) bool
	Taken() []Point
}

type Game interface {
	Reset()
	Start()
	Score() int
	TogglePause()
	IsOver() bool
}
