package main

import (
	"log"
	"os"

	"github.com/person/filehandling"
	"github.com/person/message"
	"github.com/person/models"
)

func main() {

	f, _ := os.Open("person.csv")
	defer f.Close()

	c := make(chan models.Person)
	c2 := make(chan string)

	go filehandling.ReadPerson(f, c)
	go message.Makemsg(c, c2)
	f2, err := os.Create("personEncoded.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()
	filehandling.WriteString(f2, c2)

}
