package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"os"
)

var (
	ListenPort string
)

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		fmt.Println("Failed load env Err: " + err.Error())
		panic(err)
	}

	ListenPort = os.Getenv("LISTEN_PORT")

}
func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "ようこそ、美しい世界へ")
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", ListenPort)))
}
