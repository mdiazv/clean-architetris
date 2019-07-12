package domain

type TetrisWorld struct {
	current Piece
	w, h    int
	rubble  [][]bool
}

func MakeTetrisWorld() World {
	w, h := 10, 20
	r := make([][]bool, h)
	for y := 0; y < h; y++ {
		r[y] = make([]bool, w)
	}
	return &TetrisWorld{
		w:      w,
		h:      h,
		rubble: r,
	}
}

func (w *TetrisWorld) Cleared() int {
	return 0
}

func (w *TetrisWorld) Dimensions() (int, int) {
	return w.w, w.h
}

func (w *TetrisWorld) Drop() bool {
	if w.current == nil {
		w.Spawn()
	}
	for _, p := range w.current.GetPoints() {
		if !w.IsFree(p.X, p.Y) || !w.IsFree(p.X, p.Y+1) {
			return false
		}
	}
	w.current.move(0, 1)
	return true
}

func (w *TetrisWorld) Spawn() bool {
	p := MakeSquarePiece()
	w.current = p
	return true
}

func (w *TetrisWorld) IsFree(x, y int) bool {
	if x < 0 || x >= w.w || y < 0 || y >= w.h {
		return false
	}
	return w.rubble[y][x] == false
}

func (w *TetrisWorld) Taken() []Point {
	t := make([]Point, 0)
	for y := 0; y < w.h; y++ {
		for x := 0; x < w.w; x++ {
			if !w.IsFree(x, y) {
				t = append(t, Point{X: x, Y: y})
			}
		}
	}
	t = append(t, w.current.GetPoints()...)
	return t
}
