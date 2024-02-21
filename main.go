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
	utils.DirString = GetDirNames(GetFiles())

	MakeDir(GetFiles())
}

func MakeDir(f []fs.DirEntry) {
	for _, file := range f {
		if GetXLastElements(3, file.Name()) == utils.FileExtension {
			if HasDir(file) {
				MoveFile(file)
			} else {
				err := os.Mkdir(GetXFirstElements(utils.CharacterCount, file.Name()), 0755)
				if err != nil {
					log.Fatal(err)
				}
				utils.DirString = GetDirNames(GetFiles())

			}

		}
	}
}

func HasDir(file fs.DirEntry) (values bool) {
	if strings.Contains(utils.DirString, GetXFirstElements(8, file.Name())) {
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
}

func GetDirNames(files []fs.DirEntry) (value string) {
	files, err := os.ReadDir("../")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			value += file.Name()
		}
	}
	return
}

func GetFiles() (value []fs.DirEntry) {
	files, err := os.ReadDir("../")
	//var path = "/test/"
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func GetXFirstElements(n int, f string) (value string) {
	if len(f) > n {
		value = f[0:n]
	}
	return value
}

func GetXLastElements(n int, f string) (value string) {
	if len(f) > n {
		value = f[len(f)-n:]
	}
	return
}

func ListElements(f []fs.DirEntry) {
	for _, file := range f {
		fmt.Println(file.Name(), file.IsDir())
	}
}
