package responses

import (
	"github.com/pure-project/go-pdfium/references"
)

type DestInfo struct {
	Reference references.FPDF_DEST
	PageIndex int
}

type GetDestInfo struct {
	DestInfo DestInfo
}
