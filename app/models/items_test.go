package models

import (
	"app/config"
	"fmt"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	config.ConnectDB()
	item := new(Orders_item)
	item.Id = "321"
	item.Name = "Nama item"
	item.Price = 1000
	item.ExpiredAt = time.Now()

	if item.Createitem() != nil {
		t.Errorf("Failed create item")
	}
	fmt.Println(item)
}

func TestUpdate(t *testing.T) {
	config.ConnectDB()
	item := new(Orders_item)
	item.Id = "321"
	item.Name = "Nama item"
	item.Price = 1000
	item.ExpiredAt = time.Now()

	if item.Updateitem("321") != nil {
		t.Errorf("Failed update item")
	}

	fmt.Println(item)
}

func TestDelete(t *testing.T) {
	config.ConnectDB()
	item, err := GetOneByid("321")

	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(item)
	if item.Deleteitem() != nil {
		t.Errorf(err.Error())
	}
}

func TestGet(t *testing.T) {
	config.ConnectDB()

	Orders_item, err := GetAll("item")

	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(Orders_item)
}
