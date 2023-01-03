package modules

import "fmt"

func TestFinish(result []string) bool { // Function to try if the word to find is completed or not
	for _, finish := range result { // Try for each possibility is it's _ of the letter
		if finish == "_" {
			fmt.Printf("False --> test Finish")
			return false
		}
	}
	fmt.Printf("True --> test Finish")
	return true
}
