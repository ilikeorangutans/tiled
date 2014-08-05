package main

import (
	"fmt"
	ev "github.com/ilikeorangutans/event"
	//obs "github.com/ilikeorangutans/event/observer"
	"github.com/ilikeorangutans/tiled"
	"github.com/ilikeorangutans/tiled/input"
	vp "github.com/ilikeorangutans/tiled/viewport"
	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
	"log"
)

type ScrollEvent struct {
	ev.Event
	DeltaX, DeltaY int
}

// Maps the given input event to application specific events. Can return nil if the given event is not mapped.
func MapInput(event sdl.Event) ev.Event {
	switch t := event.(type) {

	case *sdl.QuitEvent:
		// running = false
		return ev.NewEvent("application.quit", nil)

	case *sdl.MouseButtonEvent:

		log.Printf("mouse button event, button %d, state %d, type %d", t.Button, t.State, t.Type)
		if t.Button == 3 && t.State == 1 {
			//dragging = true
		}

		if t.Button == 3 && t.State == 0 {
			//dragging = false
		}

	case *sdl.MouseMotionEvent:

		//tx, ty, _ := viewport.ScreenToTile(int(t.X), int(t.Y))
		//mousePos = fmt.Sprintf("@%d/%d, @tile %d/%d", t.X, t.Y, tx, ty)

		//if dragging {
		//	viewport.MoveBy(int(-t.XRel), int(-t.YRel))
		//}

	case *sdl.KeyDownEvent:

		log.Printf("Key down, state %v, repeat %v, keysym mod %v, keysym sym %v, unicode %v, scancode %v", t.State, t.Repeat, t.Keysym.Mod, t.Keysym.Sym, t.Keysym.Unicode, t.Keysym.Scancode)
		switch t.Keysym.Sym {
		case sdl.K_q:
			return ev.NewEvent("application.quit", nil)

		case sdl.K_c:
			//renderTileCoords = !renderTileCoords
		case sdl.K_LEFT:
			return &ScrollEvent{DeltaX: -16, DeltaY: 0, Event: ev.NewEvent("viewport.scroll", nil)}
		case sdl.K_RIGHT:
			return &ScrollEvent{DeltaX: 16, DeltaY: 0, Event: ev.NewEvent("viewport.scroll", nil)}
		case sdl.K_UP:
			return &ScrollEvent{DeltaX: 0, DeltaY: -16, Event: ev.NewEvent("viewport.scroll", nil)}
		case sdl.K_DOWN:
			return &ScrollEvent{DeltaX: 0, DeltaY: 16, Event: ev.NewEvent("viewport.scroll", nil)}
		}
	}

	return nil

}

func HandleInput(ed *ev.EventDispatcher) {
	var event sdl.Event
	log.Println("-------------------------------")
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

		appEvent := MapInput(event)

		log.Printf("Got event %v", appEvent)

		if appEvent == nil {
			//log.Printf("Event type %v not mapped", event)
		} else {
			ed.Announce(appEvent)
		}
	}

}

func main() {

	ed := ev.NewEventDispatcher()
	ih := input.NewInputHandler(nil, ed)

	screenWidth := 800
	screenHeight := 600

	//m, err := tiled.LoadMap("map.tmx")
	m, err := tiled.LoadMap("elaborate_sample.tmx")
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Loaded map %s", m)

	window := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	surface := window.GetSurface()

	ttf.Init()
	font, err := ttf.OpenFont("/usr/share/fonts/truetype/freefont/FreeMono.ttf", 12)
	if err != nil {
		log.Panic(err)
	}

	viewport := vp.NewViewport(screenWidth, screenHeight, 0, 0, m.Width(), m.Height())
	ed.Subscribe("viewport.scroll", &ev.FuncListener{
		Callback: func(e ev.Event) bool {
			scrollEvent := e.(*ScrollEvent)
			viewport.MoveBy(scrollEvent.DeltaX, scrollEvent.DeltaY)
			return false
		},
	})

	mousePos := ""
	//dragging := false
	running := true
	ed.Subscribe("application.quit", &ev.FuncListener{
		Callback: func(e ev.Event) bool {
			running = false
			return false
		},
	})

	renderTileCoords := false
	for running {

		HandleInput(ed)

		surface.FillRect(&sdl.Rect{0, 0, 800, 600}, 0)

		renderingYoffset := int32(-(viewport.Y() % 32))
		renderingXoffset := int32(-(viewport.X() % 32))

		for layerIndex := range m.Layers() {
			layer := m.Layers()[layerIndex]

			x1, y1, x2, y2 := viewport.VisibleTiles()

			r := tiled.NewRect(x1, y1, (x2 - x1), (y2 - y1))
			subLayer := layer.Sub(r)

			horizontalTiles := x2 - x1 + 1
			verticalTiles := y2 - y1 + 1

			for y := 0; y < verticalTiles; y++ {
				for x := 0; x < horizontalTiles; x++ {

					p := tiled.NewPoint(x, y)
					t := subLayer.TileAt(p)

					screenX := int32(x*32) + renderingXoffset
					screenY := int32(y*32) + renderingYoffset

					rect := sdl.Rect{screenX, screenY, 31, 31}

					var color uint32 = 0xffff0000

					switch t.Type().Gid() {
					case 1:
						color = 0xff0000ff
					case 2:
						color = 0xffffff00
					case 3:
						color = 0xff008000
					case 4:
						color = 0xffaf23ff
					default:
						color = 0xff000000
					}
					//var red uint32 = uint32((255 - t.X() - t.Y()) << 16)
					//var green uint32 = uint32((t.X() * 2) << 8)
					//var blue uint32 = uint32(t.Y() * 2)
					//color = color + red + green + blue

					surface.FillRect(&rect, color)

					if renderTileCoords {
						coords := fmt.Sprintf("%d/%d", t.X(), t.Y())
						s := font.RenderText_Solid(coords, sdl.Color{R: 255, G: 255, B: 255, A: 255})
						s.Blit(&sdl.Rect{0, 0, 32, 32}, surface, &sdl.Rect{screenX, screenY, 32, 16})

						s = font.RenderText_Solid(fmt.Sprintf("%d", t.Type().Gid()), sdl.Color{R: 255, G: 255, B: 255, A: 255})
						s.Blit(&sdl.Rect{0, 0, 32, 32}, surface, &sdl.Rect{screenX, screenY + 16, 32, 16})
					}
				}
			}
		}

		numStatus = 0
		renderStatus(fmt.Sprintf("viewport: %d/%d", viewport.X(), viewport.Y()), font, surface)
		a, b, c, d := viewport.VisibleTiles()
		renderStatus(fmt.Sprintf("tiles: %d/%d -> %d/%d", a, b, c, d), font, surface)

		renderStatus(mousePos, font, surface)

		window.UpdateSurface()

		sdl.Delay(33)
	}

	window.Destroy()
}

var numStatus = 0

func renderStatus(msg string, font *ttf.Font, surface *sdl.Surface) {
	s := font.RenderText_Solid(msg, sdl.Color{R: 255, G: 255, B: 255, A: 255})
	s.Blit(&sdl.Rect{0, 0, 150, 32}, surface, &sdl.Rect{11, int32(11 + (16 * numStatus)), 150, 16})
	numStatus++
}

func MaxInt32(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func MinInt32(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
