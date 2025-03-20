package main

import "fmt"

var score int= 90

func updateScore(x  int)int{
	return score + x;
}

func main(){
	gadget :=map[string]int{
		"Laptop":10,
		"Mouse":10,
		"Keyboard":15,
	}
	// cara print satu dari dari map
	fmt.Println(gadget["Mouse"])
	// menambahkan value pada map
	gadget["Mouse"] += 5
	fmt.Println(gadget["Mouse"])
	// memberikan key baru
	gadget["Monitor"] = 5
	fmt.Println(gadget["Monitor"])
	// meremove data dari map
	delete(gadget,"Keyboard");
	fmt.Println(gadget)

	x := 10
    y := x
    y = 20

    fmt.Println("x:", x)
    fmt.Println("y:", y)

	hasil := updateScore(10)
	fmt.Println(hasil)
}
