package http

import (
	"github.com/bluele/stream"
	"io"
	"net/http"
	"path/filepath"
)

type Plugin struct{}

func (pl *Plugin) Init() error {
	return nil
}

func (pl *Plugin) Name() string {
	return "http"
}

func (pl *Plugin) Names() []string {
	return []string{"http", "https"}
}

func (pl *Plugin) FileReader(path string) (*stream.File, error) {
	// TODO allow multiple protocol names.
	rawurl := "http://" + path
	res, err := http.Get(rawurl)
	if err != nil {
		return nil, err
	}
	return stream.NewFile(filepath.Base(path), res.Body), nil
}

func (pl *Plugin) DirReader(path string) (*stream.Dir, error) {
	return nil, nil
}

func (pl *Plugin) WriteFile(fi *stream.File, path string) error {
	return nil
}

func (pl *Plugin) WriteDir(di *stream.Dir, path string) error {
	return nil
}

func (pl *Plugin) List(path string, iw io.Writer) error {
	return nil
}
