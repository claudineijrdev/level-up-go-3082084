package main

import (
	"fmt"
	"log"
)

// the number of attendees we need to serve lunch to
const consumerCount = 300

// foodCourses represents the types of resources to pass to the consumers
var foodCourses = []string{
	"Caprese Salad",
	"Spaghetti Carbonara",
	"Vanilla Panna Cotta",
}

// takeLunch is the consumer function for the lunch simulation
// Change the signature of this function as required
func takeLunch(name string, serviceChan chan string) {
	serviceChan <- name
}

// serveLunch is the producer function for the lunch simulation.
// Change the signature of this function as required
func serveLunch(course string, serviceChan chan string) {
	log.Printf("%s eats %s",course, <-serviceChan)
}

func main() {
	log.Printf("Welcome to the conference lunch! Serving %d attendees.\n", consumerCount)

	serviceChan := make(chan string)
	for _, food := range foodCourses {
		for i := 0; i < consumerCount; i++ {
			go serveLunch(fmt.Sprintf("Attendee %d",i+1), serviceChan)
			
		}
		for i := 0; i < consumerCount; i++ {
			go takeLunch(food,serviceChan)

		}
	}
}
