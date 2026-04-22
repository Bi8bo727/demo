package _const

import "fmt"

func init() {
	fmt.Println("const/Day.go")
}

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
	Myday
	Yourday
)
