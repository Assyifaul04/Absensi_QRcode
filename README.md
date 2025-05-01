# Absensi QR Code

Sistem Absensi QR Code adalah aplikasi berbasis web yang memungkinkan proses absensi menggunakan QR Code yang dapat dipindai menggunakan kamera perangkat.

## Fitur Utama

- **Tambah Data Siswa:** Menambahkan siswa baru dengan nama, NIM, dan kelas.
- **Generate QR Code:** Setiap siswa dapat memiliki QR Code yang berisi NIM mereka untuk keperluan absensi.
- **Scan QR Code:** Menggunakan kamera untuk memindai QR Code siswa saat absen.
- **Tabel Absensi:** Menampilkan daftar absensi siswa pada hari ini.
- **Export ke Excel:** Menyediakan fitur untuk mengekspor data absensi dalam format Excel.

## Teknologi yang Digunakan

- **Frontend:** HTML, CSS, Tailwind CSS, JavaScript
- **Backend:** Go (Golang)
- **Database:** Tidak disebutkan (dapat menggunakan database seperti MySQL atau PostgreSQL jika diperlukan)
- **Library:** 
  - QR Code: [qrcode.min.js](https://cdn.jsdelivr.net/npm/qrcode/build/qrcode.min.js)
  - Scan QR Code: [html5-qrcode](https://unpkg.com/html5-qrcode)
  - Export ke Excel: [xlsx.full.min.js](https://cdnjs.cloudflare.com/ajax/libs/xlsx/0.17.3/xlsx.full.min.js)

## Instalasi

1. **Clone repository ini:**

   ```bash
   git clone https://github.com/Assyifaul04/Absensi_QRcode.git
   cd Absensi_QRcode
Instalasi dependensi:

    Untuk Go: Jalankan go mod tidy untuk menginstal dependensi.

Menjalankan server:

    Jalankan server Go dengan perintah:

    go run main.go

    Akses aplikasi:

        Aplikasi akan berjalan di http://localhost:8080/ (pastikan port sesuai dengan konfigurasi yang digunakan di kode kamu).

Penggunaan

    Menambah Siswa:

        Masukkan data siswa (nama, NIM, kelas) dan klik Tambah Siswa.

    Generate QR Code:

        Setelah menambah siswa, klik tombol Generate QR untuk menghasilkan QR Code berdasarkan NIM siswa.

    Scan QR Code:

        Gunakan tombol Mulai Scan QR untuk memulai pemindaian QR Code dengan kamera.

    Export Data Absensi:

        Klik tombol Export ke Excel untuk mengekspor data absensi hari ini ke dalam format Excel.
