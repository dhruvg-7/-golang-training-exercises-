package main

import (
	"fmt"
	"log"
	"os"

	"github.com/person/processor"
	"github.com/person/readwrite"
)

func main() {

	f, _ := os.Open("person.csv")
	defer f.Close()

	c := readwrite.ReadPerson(f)
	c2 := processor.Msgdigest(c)
	f2, err := os.Create("personEncoded.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()
	err = readwrite.WriteString(f2, c2)
	if err != nil {
		fmt.Printf("Write error Got:%v", err)
	}
}
