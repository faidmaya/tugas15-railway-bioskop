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

	// 🔥 Auto-create table bioskop jika belum ada
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS bioskop (
		id SERIAL PRIMARY KEY,
		nama VARCHAR(255) NOT NULL,
		lokasi VARCHAR(255) NOT NULL,
		rating NUMERIC(2,1)
	);`

	_, err = DB.Exec(createTableSQL)
	if err != nil {
		panic("Gagal membuat table bioskop: " + err.Error())
	}

	fmt.Println("Table 'bioskop' siap digunakan!")
}
