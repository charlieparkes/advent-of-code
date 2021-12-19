package files

import (
	"fmt"
	"os"
	"path/filepath"
)

func FindPath(containing string) string {
	workingDirectory, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	lastDir := workingDirectory
	for {
		currentPath := fmt.Sprintf("%s/%s", lastDir, containing)
		fi, err := os.Stat(currentPath)
		if err == nil {
			switch mode := fi.Mode(); {
			case mode.IsDir():
				return currentPath
			}
		}
		newDir := filepath.Dir(lastDir)
		if newDir == "/" || newDir == lastDir {
			return ""
		}
		lastDir = newDir
	}
}
