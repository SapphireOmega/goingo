package main

import (
	"./ui"
)

type gameMenu struct {
	buttons ui.Element
}

func createMenu() (m gameMenu, e error) {
	onHover := func(elem *ui.Element) error {
		elem.Data.(*ui.ButtonData).Color = colorBlue
		elem.Padding.L *= 0.5
		elem.Padding.R *= 0.5
		elem.Padding.T *= 0.5
		elem.Padding.B *= 0.5
		elem.Data.(*ui.ButtonData).Text.Font = menuFontZoom
		elem.Data.(*ui.ButtonData).Text.Render(ren)
		return nil
	}

	onUnhover := func(elem *ui.Element) error {
		elem.Data.(*ui.ButtonData).Color = colorGrey
		elem.Padding.L *= 2
		elem.Padding.R *= 2
		elem.Padding.T *= 2
		elem.Padding.B *= 2
		elem.Data.(*ui.ButtonData).Text.Font = menuFont
		elem.Data.(*ui.ButtonData).Text.Render(ren)
		return nil
	}

	onPress := func(elem *ui.Element) error {
		elem.Data.(*ui.ButtonData).Color = colorDarkBlue
		return nil
	}

	button09x09 := ui.Element{
		Type: ui.Button,
		Padding: ui.Padding{0.05, 0.05, 0.1, 0.1},
		Data: new(ui.ButtonData),
	}

	button13x13 := ui.Element{
		Type: ui.Button,
		Padding: ui.Padding{0.05, 0.05, 0.1, 0.1},
		Data: new(ui.ButtonData),
	}

	button19x19 := ui.Element{
		Type: ui.Button,
		Padding: ui.Padding{0.05, 0.05, 0.1, 0.1},
		Data: new(ui.ButtonData),
	}

	*button09x09.Data.(*ui.ButtonData) = ui.ButtonData {
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

	*button13x13.Data.(*ui.ButtonData) = ui.ButtonData {
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

	*button19x19.Data.(*ui.ButtonData) = ui.ButtonData {
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

	e = button09x09.Data.(*ui.ButtonData).Text.Render(ren)
	if e != nil {return}
	e = button13x13.Data.(*ui.ButtonData).Text.Render(ren)
	if e != nil {return}
	e = button19x19.Data.(*ui.ButtonData).Text.Render(ren)
	if e != nil {return}

	m.buttons = ui.Element{
		Type: ui.Column,
		Bounds: ui.Bounds{0.0, 0.0, screenSize, screenSize},
		Padding: ui.Padding{0.2, 0.2, 0.2, 0.2},
		Data: new(ui.ColumnData),
	}
	m.buttons.Data.(*ui.ColumnData).Elems = []ui.Element{button09x09, button13x13, button19x19}

	return
}

func (m *gameMenu) update() error {
	return m.buttons.Update()
}

func (m gameMenu) draw() (e error) {
	return m.buttons.Draw(ren)
}
