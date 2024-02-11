package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	files, err := os.ReadDir("../")
	//var path = "/test/"
	if err != nil {
		log.Fatal(err)
	}

	ListElements(files)
	MakeDir(files)
}

func MakeDir(f []fs.DirEntry) {
	for _, file := range f {
		if GetXLastElements(3, file) == "pdf" {
			HasDir(file, f)
		}
	}
}

func HasDir(file fs.DirEntry, files []fs.DirEntry) {
	for _, f := range files {
		if GetXFirstElements(8, file) == GetXFirstElements(8, f) && file.IsDir() == false && file.IsDir() == true {
			fmt.Println("asd")
			err := os.Mkdir("../"+GetXFirstElements(8, f), 0755)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			MoveFile(f, GetXFirstElements(8, f))
		}
	}
}

func MoveFile(what fs.DirEntry, where string) {
	sourcePath, err := os.Getwd()
	parentPath := filepath.Dir(sourcePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(parentPath)
	//err := os.Mkdir(GetXFirstElements(), 0755)
}

func GetXFirstElements(n int, f fs.DirEntry) (value string) {
	value = ""
	if len(f.Name()) > n {
		value = f.Name()[0:n]
	}
	return
}

func GetXLastElements(n int, f fs.DirEntry) (value string) {
	value = ""
	if len(f.Name()) > n {
		value = f.Name()[len(f.Name())-n:]
	}
	return
}

func ListElements(f []fs.DirEntry) {
	for _, file := range f {
		fmt.Println(file.Name(), file.IsDir())
	}
}
