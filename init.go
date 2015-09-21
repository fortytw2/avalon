package main

import (
	"log"
	"time"
)

func main() {
	log.Println("HELLO WORLD")

	for {
		log.Println("I'm PID(1) and I'm still running!")
		time.Sleep(5 * time.Second)
	}
}
