package models

import (
	"app/config"
	"time"
)

type Users struct {
	Id          string    `json:"id" form:"id" gorm:"primaryKey"`
	Full_Name   string    `json:"full_name" form:"full_name"`
	First_Order time.Time `json:"first_order" form:"first_order"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func (user *Users) CreateUser() error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (user *Users) UpdateUser(id string) error {
	if err := config.DB.Model(&Users{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (user *Users) DeleteUser() error {
	if err := config.DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func GetOneByid(id string) (Users, error) {
	var user Users
	result := config.DB.Where("id = ?", id).First(&user)
	return user, result.Error
}

func GetAll(keywords string) ([]Users, error) {
	var users []Users
	result := config.DB.Where("id LIKE ? OR full_name LIKE ?", "%"+keywords+"%", "%"+keywords+"%").Find(&users)

	return users, result.Error
}
