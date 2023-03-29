package models

import (
	"app/config"
	"time"
)

type Orders_item struct {
	Id        string    `json:"id" form:"id" gorm:"primaryKey"`
	Name      string    `json:"name" form:"name"`
	Price     int       `json:"price" form:"price"`
	ExpiredAt time.Time `json:"expired_at" form:"expired_at"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (item *Orders_item) Createitem() error {
	if err := config.DB.Create(item).Error; err != nil {
		return err
	}
	return nil
}

func (item *Orders_item) Updateitem(id string) error {
	if err := config.DB.Model(&Orders_item{}).Where("id = ?", id).Updates(item).Error; err != nil {
		return err
	}
	return nil
}

func (item *Orders_item) Deleteitem() error {
	if err := config.DB.Delete(item).Error; err != nil {
		return err
	}
	return nil
}

func GetItemByid(id string) (Orders_item, error) {
	var item Orders_item
	result := config.DB.Where("id = ?", id).First(&item)
	return item, result.Error
}

func GetItems(keywords string) ([]Orders_item, error) {
	var Orders_item []Orders_item
	result := config.DB.Where("id LIKE ? OR name LIKE ?", "%"+keywords+"%", "%"+keywords+"%").Find(&Orders_item)

	return Orders_item, result.Error
}
