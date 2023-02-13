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

func ReadCSV(data io.Reader, c chan Person) {
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
		p := Person{Id: id, Name: record[1], Age: record[2], PhoneNumber: record[3]}
		c <- p
	}
	close(c)
}
func encodeString(s string) [16]byte {

	res := md5.Sum([]byte(s))

	return res
}
func Makemsg(c chan Person) string {
	ph := encodeString(p.PhoneNumber)
	msg := fmt.Sprintf("%d%s%s%x", p.Id, p.Name, p.Age, ph)

	msg = fmt.Sprintf("%-*s", 100, msg)
	message := encodeString(msg)
	return fmt.Sprintf("%x", message)
}

func main() {
	// f, _ := os.Open("person.csv")
	readFile, _ := os.Open("person.csv")
	c := make(chan Person)
	ReadCSV(readFile, c)

	// fmt.Println(Makemsg(p))
}
