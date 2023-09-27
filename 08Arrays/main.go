package main

import "fmt"

func main() {
	fmt.Println("Arrays Doc")
	var fruits [5]string

	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[3] = "Guava"

	fmt.Println("Fruits : ", fruits) // Fruits :  [Apple Orange  Guava ]
	// there is a blank space at index 2

	fmt.Println("Length of fruit", len(fruits)) // 5 -> dosent rely on the values stored

	var veg = [5]string{"A", "B", "C"}

	fmt.Println("veg : ", veg) // veg :  [A B C  ]

	var arr1 = [...]int{1, 2, 3} // length is caluclated from the elements
	fmt.Println(arr1, len(arr1))
	// arr1[3] = 4 // giving error out of bound

	arr2 := [5]int{1: 10, 2: 40} // initializing only at particular indexes

	fmt.Println(arr2)

}
