package implementation_cgo

/*
#cgo pkg-config: pdfium
#include "fpdf_edit.h"
#include <string.h>
*/
import "C"
import (
	"errors"
	"unsafe"

	"github.com/pure-project/go-pdfium/enums"
	"github.com/pure-project/go-pdfium/requests"
	"github.com/pure-project/go-pdfium/responses"
)

// GetPageImage returns the image of a page
func (p *PdfiumImplementation) GetPageImage(request *requests.GetPageImage) (*responses.GetPageImage, error) {
	p.Lock()
	defer p.Unlock()

	pageHandle, err := p.loadPage(request.Page)
	if err != nil {
		return nil, err
	}

	var images []responses.GetPageImageData

	objCount := C.FPDFPage_CountObjects(pageHandle.handle)
	for i := 0; i < int(objCount); i++ {
		obj := C.FPDFPage_GetObject(pageHandle.handle, C.int(i))
		objType := C.FPDFPageObj_GetType(obj)
		if enums.FPDF_PAGEOBJ(int(objType)) == enums.FPDF_PAGEOBJ_IMAGE {
			bitmap := C.FPDFImageObj_GetBitmap(obj)
			format := C.FPDFBitmap_GetFormat(bitmap)
			if enums.FPDF_BITMAP_FORMAT(format) == enums.FPDF_BITMAP_FORMAT_UNKNOWN {
				continue
			}

			var left, bottom, right, top C.float
			success := C.FPDFPageObj_GetBounds(obj, &left, &bottom, &right, &top)
			if int(success) == 0 {
				return nil, errors.New("could not get image object bounds")
			}

			width := C.FPDFBitmap_GetWidth(bitmap)
			height := C.FPDFBitmap_GetHeight(bitmap)
			stride := C.FPDFBitmap_GetStride(bitmap)
			buffer := C.FPDFBitmap_GetBuffer(bitmap)

			length := stride * height
			buf := make([]byte, int(length))
			C.memcpy(unsafe.Pointer(&buf[0]), buffer, C.size_t(length))

			images = append(images, responses.GetPageImageData{
				Position: responses.ImagePosition{
					Left:   float64(left),
					Bottom: float64(bottom),
					Right:  float64(right),
					Top:    float64(top),
				},
				Bitmap: responses.ImageBitmap{
					Format: enums.FPDF_BITMAP_FORMAT(format),
					Stride: int(stride),
					Width:  int(width),
					Height: int(height),
					Pix:    buf,
				},
			})
		}
	}

	return &responses.GetPageImage{
		Page:   pageHandle.index,
		Images: images,
	}, nil
}
