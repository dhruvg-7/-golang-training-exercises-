package filehandling

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/person/hashstring"
	"github.com/person/models"
)

func ReadPerson(data io.Reader, c chan models.Person) {
	fileScanner := bufio.NewScanner(data)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		r := fileScanner.Text()
		record := strings.Split(r, ",")
		id, err1 := strconv.Atoi(record[0])
		if err1 != nil {
			close(c)
			return
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
			println(err)
			return
		}
	}
}
