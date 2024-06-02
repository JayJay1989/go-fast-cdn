package handlers

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/kevinanielsen/go-fast-cdn/src/util"

	"github.com/gin-gonic/gin"
)

func HandleFileMetadata(c *gin.Context) {
	fileName := c.Param("filename")
	if fileName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "File name is required",
		})
		return
	}

	filePath := filepath.Join(util.ExPath, "uploads", "files", fileName)
	stat, err := os.Stat(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "File does not exist",
			})
		} else {
			log.Printf("Failed to get file %s: %s\n", fileName, err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal error",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"filename":     fileName,
		"download_url": c.Request.Host + "/api/cdn/download/file/" + fileName,
		"file_size":    stat.Size(),
	})
}
