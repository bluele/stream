package file

import (
	"github.com/bluele/stream"
	"io"
	"os"
	"path/filepath"
)

type Plugin struct{}

func (pl *Plugin) Init() error {
	return nil
}

func (pl *Plugin) Name() string {
	return ""
}

func (pl *Plugin) FileReader(path string) (*stream.File, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return stream.NewFile(filepath.Base(path), fp), nil
}

func (pl *Plugin) DirReader(path string) (*stream.Dir, error) {
	return nil, nil
}

func (pl *Plugin) WriteFile(fi *stream.File, path string) error {
	info, err := os.Stat(path)
	if err == nil && info.IsDir() {
		path += "/" + fi.Name
	} else {
		os.MkdirAll(filepath.Dir(path), 0755)
	}
	iw, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = io.Copy(iw, fi)
	return err
}

func (pl *Plugin) WriteDir(di *stream.Dir, path string) error {
	return nil
}

func (pl *Plugin) List(path string, iw io.Writer) error {
	return nil
}
