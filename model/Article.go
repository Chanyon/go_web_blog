package model

import (
	"go_blog/utils/errmsg"

	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title        string `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Desc         string `gorm:"type:varchar(200)" json:"desc"`
	Content      string `gorm:"type:longtext" json:"content"`
	Img          string `gorm:"type:varchar(100)" json:"img"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
}

func CreateArticleService(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetOneArticleInfoService(id int) (Article, int) {
	var article Article
	err := db.Where("id = ?", id).Preload("Category").First(&article).Error
	db.Model(&article).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))
	if err != nil {
		return article, errmsg.ERROR_ART_NOT_EXIST
	}
	return article, errmsg.SUCCESS
}

func GetAllArticleService(pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	// var category model.Category
	var total int64
	err := db.Select("article.id, title, img, article.created_at, article.updated_at, `desc`, comment_count, read_count, Category.name").
		Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Joins("Category").Find(&articleList).Error
	db.Model(&articleList).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCESS, total
}

// * 搜索文章标题
func SearchArticleService(title string, pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var total int64
	err := db.Select("article.id, title, img, article.created_at, article.updated_at, `desc`, comment_count, read_count, Category.name").Order("Created_At DESC").
		Where("title LIKE ?", title+"%").Joins("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
	db.Model(&articleList).Where("title LIKE ?", title+"%").Count(&total)
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCESS, total
}

// * 分类查询
func GetCategoryArticleService(id int, pageSize int, pageNum int) ([]Article, int, int64) {
	var categoryArticleList []Article
	var total int64
	err := db.Preload("Category").Where("cid =?", id).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&categoryArticleList).Error
	db.Model(&categoryArticleList).Where("cid =?", id).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return categoryArticleList, errmsg.SUCCESS, total
}

// * 修改文章
func EditArticleService(id int, data *Article) int {
	var article Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["img"] = data.Img
	maps["content"] = data.Content
	err := db.Model(&article).Where("id =?", id).UpdateColumns(&maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteArticleService(id int) int {
	var article Article
	err := db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
