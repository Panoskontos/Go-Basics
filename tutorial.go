package main

// import library that allowas you to output and input
import "fmt"




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



}

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