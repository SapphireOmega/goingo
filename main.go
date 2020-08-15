package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	screenSize = 900
)

var (
	colorGrey     = sdl.Color{R: 100, G: 100, B: 100, A: 255}
	colorWhite    = sdl.Color{R: 230, G: 230, B: 230, A: 255}
	colorBlack    = sdl.Color{R:  25, G:  25, B:  25, A: 255}
	colorBlue     = sdl.Color{R: 100, G: 100, B: 255, A: 255}
	colorDarkBlue = sdl.Color{R:  80, G:  80, B: 150, A: 255}

	menuFont *ttf.Font
	menuFontZoom *ttf.Font

	win *sdl.Window
	ren *sdl.Renderer

	state gameState
	menu  gameMenu
	board gameBoard
)

type gameState int

const (
	menuState gameState = iota
	playingState
)

func checkErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func checkErrWithMsg(err error, msg string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, msg, err)
		os.Exit(1)
	}
}

func main () {

	// INITIALIZATION

	var err error

	err = sdl.Init(sdl.INIT_VIDEO)
	checkErrWithMsg(err, "Error initializing SDL:")
	defer sdl.Quit()

	err = ttf.Init()
	checkErrWithMsg(err, "Error initializing SDL_TTF:")
	defer ttf.Quit()

	win, err = sdl.CreateWindow(
		"Go in go!",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenSize, screenSize,
		sdl.WINDOW_OPENGL,
	)
	checkErrWithMsg(err, "Error creating SDL window:")
	defer win.Destroy()

	ren, err = sdl.CreateRenderer(win, -1, sdl.RENDERER_ACCELERATED)
	checkErrWithMsg(err, "Error creating SDL renderer:")
	defer ren.Destroy()

	menuFont, err = ttf.OpenFont("test.ttf", 64)
	checkErrWithMsg(err, "Error opening font:")
	defer menuFont.Close()

	menuFontZoom, err = ttf.OpenFont("test.ttf", 70)
	checkErrWithMsg(err, "Error opening font:")
	defer menuFontZoom.Close()

	state = menuState
	menu, err = createMenu()
	checkErr(err)

	// GAME LOOP

	for {
		// check for quit event
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		// clear window
		ren.SetDrawColor(50, 50, 50, 255)
		ren.Clear()

		// update and draw the correct state
		switch state {
		case menuState:
			err = menu.update()
			checkErr(err)
			menu.draw()
		case playingState:
			keys := sdl.GetKeyboardState()
			if keys[sdl.SCANCODE_ESCAPE] == 1 {
				state = menuState
				menu, err = createMenu()
				checkErr(err)
			}
			board.update()
			board.draw()
		}

		// present the new frame
		ren.Present()
	}
}
