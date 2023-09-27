package main

import (
	"bufio"
	"fmt"
	"os"
)

func main ()  {
	welcome := "Hello World"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the reading")

	// comma ok or err ok syntax , since there is no try-catch syntax here

	input , _ := reader.ReadString('\n') // read till end of the line 
	fmt.Println("Input value is : ", input)
	fmt.Printf("Type of Input is %T", input)
	// %T works with Printf only & not with Println
	// fmt.Printf("Variable is of type : %T \n" ,input)


}
