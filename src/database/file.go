package database

import (
	"github.com/kevinanielsen/go-fast-cdn/src/models"
)

func GetFileByCheckSum(checksum []byte) models.File {
	var entries models.File

	DB.Where("checksum = ?", checksum).First(&entries)

	return entries
}

func AddFile(file models.File) (string, error) {
	result := DB.Create(&file)
	if result.Error != nil {
		return "", result.Error
	}

	return file.FileName, result.Error
}

func DeleteFile(fileName string) (string, bool) {
	var file models.File

	result := DB.Where("file_name = ?", fileName).First(&file)

	if result.Error == nil {
		DB.Delete(&file)
		return fileName, true
	} else {
		return "", false
	}
}

func RenameFile(oldFileName, newFileName string) error {
	file := models.File{}
	return DB.Model(&file).Where("file_name = ?", oldFileName).Update("file_name", newFileName).Error
}
