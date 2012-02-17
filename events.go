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

package gmd

// #include "gomacdraw/gmd.h"
import "C"

import (
	"fmt"
	"github.com/skelterjohn/go.wde"
)

func (w *Window) EventChan() (events <-chan interface{}) {
	ec := make(chan interface{})
	go func(ec chan<- interface{}) {
		for {
			println("polling event")
			e := C.getNextEvent(w.cw)
			println("got event")
			var ei interface{}
			switch e.kind {
			case C.GMDNoop:
				continue
			case C.GMDMouseDown:
				var mde wde.MouseDownEvent
				mde.X = int(e.data[0])
				mde.Y = int(e.data[1])
				mde.Button = int(e.data[2])
				ei = mde
			case C.GMDMouseUp:
				var mue wde.MouseUpEvent
				mue.X = int(e.data[0])
				mue.Y = int(e.data[1])
				mue.Button = int(e.data[2])
				ei = mue
			case C.GMDMouseDragged:
				var mue wde.MouseDraggedEvent
				mue.X = int(e.data[0])
				mue.Y = int(e.data[1])
				mue.Button = int(e.data[2])
				ei = mue
			case C.GMDMouseMoved:
				var me wde.MouseMovedEvent
				me.X = int(e.data[0])
				me.Y = int(e.data[1])
				ei = me
			case C.GMDMouseEntered:
				var me wde.MouseEnteredEvent
				me.X = int(e.data[0])
				me.Y = int(e.data[1])
				ei = me
			case C.GMDMouseExited:
				var me wde.MouseExitedEvent
				me.X = int(e.data[0])
				me.Y = int(e.data[1])
				ei = me
			case C.GMDKeyDown:
				var ke wde.KeyDownEvent
				ke.Letter = fmt.Sprintf("%c", e.data[0])
				ke.Code = int(e.data[1])
				ei = ke
			case C.GMDKeyUp:
				var ke wde.KeyUpEvent
				ke.Letter = fmt.Sprintf("%c", e.data[0])
				ke.Code = int(e.data[1])
				ei = ke
			case C.GMDResize:
				var re wde.ResizeEvent
				re.Width = int(e.data[0])
				re.Height = int(e.data[1])
				ei = re
			case C.GMDClose:
				ei = wde.CloseEvent{}
				ec <- ei
				close(ec)
				return
			}
			select {
			case ec <- ei:
			}
		}
	}(ec)
	events = ec
	return
}