package main

import "fmt"

const LoginToken string = "fsdgdhdsj" // starting the variables as capital letter in Go makes it public 

// jwtToken := 3000 throws error, as this style is only allowed inside a function

func main ()  {
	var username string = "Prateek"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n" ,username)


	var isVerified bool = false
	fmt.Println(isVerified)
	fmt.Printf("Variable is of type : %T \n" ,isVerified)

	var smallvar uint8 = 255 // had we put 256 here , it would had given error since 256 is out of bound 
	fmt.Println(smallvar)
	fmt.Printf("Variable is of type : %T \n" ,smallvar)

	var smallFloat float32 = 255.455565656565
	var largeFloat float64 = 255.455565656565
	fmt.Println(smallFloat)
	fmt.Println(largeFloat)
	fmt.Printf("Variable is of type : %T \n" ,smallFloat)
	fmt.Printf("Variable is of type : %T \n" ,largeFloat)


	var anotherVariable int // initialized with 0 default value
	fmt.Println(anotherVariable) 
	fmt.Printf("Variable is of type : %T \n" ,anotherVariable)

	//implicit way of declaring the variables
	
	var website = "myquesans.com" // go automatically fetches the type
	fmt.Println(website)
	fmt.Printf("Variable is of type : %T \n" ,website)
	// website = 2 -> will give error as go has assigned type: string to it automatically

	// no var style of creating variables

	numberOfUser := 3000 // no var keyword , no datatype
	fmt.Println(numberOfUser)
	// numberOfUser := "Prateek" or 200 will give error as go has assigned

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n" ,LoginToken)
	
}