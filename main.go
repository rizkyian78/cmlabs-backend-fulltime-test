package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("======== fizz buzz")
	fizzBuzz(100)

	fmt.Println("======== palindrome")
	// single words
	palindrome("SAIPPUAKIVIKAUPPIAS")
	palindrome("Aibohphobia")
	palindrome("Anna")
	palindrome("Civic")

	//multiple words
	palindrome("My gym")
	palindrome("No lemon, no melon")

}

func fizzBuzz(char int) {
	for i := 0; i < char; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fishbuzz", i)
		} else if i%3 == 0 {
			fmt.Println("fizz", i)
		} else if i%5 == 0 {
			fmt.Println("buzz", i)
		}
	}
}

func palindrome(char string) {
	lowerString := strings.ToLower(char)
	arr := strings.Split(strings.ToLower(lowerString), "")
	reverseArray(arr)
	fmt.Println(strings.Join(arr, "") == strings.ToLower(char))

}

func reverseArray(arr []string) {
	length := len(arr)
	for i := 0; i < length/2; i++ {
		j := length - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
}
