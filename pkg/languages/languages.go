package languages

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// LoadLanguages Loading messages file languages in */languages/* folder
func LoadLanguages() *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	var filesList []os.DirEntry
	var err error

	// Load languages filesList in languages directory
	dirPath := "languages"
	if filesList, err = os.ReadDir(dirPath); err != nil {
		log.Fatal(err)
	}

	for _, file := range filesList {
		if file.IsDir() {
			continue
		}

		if strings.HasSuffix(file.Name(), ".json") {
			continue
		}
		fmt.Println(file.Name())
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
