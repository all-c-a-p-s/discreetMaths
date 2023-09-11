package main

import "fmt"

/**
Problem: given a straight stick 12 cm long, what is the minimum number of marks that need to be made on the stick
so that these marks can be used to measure any integer distance between 1 and 12 cm
*/

func checkSolution(nums []int) bool{

	var canMeasure = make(map [int]bool, 0)
	canMeasure[1] = false
	canMeasure[2] = false
	canMeasure[3] = false
	canMeasure[4] = false
	canMeasure[5] = false
	canMeasure[6] = false
	canMeasure[7] = false
	canMeasure[8] = false
	canMeasure[9] = false
	canMeasure[10] = false
	canMeasure[11] = false
	canMeasure[12] = true//whole ruler


	for i := 0;i < len(nums);i++{
		canMeasure[nums[i]] = true
		canMeasure[12 - nums[i]] = true //starting from other end of ruler

		for j:= 0;j < len(nums);j++{
			if nums[j] > nums[i]{ //also checks that they are not the same number
				canMeasure[nums[j] - nums[i]] = true
			}
		}
	}

	for i := 1;i < 13;i++{
		if canMeasure[i] == false {
			return false
		}
	}

	return true
}

func main(){
	//no point generating 1- or 2-long lists as these clearly can't work
	a := 1
	b := 2
	c := 3
	d := 4

	for { //testing list of 3

		nums := []int {a, b, c} 

		if checkSolution(nums) == true{
			fmt.Println(nums)
			break
		}

		if c == 12 {
			if b == 11{
				if a == 10 {//all combinations tested	
					fmt.Println("impossible with 3 marks")				
					break
				}
				a++
				b = a + 1
				c = b + 1
			} else {
				b++
				c = b + 1
			}
		} else {
			c++
		}
	} 

	a = 1
	b = 2
	c = 3
	d = 4 //reset after previous loop

	for {
		nums := []int {a, b, c, d}

		if checkSolution(nums) == true {
			fmt.Println(nums)
			break
		}

		if d == 12 {
			fmt.Println(a, b, c, d)
			if c == 11{
				if b == 10 {
					if a == 9 { //all combinations tested
						fmt.Println("impossible with 4 marks")
						break
					}
					a++
					b = a + 1
					c = b + 1
					d = c + 1
				} else {
					b ++
					c = b + 1
					d = c + 1
				}

			} else {
				c++
				d = c + 1
			}
		} else {
			d++
		}


	}

}