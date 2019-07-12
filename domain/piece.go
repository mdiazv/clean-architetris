package domain

type TetrisPiece struct {
	r    int
	rots int
	pos  Point
	ps   [][]Point
}

func (tp *TetrisPiece) GetPoints() []Point {
	ps := make([]Point, len(tp.ps[tp.r]))
	for i, p := range tp.ps[tp.r] {
		ps[i] = tp.pos.plus(p)
	}
	return ps
}

func (tp *TetrisPiece) Rotate() {
	tp.r = (tp.r + 1) % tp.rots
}

func (tp *TetrisPiece) move(dx, dy int) {
	tp.pos.X += dx
	tp.pos.Y += dy
}

func MakeSquarePiece() Piece {
	return &TetrisPiece{
		rots: 1,
		ps: [][]Point{{
			{0, 0}, {1, 0},
			{0, 1}, {1, 1},
		}},
	}
}

func MakeZPiece() Piece {
	return &TetrisPiece{
		rots: 2,
		ps: [][]Point{
			{
				{0, 0}, {1, 0},
				/*   */ {1, 1}, {2, 1},
			},
			{ /*     */ {1, 0},
				{0, 1}, {1, 1},
				{0, 2},
			},
		},
	}
}
