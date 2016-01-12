package main

import (
	"github.com/bluele/stream"
	"github.com/bluele/stream/plugins/file"
	"github.com/bluele/stream/plugins/http"
	"github.com/bluele/stream/plugins/s3"
	"github.com/mitchellh/cli"
	"log"
	"os"
)

const (
	appName = "stream"
	version = "0.1.0"
)

func main() {
	stream.InitPlugins(
		[]stream.Plugin{
			&file.Plugin{},
			&s3.Plugin{},
			&http.Plugin{},
		})

	c := cli.NewCLI(appName, version)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"ls": func() (cli.Command, error) {
			return &stream.LSCommand{}, nil
		},
		"cp": func() (cli.Command, error) {
			return &stream.CopyCommand{}, nil
		},
		"cat": func() (cli.Command, error) {
			return &stream.CatCommand{}, nil
		},
	}

	status, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(status)
}
