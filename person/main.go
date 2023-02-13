package main

import (
	"crypto/md5"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Person struct {
	Id          int
	Name        string
	Age         string
	PhoneNumber string
}

func encodeString(s string) string {
	res := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", res)
}

func ReadPerson(data io.Reader, c chan Person) {
	r := csv.NewReader(data)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			if errors.Is(err, csv.ErrFieldCount) {
				log.Fatal("wrong fields")
			} else {
				log.Fatal(err)
			}
		}
		id, err1 := strconv.Atoi(record[0])
		if err1 != nil {
			continue
		}
		phone := encodeString(record[3])
		p := Person{Id: id, Name: record[1], Age: record[2], PhoneNumber: phone}
		c <- p
	}
	close(c)
}

func Makemsg(c chan Person, c2 chan string) {

	for i := range c {
		msg := fmt.Sprintf("%d%s%s%v", i.Id, i.Name, i.Age, i.PhoneNumber)
		message := fmt.Sprintf("%-*s", 100, msg)
		c2 <- message
	}
	close(c2)
}

func main() {

	f, _ := os.Open("person.csv")
	defer f.Close()

	c := make(chan Person)
	c2 := make(chan string)

	go ReadPerson(f, c)
	go Makemsg(c, c2)
	f2, err := os.Create("personEncoded.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	for i := range c2 {

		_, err := f2.WriteString(i + "\n")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(i)
	}

}
