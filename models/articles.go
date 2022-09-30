package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/unknwon/com"
)

type Article struct {
	Model

	TagID      string `json:"tag_id" gorm:"index"`
	Tag        Tag    `json:"tag"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	u4 := com.ToStr(uuid.New())
	article.ID = u4
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func ExistArticleByID(id string) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	return article.ID != ""
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

func GetArticle(id string) Article {
	var article Article
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)
	return article
}

func EditArticle(id string, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Update(data)
	return true
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(string),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true
}

func DeleteArticle(id string) bool {
	db.Where("id = ?", id).Delete(&Article{})
	return true
}
