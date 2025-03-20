package main

import (
	"crud-beego/models"
	"crud-beego/routers"
	_ "crud-beego/routers"
	"fmt"
	"log"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB // Declare global db variable

func init() {
	// koneksi dengan database
	dsn := "admin:Vektor100%@tcp(10.10.100.122:3306)/test?charset=utf8&parseTime=True"
	var err error

	// Use the global db variable instead of declaring a local one
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Gagal menghubungkan ke database: %v", err)
	}

	// check koneksinya
	sqlDB, err := db.DB() // This refers to the global db now
	if err != nil {
		log.Fatalf("❌ Gagal mendapatkan koneksi dari GORM: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("❌ Gagal ping database: %v", err)
	}

	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatalf("❌ Gagal migrasi schema: %v", err)
	}

	fmt.Println("✅ Koneksi berhasil ke database!")
}

func main() {
	/// Inject db ke dalam router
	routers.InitRouter(db)
	web.Run()
}
