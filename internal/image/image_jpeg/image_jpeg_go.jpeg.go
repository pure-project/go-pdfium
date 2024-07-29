package image_jpeg

import (
	"image"
	"image/jpeg"
	"io"
)

func Encode(w io.Writer, m *image.RGBA, o Options) error {
	return jpeg.Encode(w, m, o.Options)
}
