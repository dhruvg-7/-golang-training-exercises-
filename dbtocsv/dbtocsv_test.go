package dbtocsv

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
)

func Test_readDB(t *testing.T) {
	type args struct {
		myDb *sql.DB
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	tests := []struct {
		name    string
		args    args
		want    []student
		wantErr bool
	}{
		{
			name:    "Base case",
			args:    args{myDb: db},
			want:    []student{},
			wantErr: false,
		},
	}
	mock.ExpectQuery("SELECT .* FROM student").WithArgs().WillReturnRows(sqlmock.NewRows([]string{"Name", "0", "Phone"}))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readDB(tt.args.myDb)
			if (err != nil) != tt.wantErr {
				t.Errorf("readDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_writeio(t *testing.T) {
	type args struct {
		st []student
	}

	s := []student{
		{
			"abc",
			10,
			123,
		},
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{{
		name:    "baseCase",
		args:    args{s},
		want:    "123,abd,10\n",
		wantErr: nil,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := &bytes.Buffer{}
			got, err := writeio(tt.args.st, file)
			if err != tt.wantErr {
				t.Errorf("writeio() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotstring := fmt.Sprintf("%s", got)
			if gotstring != tt.want {
				t.Errorf("writeio() = %v, want %v", gotstring, tt.want)
			}
		})
	}
}
