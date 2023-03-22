package main
import ("fmt")

//Function Signature:defines the name of the function, its parameters, and its return types.

func greet(name string) string { //here string is the return type of the function
    // Function body
    message := "Hello, " + name + "!"
    return message
}

//Named Return Value: we name the return value as result (of type int), and return the value means return statement without specifying the variable name
func myFunction(x int, y string) (result int, txt1 string) { //multiple return value
	result = x + x
	txt1 = y + " World!"
	return
  }

func main() {
    // Function call
    greeting := greet("John")
    fmt.Println(greeting)

	a, b := myFunction(5, "Hello") //If we (for some reason) do not want to use some of the returned values, we can add an underscore (_), to omit this value.
	fmt.Println(a, b)

}