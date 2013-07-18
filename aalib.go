package aalib

/*
#cgo LDFLAGS: -L/usr/lib/x86_64-linux-gnu -laa
#include <stdlib.h>
#include <aalib.h>
// Don't use macros, use functions !!
#undef aa_putpixel
#undef aa_setpalette
#undef aa_recommendhikbd
#undef aa_recommendhimouse
#undef aa_recommendhidisplay
#undef aa_recommendlowkbd
#undef aa_recommendlowmouse
#undef aa_recommendlowdisplay
#undef aa_scrwidth
#undef aa_scrheight
#undef aa_mmwidth
#undef aa_mmheight
#undef aa_imgwidth
#undef aa_imgheight
#undef aa_image
#undef aa_text
#undef aa_attrs
*/
import "C"

import (
	"errors"
	"image"
	"image/color"
	"strings"
	"unsafe"
)

const (
	AA_NORMAL_MASK   int = C.AA_NORMAL_MASK
	AA_DIM_MASK      int = C.AA_DIM_MASK
	AA_BOLD_MASK     int = C.AA_BOLD_MASK
	AA_BOLDFONT_MASK int = C.AA_BOLDFONT_MASK
	AA_REVERSE_MASK  int = C.AA_REVERSE_MASK
	AA_ALL           int = C.AA_ALL
	AA_EIGHT         int = C.AA_EIGHT
	AA_EXTENDED      int = (C.AA_ALL | C.AA_EIGHT)
)

const (
	AA_NORMAL   int = C.AA_NORMAL
	AA_BOLD     int = C.AA_BOLD
	AA_DIM      int = C.AA_DIM
	AA_BOLDFONT int = C.AA_BOLDFONT
	AA_REVERSE  int = C.AA_REVERSE
)

type Handle struct {
	context *C.aa_context
}

type RenderParams struct {
	param *C.aa_renderparams
}

func Init(width int, height int, mask int) (*Handle, error) {
	var context *C.aa_context

	param := C.aa_defparams
	param.width = C.int(width)
	param.height = C.int(height)
	param.supported = C.int(mask)

	context = C.aa_init(&C.mem_d, &param, nil)
	if context == nil {
		return nil, errors.New("Error: aa_init")
	}

	return &Handle{context}, nil
}

func (h *Handle) Close() {
	C.aa_close(h.context)
}

func calculateBrightness(pixelColor color.Color) uint32 {
	r, g, b, _ := pixelColor.RGBA()
	var brightness uint32
	if r > g {
		if r > b {
			brightness = r
		} else {
			brightness = b
		}
	} else {
		if g > b {
			brightness = g
		} else {
			brightness = b
		}
	}

	return brightness
}

func (h *Handle) PutPixel(x int, y int, pixelColor color.Color) {
	brightness := calculateBrightness(pixelColor)
	C.aa_putpixel(h.context, C.int(x), C.int(y), C.int(brightness))
}

func (h *Handle) Puts(x int, y int, attr int, str string) {
	chars := C.CString(str)
	defer C.free(unsafe.Pointer(chars))
	C.aa_puts(h.context, C.int(x), C.int(y), uint32(attr), chars)
}

func (h *Handle) Render(rp *RenderParams, x1 int, y1 int, x2 int, y2 int) {
	var param *C.aa_renderparams
	param = C.aa_getrenderparams()
	defer C.free(unsafe.Pointer(param))
	C.aa_render(h.context, param, C.int(x1), C.int(y1), C.int(x2), C.int(y2))
}

func (handle *Handle) insertNewLine(buf string) string {
	width := handle.ScrWidth()
	height := handle.ScrHeight()

	strs := make([]string, height)
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			start := h * width
			end := start + width
			strs[h] = buf[start:end]
		}
	}

	return strings.Join(strs, "\n")
}

func (h *Handle) Resize() (int, error) {
	ret := C.aa_resize(h.context)
	if ret == 0 {
		return -1, errors.New("no resize")
	}

	return int(ret), nil
}

func (h *Handle) Flush() {
	C.aa_flush(h.context)
}

func (h *Handle) PutImage(img image.Image) {
	rect := img.Bounds()
	min, max := rect.Min, rect.Max

	for x := min.X; x < max.X; x++ {
		for y := min.Y; y < max.Y; y++ {
			pixelColor := img.At(x, y)
			h.PutPixel(x, y, pixelColor)
		}
	}
}

func (h *Handle) Text() string {
	p := C.aa_text(h.context)
	buf := C.GoString((*C.char)(unsafe.Pointer(p)))
	return h.insertNewLine(buf)
}

func (h *Handle) Attrs() string {
	p := C.aa_attrs(h.context)
	buf := C.GoString((*C.char)(unsafe.Pointer(p)))
	return h.insertNewLine(buf)
}

func (h *Handle) Image() string {
	p := C.aa_image(h.context)
	buf := C.GoString((*C.char)(unsafe.Pointer(p)))
	return h.insertNewLine(buf)
}

func (h *Handle) ScrWidth() int {
	return int(C.aa_scrwidth(h.context))
}

func (h *Handle) ScrHeight() int {
	return int(C.aa_scrheight(h.context))
}

func (h *Handle) ImgWidth() int {
	return int(C.aa_imgwidth(h.context))
}

func (h *Handle) ImgHeight() int {
	return int(C.aa_imgheight(h.context))
}
