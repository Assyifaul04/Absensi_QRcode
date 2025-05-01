package utils

import (
    "absensi-app/models"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "root:@tcp(127.0.0.1:3306)/absensi_db?charset=utf8mb4&parseTime=True&loc=Local"
    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    DB = database
    log.Println("Database connected successfully!")

    err = DB.AutoMigrate(&models.Siswa{}, &models.Absensi{})
    if err != nil {
        log.Fatal("Failed to auto-migrate tables:", err)
    }
    log.Println("Database tables migrated successfully!")
}
