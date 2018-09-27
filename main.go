package main

import (
	"fmt"
	"github.com/xyproto/purefunction"
	"os"
	"path/filepath"
	"strings"
)

const versionString = "pure 1.0.0"

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Pass one or more go source filenames as arguments.")
		os.Exit(1)
	} else if len(os.Args) == 2 && os.Args[1] == "--help" {
		fmt.Println("Can list all pure functions, given one or more Go source files")
		os.Exit(0)
	} else if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Println(versionString)
		os.Exit(0)
	}

	for _, filename := range os.Args[1:] {
		basefn := filepath.Base(filename)
		fns, err := purefunction.PureFunctions(filename)
		if err != nil {
			fmt.Println(basefn + ": " + err.Error())
			os.Exit(1)
		}
		if len(fns) == 0 {
			fmt.Println(basefn + ": no pure functions")
			continue
		}
		fmt.Println(basefn + ": " + strings.Join(fns, ", "))
	}
}
