package routers

import (
	"crud-beego/controllers"

	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) {
	// Inject db ke controller
	productController := &controllers.ProductController{DB: db}

	// Daftarkan route
	beego.Router("/products", productController, "post:CreateProduct")
	beego.Router("/getAllData",productController,"get:GetAllData")
	beego.Router("/editProduct/:id",productController,"put:UpdateProduct")
	beego.Router("/deleteProduct/:id",productController,"delete:DeleteProduct")
}
