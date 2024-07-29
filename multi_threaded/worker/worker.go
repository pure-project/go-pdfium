package worker

import (
	"github.com/pure-project/go-pdfium"
	"github.com/pure-project/go-pdfium/internal/implementation_cgo"
)

func StartWorker(config *pdfium.LibraryConfig) {
	implementation_cgo.StartPlugin(config)
}
