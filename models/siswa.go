package models

type Siswa struct {
    ID      uint   `gorm:"primaryKey"`
    Nama    string `json:"nama"`
    NIM     string `json:"nim"`
    Kelas   string `json:"kelas"`
    QRCode  string `json:"qr_code"`
}