package frame

import (
	"image"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
)

type Frame struct {
	screen.Window
	screen.Buffer
}

func Run() error {
	const W, H = 640, 480
	var frame = &Frame{}
	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(&screen.NewWindowOptions{
			Width:  W,
			Height: H,
		})
		if err != nil {
			panic(err)
		}
		b, err := s.NewBuffer(image.Point{
			X: W,
			Y: H,
		})
		if err != nil {
			panic(err)
		}
		frame.Window = w
		frame.Buffer = b
		for e := w.NextEvent(); e != nil; e = w.NextEvent() {
			switch e := e.(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}
			case paint.Event:
				w.Publish()
			}
		}
	})
	return nil
}

func (f *Frame) Upload(p []byte) {
	copy(f.Buffer.RGBA().Pix, p)
	f.Window.Upload(image.Point{}, f.Buffer, f.Buffer.Bounds())
}
