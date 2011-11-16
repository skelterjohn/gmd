package gmd

// #cgo darwin LDFLAGS: -framework gomacdraw
// #include "gomacdraw/gmd.h"
// #include "stdlib.h"
import "C"

import (
	"fmt"
	"image/draw"
	"errors"
	"unsafe"
	"runtime"
)

var appChanStart = make(chan bool)
var appChanFinish = make(chan bool)

const DoItTheRightWay = false

func init() {
	if DoItTheRightWay {
		runtime.LockOSThread()
	}
	C.initMacDraw()
	if !DoItTheRightWay {
		go func() {
			runtime.LockOSThread()
			<-appChanStart
			C.NSAppRun()
			appChanFinish <- true
		}()
	}
}

func SetAppName(name string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.setAppName(cname)
}

type Window struct {
	cw _Ctypedef_GMDWindow
	im *Image
	ec chan interface{}
}

func NewWindow() (w *Window) {
	cw := C.openWindow()
	w = &Window {
		cw : cw,
	}
	return
}

func (w *Window) SetTitle(title string) {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	C.setWindowTitle(w.cw, ctitle)
}

func (w *Window) SetSize(width, height int) {
	C.setWindowSize(w.cw, _Ctype_int(width), _Ctype_int(height))
}

func (w *Window) Size() (width, height int) {
	var rw, rh _Ctype_int
	C.getWindowSize(w.cw, &rw, &rh)
	width = int(rw)
	height = int(rh)
	return
}

func (w *Window) Show() {
	C.showWindow(w.cw)
}

func (w *Window) Screen() (im draw.Image) {
	width, height := w.Size()
	ci := C.getWindowScreen(w.cw)
	gim := &Image {
		width : width,
		height : height,
		data : make([]byte, 4*width*height),
		ci : ci,
	}
	
	ptr := unsafe.Pointer(&gim.data[0])
	
	C.setScreenData(ci, ptr)
	w.im = gim
	im = gim
	return
}

func (w *Window) FlushImage() {
	C.flushWindowScreen(w.cw)
}

func (w *Window) Close() (err error) {
	ecode := C.closeWindow(w.cw)
	if ecode != 0 {
		err = errors.New(fmt.Sprintf("error:%d", ecode))
	}
	return
}

func Run() {
	if DoItTheRightWay {
		C.NSAppRun()
	} else {
		appChanStart <- true
		//C.NSAppRun()
		<-appChanFinish
	}
}

func Stop() {
	C.NSAppStop()
}
