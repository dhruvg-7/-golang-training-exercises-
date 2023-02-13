package message

import (
	"fmt"

	"github.com/person/models"
)

func Makemsg(c chan models.Person, c2 chan string) {

	for i := range c {
		msg := fmt.Sprintf("%d%s%s%v", i.Id, i.Name, i.Age, i.PhoneNumber)
		message := fmt.Sprintf("%-*s", 100, msg)
		c2 <- message
	}
	close(c2)
}
