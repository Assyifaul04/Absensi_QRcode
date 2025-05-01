package services

import (
    "fmt"
    "os"
    "github.com/skip2/go-qrcode"
)

func GenerateQRCode(nim string) (string, error) {
    os.MkdirAll("public/qr", os.ModePerm)

    path := fmt.Sprintf("public/qr/%s.png", nim)
    err := qrcode.WriteFile(nim, qrcode.Medium, 256, path)
    if err != nil {
        return "", err
    }
    return "/" + path, nil
}
