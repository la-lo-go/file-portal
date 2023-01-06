package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var extension string
	var destFolder string

	fmt.Print("Enter the file extension you want to move: .")
	fmt.Scanln(&extension)

	fmt.Print("Enter the name of the destination folder: ")
	fmt.Scanln(&destFolder)

	if _, err := os.Stat(destFolder); os.IsNotExist(err) {
		os.Mkdir(destFolder, 0755)
	}

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == "."+extension && !info.IsDir() {
			os.Rename(path, destFolder+"/"+info.Name())
		}

		return nil
	})

	fmt.Println("All files with extension", extension, "have been moved to the folder", destFolder)
	
	fmt.Println("Press 'Enter' to exit the program")
	fmt.Scanln()
}
