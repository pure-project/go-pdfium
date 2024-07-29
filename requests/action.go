package requests

import "github.com/pure-project/go-pdfium/references"

type GetActionInfo struct {
	Document references.FPDF_DOCUMENT
	Action   references.FPDF_ACTION
}
