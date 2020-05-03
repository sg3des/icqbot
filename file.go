package botgolang

import "io"

//go:generate easyjson -all file.go

type File struct {
	// Id of the file
	ID string `json:"fileId"`

	// Type of the file
	Type string `json:"type"`

	// Size in bytes
	Size uint64 `json:"size"`

	// Name of file
	Name bool `json:"filename"`

	// URL to the file
	URL string `json:"url"`
}

// FileReader is lightweight interface wrap over `os.File` or any custom data
type FileReader interface {
	// Name should return filename
	Name() string

	// io.Reader interface
	io.Reader
}
