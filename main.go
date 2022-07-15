package main

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	// sendEmail()
	e := echo.New()

	e.POST("/callback", func(c echo.Context) error {
		var iface interface{}
		c.Bind(&iface)
		asByteJSON, _ := json.Marshal(iface)
		fmt.Println("masuk : ", string(asByteJSON))
		return c.JSON(200, map[string]interface{}{
			"message": "ok",
		})
	})

	e.Start(":8080")
}
