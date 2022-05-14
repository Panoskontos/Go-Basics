// Go is organised in packeages that contain files
package main

// import library that allowas you to output and input
import (
	"fmt"
	"time"
)





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
	// fmt.Scan(&username)
	// & is reference to memory location
	
	fmt.Printf("Hello %v welcome to the game \n", username)
	
	// Conditionals
	fmt.Printf("Enter your age?: ")
	var age int64
	// fmt.Scan(&age)
	fmt.Printf("\nYour age is: %v \n",age)

	if age >= 10 {
		fmt.Println("You are allowed to play")
	} else if age >=8 {
		fmt.Println("You need a litle more time kido")
	} else {
		fmt.Println("You are not allowed")
	}

	// Arrays
	var arr = [...]string{"hello","man","how","are","you"} 
	fmt.Println(arr[0])
	fmt.Println(arr[0:3])
	fmt.Println(len(arr))

	// Slices arrays that have dynamic size
	// Because we don't always know the size
	var slice = []int{1,2,3}
	fmt.Println(slice)
	fmt.Println(cap(slice))

	// Loops
	for i:=0;i<5;i++ {
		fmt.Println(i)
	}

	// Iterate arrays
	var arr2 = [...]int8{11,22,33,44}
	fmt.Println(arr2)
	for idx,val := range arr2{
		fmt.Println(idx,val)
	}
	// if you don't want to show index
	for _ , val := range arr2{
		fmt.Println(val)
	}

	// Go while loops are also for
	i := 0
	for i<8{
		fmt.Println(i)
		i+=1
	}


	printname(name)
	fmt.Println(mymath(2,567))
	fmt.Println(mymath2(2,567))

	// Hash Map
	var a = make(map[string]string)
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"
	a["year"] = "0"
	a["liked"] = "a lot"

	fmt.Println(a)
	for k, v := range a {        // loop with no order
		fmt.Printf("%v : %v, ", k, v)
	  }


	// Sctructs
	// Go uses structs instead of classes

	type Human struct {
		name string
		age int
		job string
	}

	var p1 Human
	fmt.Println(p1)  

	p1.age = 45
	p1.job = "Dev"
	p1.name = "Poli"
	fmt.Println(p1)  


	// Type Conversions
	var percent float64
	percent = float64(13) / float64(23)
	fmt.Println(percent)

	// Pointer (Took it from C)
	// A variable that points to the memory location of a variable
	fmt.Println(&percent)
	var value int
	// scan can assign value because it knows the memory
	// otherwise it can't
	
	
	// Go routines
	// adding go adds a new goroutine to run in parallel
	go sendTicket(1)
	
	// While i wait to input it should run the send ticket
	fmt.Scan(&value)
	
	
	// Slices again are more flexible and better to use in general
	var dynamic = []int{1,2,3}
	// adding something and changing size of array
	fmt.Println(len(dynamic))
	dynamic = append(dynamic, 4)
	fmt.Println(len(dynamic))

	


	// Import package helper
	fmt.Println(power(3,2))




}


// FUNCTIONS
// Functions you can put on top or bottom
func printname(name string){
	fmt.Printf("I got executed %v", name)
}

func mymath(a int, b int) int{
	return a+b
}

// Multiple returns
func mymath2(a int, b int) (result1 int, result2 int){
	result1 = a*b
	result2 = a-b
	return
}

// CONCURRENCY
// Goroutines
func sendTicket(tickets int){
	
// If it would take a lot of time?
time.Sleep(10*time.Second)

var tick = fmt.Sprintf("%v ticket for user",tickets)
fmt.Printf("\n sending ticket %v  \n  to email \n", tick)


}