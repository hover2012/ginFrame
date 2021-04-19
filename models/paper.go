package models

import (
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type Paper struct {
	Model
	Name string `json:"name"`
	ExmId int `json:"exm_id"`
	DetailUrl string `json:"detail_url"`
	DocUrl string `json:"doc_url"`
	PdfUrl string `json:"pdf_url"`
	StoreSubject string `json:"store_subject"`
	Xueke string `json:"xueke"`
	Xueduan string `json:"xueduan"`
}

func AddPaper(paper *Paper) bool  {
	db.Create(paper)
	log.Print(paper)
	return true
}

func GetPaper(maps interface{})(papers []Paper)  {
	db.Where(maps).Find(&papers)
	return
}

func UpdatePaper(id int , datas interface{})bool  {
	db.Model(&Paper{}).Where("id=?",id).Update(datas)
	return true
}


func (tag *Paper) BeforeCreate(scope *gorm.Scope) error  {
	scope.SetColumn("CreateOn",time.Now().Unix())
	return nil
}