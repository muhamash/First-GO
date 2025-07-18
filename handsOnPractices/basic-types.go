package main

import (
	"fmt"
	"math/cmplx"
)

var(
	ToBe bool = false
	MaxInt uint64 = 1<<64-1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main(){
	v := 1212
	fmt.Printf("type of %T value is %v\n", ToBe, ToBe)
	fmt.Printf("type of %T value is %v\n", MaxInt, MaxInt)
	fmt.Printf("type of %T value is %v\n", z, z)
	fmt.Printf("type of %T value is %v\n", v, v)
}