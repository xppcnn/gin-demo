package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/unknwon/com"
)

type Tag struct {
	Model
	// ID         string `gorm:"primary_key" json:"id" binding:"uuid"`
	// CreatedOn  int    `json:"created_on"`
	// ModifiedOn int    `json:"modified_on"`
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID != "" {
		return true
	}
	return false
}

func ExistTagByID(id string) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID != "" {
		return true
	}
	return false
}

func AddTag(name string, state int, createBy string) bool {

	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createBy,
	})
	return true
}

func EditTag(id string, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	return true
}

func DeleteTag(id string) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}
func (model *Model) BeforeCreate(scope *gorm.Scope) error {
	u4 := com.ToStr(uuid.New())
	model.ID = u4
	// model.CreatedOn = time.Now().Unix()
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (tag *Model) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}
