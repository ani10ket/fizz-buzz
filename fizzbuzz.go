package main

import "fmt"

func main() {
	//c3 and c5 are two variable which are initially 0
	//these two varibles will help us check if the incremented
	//number is divisible by 3 or 5 or both
	c3 := 0
	c5 := 0

	//we will loop throuh the numbers till 100
	for i := 1; i <= 100; i++ {

		//c3 and c5 will increment before the validation
		c3++
		c5++

		//d variable is an empty string which will help
		//concatinating fizz buzz to the numbers
		d := ""

		//if the incremented c3 is divisible by 3
		//then we will concatinate d with fizz
		if c3 == 3 {
			d += "fizz"
			c3 = 0
		}

		//if the incremented c5 is divisible by 5
		//then we will concatinate d with buzz
		if c5 == 5 {
			d += "buzz"
			c5 = 0
		}

		//if the number is divisible by both c3 and c5
		//then d is concatinated with both fizz and buzz

		//if d is empty we print the other number that
		//are not divisible both by 3 and 5
		//else we print concatinated d from above
		if d == "" {
			fmt.Println(i)
		} else {
			fmt.Println(d)
		}
	}
}
