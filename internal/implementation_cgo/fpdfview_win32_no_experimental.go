//go:build windows && !pdfium_experimental
// +build windows,!pdfium_experimental

package implementation_cgo

import (
	pdfium_errors "github.com/pure-project/go-pdfium/errors"

	"github.com/pure-project/go-pdfium/requests"
	"github.com/pure-project/go-pdfium/responses"
)

// FPDF_SetPrintMode sets printing mode when printing on Windows.
// Experimental API.
// Windows only!
func (p *PdfiumImplementation) FPDF_SetPrintMode(request *requests.FPDF_SetPrintMode) (*responses.FPDF_SetPrintMode, error) {
	return nil, pdfium_errors.ErrExperimentalUnsupported
}
