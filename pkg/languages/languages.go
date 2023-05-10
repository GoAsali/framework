package languages

import (
	"encoding/json"
	"fmt"
	filesUtils "github.com/abolfazlalz/goasali/pkg/utils/files"
	"github.com/abolfazlalz/goasali/pkg/utils/slices"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const folderPath string = "languages"

type folderLanguage struct {
	path     string
	language string
	files    []string
}

func (fl *folderLanguage) loadFiles() []os.DirEntry {
	files, err := os.ReadDir(fl.path)
	if err != nil {
		log.Fatal(err)
	}

	files = filesUtils.FilterFileExtensions(files, "json")

	fl.files = slices.Map(files, func(t os.DirEntry) string {
		return fmt.Sprintf("%s/%s", fl.path, t.Name())
	})
	fmt.Println(fl.files)

	return files
}

func getFolders() []os.DirEntry {
	dirPath := "languages"
	folders, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	folders = slices.Filter(folders, func(entry os.DirEntry) bool {
		return entry.IsDir()
	})

	return folders
}

func folderToLanguage(path string) func(entry os.DirEntry) *folderLanguage {
	return func(entry os.DirEntry) *folderLanguage {
		return &folderLanguage{
			language: entry.Name(),
			path:     filepath.Join(path, entry.Name()),
		}
	}
}

func loadMessages() {
	folders := slices.Map(getFolders(), folderToLanguage(folderPath))
	slices.Apply(folders, func(t *folderLanguage) {
		t.loadFiles()
	})
	fmt.Println(slices.Map(folders, func(t *folderLanguage) int {
		return len(t.files)
	}))
}

// LoadLanguages Loading messages file languages in */languages/* folder
func LoadLanguages() *i18n.Bundle {

	log.SetPrefix("[Language loading]")
	defer func() {
		log.Println("complete")
		log.SetPrefix("")
	}()
	loadMessages()
	log.Println("Loading languages list")
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	var filesList []os.DirEntry
	var err error

	// Load languages filesList in languages directory
	dirPath := "languages/en"
	if filesList, err = os.ReadDir(dirPath); err != nil {
		log.Fatal(err)
	}

	for _, file := range filesList {
		if file.IsDir() {
			continue
		}

		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}
		if _, err = bundle.LoadMessageFile(file.Name()); err != nil {
			filePath := filepath.Join(dirPath, file.Name())
			if _, err := bundle.LoadMessageFile(filePath); err != nil {
				log.Fatalf("Error during loading message dictionary filePath in %s: %v", filePath, err)
			}
			log.Printf("%s Language has added successfully.", file.Name())
		}
	}

	return bundle
}
