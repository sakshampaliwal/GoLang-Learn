// Maps: Maps are used to store data values in key:value pairs.A map is an unordered and changeable collection that does not allow duplicates.
// The default value of a map is nil.
//It is similar to dictionary
package main
import ("fmt")

func main() {
	ages := map[string]int{ //we have created a map ages that has string keys and int values. 
		"Alice": 25,
		"Bob": 30,
		"Charlie": 35,
	}
	fmt.Println(ages["Bob"])
	ages["Dave"] = 40 //This will add a new key "Dave" with a value of 40 to the ages map.

// The make()function is the right way to create an empty map. If you make an empty map in a different way and write to it, it will causes a runtime panic.

	var a = make(map[string]string) // The map is empty now
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"

	fmt.Println(a)
	delete(a,"year")//Removing elements is done using the delete() function.
	fmt.Println(a)

//Maps are References to hash tables, If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.
	var c = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
  b := c

  fmt.Println(c)
  fmt.Println(b)

  b["year"] = "1970"
  fmt.Println("After change to b:")

  fmt.Println(c)
  fmt.Println(b)

}	