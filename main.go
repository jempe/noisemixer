package main

import (
	"bufio"
	"crypto/rand"
	"flag"
	"fmt"
	mathrand "math/rand"
	"os"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_><=+"

var length = flag.Int("length", 25, "Max length of the strings between the characters I want to mix")

func SecureRandomString(length int) (string, error) {
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var stdinInput string
	if scanner.Scan() {
		stdinInput = scanner.Text()
	}

	for _, char := range stdinInput {
		fmt.Printf("%c", char)

		randomNumber := mathrand.Intn(*length-2) + 2

		randomStr, err := SecureRandomString(randomNumber)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("%s", randomStr)
	}

}
