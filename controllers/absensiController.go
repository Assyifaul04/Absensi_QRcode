package controllers

import (
    "absensi-app/models"
    "absensi-app/utils"
    "net/http"
    "time"

    "github.com/labstack/echo/v4"
)

func ScanQRCode(c echo.Context) error {
    type ScanRequest struct {
        NIM string `json:"nim"`
    }
    var req ScanRequest
    if err := c.Bind(&req); err != nil {
        return err
    }

    var siswa models.Siswa
    result := utils.DB.Where("nim = ?", req.NIM).First(&siswa)
    if result.Error != nil {
        return c.JSON(http.StatusNotFound, echo.Map{"message": "Siswa tidak ditemukan"})
    }

    var absensi models.Absensi
    today := time.Now().Format("2006-01-02")
    result = utils.DB.Where("nim = ? AND DATE(tanggal) = ?", req.NIM, today).First(&absensi)
    if result.Error != nil {
        absensi = models.Absensi{
            Nama:    siswa.Nama,
            NIM:     siswa.NIM,
            Kelas:   siswa.Kelas,
            Status:  "Tidak Hadir",
            Tanggal: time.Now(),
        }
        if err := utils.DB.Create(&absensi).Error; err != nil {
            return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Gagal mencatat absensi"})
        }
    }

    absensi.Status = "Hadir"
    if err := utils.DB.Save(&absensi).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Gagal memperbarui status absensi"})
    }

    return c.JSON(http.StatusOK, echo.Map{"message": "Absensi berhasil dicatat"})
}


func GetAbsensiHariIni(c echo.Context) error {
    var absensi []models.Absensi
    today := time.Now().Format("2006-01-02")
    utils.DB.Where("DATE(tanggal) = ?", today).Find(&absensi)
    return c.JSON(http.StatusOK, absensi)
}
