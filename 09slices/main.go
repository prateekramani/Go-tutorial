package main

import (
	"fmt"
	"sort"
)

func main ()  {
	fmt.Println("Slices Doc")

	var fruits = []string{"A" , "B" , "C" } // slices is array without specifing size

	fmt.Println("fruits" , fruits)
	fmt.Printf("type of fruits %T \n" , fruits)

	fruits = append(fruits, "D")

	fmt.Println("fruits" , fruits)

	fruits = append(fruits[1:])
	fmt.Println("fruits" , fruits)

	fruits = append(fruits[1:2]) // 2 is non inclusive here , [:3] is also possible
	fmt.Println("fruits" , fruits)

	// utilising the make keyword
	highScores := make([]int,4)

	highScores[0] = 234
	highScores[1] = 235
	highScores[2] = 236
	highScores[3] = 237
	// highScores[4] = 238  // this will throw error , because of no memory , size of 4

	highScores = append(highScores, 238,239,240) // append reallaocates the memory , no error in this case

	fmt.Println("highscores : ", highScores)

	sort.Ints(highScores)

	fmt.Println("highscores sorted : ", highScores)
	fmt.Println("highscores is sorted ?: ", sort.IntsAreSorted(highScores))


	// removing value from the slice based on the index

	var courses = []string {"Java" , "Python" , "JS" , "Terraform" , "Ansible"}

	fmt.Println("Courses :" , courses)
	removeIndex := 2
	// we can also use append to remove the values

	courses = append(courses[:removeIndex] , courses[removeIndex + 1:]...) // ... is varaible parameters
	fmt.Println("courses" , courses)

}