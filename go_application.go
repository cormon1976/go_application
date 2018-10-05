package sample

import "fmt"

func adder(x, y int) int {
    return x + y 
}

func subtracter(x,y int) int {
    return x - y 
}

func main() {

	if (adder(3, 4)) == 7 {
		fmt.Println("adder works")
	} else {
		fmt.Println("7 is odd")
	}
}
