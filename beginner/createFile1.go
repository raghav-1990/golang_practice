package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Create("hello.txt")
	if err != nil {
		log.Fatalf("received an error: %v", err)
	}
	defer f.Close()

	writeString := "Hello World"
	writeBytes := []byte(writeString)

	n, err := f.Write(writeBytes)
	if err != nil {
		log.Fatalf("received an error while writing: %v", err)
	}
	log.Printf("Number of bytes written is %v", n)

	file, err := os.Open("hello.txt")
	if err != nil {
		log.Fatalf("received an error while opening file : %v", err)
	}
	defer file.Close()

	f2, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("received an error while reading file : %v", err)
	}
	fmt.Println("This is the data:", string(f2))
}
