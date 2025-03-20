package repositories

import (
	"crud-beego/models"

	"github.com/beego/beego/v2/client/orm"
)

// UserRepository untuk mengelola operasi database User
type UserRepository struct {
	db orm.Ormer
}

// Constructor untuk inisialisasi repository
func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: orm.NewOrm(),
	}
}

// Mengambil semua user dari database
func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	// melakukan query untuk mengambil semua data
	_, err := r.db.QueryTable("users").All(&users)
	return users, err
}

// membuat user baru

