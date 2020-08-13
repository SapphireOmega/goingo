package main

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/gfx"
)

type boardSize int

const (
	b09x09 boardSize = 9
	b13x13 boardSize = 13
	b19x19 boardSize = 19
)

var blackTurn bool = true

type cellState int

const (
	empty cellState = iota
	black
	white
)

type gameBoard struct {
	size        boardSize
	state       [][]cellState
	boarderSize int32
	cellSize    int32
}

func createBoard(size boardSize) (b gameBoard, e error) {
	if size != b09x09 && size != b13x13 && size != b19x19 {
		e = errors.New("Error creating board: size must be b09x09, b13x13 or b19x19")
		return
	}

	b.size = size
	b.state = make([][]cellState, size)
	for i := 0; i < int(size); i++ {
		b.state[i] = make([]cellState, size)
	}

	b.boarderSize = screenSize / 4
	b.cellSize    = (screenSize - 2*b.boarderSize) / int32(size - 1)

	return
}

func (b gameBoard) update() {
	mx, my, ms := sdl.GetMouseState()

	x := (mx - b.boarderSize + b.cellSize/2) / b.cellSize
	y := (my - b.boarderSize + b.cellSize/2) / b.cellSize

	if ms == 1 && !(x >= int32(b.size) || y >= int32(b.size) || x < 0 || y < 0) && b.state[y][x] == empty {
		if blackTurn {
			b.state[y][x] = black
		} else {
			b.state[y][x] = white
		}

		blackTurn = !blackTurn
	}
}

func (b gameBoard) draw() {
	square := sdl.Rect{X: 0, Y: 0, W: screenSize, H: screenSize}

	ren.SetDrawColor(200, 180, 92, 255)
	ren.FillRect(&square)

	ren.SetDrawColor(0, 0, 0, 255)

	gridSize := (int32(b.size) - 1) * b.cellSize

	for i := int32(0); i < int32(b.size); i++ {
		tmp := b.boarderSize + i*b.cellSize
		ren.DrawLine(tmp, b.boarderSize, tmp, b.boarderSize + gridSize)
		ren.DrawLine(b.boarderSize, tmp, b.boarderSize + gridSize, tmp)
	}

	switch b.size {
	case b09x09:
		s   := int32(5)
		ma  := b.boarderSize + 2*b.cellSize - s/2
		mb  := ma            + 2*b.cellSize
		mc  := mb            + 2*b.cellSize
		tmp := sdl.Rect{ma, ma, s, s}

		ren.FillRect(&tmp)
		tmp.Y = mc
		ren.FillRect(&tmp)
		tmp.Y = ma
		tmp.X = mc
		ren.FillRect(&tmp)
		tmp.Y = mc
		ren.FillRect(&tmp)
		tmp.X = mb
		tmp.Y = mb
		ren.FillRect(&tmp)
	case b13x13:
		s   := int32(5)
		ma  := b.boarderSize + 3*b.cellSize - s/2
		mb  := ma            + 3*b.cellSize
		mc  := mb            + 3*b.cellSize
		tmp := sdl.Rect{ma, ma, s, s}

		ren.FillRect(&tmp)
		tmp.Y = mc
		ren.FillRect(&tmp)
		tmp.Y = ma
		tmp.X = mc
		ren.FillRect(&tmp)
		tmp.Y = mc
		ren.FillRect(&tmp)
		tmp.X = mb
		tmp.Y = mb
		ren.FillRect(&tmp)
	case b19x19:
		s   := int32(5)
		ma  := b.boarderSize + 3*b.cellSize - s/2
		mb  := ma            + 6*b.cellSize
		mc  := mb            + 6*b.cellSize
		m   := []int32{ma, mb, mc}
		tmp := sdl.Rect{ma, ma, s, s}

		for y := 0; y < len(m); y++ {
			for x := 0; x < len(m); x++ {
				tmp.X = m[x]
				tmp.Y = m[y]
				ren.FillRect(&tmp)
			}
		}
	}

	for by := 0; by < len(b.state); by++ {
		for bx := 0; bx < len(b.state[0]); bx++ {
			x := b.boarderSize + int32(bx)*b.cellSize
			y := b.boarderSize + int32(by)*b.cellSize

			switch b.state[by][bx] {
			case black:
				gfx.FilledCircleColor(ren, x, y, int32(float64(b.cellSize)*0.49), colorBlack)
			case white:
				gfx.FilledCircleColor(ren, x, y, int32(float64(b.cellSize)*0.49), colorWhite)
			}
		}
	}

	mx, my, _ := sdl.GetMouseState()
	if blackTurn {
		gfx.FilledCircleRGBA(ren, mx, my, int32(float64(b.cellSize)*0.49), 0, 0, 0, 100)
	} else {
		gfx.FilledCircleRGBA(ren, mx, my, int32(float64(b.cellSize)*0.49), 255, 255, 255, 100)
	}
}
