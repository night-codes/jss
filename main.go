package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/night-codes/alliance"
)

var (
	args struct {
		InputDir string `arg:"positional" help:"Input source directory (\"jss/\")."`
		Output   string `arg:"positional" help:"Output JS file (\"script.js\")."`
		Watch    string `arg:"-w" help:"Watch directory for changes."`
		Import   string `arg:"-i" help:"Specify a Scss import path."`
	}

	filelist = make(map[string]time.Time)
)

func main() {
	args.InputDir = "jss/"
	args.Output = "script.js"
	arg.MustParse(&args)

	if args.InputDir == "" || args.Output == "" {
		fmt.Println("JSS: please set input directory and output file")
		os.Exit(0)
	}

	if args.Watch == "" {
		makejs()
		os.Exit(0)
	}

	fmt.Printf("JSS started to track \"%s\"\n", args.InputDir)

	for {
		changes := false
		filepath.Walk(args.InputDir, func(path string, info os.FileInfo, err error) error {
			if err == nil && info != nil && !info.IsDir() {
				if t, ok := filelist[path]; !ok || t != info.ModTime() {
					filelist[path] = info.ModTime()
					if !ok {
						fmt.Println("Add \"" + path + "\"")
					} else {
						fmt.Println("Changed \"" + path + "\"")
					}
					changes = true
				}
			}
			return nil
		})
		if changes {
			makejs()
		}
		time.Sleep(time.Second / 10)
	}
}

func makejs() {
	if str, err := alliance.Make(args.InputDir, false); err == nil {
		err = ioutil.WriteFile(args.Output, []byte(str), 0644)
	}
}
