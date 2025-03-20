package main

import (
	// "bufio"
	"fmt"
	"strings"
)

func main(){
	// cara mendeklare variable dan assigned
	var test string = "andrean"
	fmt.Println(test)
	// cara mendeklare variable tanpa assigned
	// harus diberikan data type nya
	// akan tetap ada warning untuk variable digunakan
	var nama string
	nama = "joe biden"
	fmt.Println(nama)
	price := 500
	fmt.Printf("Ini Harganya %v \n",price)

	// cara type of
	// tapi harus menggunakan Printf
	fmt.Printf("price %T\n",price)

	// disini cara mendapatkan input user
	// harus menggunakan parsing by refrence bukan by value layaknya js
	// jadi diawali dengan &variable
	// jadi yang dikirimkan hanya alamatnya
	// contoh
	// var nama_depan string
	// var nama_belakang string
	// var umur string
	// var email string
	// fmt.Println("Siapakah nama depan mu?")
	// fmt.Scanln(&nama_depan)
	// fmt.Println("Siapakah nama belakang mu?")
	// fmt.Scanln(&nama_belakang)
	// fmt.Println("Berapakah umur mu?")
	// fmt.Scanln(&umur)
	// fmt.Println("apa nama email mu?")
	// fmt.Scanln(&email)

	// fmt.Printf("Terimakasih %v %v yang berumur %v tahun, nanti akan di informasikan kembali ke email %v",nama_depan,nama_belakang,umur,email)

	// reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Siapakah nama depan mu? ")
	// namaDepan, _ := reader.ReadString('\n')
	// namaDepan = strings.TrimSpace(namaDepan)

	// fmt.Print("Siapakah nama belakang mu? ")
	// namaBelakang, _ := reader.ReadString('\n')
	// namaBelakang = strings.TrimSpace(namaBelakang)

	// fmt.Print("Berapakah umur mu? ")
	// age, _ := reader.ReadString('\n')
	// age = strings.TrimSpace(age)

	// fmt.Print("Apa nama email mu? ")
	// emailPengguna, _ := reader.ReadString('\n')
	// emailPengguna = strings.TrimSpace(emailPengguna)

	// fmt.Printf("\nTerima kasih %v %v yang berumur %v tahun, nanti akan diinformasikan kembali ke email %v\n", namaDepan, namaBelakang, age, emailPengguna)

	// cara mendeklare empty array
	var bookings = [50]string{}
	fmt.Printf("%T \n",bookings)

	// cara mendeklare slice ( array yang tidak perlu di tentukan panjangnya)
	var pembelian []string
	// atau seperti ini 
	// var pembelian =[]string{}
	pembelian = append(pembelian, "test")
	fmt.Println(pembelian)

	// cara print array for of jika di js
	datas := [3]string{"jojo the","andrean the","jody the"}
	var firstnames []string
	for _,data := range datas{
		fmt.Println(data)
		names := strings.Fields(data)
		firstnames = append(firstnames, names[0])
	}
	fmt.Println(firstnames)
	fmt.Println(len(test))
	for _,data:=range test{
		// fmt.Println(i)
		charc := strings.Fields(string(data))
		fmt.Println(charc)
	}
	email := "Andrean@gmail.com"
	isEmail := strings.Contains(email,"@")
	fmt.Println(isEmail)
}