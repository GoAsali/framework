package multilingual

import (
	filesUtils "github.com/abolfazlalz/goasali/pkg/utils/files"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
)

type Multilingual struct {
	*i18n.Bundle
	Path string
}

// Load Load all messages from languages folder
func (m *Multilingual) Load() error {
	lastPrefix := log.Prefix()
	log.SetPrefix("[Language loading] - ")
	defer func(lastPrefix string) {
		log.Println("complete")
		log.SetPrefix(lastPrefix)
	}(lastPrefix)

	log.Println("Loading multilingual list")

	dlMaker := directoryToDirLanguage(m.Path)
	files := make([]fileLanguage, 0)

	var messages []*i18n.Message
	var err error

	for _, dir := range filesUtils.Directories(m.Path) {
		dl := dlMaker(dir)
		files = append(files, dl.files...)
	}
	if messages, err = parseFiles(files); err != nil {
		return err
	}

	if err := m.AddMessages(language.English, messages...); err != nil {
		return err
	}

	return nil
}

// ChangeLanguageDirectory Change language directory and reload messages files
func (m *Multilingual) ChangeLanguageDirectory(dirPath string) error {
	m.Path = dirPath
	return m.Load()
}
