package file

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strings"
)

func FindFile(c *gin.Context) {
	dirPath := "../uploads/"
	files, err := getJPEGFilesInDirectory(dirPath)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении файлов"})
		return
	}
	var arrayFile []string
	for _, mas := range files {
		arrayFile = append(arrayFile, mas)
	}
	c.JSON(http.StatusOK, gin.H{"files": arrayFile})
}

func getJPEGFilesInDirectory(dirPath string) ([]string, error) {
	var jpegFiles []string
	dir, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	defer func(dir *os.File) {
		err := dir.Close()
		if err != nil {
			log.Println(err.Error())
			return
		}
	}(dir)
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			fileName := strings.ToLower(fileInfo.Name())
			if strings.HasSuffix(fileName, ".jpeg") || strings.HasSuffix(fileName, ".jpg") || strings.HasSuffix(fileName, ".png") || strings.HasSuffix(fileName, ".pdf") || strings.HasSuffix(fileName, ".docx") {
				jpegFiles = append(jpegFiles, fileInfo.Name())
			}
		}
	}
	return jpegFiles, nil
}
