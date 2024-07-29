package requests

import "github.com/pure-project/go-pdfium/references"

type GetDestInfo struct {
	Document references.FPDF_DOCUMENT
	Dest     references.FPDF_DEST
}
