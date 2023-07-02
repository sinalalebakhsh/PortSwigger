package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("numbers.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for i := 0; i <= 9999; i++ {
		_, err := fmt.Fprintf(file, "%d\n", i)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("Numbers saved to numbers.txt")
}
