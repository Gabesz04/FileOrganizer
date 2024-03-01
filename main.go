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
	sourcePath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	utils.SourcePath = sourcePath

	MakeDir(GetFiles())
}

func MakeDir(f []fs.DirEntry) {
	for _, file := range f {
		if GetXLastElements(3, file.Name()) == utils.FileExtension {
			if HasDir(file) {
				MoveFile(file)
			} else {
				err := os.MkdirAll("../"+GetXFirstElements(utils.CharacterCount, file.Name()), 0777)
				if err != nil {
					log.Fatal(err)
				}
				utils.DirString = GetDirNames(GetFiles())

			}

		}
	}
}

func HasDir(file fs.DirEntry) (values bool) {
	return strings.Contains(utils.DirString, GetXFirstElements(8, file.Name()))
}

func MoveFile(fileName fs.DirEntry) {
	parentPath := filepath.Dir(utils.SourcePath)
	fmt.Print(parentPath)

	//GetRightDir(fileName)
	err := os.Rename(parentPath+"\\"+fileName.Name(), parentPath+"\\"+GetXFirstElements(utils.CharacterCount, fileName.Name()))
	if err != nil {
		log.Fatal(err)
	}
}

func GetRightDir(fileName fs.DirEntry) {
	fmt.Print(fileName.Info())
}

func GetDirNames(files []fs.DirEntry) (value string) {
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
