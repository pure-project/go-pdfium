package requests

import (
	"io"

	"github.com/pure-project/go-pdfium/references"
)

type SaveFlags uint32

const (
	SaveFlagIncremental    SaveFlags = 1 // Incremental.
	SaveFlagNoIncremental  SaveFlags = 2 // No Incremental.
	SaveFlagRemoveSecurity SaveFlags = 3 // Remove security.
)

type FPDF_SaveAsCopy struct {
	Flags      SaveFlags // The creating flags.
	Document   references.FPDF_DOCUMENT
	FilePath   *string   // A path to save the file to.
	FileWriter io.Writer // A writer to save the file to.
}

type FPDF_SaveWithVersion struct {
	Document    references.FPDF_DOCUMENT
	Flags       SaveFlags // The creating flags.
	FileVersion int       // The PDF file version. File version: 14 for 1.4, 15 for 1.5, ...
	FilePath    *string   // A path to save the file to.
	FileWriter  io.Writer // A writer to save the file to.
}
