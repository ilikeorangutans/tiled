package input

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Keep track of the state of buttons
type Buttons struct {
	mouseButtons map[uint8]bool
	keys         map[sdl.Keycode]bool
}

func (b *Buttons) OnMouseEvent(e sdl.MouseButtonEvent) {
	if e.State == sdl.PRESSED {
		b.mouseButtons[e.Button] = true
	} else {
		delete(b.mouseButtons, e.Button)
	}

}
func (b *Buttons) OnKeyboardDownEvent(e sdl.KeyDownEvent) {
	b.onKeyEvent(e.State, e.Keysym.Sym)
}

func (b *Buttons) OnKeyboardUpEvent(e sdl.KeyUpEvent) {
	b.onKeyEvent(e.State, e.Keysym.Sym)
}

func (b *Buttons) onKeyEvent(state uint8, keycode sdl.Keycode) {
	if state == sdl.PRESSED {
		b.keys[keycode] = true
	} else {
		delete(b.keys, keycode)
	}
}

func (b *Buttons) IsMouseButtonPressed(num int) bool {
	value, ok := b.mouseButtons[uint8(num)]
	if !ok {
		return false
	}
	return value
}

func (b *Buttons) IsKeyPressed(key sdl.Keycode) bool {
	value, ok := b.keys[key]
	if !ok {
		return false
	}
	return value
}
