package model

import (
	"go_blog/utils/errmsg"
)

type Profile struct {
	ID        uint   `gorm:"primarykey" json:"id"`
	Name      string `gorm:"type:varchar(20)" json:"name"`
	Desc      string `gorm:"type:varchar(200)" json:"desc"`
	QqChat    string `gorm:"type:varchar(200)" json:"qq_chat"`
	WeChat    string `gorm:"type:varchar(200)" json:"wei_chat"`
	WeiBo     string `gorm:"type:varchar(200)" json:"wei_bo"`
	BiLi      string `gorm:"type:varchar(200)" json:"bili"`
	Email     string `gorm:"type:varchar(200)" json:"email"`
	Img       string `gorm:"type:varchar(200)" json:"img"`
	Avatar    string `gorm:"type:varchar(200)" json:"avatar"`
	IcpRecord string `gorm:"type:varchar(200)" json:"icp_record"`
}

func GetProfileService(id int) (Profile, int) {
	var profile Profile
	err := db.Where("id =?", id).First(&profile).Error
	if err != nil {
		return profile, errmsg.ERROR
	}
	return profile, errmsg.SUCCESS
}

func UpdateProfileService(id int, data *Profile) int {
	var profile Profile
	err := db.Model(&profile).Where("id =?", id).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func CreatProfileService(data *Profile) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
