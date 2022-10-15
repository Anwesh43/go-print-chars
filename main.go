package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func scanLines(ch chan []string) {
	scanner := bufio.NewScanner(os.Stdin)
	words := make([]string, 0)
	for scanner.Scan() {
		word := scanner.Text()
		if word == "quit" {
			break
		}
		words = append(words, word)
	}
	ch <- words
}

func printWords(words []string, ch chan string) {
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			fmt.Print(string(word[i]))
			time.Sleep(250 * time.Millisecond)
		}
		fmt.Println()
	}
	ch <- "done"
}

func main() {
	ch1 := make(chan []string)
	ch2 := make(chan string)
	fmt.Println("Enter words")
	go scanLines(ch1)
	words := <-ch1
	fmt.Println("Typing words")
	fmt.Println("______________________________________________________XXXXX____________________________________________")
	fmt.Println("______________________________________________________XXXXX____________________________________________")
	fmt.Println("______________________________________________________XXXXX____________________________________________")
	fmt.Println("______________________________________________________XXXXX____________________________________________")
	time.Sleep(2 * time.Second)
	go printWords(words, ch2)
	<-ch2
	fmt.Println("Done printing words")
}
