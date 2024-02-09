package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {
	files, err := os.ReadDir("../")
	//var path = "/test/"
	if err != nil {
		log.Fatal(err)
	}

	listElements(files)
	makeDir(files)
}

func makeDir(f []fs.DirEntry) {
	stringEnd := ""
	for _, file := range f {
		stringEnd = file.Name()[len(file.Name())-3:]
		if stringEnd == "pdf" {
			HasDir(file, f)
		}
	}
}

func HasDir(file fs.DirEntry, files []fs.DirEntry) {
	for _, f := range files {
		if GetXFirstElements(8, f) == GetXFirstElements(8, file) && f.IsDir() == false {
			fmt.Println("OTTVAGYUNK")
		}
	}
}

func GetXFirstElements(n int, f fs.DirEntry) string {
	if len(f.Name()) > n {
		return f.Name()[0:n]
	} else {
		return ""
	}
}

func GetXLastElements(n int, f fs.DirEntry) string {
	if len(f.Name()) > n {
		return f.Name()[len(f.Name())-n:]
	} else {
		return ""
	}
}

func listElements(f []fs.DirEntry) {
	for _, file := range f {
		fmt.Println(file.Name(), file.IsDir())
	}
}
