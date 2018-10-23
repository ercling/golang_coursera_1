package main

import (
	"fmt"
)

func main() {
	printFormattedLine(0,"zero")
	printFormattedLine(1,"one_1")
	printFormattedLine(1,"one_2")
	printFormattedLine(2,"two")

	//out := os.Stdout
	//if !(len(os.Args) == 2 || len(os.Args) == 3) {
	//	panic("usage go run main.go . [-f]")
	//}
	//path := os.Args[1]
	//printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	//err := dirTree(out, path, printFiles)
	//if err != nil {
	//	panic(err.Error())
	//}
}

//func dirTree(out io.Writer, path string, printFiles bool) error {
//	fullPath, _ := filepath.Abs(path)
//	var files []string
//
//	err := filepath.Walk(fullPath, func(path string, info os.FileInfo, err error) error {
//		files = append(files, path)
//		return nil
//	})
//	if err != nil {
//		panic(err)
//	}
//
//	return nil
//}

func printFormattedLine(position int, line string){
	var outputLine string

	if (position > 0){
		for i:=0; i<position; i++ {
			outputLine += "│	"
		}
	}
	outputLine += "├───"+line
	fmt.Println(outputLine)
}