package main

import (
	"fmt"
	"time"
)

func  main () {
	fmt.Println("hello world")
	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // in the doc , its mentioned to use this time 01-02-2006 15:04:05 Monday as standard 
	createDate := time.Date(2020,time.April,10,20,3,0,0,time.UTC)

	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday")) 
}