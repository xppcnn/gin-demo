package article

import (
	"github.com/xppcnn/gin-demo/global"
	"github.com/xppcnn/gin-demo/models"
)

type TagService struct {
}

func (t *TagService) GetTags(pageNum int, pageSize int, maps interface{}) (tags []models.Tag) {
	global.Db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func (t *TagService) GetTagTotal(maps interface{}) (count int) {
	global.Db.Model(&models.Tag{}).Where(maps).Count(&count)
	return
}

func (t *TagService) ExistTagByName(name string) bool {
	var tag models.Tag
	global.Db.Select("id").Where("name = ?", name).First(&tag)
	return tag.ID != ""
}

func (t *TagService) ExistTagByID(id string) bool {
	var tag models.Tag
	global.Db.Select("id").Where("id = ?", id).First(&tag)
	return tag.ID != ""
}

func (t *TagService) AddTag(name string, state int, createBy string) bool {

	global.Db.Create(&models.Tag{
		Name:      name,
		State:     state,
		CreatedBy: createBy,
	})
	return true
}

func (t *TagService) EditTag(id string, data interface{}) bool {
	global.Db.Model(&models.Tag{}).Where("id = ?", id).Updates(data)
	return true
}

func (t *TagService) DeleteTag(id string) bool {
	global.Db.Where("id = ?", id).Delete(&models.Tag{})
	return true
}
