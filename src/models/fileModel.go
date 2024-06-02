package models

import "gorm.io/gorm"

type File struct {
	gorm.Model

	FileName string `json:"file_name"`
	Checksum []byte `json:"checksum"`
}
