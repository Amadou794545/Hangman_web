package modules

func TestFinish(result []string) bool { // Function to try if the word to find is completed or not
	for _, finish := range result { // Try for each possibility is it's _ of the letter
		if finish == "_" {
			return false
		}
	}
	return true
}
