package input

import (
	ev "github.com/ilikeorangutans/event"
	"github.com/veandco/go-sdl2/sdl"
)

type InputMapper interface {
	MapEvent(sdl.Event) ev.Event
}

func NewInputHandler(inputMapper InputMapper, ed *ev.EventDispatcher) *InputHandler {

	return &InputHandler{
		buttons:         &Buttons{},
		inputMapper:     inputMapper,
		eventDispatcher: ed,
	}
}

type InputHandler struct {
	eventDispatcher *ev.EventDispatcher
	inputMapper     InputMapper
	buttons         *Buttons
}

func (ih *InputHandler) HandleInput() ev.Event {
	var event sdl.Event

	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

		mappedEvent := ih.inputMapper.MapEvent(event)
		ih.eventDispatcher.Announce(mappedEvent)
	}
	return nil
}
func (ih *InputHandler) SwitchMapper(im InputMapper) {
	ih.inputMapper = im
}
