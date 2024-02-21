package main

import (
	"fmt"
	"io/fs"
	"log"
	"main/utils"
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
		if GetXLastElements(3, file) == utils.FileExtension {
			HasDir(file, f)
		}
	}
}

func HasDir(file fs.DirEntry, files []fs.DirEntry) {
	for _, f := range files {
		if GetXFirstElements(utils.CharacterCount, file) == GetXFirstElements(utils.CharacterCount, f) && !f.IsDir() && file.IsDir() {
			err := os.Mkdir("../"+GetXFirstElements(utils.CharacterCount, f), 0755)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			MoveFile(f, GetXFirstElements(utils.CharacterCount, f))
		}
	}
}

func MoveFile(what fs.DirEntry, where string) {
	sourcePath, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}
	parentPath := filepath.Dir(sourcePath)
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
