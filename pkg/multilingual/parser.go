package multilingual

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"io/ioutil"
	"os"
)

func convertRaw(name string, raw interface{}) ([]*i18n.Message, error) {
	messages := make([]*i18n.Message, 0)
	switch data := raw.(type) {
	case map[string]string:
		fmt.Println(data)
		break
	case map[string]interface{}:
		for k, v := range data {
			key := fmt.Sprintf("%s.%s", name, k)
			switch value := v.(type) {
			case string:
				msg := &i18n.Message{
					ID:    key,
					Other: value,
				}
				messages = append(messages, msg)
				break
			case interface{}:
				msg, err := convertRaw(key, value)
				if err != nil {
					return nil, err
				}
				messages = append(messages, msg...)
				break
			}
		}
		break
	}

	return messages, nil
}

func parseFile(file fileLanguage) ([]*i18n.Message, error) {
	jsonFile, err := os.Open(file.path)
	if err != nil {
		return nil, err
	}
	fileContent, _ := ioutil.ReadAll(jsonFile)

	var raw interface{}

	if err = json.Unmarshal(fileContent, &raw); err != nil {
		return nil, err
	}

	messages, err := convertRaw(file.name, raw)

	if err != nil {
		return nil, err
	}

	return messages, nil
}

func parseFiles(files []fileLanguage) ([]*i18n.Message, error) {
	messages := make([]*i18n.Message, 0)
	for _, file := range files {
		if result, err := parseFile(file); err != nil {
			return nil, err
		} else {
			messages = append(messages, result...)
		}
	}
	return messages, nil
}
