package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name         string `json:"name" form:"name"`
	Status       string `json:"status" form:"status"`
	Categories   string `json:"categories" form:"categories"`
	Descriptions string `json:"descriptions" form:"descriptions"`
}
