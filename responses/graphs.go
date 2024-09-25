package responses

import "github.com/pure-project/go-pdfium/enums"

type GetPageImage struct {
	Page   int                // The page this images came from (0-index based).
	Images []GetPageImageData // A list of images of a page.
}

type GetPageImageData struct {
	Position ImagePosition
	Bitmap   ImageBitmap
}

type ImagePosition struct {
	Left   float64
	Bottom float64
	Right  float64
	Top    float64
}

type ImageBitmap struct {
	Format enums.FPDF_BITMAP_FORMAT
	Stride int
	Width  int
	Height int
	Pix    []byte
}

type GetPagePath struct {
	Page  int
	Paths []GetPagePathData // A list of path of a page
}

type GetPagePathData struct {
	Position PathPosition
	Segments []PathSegment // A list of segment of a path
}

type PathPosition = ImagePosition

type PathSegment struct {
	Type  enums.FPDF_SEGMENT
	Point PathPoint
	Close bool
}

type PathPoint struct {
	X float64
	Y float64
}
