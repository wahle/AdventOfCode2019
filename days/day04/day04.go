package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wahle/AdventOfCode2019/password"
)

func main() {
	//sinput := "231832-767346"
	//passWords := createInputs(input)
	//fmt.Printf("\nPassWord: %v\n", passWords)
	numberOfPasswords := 0
	for i := 231832; i <= 767346; i++ {
		tempPass := password.New(strconv.Itoa(i))
		//fmt.Printf("\nPassWord: %v\n", tempPass)
		if tempPass.Verify() {
			numberOfPasswords++
		}
	}
	fmt.Printf("\nNumber Of Passwords: %v\n", numberOfPasswords)

	nick := "hello"
	fmt.Printf("\nNumber Of Passwords: %v\n", nick)
}

func createInputs(input string) []password.Password {
	var passWords []password.Password
	//convert input into two arrays of ints
	for _, eachPass := range strings.Split(input, "-") {
		passInt := password.New(eachPass)
		passWords = append(passWords, passInt)
	}
	return passWords
}
