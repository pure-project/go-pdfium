//go:build windows
// +build windows

package implementation_cgo

// #cgo pkg-config: pdfium
// #include "fpdfview.h"
import "C"
import (
	"errors"

	"github.com/pure-project/go-pdfium/requests"
	"github.com/pure-project/go-pdfium/responses"
)

// FPDF_RenderPage renders contents of a page to a device (screen, bitmap, or printer).
// This feature does not work on multi-threaded usage as you will need to give a device handle.
// Windows only!
func (p *PdfiumImplementation) FPDF_RenderPage(request *requests.FPDF_RenderPage) (*responses.FPDF_RenderPage, error) {
	p.Lock()
	defer p.Unlock()

	pageHandle, err := p.loadPage(request.Page)
	if err != nil {
		return nil, err
	}

	hdc, ok := request.DC.(C.HDC)
	if !ok {
		return nil, errors.New("DC is not of type C.HDC")
	}

	C.FPDF_RenderPage(hdc, pageHandle.handle, C.int(request.StartX), C.int(request.StartY), C.int(request.SizeX), C.int(request.SizeY), C.int(request.Rotate), C.int(request.Flags))

	return &responses.FPDF_RenderPage{}, nil
}
