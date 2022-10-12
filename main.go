package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {

	kalimat := "katak"
	result := isPalindrome(kalimat)

	if !result {
		fmt.Println("bukan palindrome")
	} else {
		fmt.Println(fmt.Sprintf("kata %s adalah palindrome", kalimat))
	}

	huruf := "a"
	huruf2 := "kodok"
	huruf3 := "k"

	resultKonsonan, err := isKonsonan(huruf)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fmt.Sprintf("huruf %s merupakan : %s", huruf, resultKonsonan))
	}

	resultKonsonan2, err := isKonsonan(huruf2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fmt.Sprintf("huruf %s merupakan : %s", huruf2, resultKonsonan2))
	}

	resultKonsonan3, err := isKonsonan(huruf3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fmt.Sprintf("huruf %s merupakan : %s", huruf3, resultKonsonan3))
	}

}

func isPalindrome(kata string) bool {

	for i := 0; i < len(kata); i++ {
		j := len(kata) - 1 - i

		if strings.ToLower(string(kata[i])) != strings.ToLower(string(kata[j])) {
			return false
		}
	}

	return true
}

func isKonsonan(huruf string) (string, error) {

	if len(huruf) > 1 {
		return "", errors.New("error, kata harus satu huruf")
	}

	vocals := []string{"A", "I", "U", "E", "O", "a", "i", "u", "e", "o"}

	for _, v := range vocals {
		if huruf == v {
			return "vocal", nil
		}
	}

	return "konsonan", nil
}
