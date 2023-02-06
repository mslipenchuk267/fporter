package fs

import (
	"fmt"
)

type FileSystem struct {
	Procedure string
}

func (fs FileSystem) GetFiles(myDir string) []string {
	fmt.Println()
	var x []string = []string{"yo", "man"}
	return x
}
