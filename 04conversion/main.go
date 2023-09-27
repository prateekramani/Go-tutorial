package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to the app")
	fmt.Println("Give the number between 1-5")
	
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	
	fmt.Println("the input value is : ", input)

	rating, err := strconv.ParseFloat(strings.TrimSpace(input) , 64)
	
	if err != nil {
		fmt.Println("error : ", err)
	} else {
		fmt.Println("Final Rating is : ", rating + 1)
	}

}
