package models

import (
	"app/config"
	"fmt"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	config.ConnectDB()
	user := new(Users)
	user.Id = "123"
	user.Full_Name = "Nama saya"
	user.First_Order = time.Now()

	if user.CreateUser() != nil {
		t.Errorf("Failed create user")
	}
	fmt.Println(user)
}

func TestUpdate(t *testing.T) {
	config.ConnectDB()
	user := new(Users)
	user.Id = "123"
	user.Full_Name = "Nama saya"
	user.First_Order = time.Now()

	if user.UpdateUser("123") != nil {
		t.Errorf("Failed update user")
	}

	fmt.Println(user)
}

func TestDelete(t *testing.T) {
	config.ConnectDB()
	user, err := GetOneByid("123")

	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(user)
	if user.DeleteUser() != nil {
		t.Errorf(err.Error())
	}
}

func TestGet(t *testing.T) {
	config.ConnectDB()

	users, err := GetAll("Nama")

	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(users)
}
