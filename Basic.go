package main //This line declares that this Go source file is part of the main package, which is a special package that indicates that this is the entry point of a standalone executable program.

import ("fmt") //This line imports the fmt package, which provides functions for formatting and printing output to the console.

func main() { //declares a function named main, which is the entry point of the program.
  var student1 string = "John"
  var student_roll int = 12
  x:=5 //declare a new variable called 'x' with an inferred type of integer and assighn 2 to it
  //:= Can only be used inside functions
  var y int

  fmt.Println(x)
  fmt.Println(student1)
  fmt.Println(student_roll)

  fmt.Println(y)
  y=34 //In Go, the assignment operator "=" is used to assign a value to a pre-declared variable
  fmt.Println(y)

  var a,b,c int = 1,2,3
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  
  //Arrays
  var arr1=[...]int{1,2,3,4}
  var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  fmt.Println(arr1)
  fmt.Println(cars)
  fmt.Println(cars[2])//Accessing array element
  cars[2]="mercedez" //changing array values
  fmt.Println(cars[2])

  arr2 := [5]int{} //not initialized
  arr3 := [5]int{1,2} //partially initialized
  arr4 := [5]int{1,2,3,4,5} //fully initialized

  fmt.Println(arr2)
  fmt.Println(arr3)
  fmt.Println(arr4)

  arr5 := [5]int{1:10,2:40}// initializes only the second and third elements
  fmt.Println(arr5)

  //Slices
  // Like arrays, slices are also used to store multiple values of the same type in a single variable.
  // it's important to note that once an array is created with a fixed length, its length cannot be changed. So while you can use ... to create an array of a variable length, it's still a fixed-size data structure and its size cannot be changed at runtime.

  myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
  fmt.Println(len(myslice2))//returns the number of elements currently present in the slice, which is 4.
  fmt.Println(cap(myslice2))// if we create a slice with a length of 3 and a capacity of 5, it means that the slice can hold up to 5 elements, but only 3 elements are currently stored in the slice.
  fmt.Println(myslice2)

  array1 := [6]int{10, 11, 12, 13, 14,15}
  myslice := array1[2:4]

  fmt.Printf("myslice = %v\n", myslice)
  //make() function to create slice:
  myslice1 := make([]int, 5, 10)
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))

  temperature := 14
  if (temperature > 15) {
    fmt.Println("It is warm out there")
  } else { //don't put else on next line bcoz go treats a newline character immediately following an if or else statement as the end of that statement.
    fmt.Println("It is cold out there")
  }

  time := 22
  if time < 10 {
    fmt.Println("Good morning.")
  } else if time < 20 {
    fmt.Println("Good day.")
  } else {
    fmt.Println("Good evening.")
  }

  day := 4

  switch day {
  case 1:
    fmt.Println("Monday")
  case 2:
    fmt.Println("Tuesday")
  case 3:
    fmt.Println("Wednesday")
  default:
    fmt.Println("Any other day")
  }
  //All the case values should have the same type as the switch expression. Otherwise, the compiler will raise an error

  //Multi-Case Switch
  day1 := 5

   switch day1 {
   case 1,3,5:
    fmt.Println("Odd weekday")
   case 2,4:
     fmt.Println("Even weekday")
   case 6,7:
    fmt.Println("Weekend")
  default:
    fmt.Println("Invalid day of day number")
  }

  //for loop is the only loop available in Go.

  for i:=0; i < 5; i++ {
    if i == 3 {
      continue
    }
   fmt.Println(i)
  }

  numbers := [5]int{1, 2, 3, 4, 5}
  for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
  }

  fruits := [3]string{"apple", "banana", "orange"}
  for _, fruit := range fruits {
    fmt.Println(fruit)
  }
  //we don't need the index of each element, we use the blank identifier _ to ignore it.
}

/*
In Go, statements are separated by ending a line (hitting the Enter key) or by a semicolon ";".

Hitting the Enter key adds ";" to the end of the line implicitly (does not show up in the source code).
The left curly bracket { cannot come at the start of a line.

You can write more compact code, like shown below (this is not recommended because it makes the code more difficult to read):
package main; import ("fmt"); func main() { fmt.Println("Hello World!");}
*/
