package single_threaded

import (
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"

	"github.com/pure-project/go-pdfium"
	"github.com/pure-project/go-pdfium/internal/implementation_cgo"
)

var singleThreadedMutex = &sync.Mutex{}

type Config struct {
	LibraryConfig *pdfium.LibraryConfig
}

// Init will initialize pdfium library.
// Every pool will keep track of its own instances and the documents that
// belong to those instances. When you close it, it will clean up the resources
// of that pool. Underwater every pool/instance uses the same mutex to ensure
// thread safety in PDFium across pools/instances/documents.
func Init(config Config) {
	singleThreadedMutex.Lock()
	defer singleThreadedMutex.Unlock()

	// Init the PDFium library.
	implementation_cgo.InitLibrary(config.LibraryConfig)
}

func Fini() {
	singleThreadedMutex.Lock()
	defer singleThreadedMutex.Unlock()

	implementation_cgo.DestroyLibrary()
}

// NewInstance will return a unique PDFium instance that keeps track of its
// own documents. When you close it, it will clean up all resources of this
// instance.
func NewInstance() (pdfium.Pdfium, error) {
	newInstance := &pdfiumInstance{
		pdfium: implementation_cgo.Pdfium.GetInstance(),
		lock:   &sync.Mutex{},
	}

	instanceRef := uuid.New()
	newInstance.instanceRef = instanceRef.String()

	return newInstance, nil
}

type pdfiumInstance struct {
	pdfium      *implementation_cgo.PdfiumImplementation
	instanceRef string
	closed      bool
	lock        *sync.Mutex
}

// Close will close the instance and will clean up the underlying PDFium resources
// by calling i.pdfium.Close().
func (i *pdfiumInstance) Close() (err error) {
	i.lock.Lock()
	defer i.lock.Unlock()

	if i.closed {
		return errors.New("instance is already closed")
	}

	defer func() {
		if panicError := recover(); panicError != nil {
			err = fmt.Errorf("panic occurred in %s: %v", "NewDocumentFromReader", panicError)
		}
	}()

	// Close underlying instance. That will close all docs.
	err = i.pdfium.Close()
	if err != nil {
		return err
	}

	// Remove references.
	i.pdfium = nil
	i.closed = true

	return nil
}

// Kill is the same as Close on single-threaded usage.
func (i *pdfiumInstance) Kill() (err error) {
	return i.Close()
}

func (i *pdfiumInstance) GetImplementation() interface{} {
	return i.pdfium
}
