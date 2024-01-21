

package main

import (
	"fmt"
)

func demo(a int) {
	for i := 0; i < a; i++ {
		fmt.Println("Hello")
	}
}

func convertToInt(value interface{}) int {
	if v, ok := value.(int); ok {
		return v
	}
	panic(fmt.Sprintf("Expected an integer, but got: %v", value))
	//panic(fmt.Printf("Expected an integer, but got: %v", value))
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	demo(convertToInt("vj"))

}
