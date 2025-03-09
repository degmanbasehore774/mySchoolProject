func main() {
	// Create a new array of integers with length 5
	arr := make([]int, 5)

	// Populate the array with random numbers between 1 and 10
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(10) + 1
	}

	// Print out the array
	fmt.Println(arr)
}