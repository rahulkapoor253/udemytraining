package main

import "fmt"

func res(temp int) int {

	return  temp + 50

}

func main() {

res := res(10)	//here res stores the returned result;

	fmt.Printf("%v\n",res)

}
