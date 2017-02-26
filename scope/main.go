package main

import "fmt"

var x string = "rahul"

func main() {

	fmt.Printf("%v\n",x)
	func1()
	temp := 10//this temp var local to main;
	fmt.Printf("%v\n",temp)
	func1()

}
func func1() {

	fmt.Printf("%v\n",x)
	//fmt.Printf("%v\n",temp)
}

