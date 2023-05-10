package files

import (
	"errors"
	"github.com/abolfazlalz/goasali/pkg/utils/slices"
	"os"
	"strings"
)

var (
	PathNotFile = errors.New("the entered path does not belong to a file")
)

type SupportedFiles interface {
	string
}

func CheckFileExtension(file os.DirEntry, extension string) (bool, error) {
	if file.IsDir() {
		return false, PathNotFile
	}
	return strings.HasSuffix(file.Name(), extension), nil
}

func FilterFileExtensions(files []os.DirEntry, extension string) []os.DirEntry {
	return slices.Filter(files, func(file os.DirEntry) bool {
		f, e := CheckFileExtension(file, extension)
		return f && e != nil
	})
}
