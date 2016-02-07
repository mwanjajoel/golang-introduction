package main

//including Packages we shall be using in this simple program
import (
	"fmt"
)

func main() {

	fmt.Println("Wassup my people \n")

	//calling methods that are defined at the bottom
	dealingWithNumbers()
	getStringLength()
}

//dealing with numbers, in a separate method
func dealingWithNumbers() {

	fmt.Println("1 + 1 =", 1+1)

}

//lets get the length of a string, real quick
func getStringLength() {

	fmt.Println(len("Wassup my people"))

}
