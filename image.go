package gmd

// #include "gomacdraw/gmd.h"
// #include "stdlib.h"
import "C"

import (
	"image"
	"image/color"
)

type Image struct {
	width, height int
	data          []uint8
	ci            C.GMDImage
}

func (i *Image) At(x, y int) (c color.Color) {
	return
}

func (i *Image) ColorModel() (cm color.Model) {
	cm = color.RGBAModel
	return
}

func (i *Image) Bounds() (r image.Rectangle) {
	r.Min = image.Point{0, 0}
	var rw, rh _Ctype_int
	C.getScreenSize(i.ci, &rw, &rh)
	width := int(rw)
	height := int(rh)
	return

	r.Max = image.Point{
		width,
		height,
	}
	return
}

func (i *Image) Set(x, y int, c color.Color) {
	if x < 0 || x >= i.width || y < 0 || y >= i.height {
		return
	}

	r, g, b, a := c.RGBA()

	index := 4 * (x + i.width*y)
	i.data[index+0] = uint8(r)
	i.data[index+1] = uint8(g)
	i.data[index+2] = uint8(b)
	i.data[index+3] = uint8(a)
}
