package main

// import library that allowas you to output and input 
import  "fmt"

func main() {
	// print line
	fmt.Println("Welcome to my game")
	
	// Variables and data types
	var name string = "Panos"
	// name = 23 will result in error
	var id int = 4
	var id2 uint = 4
	var id3 float64 = 4.23
	var id4 bool = true

	// if you are bored and don't want to define type
	name2 := "Giorgos"


	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(id)
	fmt.Println(id2)
	fmt.Println(id3)
	fmt.Println(id4)


	// Outputing vaules use printf and %v = value
	// see notes for string formats
	fmt.Printf("Hello %v %v %10.2f, how are you? \n", name, id,id3)
	
	// Get user input
	var username string 
	fmt.Scan(&username)
	// & is reference to memory location
	fmt.Printf("Hello %v welcome to the game \n", username)
	
	fmt.Printf("Enter your age?: ")
	var age int64
	fmt.Scan(&age)
	fmt.Printf("\nYour age is: %v",age)


}
