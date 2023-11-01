package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AddScienceBook(c *gin.Context) {
	formFile, _ := c.FormFile("file")
	imageHeader, _ := c.FormFile("image")
	filePath := "../uploads/" + formFile.Filename
	filePathImage := "../uploads/" + imageHeader.Filename
	if err := c.SaveUploadedFile(formFile, filePath); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении файла"})
		return
	}
	if err := c.SaveUploadedFile(formFile, filePathImage); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении файла"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "успешно добавлено!"})
}

func AddHumanitarianBook(c *gin.Context) {
	fileHeader, _ := c.FormFile("file")
	imageHeader, _ := c.FormFile("image")
	filePath := "../uploads/" + fileHeader.Filename
	imageDest := "../uploads/" + imageHeader.Filename
	if err := c.SaveUploadedFile(fileHeader, filePath); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении файла"})
		return
	}
	if err := c.SaveUploadedFile(imageHeader, imageDest); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении файла 'image'"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Успешно добавлено!"})
}
