package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strings"
	"time"
)

func UploadFile(c *gin.Context, key string, dst string) (fileName string, error error) {
	file, _ := c.FormFile(key)
	log.Println(file.Filename)

	dst = fmt.Sprintf("/uploads/%s", dst)
	fileNameSplit := strings.Split(file.Filename, ".")
	fileEx := ""
	if len(fileNameSplit) > 1 {
		fileEx = "." + fileNameSplit[len(fileNameSplit)-1]
	}
	fileName = fmt.Sprintf("%s/%d%s", dst, time.Now().Unix(), fileEx)

	err := os.MkdirAll(dst, os.ModePerm)
	if err != nil {
		return "", err
	}

	// Upload the file to specific dst.
	err = c.SaveUploadedFile(file, "public"+fileName)
	log.Println(err)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func CheckFileExists(path string) bool {
	if _, err := os.Stat("public" + path); err == nil {
		return true
	}
	return false
}
