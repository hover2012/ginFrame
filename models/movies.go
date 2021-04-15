package models

import (
	"github.com/jinzhu/gorm"
	"time"
)
type MovieModel struct {
	Model

	Title       string
	Subtitle    string
	Other       string
	Description string
	Year        string
	Area        string
	Tag         string
	Star        string
	Comment     string
	Quote       string
}

func AddMovie(movie *MovieModel) bool {
	db.Create(movie)
	//log.Printf("success")
	return true
}

func (tag *MovieModel) BeforeCreate(scope *gorm.Scope) error  {
	scope.SetColumn("CreateOn",time.Now().Unix())
	return nil
}
