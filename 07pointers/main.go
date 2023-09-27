package main

import "fmt"

func main ()  {
	fmt.Println("Hello World")

	var ptr *int
	fmt.Println("ptr" ,ptr) // default value is nil

	myNumber := 23

	// for referencing we use "&"
	var ptrTomyNumber = &myNumber // creating a pointer referencing to muNumber
	fmt.Println("value of actual pointer is " , ptrTomyNumber) // pointer to a direct reference to memory location
	fmt.Println("value of actual pointer is " , *ptrTomyNumber) // * means I want to see inside pointer

	// *ptr -> means value inside ptr
	*ptrTomyNumber = *ptrTomyNumber * 2
	fmt.Println("new value is " , myNumber)

	// pointer gives the gurantee that the operation is going to be performed on the actual value and not on the copies on the value
}