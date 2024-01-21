package main

import (
	"errors"
	"fmt"
)

func main() {
	sendsSoFar := 430
	const sendsToAdd = 25
	incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("you've sent", sendsSoFar, "messages")

// ignoring return values
	firstName, _ := getNames()
	fmt.Println("Welcome to Textio,", firstName)
}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	sendsSoFar = sendsSoFar + sendsToAdd
	return sendsSoFar
}

func getNames() (string, string) {
	return "John", "Doe"
}

//Named Returns
func calculator(a, b int) (mul, div int, err error) {
	if b == 0 {
		return 0, 0, errors.New("Can't divide by zero")
	}
	mul = a * b
	div = a / b
	return mul, div, nil
}