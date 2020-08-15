package main

import (
	"./ui"
)

type gameMenu struct {
	button09x09 ui.Element
	button13x13 ui.Element
	button19x19 ui.Element
}

func createMenu() (m gameMenu, e error) {
	x  := float64(screenSize/5)
	w  := float64(screenSize) - 2*x
	y1 := float64(screenSize/5)
	y2 := float64(screenSize/5*2)
	y3 := float64(screenSize/5*3)
	h  := float64(screenSize/5)

	onHover := func(elem *ui.Element) error {
		elem.Data.(*ui.ButtonData).Color = colorBlue
		elem.Padding.L *= 0.5
		elem.Padding.R *= 0.5
		elem.Padding.T *= 0.5
		elem.Padding.B *= 0.5
		return nil
	}

	onUnhover := func(elem *ui.Element) error {
		elem.Data.(*ui.ButtonData).Color = colorGrey
		elem.Padding.L *= 2
		elem.Padding.R *= 2
		elem.Padding.T *= 2
		elem.Padding.B *= 2
		return nil
	}

	onPress := func(elem *ui.Element) error {
		elem.Data.(*ui.ButtonData).Color = colorDarkBlue
		return nil
	}

	m.button09x09 = ui.Element{
		ui.Button,
		ui.Bounds{x, y1, w, h},
		ui.Padding{0.05, 0.05, 0.1, 0.1},
		new(ui.ButtonData),
	}

	m.button13x13 = ui.Element{
		ui.Button,
		ui.Bounds{x, y2, w, h},
		ui.Padding{0.05, 0.05, 0.1, 0.1},
		new(ui.ButtonData),
	}

	m.button19x19 = ui.Element{
		ui.Button,
		ui.Bounds{x, y3, w, h},
		ui.Padding{0.05, 0.05, 0.1, 0.1},
		new(ui.ButtonData),
	}

	*m.button09x09.Data.(*ui.ButtonData) = ui.ButtonData {
		Text: ui.Text {Text: "09x09", Size: 1.0, Color: colorWhite, Font: menuFont },
		Color: colorGrey,
		OnHover: onHover,
		OnUnhover: onUnhover,
		OnPress: onPress,
		OnRelease: func(elem *ui.Element) error {
			var err error
			state = playingState
			board, err = createBoard(b09x09)
			return err
		},
	}

	*m.button13x13.Data.(*ui.ButtonData) = ui.ButtonData {
		Text: ui.Text {Text: "13x13", Size: 1.0, Color: colorWhite, Font: menuFont},
		Color: colorGrey,
		OnHover: onHover,
		OnUnhover: onUnhover,
		OnPress: onPress,
		OnRelease: func(elem *ui.Element) error {
			var err error
			state = playingState
			board, err = createBoard(b13x13)
			return err
		},
	}

	*m.button19x19.Data.(*ui.ButtonData) = ui.ButtonData {
		Text: ui.Text {Text: "19x19", Size: 1.0, Color: colorWhite, Font: menuFont},
		Color: colorGrey,
		OnHover: onHover,
		OnUnhover: onUnhover,
		OnPress: onPress,
		OnRelease: func(elem *ui.Element) error {
			var err error
			state = playingState
			board, err = createBoard(b19x19)
			return err
		},
	}

	e = m.button09x09.Data.(*ui.ButtonData).Text.Render(ren)
	if e != nil {return}
	e = m.button13x13.Data.(*ui.ButtonData).Text.Render(ren)
	if e != nil {return}
	e = m.button19x19.Data.(*ui.ButtonData).Text.Render(ren)
	if e != nil {return}

	return
}

func (m *gameMenu) update() (e error) {
	e = m.button09x09.Update()
	if e != nil {return}
	e = m.button13x13.Update()
	if e != nil {return}
	e = m.button19x19.Update()
	if e != nil {return}
	return
}

func (m gameMenu) draw() (e error) {
	e = m.button09x09.Draw(ren)
	if e != nil {return}
	e = m.button13x13.Draw(ren)
	if e != nil {return}
	m.button19x19.Draw(ren)
	if e != nil {return}
	return
}
