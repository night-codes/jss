package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/night-codes/alliance"
)

var (
	inputDir   = flag.String("i", "jss", "Input source directory")
	outputFile = flag.String("o", "script.js", "Output JS file.")
	filelist   = make(map[string]time.Time)
)

func main() {
	flag.Parse()
	if *inputDir == "" || *outputFile == "" {
		fmt.Println("JSS: please set input directory and output file")
		flag.PrintDefaults()
		os.Exit(0)
	}

	fmt.Printf("JSS started to track \"%s\"\n", *inputDir)

	for {
		changes := false
		filepath.Walk(*inputDir, func(path string, info os.FileInfo, err error) error {
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
	if str, err := alliance.Make(*inputDir, false); err == nil {
		err = ioutil.WriteFile(*outputFile, []byte(str), 0644)
	}
}
