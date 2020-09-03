package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  int    `json:"created_by"`
	ModifiedBy int    `json:"modified_by"`
	State      int    `json:"state"`
	
	TagId int `json:"tag_id" gorm:"index"`
	Tag Tag `json:"tag"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error  {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error  {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func ExistArticleById(id int ) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	return article.ID > 0
}