package model

import (
	"go_blog/utils/errmsg"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

func CreateCategoryService(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 分类名检查
func CheckCategory(name string) int {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATE_NAME_USED
	}
	return errmsg.SUCCESS
}

func GetCategoryInfoService(id int) (Category, int) {
	var cate Category
	db.Where("id = ?", id).First(&cate)
	return cate, errmsg.SUCCESS
}

// 查询分类列表[]Category
func GetCategoryService(pageSize int, pageNum int) ([]Category, int64) {
	var cate []Category
	var total int64
	err := db.Find(&cate).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	db.Model(&cate).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cate, total
}

func EditCategoryService(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err := db.Model(&cate).Where("id =?", id).Updates(&maps).Error
	if err != nil {
		return errmsg.ERROR
	} else {
		return errmsg.SUCCESS
	}
}

func DeleteCategoryService(id int) int {
	var cate Category
	err := db.Where("id =?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	} else {
		return errmsg.SUCCESS
	}
}
