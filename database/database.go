package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		panic("DATABASE_URL tidak ditemukan! Pastikan sudah diset di Railway.")
	}

	var err error
	DB, err = sql.Open("postgres", databaseUrl)
	if err != nil {
		panic("Gagal membuka koneksi: " + err.Error())
	}

	err = DB.Ping()
	if err != nil {
		panic("Gagal connect ke database: " + err.Error())
	}

	fmt.Println("Berhasil connect ke database!")
}
