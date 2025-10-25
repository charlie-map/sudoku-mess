package utils

import (
	"log"
	"math/rand"
	"strconv"
	"time"
)

func StringToNumber(input string) int {
	output, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalf("Error converting string: %v", err)
	}

	return output
}

// This stuff is all from my friend Rocket Racoon (Gemini).
func GetRandomNumber(min int, max int) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// This little formula does the trick. 'Intn' returns a number between 0
	// and (max - min + 1), and then we just shift it up by 'min'.
	return r.Intn(max-min+1) + min
}

func RandomNumbersList(n int) []int {
	if n <= 0 {
		return nil
	}

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = i + 1
	}

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// This little line does all the heavy lifting for the shuffle.
	// It uses the Fisher-Yates algorithm, which is guaranteed to be random.
	r.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	return numbers
}
