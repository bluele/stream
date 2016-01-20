package stream

import (
	"fmt"
	"io"
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

type CatCommand struct{}

func (cm *CatCommand) Help() string {
	return "cat command"
}

func (cm *CatCommand) Run(args []string) int {
	if len(args) < 1 {
		return 1
	}

	if hasStdin() {
		return cm.writeStream(args[0])
	} else {
		for _, arg := range args {
			if ret := cm.readStream(arg); ret != 0 {
				return ret
			}
		}
	}
	return 0
}

func (cm *CatCommand) writeStream(path string) int {
	dst, dstPath, ok := getPlugin(path)
	if !ok {
		return 1
	}
	if err := dst.WriteFile(NewFile("", os.Stdin), dstPath); err != nil {
		fmt.Println(err)
	}
	return 0
}

func (cm *CatCommand) readStream(path string) int {
	src, srcPath, ok := getPlugin(path)
	if !ok {
		return 1
	}
	reader, err := src.FileReader(srcPath)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	if _, err := io.Copy(os.Stdout, reader); err != nil {
		fmt.Println(err)
	}
	return 0
}

func (cm *CatCommand) Synopsis() string {
	return "Print \"Cat!\""
}
