package errors

import "errors"

var (
	ErrSuccess                  = errors.New("0: success")
	ErrUnknown                  = errors.New("1: unknown error")
	ErrFile                     = errors.New("2: unable to read file")
	ErrFormat                   = errors.New("3: incorrect format")
	ErrPassword                 = errors.New("4: invalid password")
	ErrSecurity                 = errors.New("5: invalid encryption")
	ErrPage                     = errors.New("6: incorrect page")
	ErrUnexpected               = errors.New("unexpected error")
	ErrExperimentalUnsupported  = errors.New("this functionality is only supported when using the pdfium_experimental build flag, see https://github.com/pure-project/go-pdfium#experimental for more information")
	ErrWindowsUnsupported       = errors.New("this functionality is Windows only")
	ErrUnsupportedOnWebassembly = errors.New("this functionality is not supported on Webassembly")
)
