package main

/**
Problem: given a regular polygon with n sides, how many unique (non-congruent) triangles can be made
by connecting any 3 vertices in the polygon
*/

import (
	"fmt"
	"strconv"
)

type triangle struct {
	//represent each triangle as only the "gap" between the vertices
	v1 int
	v2 int 
	v3 int
}

func findSolutions(target int) int {
	hashSet := map[triangle]bool{}

	for i := 1;i < target - 1;i++{ //must leave at least 2 gap for the last two vertices
		for j := 1;j < target - i;j++{
			vLast := target - i - j
			v1 := min(i, j, vLast)
			v3 := max(i, j, vLast)
			v2 := target - v1 - v3 //sorted in ascending order so that repetitions can be detected

			t := triangle {v1: v1, v2: v2, v3: v3}
			hashSet[t] = true
		}
		
	}

	for t, _ := range hashSet {
		fmt.Println(t)
	} 
	return len(hashSet)
}

func main(){
	var polygonSides string
	
	fmt.Println("Enter polygon sides: ")
	fmt.Scanln(&polygonSides)

	n, err := strconv.Atoi(polygonSides)
	if err != nil {
		panic(err)
	}
	
	if n < 3 {
		panic("Polygon must have at least 3 sides")
	}

	triangles := findSolutions(n)
	fmt.Println(triangles)
}