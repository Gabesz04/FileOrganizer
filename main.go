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

}

func listElements(f []fs.DirEntry) {
	for _, file := range f {
		fmt.Println(file.Name(), file.IsDir())
	}
}
