package util

import (
	"io"
)

const(
	LogagentPort = "8090"
)

type ReaderCloserGetter interface {
	GetReaderCloser() (io.ReadCloser,error)
}
