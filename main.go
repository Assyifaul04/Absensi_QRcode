package main

import (
    "absensi-app/routes"
    "absensi-app/utils"
    "github.com/labstack/echo/v4"
    "log"
)

func main() {
    utils.ConnectDatabase()
    e := echo.New()

    e.Static("/", "public")

    routes.SetupRoutes(e)

    e.GET("/", func(c echo.Context) error {
        return c.File("public/index.html")
    })

    log.Println("Server berjalan di port 8080")

    e.Logger.Fatal(e.Start(":8080"))
}
