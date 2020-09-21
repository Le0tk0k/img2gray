package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Le0tk0k/img2gray"
)

var rm = flag.Bool("r", false, "Remove sorce file")

func getFileNameWithoutExt(file string) string {
	return file[:len(file)-len(filepath.Ext(file))]
}

func main() {
	flag.Parse()
	src := flag.Arg(0)
	dst := getFileNameWithoutExt(src) + "_gray" + filepath.Ext(src)

	err := img2gray.ToGray(src, dst, *rm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}
