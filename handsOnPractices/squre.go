package main

import (
	"fmt"
	"math"
)


func main(){
	const pi = 3.14;
	var x,y int = 39, 49
	var f float64 = math.Sqrt(float64(x*x + y*y));
	z := uint(f)
	fmt.Println(x,y,f,z, pi)
	fmt.Printf("%v is type %T", pi, pi)
}