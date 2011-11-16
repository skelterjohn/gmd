package gmd

// #include "gomacdraw/gmd.h"
import "C"

import (
	"fmt"
)

type Event int

type MouseEvent struct {
	Event
	X, Y int
}

type MouseMovedEvent MouseEvent

type MouseButtonEvent struct {
	MouseEvent
	Button int
}

type MouseDownEvent MouseButtonEvent
type MouseUpEvent MouseButtonEvent
type MouseDraggedEvent MouseButtonEvent

type MouseEnteredEvent MouseEvent
type MouseExitedEvent MouseEvent

type KeyEvent struct {
	Code int
	Letter string
}

type KeyDownEvent KeyEvent
type KeyUpEvent KeyEvent
type KeyPressEvent KeyEvent

type ResizeEvent struct {
	Width, Height int
}

type CloseEvent struct {}

func (w *Window) EventChan() (events <-chan interface{}) {
	ec := make(chan interface{})
	go func() {
		for {
			println("polling event")
			e := C.getNextEvent(w.cw)
			println("got event")
			var ei interface{}
			switch e.kind {
			case C.GMDNoop:
				continue
			case C.GMDMouseDown:
				var mde MouseDownEvent
				mde.X = int(e.data[0])
				mde.Y = int(e.data[1])
				mde.Button = int(e.data[2])
				ei = mde
			case C.GMDMouseUp:
				var mue MouseUpEvent
				mue.X = int(e.data[0])
				mue.Y = int(e.data[1])
				mue.Button = int(e.data[2])
				ei = mue
			case C.GMDMouseDragged:
				var mue MouseDraggedEvent
				mue.X = int(e.data[0])
				mue.Y = int(e.data[1])
				mue.Button = int(e.data[2])
				ei = mue
			case C.GMDMouseMoved:
				var me MouseMovedEvent
				me.X = int(e.data[0])
				me.Y = int(e.data[1])
				ei = me
			case C.GMDMouseEntered:
				var me MouseEnteredEvent
				me.X = int(e.data[0])
				me.Y = int(e.data[1])
				ei = me
			case C.GMDMouseExited:
				var me MouseExitedEvent
				me.X = int(e.data[0])
				me.Y = int(e.data[1])
				ei = me
			case C.GMDKeyDown:
				var ke KeyDownEvent
				ke.Letter = fmt.Sprintf("%c", e.data[0])
				ke.Code = int(e.data[1])
				ei = ke
			case C.GMDKeyUp:
				var ke KeyUpEvent
				ke.Letter = fmt.Sprintf("%c", e.data[0])
				ke.Code = int(e.data[1])
				ei = ke
			case C.GMDResize:
				var re ResizeEvent
				re.Width = int(e.data[0])
				re.Height = int(e.data[1])
				ei = re
			case C.GMDClose:
				ei = CloseEvent{}
			}
			select {
			case ec <- ei:
			}
		}
	}()
	events = ec
	return
}