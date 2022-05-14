package main

// import library that allowas you to output and input 
import  "fmt"

func main() {
	// print line
	fmt.Println("Welcome to my game")

	var name string = "Panos"
	// name = 23 will result in error
	var id int = 4
	var id2 uint = 4
	var id3 float64 = 4.23
	var id4 bool = true



	fmt.Println(name)
	fmt.Println(id)
	fmt.Println(id2)
	fmt.Println(id3)
	fmt.Println(id4)
}
