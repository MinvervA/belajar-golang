package controllers

// import (
// 	"crud-beego/repositories" // ✅ Gunakan nama module yang benar
// 	"encoding/json"
// 	"fmt"

// 	"github.com/beego/beego/v2/server/web"
// )

// // UserController menangani request terkait user
// type UserController struct {
// 	web.Controller
// 	repo *repositories.UserRepository // ✅ Gunakan repositories.UserRepository
// }

// // Constructor untuk inisialisasi repository di controller
// func (c *UserController) Prepare() {
//     c.repo = repositories.NewUserRepository() // ✅ Panggil constructor dari repositories
// }

// // @router / [get]
// func (c *UserController) GetAllUsers() {
//     users, err := c.repo.GetAllUsers()
//     if err != nil {
//         fmt.Println("Error Query:", err)
//         c.Ctx.Output.SetStatus(500)
//         c.Data["json"] = map[string]string{"error": "Gagal mengambil data"}
//     } else {
//         // fmt.Println("Data Users:", users)
//         c.Data["json"] = users
//     }
//     c.ServeJSON()
// }

// type Test struct{
// 	Nama string `json:"name"`
// 	Umur string `json:"umur"`

// }

// func (c *UserController) CreateUsers() {
//     var user Test

//     // nama := c.GetString("name")
// 	fmt.Println(">>>>")
// 	// fmt.Println(nama)
// 	rawBody := c.Ctx.Input.RequestBody
// 	// fmt.Println(rawBody)
//     err := json.Unmarshal(rawBody, &user)
//     if err != nil {
//         c.Ctx.Output.SetStatus(400)
//         c.Data["json"] = map[string]string{"error": "Invalid input"}
//         c.ServeJSON()
//         return
//     }
// 	// c.Data["json"] =
//     // ✅ Data berhasil diparsing
//     fmt.Printf("%+v\n", user)

// }

