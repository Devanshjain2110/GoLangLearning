package main

import "fmt"

func main(){
	if 9%2 == 0{
		fmt.Println("This is an even integer")
	} else {
		fmt.Println("This is an odd integer")
	}

	if n := 9; n<0 {
		fmt.Println("Hey, this seems to be a negative number")
	} else if n < 10 {
		fmt.Println("This numbers seems to be a single digit")
	} else {
		fmt.Println("This is definitely a bigger number")
	}
}