package main

import (
	"fmt"
	"slices"
)

func main(){

	// we also have dynamic array in go that is slice
	// animals := [2]string{"cat", "dog"} // this is array
	animals := []string{"cat", "dog", "fish"} // this is slice
	fmt.Println(animals)

	animals = slices.Delete(animals, 0,1)
	fmt.Println(animals)

	animals = append(animals, "lion")


	for i := 0; i < len(animals); i++ {
		fmt.Println(animals[i])
	}
	fmt.Println(animals)
}