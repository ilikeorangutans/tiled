package main

import (
	"fmt"
	"github.com/ilikeorangutans/tiled"
	vp "github.com/ilikeorangutans/tiled/viewport"
	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
	"log"
)

func main() {

	screenWidth := 800
	screenHeight := 600

	m, err := tiled.LoadMap("map.tmx")
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

	var event sdl.Event

	viewport := vp.NewViewport(screenWidth, screenHeight, 0, 0, m.Width(), m.Height())

	dragging := false
	running := true
	for running {

		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

			switch t := event.(type) {

			case *sdl.QuitEvent:
				running = false

			case *sdl.MouseButtonEvent:
				log.Printf("Mouse button %d , state %d", t.Button, t.State)

				if t.Button == 3 && t.State == 1 {
					dragging = true
				}

				if t.Button == 3 && t.State == 0 {
					dragging = false
				}

			case *sdl.MouseMotionEvent:

				if dragging {
					viewport.MoveBy(int(-t.XRel), int(-t.YRel))

					//log.Printf("viewport %d/%d", viewport.X(), viewport.Y())
				}

			case *sdl.KeyDownEvent:

				switch t.Keysym.Sym {
				case sdl.K_q:
					running = false
				case sdl.K_LEFT:
					viewport.MoveBy(-16, 0)

				case sdl.K_RIGHT:
					viewport.MoveBy(16, 0)
				case sdl.K_UP:
					viewport.MoveBy(0, -16)
				case sdl.K_DOWN:
					viewport.MoveBy(0, 16)
				}

				//log.Printf("Viewport is now at %d/%d", viewport.X(), viewport.Y())
			}
		}

		surface.FillRect(&sdl.Rect{0, 0, 800, 600}, 0)

		renderingYoffset := int32(-(viewport.Y() % 32))
		renderingXoffset := int32(-(viewport.X() % 32))

		for layerIndex := range m.Layers() {
			layer := m.Layers()[layerIndex]

			x1, y1, x2, y2 := viewport.VisibleTiles()

			r := tiled.NewRect(x1, y1, (x2 - x1), (y2 - y1))
			//log.Printf("viewport at %d/%d, rendering tiles %d/%d -> %d/%d with offset %d/%d", viewport.X(), viewport.Y(), r.X(), r.Y(), x2, y2, renderingXoffset, renderingYoffset)
			subLayer := layer.Sub(r)

			horizontalTiles := x2 - x1
			verticalTiles := y2 - y1

			for y := 0; y < verticalTiles; y++ {
				for x := 0; x < horizontalTiles; x++ {

					p := tiled.NewPoint(x, y)
					t := subLayer.TileAt(p)

					//if y < 2 && x < 2 {
					//	log.Printf("Rendering first tile %d/%d at point %d/%d (x/y %d/%d), layer starts at %d/%d", t.X(), t.Y(), p.X(), p.Y(), x, y, subLayer.X(), subLayer.Y())
					//}

					screenX := int32(x*32) + renderingXoffset
					screenY := int32(y*32) + renderingYoffset

					rect := sdl.Rect{screenX, screenY, 31, 31}

					var color uint32 = 0xffff0000

					var red uint32 = uint32((t.X() * 2) << 8)
					var blue uint32 = uint32(t.Y() * 2)
					color = color + red + blue

					surface.FillRect(&rect, color)

					//coords = fmt.Sprintf("%d/%d", x, y)
					//s = font.RenderText_Solid(coords, sdl.Color{R: 255, G: 255, B: 255, A: 255})
					//s.Blit(&sdl.Rect{0, 0, 32, 32}, surface, &sdl.Rect{screenX, screenY + 16, 32, 16})
				}
			}
		}

		numStatus = 0
		renderStatus(fmt.Sprintf("viewport: %d/%d", viewport.X(), viewport.Y()), font, surface)
		a, b, c, d := viewport.VisibleTiles()
		renderStatus(fmt.Sprintf("tiles: %d/%d -> %d/%d", a, b, c, d), font, surface)

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
