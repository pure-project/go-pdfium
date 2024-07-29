//go:build !pdfium_experimental
// +build !pdfium_experimental

package implementation_cgo

import (
	pdfium_errors "github.com/pure-project/go-pdfium/errors"
	"github.com/pure-project/go-pdfium/requests"
	"github.com/pure-project/go-pdfium/responses"
)

// FPDF_RenderPageBitmapWithColorScheme_Start starts to render page contents to a device independent bitmap progressively with a specified color scheme for the content.
// Not supported on multi-threaded usage.
// Experimental API.
func (p *PdfiumImplementation) FPDF_RenderPageBitmapWithColorScheme_Start(request *requests.FPDF_RenderPageBitmapWithColorScheme_Start) (*responses.FPDF_RenderPageBitmapWithColorScheme_Start, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}
