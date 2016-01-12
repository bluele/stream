package stream

import (
	"fmt"
	"os"
	"path/filepath"
)

type LSCommand struct{}

func (cm *LSCommand) Help() string {
	return "ls command"
}

func (cm *LSCommand) Run(args []string) int {
	if len(args) == 0 {
		return 1
	}
	if plugin, path, ok := getPlugin(args[0]); ok {
		plugin.List(path, os.Stdout)
		return 0
	}
	return 1
}

func (cm *LSCommand) Synopsis() string {
	return "Print \"LS!\""
}

type CopyCommand struct{}

func (cm *CopyCommand) Help() string {
	return "cp command"
}

func (cm *CopyCommand) Run(args []string) int {
	if len(args) < 2 {
		return 1
	}
	src, srcPath, ok := getPlugin(args[0])
	if !ok {
		return 1
	}
	dst, dstPath, ok := getPlugin(args[1])
	if !ok {
		return 1
	}
	if dstSize := len(dstPath); dstSize > 0 && dstPath[dstSize-1] == '/' {
		dstPath += filepath.Base(srcPath)
	}

	fmt.Printf("cp %v to %v\n", srcPath, dstPath)

	reader, err := src.FileReader(srcPath)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	if err = dst.WriteFile(reader, dstPath); err != nil {
		fmt.Println(err)
	}

	return 0
}

func (cm *CopyCommand) Synopsis() string {
	return "Print \"Copy!\""
}

func getPlugin(name string) (Plugin, string, bool) {
	scheme, path, err := parseProtocol(name)
	if err != nil {
		return nil, "", false
	}
	pg, ok := plugins[scheme]
	if !ok {
		return nil, "", false
	}
	return pg, path, true
}
