package implementation_cgo

/*
#cgo pkg-config: pdfium
#include "fpdf_edit.h"
#include <string.h>
*/
import "C"
import (
	"unsafe"

	"github.com/pure-project/go-pdfium/enums"
	"github.com/pure-project/go-pdfium/requests"
	"github.com/pure-project/go-pdfium/responses"
)

// GetPageImage returns all the image of a page
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
			var left, bottom, right, top C.float
			success := C.FPDFPageObj_GetBounds(obj, &left, &bottom, &right, &top)
			if int(success) == 1 {
				image := responses.GetPageImageData{
					Position: responses.ImagePosition{
						Left:   float64(left),
						Bottom: float64(bottom),
						Right:  float64(right),
						Top:    float64(top),
					},
				}

				if request.Bitmap {
					bitmap := C.FPDFImageObj_GetBitmap(obj)
					format := C.FPDFBitmap_GetFormat(bitmap)
					if enums.FPDF_BITMAP_FORMAT(format) != enums.FPDF_BITMAP_FORMAT_UNKNOWN {
						width := C.FPDFBitmap_GetWidth(bitmap)
						height := C.FPDFBitmap_GetHeight(bitmap)
						stride := C.FPDFBitmap_GetStride(bitmap)
						buffer := C.FPDFBitmap_GetBuffer(bitmap)

						length := stride * height
						buf := make([]byte, int(length))
						C.memcpy(unsafe.Pointer(&buf[0]), buffer, C.size_t(length))

						image.Bitmap = responses.ImageBitmap{
							Format: enums.FPDF_BITMAP_FORMAT(format),
							Stride: int(stride),
							Width:  int(width),
							Height: int(height),
							Pix:    buf,
						}
					}
				}

				images = append(images, image)
			}
		}
	}

	return &responses.GetPageImage{
		Page:   pageHandle.index,
		Images: images,
	}, nil
}

// GetPagePath returns all the path of a page
func (p *PdfiumImplementation) GetPagePath(request *requests.GetPagePath) (*responses.GetPagePath, error) {
	p.Lock()
	defer p.Unlock()

	pageHandle, err := p.loadPage(request.Page)
	if err != nil {
		return nil, err
	}

	var paths []responses.GetPagePathData

	objCount := C.FPDFPage_CountObjects(pageHandle.handle)
	for i := 0; i < int(objCount); i++ {
		obj := C.FPDFPage_GetObject(pageHandle.handle, C.int(i))
		objType := C.FPDFPageObj_GetType(obj)
		if enums.FPDF_PAGEOBJ(int(objType)) == enums.FPDF_PAGEOBJ_PATH {
			var left, bottom, right, top C.float
			C.FPDFPageObj_GetBounds(obj, &left, &bottom, &right, &top)
			segCount := C.FPDFPath_CountSegments(obj)

			path := responses.GetPagePathData{
				Position: responses.PathPosition{
					Left:   float64(left),
					Bottom: float64(bottom),
					Right:  float64(right),
					Top:    float64(top),
				},
				Segments: make([]responses.PathSegment, 0, int(segCount)),
			}

			for j := 0; j < int(segCount); j++ {
				seg := C.FPDFPath_GetPathSegment(obj, C.int(j))
				segType := C.FPDFPathSegment_GetType(seg)
				isClose := C.FPDFPathSegment_GetClose(seg)
				var x, y C.float
				C.FPDFPathSegment_GetPoint(seg, &x, &y)

				path.Segments = append(path.Segments, responses.PathSegment{
					Type: enums.FPDF_SEGMENT(segType),
					Point: responses.PathPoint{
						X: float64(x),
						Y: float64(y),
					},
					Close: isClose == 1,
				})
			}

			paths = append(paths, path)
		}
	}

	return &responses.GetPagePath{
		Page:  pageHandle.index,
		Paths: paths,
	}, nil
}
