package main

import (
	"bufio"
	"crypto/rand"
	"flag"
	"fmt"
	"os"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_><=+"

var length = flag.Int("length", 40, "Max length of the strings between the characters I want to mix")

func secureRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes) // Read cryptographically secure random bytes
	if err != nil {
		return "", err
	}

	for i := 0; i < length; i++ {
		bytes[i] = charset[bytes[i]%byte(len(charset))] // Map bytes to charset
	}

	return string(bytes), nil
}

func randomNumber(max int) (int, error) {
	if max < 2 {
		return 2, nil
	}

	// Create a single byte buffer
	b := make([]byte, 1)

	for {
		// Read a random byte
		_, err := rand.Read(b)
		if err != nil {
			return 0, err
		}

		// Calculate the number within our range
		n := int(b[0])%(max-1) + 2

		// If the number is within our desired range, return it
		if n <= max {
			return n, nil
		}
		// If not, try again
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var stdinInput string
	if scanner.Scan() {
		stdinInput = scanner.Text()

		// Add a random character at the beginning so that the first character is not part of the input
		randomCharacter, err := secureRandomString(1)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		stdinInput = randomCharacter + stdinInput
	}

	for _, char := range stdinInput {
		fmt.Printf("%c", char)

		randomInt, err := randomNumber(*length)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		randomStr, err := secureRandomString(randomInt)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("%s", randomStr)
	}

}
