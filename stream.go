package stream

import (
	"fmt"
	"io"
	"net/url"
)

func Print(msg string) {
	fmt.Println(msg)
}

type Plugin interface {
	Init() error
	Name() string

	FileReader(string) (*File, error)
	DirReader(string) (*Dir, error)
	WriteFile(*File, string) error
	WriteDir(*Dir, string) error
	List(string, io.Writer) error
}

var plugins map[string]Plugin

func InitPlugins(pgs []Plugin) {
	plugins = make(map[string]Plugin)
	for _, pg := range pgs {
		if err := pg.Init(); err != nil {
			panic(err)
		}
		plugins[pg.Name()] = pg
	}
}

func parseProtocol(path string) (string, string, error) {
	u, err := url.Parse(path)
	if err != nil {
		return "", "", err
	}
	return u.Scheme, u.Host + u.Path, nil
}

type Dir struct {
	Dirs  []*Dir
	Files []*File
}

type File struct {
	Name string
	io.ReadCloser
}

func NewFile(name string, ir io.ReadCloser) *File {
	return &File{name, ir}
}
