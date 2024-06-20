package main

import (
	"log"
	"strings"
	"time"
)

const delay = 700 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
func print(msg string) {
	log.Println(msg)
	time.Sleep(delay)
}

// slowDown takes the given string and repeats its characters
// according to their index in the string.
func slowDown(msg string) {
	words := strings.Split(msg, " ")
	var newMessage []string
	for _, word := range words{
		newWord := make([]string, len(word))
		for i, l := range word {
			newWord = append(newWord, strings.Repeat(string(l),i+1))  
		}
		newMessage = append(newMessage, strings.Join(newWord,""))
	}
 	print(strings.Join(newMessage, " "))
}

func main() {
	msg := "Time to learn about Go strings!"
	slowDown(msg)
}