package main

import (
	// "fmt"
	"fmt"
	"time"
)

func day() string{
	day := time.Now().Weekday()

	switch day{
	case time.Monday:
		return "Monday"
	case time.Tuesday:
		return "Tuesday"
	case time.Wednesday:
		return "Wednesday"
	case time.Thursday:
		return "Thursday"
	case time.Friday:
		return "Friday"
	case time.Saturday:
		return "Saturday"
	case time.Sunday:
		return "Sunday"
	default:
		return "Not found"
	}
}



func main(){
	fmt.Println(day())
}