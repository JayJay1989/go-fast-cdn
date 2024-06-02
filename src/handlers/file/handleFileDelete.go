package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinanielsen/go-fast-cdn/src/database"
	"github.com/kevinanielsen/go-fast-cdn/src/util"
)

func HandleFileDelete(c *gin.Context) {
	fileName := c.Param("filename")
	if fileName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "File name is required",
		})
		return
	}

	deletedFileName, success := database.DeleteFile(fileName)
	if !success {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "File not found",
		})
		return
	}

	err := util.DeleteFile(deletedFileName, "file")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Failed to delete file",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "File deleted successfully",
		"fileName": deletedFileName,
	})
}
