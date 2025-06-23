package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func swap(x, y string)(string, string){
	return x, y
}

func split(sum int)(x, y int){
	x = sum * 4 / 9
	y = sum - x
	return
}

var javascript, typescript bool
const i float32 = 100.20

func main(){

	sayHello()

	a := add(1212, 43534)
	c, d := swap("hello", "world")
	fmt.Println(add(2, 5), a, c, d)
	fmt.Println(split(30))

	fmt.Println("testing variables :" ,i , javascript, typescript)
	fmt.Printf("testing variables : %T %T %T" ,i , javascript, typescript)
}

func sayHello(){
	fmt.Println("Say hello function!!")
}