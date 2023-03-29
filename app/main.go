package main

import (
	"app/config"
	"app/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDB()
	route := echo.New()

	// user method
	route.POST("user/create", func(c echo.Context) error {
		user := new(model.Users)
		c.Bind(user)
		response := new(Response)
		if user.CreateUser() != nil {
			response.ErrorCode = 10
			response.Message = "Failed"
		} else {
			response.ErrorCode = 0
			response.Message = "Success"
			response.Data = *user
		}
		return c.JSON(http.StatusOK, response)
	})

	route.PUT("user/update_user/:id", func(c echo.Context) error {
		user := new(model.Users)
		c.Bind(user)
		response := new(Response)
		if user.UpdateUser(c.Param("id")) != nil {
			response.ErrorCode = 10
			response.Message = "Failed"
		} else {
			response.ErrorCode = 0
			response.Message = "Success"
			response.Data = *user
		}
		return c.JSON(http.StatusOK, response)
	})

	route.DELETE("user/delete_user/:id", func(c echo.Context) error {
		user, _ := model.Users.GetOneByid(c.Param("id"))
		response := new(Response)

		if user.DeleteUser() != nil { // method update user
			response.ErrorCode = 10
			response.Message = "Failed"
		} else {
			response.ErrorCode = 0
			response.Message = "Success"
		}
		return c.JSON(http.StatusOK, response)
	})

	route.GET("user/get_users", func(c echo.Context) error {
		response := new(Response)
		users, err := model.Users.GetAll(c.QueryParam("keywords"))
		if err != nil {
			response.ErrorCode = 10
			response.Message = "Failed"
		} else {
			response.ErrorCode = 0
			response.Message = "Success"
			response.Data = users
		}
		return c.JSON(http.StatusOK, response)
	})

	// item method
	route.POST("item/create", func(c echo.Context) error {
		item := new(model.Orders_item)
		c.Bind(item)
		response := new(Response)
		if item.Createitem() != nil {
			response.ErrorCode = 10
			response.Message = "Failed"
		} else {
			response.ErrorCode = 0
			response.Message = "Success"
			response.Data = *item
		}
		return c.JSON(http.StatusOK, response)
	})

	route.PUT("item/update_item/:id", func(c echo.Context) error {
		item := new(model.Orders_item)
		c.Bind(item)
		response := new(Response)
		if item.Updateitem(c.Param("id")) != nil {
			response.ErrorCode = 10
			response.Message = "Failed"
		} else {
			response.ErrorCode = 0
			response.Message = "Success"
			response.Data = *item
		}
		return c.JSON(http.StatusOK, response)
	})

	route.DELETE("item/delete_item/:id", func(c echo.Context) error {
		item, _ := model.Orders_item.GetItemByid(c.Param("id"))
		response := new(Response)

		if item.Deleteitem() != nil {
			response.ErrorCode = 10
			response.Message = "Failed"
		} else {
			response.ErrorCode = 0
			response.Message = "Success"
		}
		return c.JSON(http.StatusOK, response)
	})

	route.GET("item/get_Orders_item", func(c echo.Context) error {
		response := new(Response)
		Orders_item, err := model.Orders_item.GetItems(c.QueryParam("keywords"))
		if err != nil {
			response.ErrorCode = 10
			response.Message = "Failed"
		} else {
			response.ErrorCode = 0
			response.Message = "Success"
			response.Data = Orders_item
		}
		return c.JSON(http.StatusOK, response)
	})

	// history method
	route.POST("history/create", func(c echo.Context) error {
		history := new(model.Order_histories)
		c.Bind(history)
		response := new(Response)
		if history.Createhistory() != nil {
			response.ErrorCode = 10
			response.Message = "Failed"
		} else {
			response.ErrorCode = 0
			response.Message = "Success"
			response.Data = *history
		}
		return c.JSON(http.StatusOK, response)
	})

	route.PUT("history/update_history/:id", func(c echo.Context) error {
		history := new(model.Order_histories)
		c.Bind(history)
		response := new(Response)
		if history.Updatehistory(c.Param("id")) != nil {
			response.ErrorCode = 10
			response.Message = "Failed"
		} else {
			response.ErrorCode = 0
			response.Message = "Success"
			response.Data = *history
		}
		return c.JSON(http.StatusOK, response)
	})

	route.DELETE("history/delete_history/:id", func(c echo.Context) error {
		history, _ := model.Order_histories.GetHistoryByid(c.Param("id"))
		response := new(Response)

		if history.Deletehistory() != nil { // method update history
			response.ErrorCode = 10
			response.Message = "Failed"
		} else {
			response.ErrorCode = 0
			response.Message = "Success"
		}
		return c.JSON(http.StatusOK, response)
	})

	route.GET("history/get_Order_histories", func(c echo.Context) error {
		response := new(Response)
		Order_histories, err := model.Order_histories.GetHistories(c.QueryParam("id"))
		if err != nil {
			response.ErrorCode = 10
			response.Message = "Failed"
		} else {
			response.ErrorCode = 0
			response.Message = "Success"
			response.Data = Order_histories
		}
		return c.JSON(http.StatusOK, response)
	})

	route.Start(":9000")
}

type Response struct {
	ErrorCode int         `json:"error_code" form:"error_code"`
	Message   string      `json:"message" form:"message"`
	Data      interface{} `json:"data"`
}
