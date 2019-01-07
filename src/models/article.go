package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model
	TagId      int    `json:"tag_id" gorm:"index"`
	Tag        Tag    `json:"tag"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

func ExistArticleByID(id int) bool {
	var article Article
	err := db.Select("id").Where("id = ?", id).First(&article).Error
	if err != nil {
		return false
	} else if article.ID > 0 {
		return true
	} else {
		return false
	}
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Model(&Article{}).Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

func GetArticleByID(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)
	return
}

func EditArticle(id int, data interface{}) bool {
	err := db.Model(&Article{}).Where("id = ?", id).Update(data).Error
	if err != nil {
		return false
	} else {
		return true
	}
}

func DeleteArticle(id int) bool {
	err := db.Where("id = ?", id).Delete(&Article{}).Error
	if err != nil {
		return false
	} else {
		return true
	}
}

func AddArticle(data map[string]interface{}) bool {
	var article = &Article{
		TagId:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	}
	err := db.Create(article).Error
	if err != nil {
		return false
	} else {
		return true
	}
}
