package cli

import (
	"os"
	"strings"
)

func GetArgs() []string {
	return os.Args[1:]
}

func GetArgsFromKey(key string) string {
	for _, arg := range GetArgs() {
		keyValue := strings.Split(arg, "=")
		if keyValue[0] == key {
			return keyValue[1]
		}
	}

	return ""
}

func HasArgsKey(key string) bool {
	return GetArgsFromKey(key) != ""
}
