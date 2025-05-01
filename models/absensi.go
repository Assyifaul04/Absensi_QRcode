package models

import "time"

type Absensi struct {
    ID      uint      `gorm:"primaryKey"`
    Nama    string    `json:"nama"`
    NIM     string    `json:"nim"`
    Kelas   string    `json:"kelas"`
    Status  string    `json:"status"`
    Tanggal time.Time `json:"tanggal"`
}