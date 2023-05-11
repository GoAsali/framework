package multilingual

import (
	"fmt"
	filesUtils "github.com/abolfazlalz/goasali/pkg/utils/files"
	"github.com/abolfazlalz/goasali/pkg/utils/slices"
	"golang.org/x/text/language"
	"os"
	"path/filepath"
	"strings"
)

type fileLanguage struct {
	language language.Tag
	name     string
	path     string
}

type DirectoryLanguage struct {
	path     string
	language string
	files    []fileLanguage
}

// LoadFiles Load json files in a directory
func (fl *DirectoryLanguage) loadFiles() ([]os.DirEntry, error) {
	files, err := os.ReadDir(fl.path)
	if err != nil {
		return nil, err
	}

	files = filesUtils.FilterFileExtensions(files, "json")

	fl.files = slices.Map(files, func(t os.DirEntry) fileLanguage {
		name := t.Name()
		return fileLanguage{
			language: language.MustParse(fl.language),
			path:     fmt.Sprintf("%s/%s", fl.path, name),
			name:     strings.Split(name, ".")[0],
		}
	})

	return files, nil
}

func NewDirectoryLanguage(language string, path string) *DirectoryLanguage {
	return &DirectoryLanguage{
		language: language,
		path:     path,
	}
}

func directoryToDirLanguage(path string) func(entry os.DirEntry) *DirectoryLanguage {
	return func(entry os.DirEntry) *DirectoryLanguage {
		dl := NewDirectoryLanguage(entry.Name(), filepath.Join(path, entry.Name()))
		if _, err := dl.loadFiles(); err != nil {
			return nil
		}
		return dl
	}
}
