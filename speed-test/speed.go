package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"sudoko-speed.charlie.city/practice"
	"sudoko-speed.charlie.city/utils"
)

func main() {
	if len(os.Args) != 4 {
		panic("Must pass in 3 arguments for: # of tests, width, and height")
	}

	numberOfTests := utils.StringToNumber(os.Args[1])

	width := utils.StringToNumber(os.Args[2])
	height := utils.StringToNumber(os.Args[3])

	fmt.Printf("Building %d tests... %dx%d\n", numberOfTests, width, height)
	fmt.Printf("Ready?")

	var totalTime time.Duration = 0

	reader := bufio.NewReader(os.Stdin)

	for range numberOfTests {
		output, answer := practice.BuildMissingNumberTest(width, height)

		fmt.Printf("\n%s", output)

		startTime := time.Now()

		for {
			fmt.Print("Answer: ")

			// Read the input until the user presses Enter ('\n').
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("\n[ERROR] Failed to read input. Aborting test.")
				return
			}

			// Clean up the input: trim whitespace and newline characters.
			input = strings.TrimSpace(input)

			// Try to convert the cleaned input string into an integer.
			code, err := strconv.Atoi(input)

			// Check if the input was a valid number AND if it matches the secret code.
			if err == nil && code == answer {
				// SUCCESS! Break out of the loop and stop the timer.
				break
			}

			// If we reach here, the input was wrong or not a number.
			fmt.Printf("Incorrect.\n")
		}

		endTime := time.Now()

		totalTime = totalTime + endTime.Sub(startTime)
	}

	fmt.Printf("Well done! Your average response time was %v ms\n", totalTime/time.Duration(numberOfTests))
}
