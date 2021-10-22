package user

import (
	"errors"

	"gorm.io/gorm"
	"houserqu.com/gin-starter/core"
)

func GetModelByID(id int) (user User, err error) {
	err = core.Mysql.Take(&user, id).Error
	return
}

func GetModelAll(page int, size int, where interface{}) (user []User, err error) {
	if size <= 0 {
		size = 20
	}

	if page <= 0 {
		size = 1
	}

	err = core.Mysql.Where(where).Limit(size).Offset((page - 1) * size).Find(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
		user = []User{}
	}
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
