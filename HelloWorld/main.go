package main

//All imports for the code
import (
	"fmt"
)

//All variables
var (
	messageTrue       = "Hello world!"
	messageFalse      = "Something is wrong"
	messageCheck bool = true
)

//Quick how a function is built up
func add(a, b int) int {
	return a + b
}

//Main function to run program
func main() {
	if messageCheck {
		fmt.Println(messageTrue)
	} else {
		fmt.Println(messageFalse)
	}

	fmt.Println("\nPress enter to exit...")
	fmt.Scanln()
}
