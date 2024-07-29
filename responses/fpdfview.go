package responses

import (
	"github.com/pure-project/go-pdfium/enums"
	"github.com/pure-project/go-pdfium/references"
	"github.com/pure-project/go-pdfium/structs"
)

type FPDF_LoadDocument struct {
	Document references.FPDF_DOCUMENT
}

type FPDF_LoadMemDocument struct {
	Document references.FPDF_DOCUMENT
}

type FPDF_LoadMemDocument64 struct {
	Document references.FPDF_DOCUMENT
}

type FPDF_LoadCustomDocument struct {
	Document references.FPDF_DOCUMENT
}

type FPDF_CloseDocument struct{}

type FPDF_GetLastErrorError int

const (
	FPDF_GetLastErrorErrorSuccess        FPDF_GetLastErrorError = 0 // Error code: Success, which means no error.
	FPDF_GetLastErrorErrorUnknown        FPDF_GetLastErrorError = 1 // Error code: Unknown error.
	FPDF_GetLastErrorErrorFile           FPDF_GetLastErrorError = 2 // Error code: File access error, which means file cannot be found or be opened.
	FPDF_GetLastErrorErrorFormat         FPDF_GetLastErrorError = 3 // Error code: Data format error.
	FPDF_GetLastErrorErrorPassword       FPDF_GetLastErrorError = 4 // Error code: Incorrect password error.
	FPDF_GetLastErrorErrorSecurity       FPDF_GetLastErrorError = 5 // Error code: Unsupported security scheme error.
	FPDF_GetLastErrorErrorInvalidLicense FPDF_GetLastErrorError = 6 // Error code: License authorization error.
)

type FPDF_GetLastError struct {
	Error FPDF_GetLastErrorError
}

type FPDF_SetSandBoxPolicy struct{}

type FPDF_LoadPage struct {
	Page references.FPDF_PAGE
}

type FPDF_ClosePage struct{}

type FPDF_GetFileVersion struct {
	FileVersion int // The numeric version of the file: 14 for 1.4, 15 for 1.5, ...
}

type FPDF_GetDocPermissions struct {
	DocPermissions                      uint32 // A 32-bit integer which indicates the permission flags. Please refer to "TABLE 3.20 User access permissions" in PDF Reference 1.7 P123 for detailed description. If the document is not protected, 0xffffffff (4294967295) will be returned.
	PrintDocument                       bool   // Bit position 3: (Security handlers of revision 2) Print the document, (Security handlers of revision 3 or greater) Print the document (possibly not at the highest quality level, depending on whether PrintDocumentAsFaithfulDigitalCopy (bit 12) is also set).
	ModifyContents                      bool   // Bit position 4: Modify the contents of the document by operations other than those controlled by AddOrModifyTextAnnotations (bit 6), FillInExistingInteractiveFormFields (bit 9), and AssembleDocument (bit 11).
	CopyOrExtractText                   bool   // Bit position 5: (Security handlers of revision 2) Copy or otherwise extract  text and graphics from the document, including extracting text and graphics (in support of accessibility to users with disabilities or for other purposes). (Security handlers of revision 3 or greater) Copy or otherwise extract text and graphics from the document by operations other than that controlled by ExtractTextAndGraphics (bit 10).
	AddOrModifyTextAnnotations          bool   // Bit position 6: Add or modify text annotations
	FillInInteractiveFormFields         bool   // Bit position 6: fill in interactive form fields
	CreateOrModifyInteractiveFormFields bool   // Bit position 6 & 4: create or modify interactive form fields (including signature fields).
	FillInExistingInteractiveFormFields bool   // Bit position 9: (Security handlers of revision 3 or greater) Fill in existing interactive form fields (including signature fields), even if FillInInteractiveFormFields (bit 6) is clear.
	ExtractTextAndGraphics              bool   // Bit position 10: (Security handlers of revision 3 or greater) Extract text and graphics (in support of accessibility to users with disabilities or for other purposes).
	AssembleDocument                    bool   // Bit position 11: (Security handlers of revision 3 or greater) Assemble the  document (insert, rotate, or delete pages and create bookmarks or thumbnail images), even if ModifyContents (bit 4) is clear.
	PrintDocumentAsFaithfulDigitalCopy  bool   // Bit position 12: (Security handlers of revision 3 or greater) Print the document to a representation from which a faithful digital copy of the PDF content could be generated. When this bit is clear (and PrintDocument (bit 3) is set), printing is limited to a low-level representation of the appearance, possibly of degraded quality.
}

type FPDF_GetDocUserPermissions struct {
	DocUserPermissions                  uint32 // A 32-bit integer which indicates the permission flags. Please refer to "TABLE 3.20 User access permissions" in PDF Reference 1.7 P123 for detailed description. If the document is not protected, 0xffffffff (4294967295) will be returned.
	PrintDocument                       bool   // Bit position 3: (Security handlers of revision 2) Print the document, (Security handlers of revision 3 or greater) Print the document (possibly not at the highest quality level, depending on whether PrintDocumentAsFaithfulDigitalCopy (bit 12) is also set).
	ModifyContents                      bool   // Bit position 4: Modify the contents of the document by operations other than those controlled by AddOrModifyTextAnnotations (bit 6), FillInExistingInteractiveFormFields (bit 9), and AssembleDocument (bit 11).
	CopyOrExtractText                   bool   // Bit position 5: (Security handlers of revision 2) Copy or otherwise extract  text and graphics from the document, including extracting text and graphics (in support of accessibility to users with disabilities or for other purposes). (Security handlers of revision 3 or greater) Copy or otherwise extract text and graphics from the document by operations other than that controlled by ExtractTextAndGraphics (bit 10).
	AddOrModifyTextAnnotations          bool   // Bit position 6: Add or modify text annotations
	FillInInteractiveFormFields         bool   // Bit position 6: fill in interactive form fields
	CreateOrModifyInteractiveFormFields bool   // Bit position 6 & 4: create or modify interactive form fields (including signature fields).
	FillInExistingInteractiveFormFields bool   // Bit position 9: (Security handlers of revision 3 or greater) Fill in existing interactive form fields (including signature fields), even if FillInInteractiveFormFields (bit 6) is clear.
	ExtractTextAndGraphics              bool   // Bit position 10: (Security handlers of revision 3 or greater) Extract text and graphics (in support of accessibility to users with disabilities or for other purposes).
	AssembleDocument                    bool   // Bit position 11: (Security handlers of revision 3 or greater) Assemble the  document (insert, rotate, or delete pages and create bookmarks or thumbnail images), even if ModifyContents (bit 4) is clear.
	PrintDocumentAsFaithfulDigitalCopy  bool   // Bit position 12: (Security handlers of revision 3 or greater) Print the document to a representation from which a faithful digital copy of the PDF content could be generated. When this bit is clear (and PrintDocument (bit 3) is set), printing is limited to a low-level representation of the appearance, possibly of degraded quality.
}
type FPDF_GetSecurityHandlerRevision struct {
	SecurityHandlerRevision int // The revision number of security handler. Please refer to key "R" in "TABLE 3.19 Additional encryption dictionary entries for the standard security handler" in PDF Reference 1.7 P122 for detailed description. If the document is not protected, -1 will be returned.
}

type FPDF_GetPageCount struct {
	PageCount int // The amount of pages of the document.
}

type FPDF_GetPageWidth struct {
	Page  int     // The page this size came from (0-index based).
	Width float64 // The width of the page in points. One point is 1/72 inch (around 0.3528 mm).
}

type FPDF_GetPageHeight struct {
	Page   int     // The page this size came from (0-index based).
	Height float64 // The height of the page in points. One point is 1/72 inch (around 0.3528 mm).
}

type FPDF_GetPageSizeByIndex struct {
	Page   int     // The page this size came from (0-index based).
	Width  float64 // The width of the page in points. One point is 1/72 inch (around 0.3528 mm).
	Height float64 // The height of the page in points. One point is 1/72 inch (around 0.3528 mm).
}

type FPDF_DocumentHasValidCrossReferenceTable struct {
	DocumentHasValidCrossReferenceTable bool
}

type FPDF_GetTrailerEnds struct {
	TrailerEnds []int
}

type FPDF_GetPageWidthF struct {
	PageWidth float32
}

type FPDF_GetPageHeightF struct {
	PageHeight float32
}

type FPDF_GetPageBoundingBox struct {
	Rect structs.FPDF_FS_RECTF
}

type FPDF_GetPageSizeByIndexF struct {
	Size structs.FPDF_FS_SIZEF
}

type FPDF_RenderPageBitmap struct{}

type FPDF_RenderPageBitmapWithMatrix struct{}

type FPDF_DeviceToPage struct {
	PageX float64
	PageY float64
}

type FPDF_PageToDevice struct {
	DeviceX int
	DeviceY int
}

type FPDFBitmap_Create struct {
	Bitmap references.FPDF_BITMAP
}

type FPDFBitmap_CreateEx struct {
	Bitmap references.FPDF_BITMAP
}

type FPDFBitmap_GetFormat struct {
	Format enums.FPDF_BITMAP_FORMAT
}

type FPDFBitmap_FillRect struct{}

type FPDFBitmap_GetBuffer struct {
	Buffer []byte
}

type FPDFBitmap_GetWidth struct {
	Width int
}

type FPDFBitmap_GetHeight struct {
	Height int
}

type FPDFBitmap_GetStride struct {
	Stride int
}

type FPDFBitmap_Destroy struct{}

type FPDF_VIEWERREF_GetPrintScaling struct {
	PreferPrintScaling bool
}

type FPDF_VIEWERREF_GetNumCopies struct {
	NumCopies int
}

type FPDF_VIEWERREF_GetPrintPageRange struct {
	PageRange references.FPDF_PAGERANGE
}

type FPDF_VIEWERREF_GetPrintPageRangeCount struct {
	Count uint64
}

type FPDF_VIEWERREF_GetPrintPageRangeElement struct {
	Value int
}

type FPDF_VIEWERREF_GetDuplex struct {
	DuplexType enums.FPDF_DUPLEXTYPE
}

type FPDF_VIEWERREF_GetName struct {
	Value string
}

type FPDF_CountNamedDests struct {
	Count uint64
}

type FPDF_GetNamedDestByName struct {
	Dest references.FPDF_DEST
}

type FPDF_GetNamedDest struct {
	Dest references.FPDF_DEST
	Name string
}

type FPDF_GetXFAPacketCount struct {
	Count int
}

type FPDF_GetXFAPacketName struct {
	Index int
	Name  string
}

type FPDF_GetXFAPacketContent struct {
	Index   int
	Content []byte
}

type FPDF_SetPrintMode struct {
}

type FPDF_RenderPage struct {
}
