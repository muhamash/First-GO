package main

import "fmt"

func main(){
	name := "ashraful"
	nage := 23
	fmt.Println("Hello Go:", name, nage)
	fmt.Printf("Hello Go: %s and age: %d\n", name, nage)

	const aname, age = "Kim", 22
	fmt.Println(aname, "is", age, "years old.")
	fmt.Printf("%s is %d years old; \t and %T and %T", aname, age, aname, age)

	// string literal
	fmt.Println(`
	UTF-8 saves space. 
	In UTF-8, common characters like “C” take 8 bits, 
	while rare characters like “💕” take 32 bits. 
	Other characters take 16 or 24 bits. 
	A blog post like this one takes about 
	four times less space in UTF-8 
	than it would in UTF-32. 
	So it loads four times faster.`)
}