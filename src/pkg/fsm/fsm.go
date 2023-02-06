package fs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type FileSystemManager struct {
	Procedure string
}

func (fsm FileSystemManager) check(e error) {
	if e != nil { // if error
		// stop execution, run any deffered functions, non-zero exit
		panic(e)
	}
}

func (fsm FileSystemManager) GetFilePaths(dirPath string) []string {
	files := []string{}

	filepath.Walk(dirPath, func(path string, f os.FileInfo, err error) error {

		f, err = os.Stat(path)

		// If no error
		if err != nil {
			log.Fatal(err)
		}

		// File & Folder Mode
		f_mode := f.Mode()

		// Make sure item is file
		if f_mode.IsRegular() {

			// Print File Contents
			// fsm.PrintFileContents(path)

			// Append to Files Array
			files = append(files, path)
		}

		return nil
	})

	return files
}

func (fsm FileSystemManager) GetFileContents(filePath string) []byte {
	dat, err := os.ReadFile(filePath)
	fsm.check(err)
	return dat
}

func (fsm FileSystemManager) PrintFileContents(filePath string) {
	dat, err := os.ReadFile(filePath)
	fsm.check(err)
	fmt.Println("File Contents:", filePath)
	fmt.Println(string(dat))
}

func (fsm FileSystemManager) WriteFiles(filePaths []string, outputDir string) {
	for _, filePath := range filePaths {
		var dat []byte = fsm.GetFileContents(filePath)
		var fileName string = filepath.Base(filePath)
		var outFilePath string = filepath.Join(outputDir, fileName)
		err := os.WriteFile(outFilePath, dat, 0644)
		fsm.check(err)
	}
}
