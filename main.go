package main

import (
	"fmt"
	"io/fs"
	"log"
	"main/utils"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	files, err := os.ReadDir("../")
	//var path = "/test/"
	if err != nil {
		log.Fatal(err)
	}
	utils.DirString = GetDirNames(files)
	ListElements(files)
	MakeDir(files)
}

func MakeDir(f []fs.DirEntry) {
	for _, file := range f {
		if GetXLastElements(3, file) == utils.FileExtension {
			if HasDir(file) {
				MoveFile()
			} else {
				MakeDir()
			}

		}
	}
}

func HasDir(file fs.DirEntry) (values bool) {
	if strings.Contains(utils.DirString, GetXFirstElements(8, file)) {
		return true
	}
	return false
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

func GetDirNames(files []fs.DirEntry) (value string) {
	for _, file := range files {
		if file.IsDir() {
			value += file.Name()
		}
	}
	return
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
