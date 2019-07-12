package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	p := MakeZPiece()

	expected := [][]Point{
		{
			/*   */ {1, 0},
			{0, 1}, {1, 1},
			{0, 2},
		},
		{
			{0, 0}, {1, 0},
			/*   */ {1, 1}, {2, 1},
		},
		{
			/*   */ {1, 0},
			{0, 1}, {1, 1},
			{0, 2},
		},
		{
			{0, 0}, {1, 0},
			/*   */ {1, 1}, {2, 1},
		},
	}

	for _, e := range expected {
		p.Rotate()
		assert.Equal(t, e, p.GetPoints())
	}
}

func TestMove(t *testing.T) {
	p := MakeSquarePiece().(*TetrisPiece)

	moves := []Point{
		{0, 0},
		{5, 0},
		{0, 10},
		{-1, 1},
	}

	expected := [][]Point{
		{
			{0, 0}, {1, 0},
			{0, 1}, {1, 1},
		},
		{
			{5, 0}, {6, 0},
			{5, 1}, {6, 1},
		},
		{
			{5, 10}, {6, 10},
			{5, 11}, {6, 11},
		},
		{
			{4, 11}, {5, 11},
			{4, 12}, {5, 12},
		},
	}
	for i, e := range expected {
		p.move(moves[i].X, moves[i].Y)
		assert.Equal(t, e, p.GetPoints())
	}
}
