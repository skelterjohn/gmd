/*
   Copyright 2012 John Asmuth

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"github.com/skelterjohn/gmd"
	"github.com/skelterjohn/go.wde"
	"image/color"
	"image/draw"
	"time"
	"runtime"
)

func main() {
	if false {
		runtime.LockOSThread()
	}

	var dw wde.Window

	gmd.SetAppName("gmd")

	x := func() {
		w := gmd.NewWindow()
		w.SetTitle("hi!")
		w.SetSize(100, 100)
		w.Show()

		dw = w

		events := dw.EventChan()

		var s draw.Image = dw.Screen()

		done := make(chan bool)

		go func() {
			for ei := range events {
				switch e := ei.(type) {
				case wde.MouseDownEvent:
					println("md", e.X, e.Y, e.Button)
					dw.Close()
				case wde.MouseUpEvent:
					println("mu", e.X, e.Y, e.Button)
				case wde.MouseMovedEvent:
					println("mv", e.X, e.Y)
				case wde.MouseDraggedEvent:
					println("mdr", e.X, e.Y, e.Button)
				case wde.MouseEnteredEvent:
					println("men", e.X, e.Y)
				case wde.MouseExitedEvent:
					println("mex", e.X, e.Y)
				case wde.KeyDownEvent:
					println("kd", e.Letter)
				case wde.KeyUpEvent:
					println("ku", e.Letter)
				case wde.CloseEvent:
					println("close")
				case wde.ResizeEvent:
					println("resize")
					s = dw.Screen()
				}
			}
			println("end of events")
			done <- true
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
			select {
			case <-time.After(1e9):
			case <-done:
				return
			}
		}
	}
	go x()
	gmd.Run()
	println("done")
}
