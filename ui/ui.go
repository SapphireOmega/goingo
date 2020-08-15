package ui

import (
	"errors"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Text
type Text struct {
	Text  string
	Size  float64
	Color sdl.Color
	Font *ttf.Font

	texture *sdl.Texture
	rect     sdl.Rect
}

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
	Text    Text
	Color   sdl.Color

	OnPress   ButtonAction
	OnRelease ButtonAction
	OnHover   ButtonAction
	OnUnhover ButtonAction

	rect sdl.Rect

	pressed bool
	hover   bool
}

func (t *Text) Render(ren *sdl.Renderer) error {
	var err error
	var surface *sdl.Surface

	surface, err = t.Font.RenderUTF8Solid(t.Text, t.Color)
	if err != nil {
		return errors.New("Error rendering button text \"" + t.Text + "\": " + err.Error())
	}
	t.texture, err = ren.CreateTextureFromSurface(surface)
	surface.Free()
	if err != nil {
		return errors.New(
			"Error creating texture from surface of button text \"" + t.Text + "\": " + err.Error(),
		)
	}
	_, _, t.rect.W, t.rect.H, err = t.texture.Query()
	if err != nil {
		return errors.New("Error querying texture of button text \"" + t.Text + "\": " + err.Error())
	}
	t.rect.W = int32(float64(t.rect.W) * t.Size)
	t.rect.H = int32(float64(t.rect.H) * t.Size)

	return nil
}

func (t Text) Draw(ren *sdl.Renderer) error {
	ren.SetDrawColor(t.Color.R, t.Color.G, t.Color.B, t.Color.A)
	err := ren.Copy(t.texture, nil, &t.rect)
	if err != nil {
		return errors.New("Error drawing text: " + err.Error())
	}

	return nil
}

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
func (b *ButtonData) Draw(ren *sdl.Renderer) error {
	ren.SetDrawColor(b.Color.R, b.Color.G, b.Color.B, b.Color.A)
	ren.FillRect(&b.rect)

	b.Text.rect.X = b.rect.X + b.rect.W/2 - b.Text.rect.W/2
	b.Text.rect.Y = b.rect.Y + b.rect.H/2 - b.Text.rect.H/2

	return b.Text.Draw(ren)
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
