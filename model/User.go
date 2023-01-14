package model

import (
	"go_blog/utils/errmsg"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=5,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=16" label:"密码"`
	Role     int    `gorm:"type:int;default:2" json:"role" validate:"required,gte=2" label:"角色id"`
}

// service

func CreateUserService(data *User) int {
	data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	// 发生错误
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// check
func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// update 更新用户信息时检查
func CheckUpUser(id int, name string) (code int) {
	var user User
	db.Select("id,username").Where("username = ?", name).First(&user)
	// if user.ID == uint(id) {
	// 	return errmsg.SUCCESS
	// }
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 通过id查询用户
func GetUserInfoService(id int) (User, int) {
	var user_data User
	err := db.Limit(1).Where("ID = ?", id).Find(&user_data).Error
	if err != nil {
		return user_data, errmsg.ERROR
	}
	return user_data, errmsg.SUCCESS
}

// 查询用户列表
func GetUsersService(username string, pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64
	if username != "" {
		db.Select("id,username,role,created_at").Where(
			"username Like ?", username+"%",
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
		db.Model(&users).Where("username LIKE ?", username+"%").Count(&total)
		// return users, total
	} else {
		db.Select("id,username,role,created_at").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
		db.Model(&users).Count(&total)
	}
	if err != nil {
		return users, 0
	}
	return users, total
}

// 编辑用户信息
func EditUserService(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role

	err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 修改密码
func ChangePasswordService(id int, data *User) int {
	data.Password = ScryptPw(data.Password)
	err = db.Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户
func DeleteUserService(id int) int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func ScryptPw(password string) string {
	const cost = 10
	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}
	return string(HashPw)
}

// 后台验证
func CheckLogin(username string, password string) (User, int) {
	var user User
	var passwordErr error
	db.Where("username = ?", username).First(&user)
	passwordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_FOUND
	}
	if passwordErr != nil {
		return user, errmsg.ERROR_USER_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return user, errmsg.ERROR_USER_NOT_FOUND
	}
	return user, errmsg.SUCCESS
}

func CheckLoginFront(username string, password string) (User, int) {
	var user User
	var passwordErr error
	db.Where("username = ?", username).First(&user)
	passwordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_FOUND
	}
	if passwordErr != nil {
		return user, errmsg.ERROR_USER_PASSWORD_WRONG
	}
	return user, errmsg.SUCCESS
}
