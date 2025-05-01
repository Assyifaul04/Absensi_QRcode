package controllers

import (
    "absensi-app/models"
    "absensi-app/services"
    "absensi-app/utils"
    "net/http"

    "github.com/labstack/echo/v4"
)

func CreateSiswa(c echo.Context) error {
    var siswa models.Siswa
    if err := c.Bind(&siswa); err != nil {
        return err
    }

    qrPath, err := services.GenerateQRCode(siswa.NIM)
    if err != nil {
        return err
    }
    siswa.QRCode = qrPath

    if err := utils.DB.Create(&siswa).Error; err != nil {
        return err
    }

    return c.JSON(http.StatusOK, siswa)
}

func GetSiswaList(c echo.Context) error {
    var siswa []models.Siswa
    utils.DB.Find(&siswa)
    return c.JSON(http.StatusOK, siswa)
}
