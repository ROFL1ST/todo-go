package main

import (
	"encoding/json"
	"net/http"

	"github.com/ROFL1ST/todo-go-api/database"
	"github.com/labstack/echo"
)

type CreateReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {
	db := database.InitDb()
	defer db.Close()
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"hello": "world",
		})
	})
	e.POST("/todos", func(c echo.Context) error {
		var req CreateReq
		json.NewDecoder(c.Request().Body).Decode(&req)
		_, err := db.Exec(
			"INSERT INTO todos(title, description, complete) VALUES(?, ?, 0)",
			req.Title,
			req.Description,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,  err.Error())
		}
		return c.JSON(http.StatusOK, "OK")
	})
	e.Logger.Fatal(e.Start(":9000"))
}
