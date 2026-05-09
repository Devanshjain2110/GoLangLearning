package main
import (
	"fmt"
	"time"
)

func main(){
	i := 3
	switch  i {
	case 1: 
	fmt.Println("Number is 1")
	case 2:
	fmt.Println("Number is 2")
	case 3:
	fmt.Println("Number is 3")
	}

	 switch time.Now().Weekday() {
    case time.Saturday, time.Sunday,:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

	whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
}
  whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}