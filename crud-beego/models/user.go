package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// // Struktur model User
// type User struct {
// 	Id    int    `json:"id" orm:"auto"`
// 	Employee string `json:"employee_id"`
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// 	Password string `json:"password"`
// 	Picture string `json:"picture"`
// 	Remember_Token string `json:"remember_token"`
// }

// User struct untuk representasi tabel database
type User struct {
	Id             int       `orm:"auto;pk" json:"id"`
	EmployeeId     string    `orm:"size(50);null" json:"employee_id"`
	Name           string    `orm:"size(191);null" json:"name"`
	Email          string    `orm:"size(191);unique" json:"email"`
	Password       string    `orm:"size(191)" json:"password"`
	Picture        string    `orm:"type(text);null" json:"picture"`
	RememberToken  string    `orm:"size(100);null" json:"remember_token"`
	CreatedAt      time.Time `orm:"auto_now_add;type(timestamp)" json:"created_at"`
	UpdatedAt      time.Time `orm:"auto_now;type(timestamp)" json:"updated_at"`
	DeletedAt      time.Time `orm:"null;type(timestamp)" json:"deleted_at"`
	GolonganUser   int       `orm:"null" json:"golongan_user"`
	Cabang         int       `orm:"null" json:"cabang"`
	PhoneCountryCode string  `orm:"size(3);null" json:"phone_country_code"`
	Phone          string    `orm:"size(20);null" json:"phone"`
	Status         string    `orm:"size(1);null" json:"status"`
	ChatId         int       `orm:"size(11);null" json:"chat_id"`
}

func (u *User) TableName() string{
	return "users" // Change t
}

func init() {
    orm.RegisterModel(new(User))
}
