package model

import (
	"go_blog/utils/errmsg"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId    int    `gorm:"type:int;not null" json:"user_id"`
	ArticleId int    `gorm:"type:int;not null" json:"article_id"`
	Title     string `gorm:"type:varchar(500);not null" json:"article_title"`
	Username  string `gorm:"type:varchar(20);not null" json:"user_username"`
	Content   string `gorm:"type:varchar(500);not null" json:"content"`
	Status    int8   `gorm:"type:int;default:2" json:"status"`
}

func CreateCommentService(data *Comment) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetCommentService(id int) (Comment, int) {
	var comment Comment
	err := db.Where("id=?", id).First(&comment).Error
	if err != nil {
		return comment, errmsg.ERROR
	}
	return comment, errmsg.SUCCESS
}

// 查询所有评论
func GetCommentListService(pageSize int, pageNum int) ([]Comment, int64, int) {
	var comments []Comment
	var total int64
	db.Find(&comments).Count(&total)
	err := db.Model(&comments).Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").
		Select("comment.id, article.title,user_id,article_id, user.username, comment.content, comment.status,comment.created_at,comment.deleted_at").
		Joins("LEFT JOIN article ON comment.article_id = article.id").
		Joins("LEFT JOIN user ON comment.user_id = user.id").
		Scan(&comments).Error

	if err != nil {
		return comments, 0, errmsg.ERROR
	}
	return comments, total, errmsg.SUCCESS
}

func GetCommentListFrontService(articleId int, pageSize int, pageNum int) ([]Comment, int64, int) {
	var comments []Comment
	var total int64
	db.Find(&Comment{}).Where("article_id = ?", articleId).Where("status = ?", 1).Count(&total)
	err := db.Model(&Comment{}).Limit(pageSize).Offset((pageNum-1)*pageSize).Order("Created_At DESC").
		Select("comment.id, article.title,user_id,article_id, user.username, comment.content, comment.status,comment.created_at,comment.deleted_at").
		Joins("LEFT JOIN article ON comment.article_id = article.id").
		Joins("LEFT JOIN user ON comment.user_id = user.id").
		Where("article_id = ?", articleId).Where("status = ?", 1).
		Scan(&comments).Error
	if err != nil {
		return comments, 0, errmsg.ERROR
	}
	return comments, total, errmsg.SUCCESS
}

func GetCommentCountService(id int) int64 {
	var comment Comment
	var total int64
	db.Find(&comment).Where("article_id = ?", id).Where("status = ?", 1).Count(&total)
	return total
}

func DeleteCommentService(id int) int {
	var comment Comment
	err := db.Where("id = ?", id).Delete(&comment).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 通过评论
func CheckCommentService(id int, data *Comment) int {
	var comment Comment
	var res Comment
	var article Article
	var maps = make(map[string]interface{})
	maps["status"] = data.Status

	err := db.Model(&comment).Where("id =?", id).Updates(maps).First(&res).Error
	db.Model(&article).Where("id =?", res.ArticleId).
		UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1))
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 隐藏
func UncheckCommentService(id int, data *Comment) int {
	var comment Comment
	var res Comment
	var article Article
	var maps = make(map[string]interface{})
	maps["status"] = data.Status
	err := db.Model(&comment).Where("id =?", id).Updates(maps).First(&res).Error
	db.Model(&article).Where("id =?", res.ArticleId).
		UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1))
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
