package controllers

import (
	"crud-beego/models"
	"encoding/json"
	"fmt"

	"github.com/beego/beego/v2/server/web"
	"gorm.io/gorm"
)

// untuk menangani request
type ProductController struct{
	web.Controller
	DB *gorm.DB
}

// create product
func (c *ProductController) CreateProduct(){
	c.EnableRender = false
	//  cara check value yang dikirim dari post
	fmt.Println("Request Body:", string(c.Ctx.Input.RequestBody))

	var product models.Product
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &product); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Invalid data"}
		c.ServeJSON()
		return
	}

	// ✅ Log data yang diterima di console
	fmt.Printf("Received data: %+v\n", product)

	// mengirim data dengan format json
	c.Data["json"] = map[string]interface{}{
		"isSuccess":true,
		"data":map[string]interface{}{
			"nama":product.Name,
			"price":product.Price,
		},
	}
	c.ServeJSON()

	// db := c.GetDB()
	if err := c.DB.Create(&product).Error; err != nil{
		c.Ctx.Output.SetStatus(500)
		c.Data["json"]=map[string]string{"error":"failed to create Product"}
		c.ServeJSON()
		return
	}
		// mengirim data dengan format json
		c.Data["json"] = map[string]interface{}{
			"isSuccess":true,
			"message":"product success created!",
			"data":map[string]interface{}{
				"nama":product.Name,
				"price":product.Price,
			},
		}
}

func (c *ProductController) GetAllData(){
	c.EnableRender = false
	var products []models.Product
	if err := c.DB.Find(&products).Error; err != nil{
		fmt.Println("Gagal mendapatkan data")
		return
	}
	fmt.Println("Daftar Product :")
	for _,product := range products{
		fmt.Printf("ID: %d , Nama Product : %s, Price : %f \n",product.ID,product.Name,product.Price)
	}

	c.Data["json"] = map[string]interface{}{
		"isSucces":true,
		"message":"Berhasil mendapat kan semua data",
		"data":products,
	}
	c.ServeJSON()
	// fmt.Println(products)
}

func (c *ProductController) UpdateProduct(){
	c.EnableRender=false
	// cara mengambil id pada query (seperti req param)
	id, err := c.GetInt(":id")
	fmt.Println("ini response nya : ",id)
	// var product models.Product
	if err != nil{
		c.Ctx.Output.SetStatus(400)
		c.Data["json"]=map[string]string{"error":"berikan id benar!"}
		c.ServeJSON()
		return
	}

	//  cara check value yang dikirim dari post
	fmt.Println("Request Body:", string(c.Ctx.Input.RequestBody))

	var newProduct models.Product
	if err := json.Unmarshal(c.Ctx.Input.RequestBody,&newProduct); err!= nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"]=map[string]string{"error":"tidak mendapatkan input"}
		c.ServeJSON()
		return
	}
	fmt.Println(newProduct)

	if err := c.DB.Model(&models.Product{}).Where("id = ?",id).Updates(&models.Product{Name:newProduct.Name,Price:newProduct.Price}).Error;err != nil{
	// if err := c.DB.Model(&models.Product{}).Where("id",1).Update("price",newProduct.Price).Error;err != nil{
		c.Ctx.Output.SetStatus(400)
		c.Data["json"]=map[string]string{"error":"tidak berhasil update"}
		fmt.Println("tidak berhasil update")
		c.ServeJSON()
		fmt.Println("ini error :",err)
		return
	}
	fmt.Println("disini ku menunggu")
	// var product models.Product
	// if err := c.DB.First(&product,id).Error;err != nil{
	// 	c.Ctx.Output.SetStatus(400)
	// 	c.Data["json"]=map[string]string{"error":"tidak menemukan data tersebut!"}
	// 	c.ServeJSON()
	// 	return
	// }

	// if newProduct.Name != "" {
		// product.Name = newProduct.Name
	// }

	// product.Price = newProduct.Price
}

func (c *ProductController) DeleteProduct(){
	c.EnableRender=false
	// cara mengambil id pada query (seperti req param)
	id, err := c.GetInt(":id")
	if err != nil{
		c.Ctx.Output.SetStatus(400)
		c.Data["json"]=map[string]string{"error":"berikan id benar!"}
		c.ServeJSON()
		return
	}
	fmt.Println("berikut ini id nya",id)
	var product models.Product
	if err := c.DB.Delete(&product,id).Error;err!=nil{
		c.Ctx.Output.SetStatus(400)
		c.Data["json"]=map[string]string{"error":"tidak menemukan product berdasarkan id"}
		c.ServeJSON()
		return
	}
	c.Data["json"]=map[string]interface{}{
		"isSuccess":true,
		"message":"Berhasil Menghapus Product!",
	}
	c.ServeJSON()
	// cara membuka kembali deleted_At
	// db.Model(&product).Where("id = ?", productID).Update("deleted_at", nil)
}