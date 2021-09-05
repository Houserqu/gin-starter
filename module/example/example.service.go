package example

import (
	"houserqu.com/gin-starter/internal"
)

func GetModelByID(id int) (example Example, err error) {
	err = internal.DB.Take(&example, id).Error
	return
}

func GetModelAll() (examples []Example, err error) {
	err = internal.DB.Find(&examples).Error
	return
}

func DelModel(id int) (err error) {
	err = internal.DB.Delete(&Example{}, id).Error
	return
}

func CreateModel(example *Example) (err error) {
	err = internal.DB.Create(&example).Error
	return
}

func UpdateModel(id int, name string, email string) (example Example, err error) {
	example, err = GetModelByID(id)
	if err != nil {
		return
	}

	err = internal.DB.Model(&example).Updates(map[string]interface{}{"name": name, "email": email}).Error
	return
}
