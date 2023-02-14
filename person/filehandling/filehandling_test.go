package filehandling

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/person/hashstring"
	"github.com/person/models"
)

func TestReadPerson(t *testing.T) {
	type args struct {
		data io.Reader
		c    chan models.Person
	}
	p := models.Person{

		Id:          16,
		Name:        "Sashenka",
		Age:         "12",
		PhoneNumber: "9258568864",
	}

	mys := "16,Sashenka,12,9258568864\n"
	s := strings.NewReader(mys)

	tests := []struct {
		name string
		args args
		want models.Person
	}{
		{
			name: "ReadPersonBaseCase",
			args: args{s, make(chan models.Person)},
			want: p,
		},
		{
			name: "ReadPersonATOIerrorCase",
			args: args{strings.NewReader("h,23,23,89"), make(chan models.Person)},
			want: models.Person{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go ReadPerson(tt.args.data, tt.args.c)
			for got := range tt.args.c {
				tt.want.PhoneNumber = hashstring.Md5(tt.want.PhoneNumber)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("testFailed got %v, want %v", got, tt.want)
				}
			}
		})
	}
}
