package handlers

import (
	"crypto/md5"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/kevinanielsen/go-fast-cdn/src/database"
	"github.com/kevinanielsen/go-fast-cdn/src/models"
	"github.com/kevinanielsen/go-fast-cdn/src/util"
)

func HandleFileUpload(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	newName := c.PostForm("filename")

	if err != nil {
		c.String(http.StatusBadRequest, "Failed to read file: %s", err.Error())
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.String(http.StatusBadRequest, "Failed to open file: %s", err.Error())
		return
	}
	defer file.Close()

	fileBuffer := make([]byte, 512)
	_, err = file.Read(fileBuffer)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to read file: %s", err.Error())
		return
	}
	fileType := http.DetectContentType(fileBuffer)

	allowedMimeTypes := map[string]bool{
		"application/vnd.rar":         true,
		"application/zip":             true,
		"application/x-7z-compressed": true,
		"XML":                         true,
		"application/octet-stream":    true,
		"application/vnd.microsoft.portable-executable": true,
		"application/x-ms-installer":                    true,
	}

	if !allowedMimeTypes[fileType] {
		c.String(http.StatusBadRequest, "Invalid file type: %s", fileType)
		return
	}

	fileHashBuffer := md5.Sum(fileBuffer)
	var filename string
	if newName == "" {
		filename = fileHeader.Filename
	} else {
		filename = newName + filepath.Ext(fileHeader.Filename)
	}

	filteredFilename, err := util.FilterFilename(filename)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	doc := models.File{
		FileName: filteredFilename,
		Checksum: fileHashBuffer[:],
	}

	docInDatabase := database.GetFileByCheckSum(fileHashBuffer[:])
	if len(docInDatabase.Checksum) > 0 {
		c.JSON(http.StatusConflict, "File already exists")
		return
	}

	savedFileName, err := database.AddFile(doc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = c.SaveUploadedFile(fileHeader, util.ExPath+"/uploads/files/"+savedFileName)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to save file: %s", err.Error())
		return
	}

	body := gin.H{
		"file_url": c.Request.Host + "/download/files/" + savedFileName,
	}

	c.JSON(http.StatusOK, body)
}
