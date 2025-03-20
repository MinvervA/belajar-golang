package main

import "fmt"

func main(){
	fmt.Println("Hello World!")
}

// macam macam data type golang
// boolean
var isActive bool = true;

// integer
// int8	8 bit	-128 to 127
// int16	16 bit	-32,768 to 32,767
// int32	32 bit	-2,147,483,648 to 2,147,483,647
// int64	64 bit	-9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
// uint8	8 bit	0 to 255
// uint16	16 bit	0 to 65,535
// uint32	32 bit	0 to 4,294,967,295
// uint64	64 bit	0 to 18,446,744,073,709,551,615

// float
// float32	32 bit	±1.18 x 10⁻³⁸ to ±3.4 x 10³⁸
// float64	64 bit	±2.23 x 10⁻³⁰⁸ to ±1.8 x 10³⁰⁸

// complex
var c complex64 = 1 + 2i

// string
var name string = "Golang"

// array
var arr [5]int = [5]int{1,2,3,4,5}

// slice
// numbers :=[]int{1,2,3,4,5}

//map
person :=map[string]string{
	"name":"Andrean",
	"age":"25"
}