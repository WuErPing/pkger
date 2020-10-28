package embed

import (
	"github.com/WuErPing/pkger/here"
	"github.com/WuErPing/pkger/pkging"
)

type File struct {
	Info   *pkging.FileInfo `json:"info"`
	Here   here.Info        `json:"her"`
	Path   here.Path        `json:"path"`
	Data   []byte           `json:"data"`
	Parent here.Path        `json:"parent"`
}
