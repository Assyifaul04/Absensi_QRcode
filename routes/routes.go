package routes

import (
    "absensi-app/controllers"
    "github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
    e.POST("/siswa", controllers.CreateSiswa)
    e.GET("/siswa-list", controllers.GetSiswaList)

    e.POST("/scan", controllers.ScanQRCode)
    e.GET("/absensi/hariini", controllers.GetAbsensiHariIni)
}
