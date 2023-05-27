package multilingual

import (
	filesUtils "github.com/abolfazlalz/goasali/pkg/utils/files"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"log"
)

type Multilingual struct {
	*i18n.Bundle
	Path    string
	Modules []string
}

func New() *Multilingual {
	return &Multilingual{
		Bundle:  i18n.NewBundle(language.English),
		Path:    "languages",
		Modules: make([]string, 0),
	}
}

// Load Load all messages.json from languages folder
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

	var messagesMap map[language.Tag][]*i18n.Message
	var err error

	for _, dir := range filesUtils.Directories(m.Path) {
		dl := dlMaker(dir)
		files = append(files, dl.files...)
	}
	if messagesMap, err = parseFiles(files); err != nil {
		return err
	}

	for lang, messages := range messagesMap {
		if err := m.AddMessages(lang, messages...); err != nil {
			return err
		}
	}

	return nil
}

func (m *Multilingual) AddModule(path string) {
	m.Modules = append(m.Modules, path)
}

// ChangeLanguageDirectory Change language directory and reload messages.json files
func (m *Multilingual) ChangeLanguageDirectory(dirPath string) error {
	m.Path = dirPath
	return m.Load()
}
