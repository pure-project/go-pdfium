package requests

type GetPageImage struct {
	Page   Page
	Bitmap bool // Weather to get bitmap data.
}

type GetPagePath struct {
	Page Page
}
