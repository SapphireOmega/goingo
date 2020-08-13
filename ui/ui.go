package ui

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// The bounds of a element
type Bounds struct {
	X, Y, W, H float64
}

// The padding to the bounds of the element as a number between 0 and 1
type Padding struct {
	L, R, T, B float64
}

// An "enum" of all the different element types
type ElementType int

const (
	Button ElementType = iota
)

// A ui element
type Element struct {
	Type    ElementType
	Bounds  Bounds
	Padding Padding
	Data    interface{} // Data specific to different element types
}

// A function that gets called when a certain button event occurs
type ButtonAction func(*Element) error

// Button specific data
type ButtonData struct {
	Text string

	BgColor   sdl.Color
	TextColor sdl.Color

	Font *ttf.Font

	OnPress   ButtonAction
	OnRelease ButtonAction
	OnHover   ButtonAction
	OnUnhover ButtonAction

	rect sdl.Rect

	pressed bool
	hover   bool
}

//type uiColumn struct {
//	elems []uiElement
//}

// BUTTON METHODS

//func CreateButtonData(
//	text string,
//	bgColor, textColor sdl.Color,
//) (b ButtonData) {
//	b.Text      = text
//	b.BgColor   = bgColor
//	b.TextColor = textColor
//	return
//}
//
//func (b *ButtonData) setOnPress(action ButtonAction) {
//	b.OnPress = action
//}
//
//func (b *ButtonData) setOnRelease(action ButtonAction) {
//	b.OnRelease = action
//}
//
//func (b *ButtonData) setOnHover(action ButtonAction) {
//	b.OnHover = action
//}
//
//func (b *ButtonData) setOnUnhover(action ButtonAction) {
//	b.OnUnhover = action
//}

// Update a button
func (b *ButtonData) Update(elem *Element) (e error) {
	pl := elem.Bounds.W * elem.Padding.L
	pr := elem.Bounds.W * elem.Padding.R
	pt := elem.Bounds.H * elem.Padding.T
	pb := elem.Bounds.H * elem.Padding.B

	b.rect.X = int32(elem.Bounds.X + pl)
	b.rect.Y = int32(elem.Bounds.Y + pt)
	b.rect.W = int32(elem.Bounds.W - pl - pr)
	b.rect.H = int32(elem.Bounds.H - pt - pb)

	mx, my, ms := sdl.GetMouseState()

	bl := b.rect.X
	br := b.rect.X + b.rect.W
	bt := b.rect.Y
	bb := b.rect.Y + b.rect.H

	if mx >= bl && mx <= br && my >= bt && my <= bb {
		if ms == 1 && !b.pressed {
			if b.OnPress != nil {
				e = b.OnPress(elem)
			}
			b.pressed = true
			b.hover   = false
		} else if ms == 0 && b.pressed {
			if b.OnRelease != nil {
				e = b.OnRelease(elem)
			}
			b.pressed = false
			b.hover   = true
		} else if ms == 0 && !b.hover {
			if b.OnHover != nil {
				e = b.OnHover(elem)
			}
			b.hover = true
		}
	} else {
		if b.hover {
			if b.OnUnhover != nil {
				e = b.OnUnhover(elem)
			}
			b.hover = false
		}
	}

	return
}

// Draw a button
func (b ButtonData) Draw(ren *sdl.Renderer) (e error) {
	var err error
	var surface *sdl.Surface
	var texture *sdl.Texture
	var textRect sdl.Rect

	ren.SetDrawColor(b.BgColor.R, b.BgColor.G, b.BgColor.B, b.BgColor.A)
	ren.FillRect(&b.rect)

	surface, err = b.Font.RenderUTF8Solid(b.Text, b.TextColor)
	if err != nil {
		e = errors.New("Error rendering button text \"" + b.Text + "\": " + err.Error())
		return
	}
	texture, err = ren.CreateTextureFromSurface(surface)
	surface.Free()
	if err != nil {
		e = errors.New(
			"Error creating texture from surface of button text \"" + b.Text + "\": " + err.Error(),
		)
		return
	}
	_, _, textRect.W, textRect.H, err = texture.Query()
	if err != nil {
		e = errors.New("Error querying texture of button text \"" + b.Text + "\": " + err.Error())
		return
	}
	textRect.X = b.rect.X + b.rect.W/2 - textRect.W/2
	textRect.Y = b.rect.Y + b.rect.H/2 - textRect.H/2

	return ren.Copy(texture, nil, &textRect)
}

// Update element
func (e *Element) Update() error {
	switch e.Type {
	case Button:
		return e.Data.(*ButtonData).Update(e)
	default:
		return errors.New("Error updating element: unknown element type: " + string(e.Type))
	}
}

// Draw element
func (e *Element) Draw(ren *sdl.Renderer) error {
	switch e.Type {
	case Button:
		return e.Data.(*ButtonData).Draw(ren)
	default:
		return errors.New("Error drawing element: unknown element type: " + string(e.Type))
	}
}
