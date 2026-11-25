package main

import (
	"fmt"
	"strings"
	"sort"
)



func add(i uint32, j uint32) uint32 {
	return i + j
} 

func main() {
	fmt.Println("Hello, World")

	//strings

	var nameOne string = "Ishema"
	var nameTwo = "Gigi"
	nameThree := "Nigga"

	fmt.Println(nameOne, nameThree, nameTwo)


	//integers

	var number1 int = 234
	var number2 = 4444
	number3 := 333

	fmt.Println(number1, number2, number3)


	//floats  we must declare explicitly the number of bits to use either `float32 or `float64`

     var score1 float32 = 25.98
	 var score2 float64 = 6782.578

     fmt.Println(score1, score2)


	 // How fmt works and the resulting methods

     // Print()
	 fmt.Print("Hello, ")
	 fmt.Print("World")

	 //Println
	 fmt.Println("Hello ninjas")
	 fmt.Println("My score in math is", score1 , "out of 30")
    
	 //Printf(formatted Strings)
	 fmt.Printf("My name is %v", nameOne)  // %v format specifier for variables
     // %T gives the data type of the passed variable
	 // %f used for floats
	 // %q used for strings


	 //Sprintf(Saving formatted strings)

	 number := fmt.Sprintf("My Home town is %v km away", score2)
	 fmt.Println(number)



	 // Arrays

	var ages = [3]int{20, 29, 24}
	 age := [4]float64{40.0, 30.89, 49.333}


	 fmt.Println(ages[2])
	 fmt.Println(age)


	 // Slices

	 var scores = []int{39, 47, 56, 78}
	 scores = append(scores, 34)

	 
    // Slice Ranges 
	
	var rangeOne = scores[1:2]
	var rangeTwo = scores[2:]
	var rangeThree = scores[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	// The STD package

     //Strings

	 greeting := "Holla wassup everyone"

	 fmt.Println(strings.ReplaceAll(greeting, "Holla", "Hi"))
	 fmt.Println(strings.ToUpper(greeting))
	 fmt.Println(strings.ToLower(greeting))
    

	 // Sorting 

	 newAges := []int{34, 5, 54, 54, 35, 33, 53, 56}
	 sort.Ints(newAges)
	 fmt.Println(newAges)

	 fmt.Println(sort.SearchInts(newAges, 44))
	  fmt.Println(sort.SearchInts(newAges, 54))
    

	// Loops 

	x := 0

	for x < 5 {
		fmt.Println(x)
		x++
	}

	for i:= 0; i < len(newAges); i++ {
		fmt.Println("Todays age is:", newAges[i])
	}




	
	// conditionals

	cd := 45 

	if cd <= 50 {
		fmt.Println("You have normal cd")
	} else if cd > 50 {
        fmt.Println("You have a lot of cd.") 
	}



	// functions

	sum := add(10, 75)
	fmt.Println("The sum is", sum)

	fn1, sn1 := getInitials("Mario Luigi")
    fmt.Println(fn1, sn1)

	// Maps

	  menu := map[string] float64 {
		 "soup" : 3.54,
		 "pie" : 5.78,
		 "rice" : 6.79,
		 "ice cream" : 3.35,
	  }

	  fmt.Println(menu["pie"])


	     // looping into maps 

		 for k, v := range menu {
			fmt.Printf("The price of %v is %v \n", k, v)
		 }

    // Import from the greeting
	sayHello("Ishema")
}


func getInitials(name string) (string, string) {
     strings.ToUpper(name)
	 names := strings.Split(name, " ")
	 
	 var initials []string

	 for _,v := range names {
		initials = append(initials, v[:1])
	 }

	 if len(initials) > 1 {
		return initials[0], initials[1]
	 }

	 return initials[0], "_"
}







