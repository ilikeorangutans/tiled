package input

import (
	"github.com/stretchr/testify/assert"
	"github.com/veandco/go-sdl2/sdl"
	"testing"
)

func TestMouseButtons(t *testing.T) {

	b := &Buttons{
		mouseButtons: make(map[uint8]bool),
	}

	mouseEvent := sdl.MouseButtonEvent{
		Type:   sdl.MOUSEBUTTONDOWN,
		Button: 1,
		State:  sdl.PRESSED,
	}

	b.OnMouseEvent(mouseEvent)
	assert.True(t, b.IsMouseButtonPressed(1), "First button should be pressed")
	assert.False(t, b.IsMouseButtonPressed(2), "Second button should not be pressed")

	mouseEvent.State = 0

	b.OnMouseEvent(mouseEvent)
	assert.False(t, b.IsMouseButtonPressed(1), "First button should be released")

}

func TestKeyboard(t *testing.T) {

	b := &Buttons{
		keys: make(map[sdl.Keycode]bool),
	}

	event := sdl.KeyDownEvent{
		Type: sdl.KEYDOWN,
		Keysym: sdl.Keysym{
			Sym: sdl.K_q,
		},
		State:  sdl.PRESSED,
		Repeat: 0,
	}

	b.OnKeyboardDownEvent(event)
	assert.True(t, b.IsKeyPressed(sdl.K_q), "First button should be pressed")

	event.State = 0

	b.OnKeyboardDownEvent(event)
	assert.False(t, b.IsKeyPressed(sdl.K_q), "First button should be released")

}
