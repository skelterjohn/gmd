package main

import (
	"exp/gui"
	"github.com/skelterjohn/gmd"
	"image/color"
	"image/draw"
	"time"
)

func main() {

	var dw gui.Window

	gmd.SetAppName("gmd")

	x := func() {
		w := gmd.NewWindow()
		w.SetTitle("hi!")
		w.SetSize(100, 100)
		w.Show()

		dw = w

		events := dw.EventChan()

		var s draw.Image = dw.Screen()

		go func() {
			for ei := range events {
				switch e := ei.(type) {
				case gmd.MouseDownEvent:
					println("md", e.X, e.Y, e.Button)
				case gmd.MouseUpEvent:
					println("mu", e.X, e.Y, e.Button)
				case gmd.MouseMovedEvent:
					println("mv", e.X, e.Y)
				case gmd.MouseDraggedEvent:
					println("mdr", e.X, e.Y, e.Button)
				case gmd.MouseEnteredEvent:
					println("men", e.X, e.Y)
				case gmd.MouseExitedEvent:
					println("mex", e.X, e.Y)
				case gmd.KeyDownEvent:
					println("kd", e.Letter)
				case gmd.KeyUpEvent:
					println("ku", e.Letter)
				case gmd.CloseEvent:
					println("close")
				case gmd.ResizeEvent:
					println("resize")
					s = dw.Screen()
				}
			}
		}()

		for i := 0; ; i++ {
			for x := 0; x < 100; x++ {
				for y := 0; y < 100; y++ {
					var r uint8
					if x > 50 {
						r = 255
					}
					var g uint8
					if y >= 50 {
						g = 255
					}
					var b uint8
					if y < 25 || y >= 75 {
						b = 255
					}
					if i%2 == 1 {
						r = 255 - r
					}

					if y > 90 {
						r = 255
						g = 255
						b = 255
					}

					if x == y {
						r = 100
						g = 100
						b = 100
					}

					s.Set(x, y, color.RGBA{r, g, b, 255})
				}
			}
			dw.FlushImage()
			<-time.After(1e9)
		}
	}
	go x()
	gmd.Run()
	println("done")
}
