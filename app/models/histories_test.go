package models

import (
	"app/config"
	"fmt"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	config.ConnectDB()
	history := new(Order_histories)
	history.Id = "234"
	history.User_id = new(Users)
	history.Order_item_id = new(Orders_item)
	history.Descriptions = "ini history"

	if history.Createhistory() != nil {
		t.Errorf("Failed create history")
	}
	fmt.Println(history)
}

func TestUpdate(t *testing.T) {
	config.ConnectDB()
	history := new(Order_histories)
	history.Id = "234"
	history.User_id = new(Users)
	history.Order_item_id = new(Orders_item)
	history.Descriptions = "ini history"

	if history.Updatehistory("234") != nil {
		t.Errorf("Failed update history")
	}

	fmt.Println(history)
}

func TestDelete(t *testing.T) {
	config.ConnectDB()
	history, err := GetHistoryByid("234")

	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(history)
	if history.Deletehistory() != nil {
		t.Errorf(err.Error())
	}
}

func TestGet(t *testing.T) {
	config.ConnectDB()

	Orders_history, err := GetHistories("history")

	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(Orders_history)
}
