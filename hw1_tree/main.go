package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	fullPath, _ := filepath.Abs(path)
	//var files []string

	err := filepath.Walk(fullPath, func(path string, info os.FileInfo, err error) error {
		formattedName := info.Name()
		if !info.IsDir() {
			formattedName += " ("
			if info.Size() == 0 {
				formattedName += "empty)"
			} else {
				formattedName += strconv.FormatInt(info.Size(), 10)+"b)"
			}
		}
		relativePath := strings.Replace(path, fullPath+"/", "", -1)
		if fullPath == relativePath{
			return nil
		}
		position := strings.Count(relativePath, "/")
		printFormattedLine(position, formattedName)
		return nil
	})
	if err != nil {
		panic(err)
	}

	return nil
}

func printFormattedLine(position int, line string, lastInGroup bool){
	var outputLine string

	if (position > 0){
		for i:=0; i<position; i++ {
			outputLine += "│	"
		}
	}
	outputLine += "├───"+line
	fmt.Println(outputLine)
}