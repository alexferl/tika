package main

import (
	"fmt"
	"log"
	"os"

	"github.com/admiralobvious/tika"
)

func main() {
	c := tika.NewClient(&tika.Options{Url: "http://localhost:9998"})

	hi, err := c.Hello()
	if err != nil {
		log.Fatalf("Error getting hello: %v", err)
	}

	fmt.Printf("Server replied: %s", hi)

	file, err := os.Open("foo.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	c.Document = file
	c.DocumentName = file.Name()

	r, err := c.Meta().Json()
	if err != nil {
		log.Fatalf("Error getting metadata: %v", err)
	}

	fmt.Printf("File metadata: %s\n", r)
}
