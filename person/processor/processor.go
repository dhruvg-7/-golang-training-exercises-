package processor

import (
	"crypto/md5"
	"fmt"

	"github.com/person/models"
)

func Msgdigest(c <-chan models.Person) <-chan string {

	c2 := make(chan string)

	go func() {
		defer close(c2)
		for i := range c {
			i.PhoneNumber = fmt.Sprintf("%x", md5.Sum([]byte(i.PhoneNumber)))
			msg := fmt.Sprintf("%d%s%s%v", i.Id, i.Name, i.Age, i.PhoneNumber)
			message := fmt.Sprintf("%-*s", 100, msg)
			c2 <- message
		}
	}()
	return c2
}
