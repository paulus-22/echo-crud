package models

import (
	"app/config"
	"app/models"
	"time"
)

type Order_histories struct {
	Id            string `json:"id" form:"id" gorm:"primaryKey"`
	User_id       models.Users
	Order_item_id models.Orders_item
	Descriptions  string `json:"descriptions" form:"descriptions"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (history *Order_histories) Createhistory() error {
	if err := config.DB.Create(history).Error; err != nil {
		return err
	}
	return nil
}

func (history *Order_histories) Updatehistory(id string) error {
	if err := config.DB.Model(&Order_histories{}).Where("id = ?", id).Updates(history).Error; err != nil {
		return err
	}
	return nil
}

func (history *Order_histories) Deletehistory() error {
	if err := config.DB.Delete(history).Error; err != nil {
		return err
	}
	return nil
}

func GetOneByid(id string) (Order_histories, error) {
	var history Order_histories
	result := config.DB.Where("id = ?", id).First(&history)
	return history, result.Error
}

func GetAll(id string) ([]Order_histories, error) {
	var Order_histories []Order_histories
	result := config.DB.Where("id LIKE ? ", "%"+id+"%").Find(&Order_histories)

	return Order_histories, result.Error
}
