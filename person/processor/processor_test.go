package processor

import (
	"reflect"
	"testing"

	"github.com/person/models"
)

func TestMsgdigest(t *testing.T) {
	type args struct {
		c <-chan models.Person
	}
	p := []models.Person{
		{
			Id:          4,
			Name:        "test2",
			Age:         "14",
			PhoneNumber: "123123",
		},
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "MsgDigestBaseCase",
			want: "4test2144297f44b13955235245b2497399d7a93                                                            ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputC := make(chan models.Person)
			go func() {
				defer close(inputC)
				for i := range p {
					inputC <- p[i]
				}
			}()
			c := Msgdigest(inputC)
			for got := range c {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Msgdigest() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
