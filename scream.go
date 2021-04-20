package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	log.SetOutput(os.Stdout)

	err := os.Setenv("INTERVAL", "2000")
	if err != nil {
		log.Fatal(err)
	}
	n := os.Getenv("INTERVAL")
	i, err := strconv.Atoi(n)
	if err != nil {
		log.Fatal(err)
	}

	for {
		log.Print("SCREAM")
		time.Sleep(time.Duration(i) * time.Millisecond)
	}
}
