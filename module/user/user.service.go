package user

import (
	"houserqu.com/gin-starter/core"
)

func GetModelByID(id int) (user User, err error) {
	err = core.Mysql.Take(&user, id).Error
	return
}

func GetModelAll() (user []User, err error) {
	err = core.Mysql.Find(&user).Error
	return
}

func DelModel(id int) (err error) {
	err = core.Mysql.Delete(&User{}, id).Error
	return
}

func CreateModel(user *User) (err error) {
	err = core.Mysql.Create(&user).Error
	return
}

func UpdateModel(id int, name string, email string) (user User, err error) {
	user, err = GetModelByID(id)
	if err != nil {
		return
	}

	err = core.Mysql.Model(&user).Updates(map[string]interface{}{"name": name, "email": email}).Error
	return
}
