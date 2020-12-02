package main

import "fmt"

type subscriber struct {
	name string
	rate int
	active bool
}
func main(){
	var s1 subscriber

	s1.name = "kim"
	s1.rate = 5000
	s1.active = false

	fmt.Printf("%s\n",s1.name)
	fmt.Println(s1.rate)
	fmt.Println(s1.active)
}

