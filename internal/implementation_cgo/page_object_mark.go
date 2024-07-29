package implementation_cgo

// #cgo pkg-config: pdfium
// #include "fpdf_edit.h"
import "C"
import (
	"github.com/pure-project/go-pdfium/references"

	"github.com/google/uuid"
)

func (p *PdfiumImplementation) registerPageObjectMark(pageObjectMark C.FPDF_PAGEOBJECTMARK) *PageObjectMarkHandle {
	ref := uuid.New()
	handle := &PageObjectMarkHandle{
		handle:    pageObjectMark,
		nativeRef: references.FPDF_PAGEOBJECTMARK(ref.String()),
	}

	p.pageObjectMarkRefs[handle.nativeRef] = handle

	return handle
}
