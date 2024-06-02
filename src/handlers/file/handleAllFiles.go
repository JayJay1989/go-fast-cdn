package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevinanielsen/go-fast-cdn/src/database"
	"github.com/kevinanielsen/go-fast-cdn/src/models"
)

func HandleAllFiles(c *gin.Context) {
	var entries []models.File

	database.DB.Find(&entries, &models.File{})

	c.JSON(http.StatusOK, entries)
}
