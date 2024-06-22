package main

import (
	"flag"
	"log"
	"sync"
)

var messages = []string{
	"Hello!",
	"How are you?",
	"Are you just going to repeat what I say?",
	"So immature",
	"Stop copying me!",
}

// repeat concurrently prints out the given message n times
func repeat(n int, message string ) {
	var wg sync.WaitGroup	
	wg.Add(n)
	
	for i := 0; i < n; i++ {
		go func (i int) {
			log.Printf("[G%d]:%s",i,message)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func repeatWithChan(n int, message string) {
	c := make(chan string)
	for i := 0; i < n; i++ {
		go func (i int) {
			c <- message
		}(i)
	} 
			
	for i := 0; i < n; i++ {
		log.Printf("[G%d]:%s",i,<-c)
		
	}
}

func main() {
	factor := flag.Int64("factor", 0, "The fan-out factor to repeat by")
	flag.Parse()
	for _, m := range messages {
		log.Printf("[Main]:%s",m)
		repeatWithChan(int(*factor),m)
	}
}
