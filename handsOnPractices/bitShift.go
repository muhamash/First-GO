package main

import "fmt"

const (
	bigInt = 1 << 100
	small = bigInt >> 99
)

func needInt(x int) int {
	return x * 10 + 1;
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main(){
	fmt.Println(needInt((small)))
	// fmt.Println(needInt((bigInt)))
	fmt.Println(needFloat((small)))
	fmt.Println(needFloat((bigInt)))
}