package files

import (
	"errors"
	"github.com/abolfazlalz/goasali/pkg/utils/slices"
	"log"
	"os"
	"strings"
)

var (
	PathNotFile = errors.New("the entered path does not belong to a file")
)

func CheckFileExtension(file os.DirEntry, extension string) (bool, error) {
	if file.IsDir() {
		return false, PathNotFile
	}
	return strings.HasSuffix(file.Name(), "."+extension), nil
}

// FilterFileExtensions Filter list of files as __DirEntry__ by their extensions format
func FilterFileExtensions(files []os.DirEntry, extension string) []os.DirEntry {
	return slices.Filter(files, func(file os.DirEntry) bool {
		f, e := CheckFileExtension(file, extension)
		return f && e == nil
	})
}

// Directories Return list of directories in a directory
func Directories(dirPath string) []os.DirEntry {
	directories, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	directories = slices.Filter(directories, func(entry os.DirEntry) bool {
		return entry.IsDir()
	})

	return directories
}
