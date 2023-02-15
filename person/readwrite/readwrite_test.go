package readwrite

import (
	"bytes"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/person/models"
)

func TestReadPerson(t *testing.T) {
	type args struct {
		data io.Reader
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
			args: args{s},
			want: p,
		},
		{
			name: "ReadPersonATOIerrorCase",
			args: args{strings.NewReader("h,23,23,89")},
			want: models.Person{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ReadPerson(tt.args.data)
			for got := range c {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("testFailed got %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestWriteString(t *testing.T) {
	tests := []struct {
		name    string
		input1  string
		wantW   string
		wantErr bool
		w       *bytes.Buffer
	}{
		{
			name:   "WriteStringBaseCase",
			input1: "test1",
			wantW:  "test1\n",
			w:      &bytes.Buffer{},
		},
		// {
		// 	name:    "WriteStringErrCase",
		// 	input1:  "",
		// 	wantW:   "test1\n",
		// 	wantErr: true,
		// 	w:       &bytes.Buffer{},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c := make(chan string)
			go func() {
				c <- tt.input1
				close(c)
			}()
			err := WriteString(tt.w, c)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteString() = %v, want %v", err, tt.wantErr)
			}
			if gotW := tt.w.String(); gotW != tt.wantW {
				t.Errorf("WriteString() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
