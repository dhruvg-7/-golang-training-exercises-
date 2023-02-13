package filehandling

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"strconv"

	"github.com/person/hashstring"
	"github.com/person/models"
)

func ReadPerson(data io.Reader, c chan models.Person) {
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
		phone := hashstring.Md5(record[3])
		p := models.Person{Id: id, Name: record[1], Age: record[2], PhoneNumber: phone}
		c <- p
	}
	close(c)
}

func WriteString(w io.Writer, c chan string) {
	for i := range c {

		_, err := w.Write([]byte(i + "\n"))
		if err != nil {
			log.Fatal(err)
		}
	}
}
