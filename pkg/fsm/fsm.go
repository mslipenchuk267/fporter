package fs

import (
	"log"
	"os"
	"path/filepath"
)

type FileSystemManager struct {
	Procedure string
}

func (fsm FileSystemManager) GetFiles(dirPath string) []string {
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

			// Append to Files Array
			files = append(files, path)
		}

		return nil
	})

	return files
}
