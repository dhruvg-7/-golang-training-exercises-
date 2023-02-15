package readwrite

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/person/models"
)

func ReadPerson(data io.Reader) <-chan models.Person {

	fileScanner := bufio.NewScanner(data)
	c := make(chan models.Person)
	go func() {
		defer close(c)
		for fileScanner.Scan() {
			r := fileScanner.Text()
			record := strings.Split(r, ",")
			id, err1 := strconv.Atoi(record[0])
			if err1 != nil {
				return
			}
			p := models.Person{Id: id, Name: record[1], Age: record[2], PhoneNumber: record[3]}
			c <- p
		}
	}()
	return c
}

func WriteString(w io.Writer, c <-chan string) error {
	for i := range c {

		_, err := w.Write([]byte(i + "\n"))
		if err != nil {
			return err
		}
	}
	return nil
}
