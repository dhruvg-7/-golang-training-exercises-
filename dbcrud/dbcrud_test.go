package dbcrud

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
)

func Test_insertRecord(t *testing.T) {
	type args struct {
		db *newDB
		s  Student
	}
	st := []Student{
		{
			RollNo: 1,
			Name:   "Test1",
			Age:    12,
		},
		{
			RollNo: 2,
			Name:   "Test2",
			Age:    999,
		},
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("Error openinig mock db reason: %v", err.Error())
	}
	dbobj := &newDB{db}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		mockcall *sqlmock.ExpectedExec
	}{
		{
			name:     "baseCase",
			args:     args{dbobj, st[0]},
			wantErr:  false,
			mockcall: (mock.ExpectExec(`insert into Student`).WithArgs(st[0].RollNo, st[0].Name, st[0].Age).WillReturnResult(sqlmock.NewResult(1, 1))),
		},
		{
			name:     "baseCase2",
			args:     args{dbobj, st[1]},
			wantErr:  true,
			mockcall: (mock.ExpectExec(`insert into Student`).WithArgs(st[0].RollNo, st[0].Name, st[0].Age).WillReturnError(fmt.Errorf("INSERT ERROR"))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := insertRecord(tt.args.db, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("insertRecord() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_updateRecord(t *testing.T) {
	type args struct {
		db *newDB
		s  Student
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	dbobj := &newDB{db}

	st := []Student{
		{
			RollNo: 1,
			Name:   "Test1",
			Age:    12,
		},
		{
			RollNo: 8,
			Name:   "Test2",
			Age:    999,
		},
	}

	tests := []struct {
		name     string
		args     args
		wantErr  bool
		mockcall *sqlmock.ExpectedExec
	}{
		{
			name: "UpdateBaseCase",
			args: args{dbobj,
				st[0],
			},
			mockcall: (mock.ExpectExec(`UPDATE Student`).WithArgs(st[0].Name, st[0].Age, st[0].RollNo).WillReturnResult(sqlmock.NewResult(0, 1))),
		},
		{
			name:     "UpdateErrorCase",
			args:     args{dbobj, st[1]},
			wantErr:  true,
			mockcall: (mock.ExpectExec(`UPDATE Student`).WithArgs(st[0].Name, st[0].Age, st[0].RollNo).WillReturnError(fmt.Errorf("SQL ERROR"))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := updateRecord(tt.args.db, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("updateRecord() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_deleteRecord(t *testing.T) {
	type args struct {
		db     *newDB
		rollno int
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	dbobj := &newDB{db}
	defer db.Close()
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		mockcall *sqlmock.ExpectedExec
	}{
		{
			name:     "DeleteBaseCase",
			args:     args{dbobj, 5},
			mockcall: (mock.ExpectExec(`DELETE FROM`).WithArgs(5).WillReturnResult(sqlmock.NewResult(0, 1))),
		},
		{
			name:     "DeleteErrorCase",
			args:     args{dbobj, 5},
			mockcall: (mock.ExpectExec(`DELETE FROM`).WithArgs(5).WillReturnError(fmt.Errorf("DELETE ERROR"))),
		},
	}
	for _, tt := range tests {
		mock.ExpectExec(`DELETE FROM`).WithArgs(tt.args.rollno).WillReturnResult(sqlmock.NewResult(0, 1))
		t.Run(tt.name, func(t *testing.T) {
			if err := deleteRecord(tt.args.db, tt.args.rollno); (err != nil) != tt.wantErr {
				t.Errorf("deleteRecord() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_readRecord(t *testing.T) {
	type args struct {
		db     *newDB
		rollno int
	}
	st := []Student{
		{
			RollNo: 2,
			Name:   "TEST1",
			Age:    99,
		},
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("Error openinig mock db reason: %v", err.Error())
	}
	dbobj := &newDB{db}

	tests := []struct {
		name     string
		args     args
		want     Student
		wantErr  bool
		mockcall *sqlmock.ExpectedQuery
	}{
		{
			name:     "ReadBaseCase",
			args:     args{dbobj, 2},
			want:     st[0],
			mockcall: (mock.ExpectQuery(`SELECT .* FROM`).WithArgs(2).WillReturnRows(mock.NewRows([]string{"rollno", "name", "age"}).AddRow(2, "TEST1", 99))),
		},
		{
			name:     "ReadErrorCase",
			args:     args{dbobj, 2},
			want:     Student{},
			wantErr:  true,
			mockcall: (mock.ExpectQuery(`SELECT .* FROM`).WithArgs(2).WillReturnError(fmt.Errorf("ROW NOT FOUND"))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readRecord(tt.args.db, tt.args.rollno)
			if (err != nil) != tt.wantErr {
				t.Errorf("readRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readAllRecord(t *testing.T) {
	type args struct {
		db *newDB
	}
	st := []Student{
		{
			RollNo: 2,
			Name:   "TEST1",
			Age:    1,
		},
		{
			RollNo: 3,
			Name:   "TEST2",
			Age:    2,
		},
		{
			RollNo: 4,
			Name:   "TEST3",
			Age:    3,
		},
	}
	// res := []string{
	// 	"2", "TEST1", "99", "3", "TEST2", "2", "4", "TEST3", "3",
	// }
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("Error openinig mock db reason: %v", err.Error())
	}
	dbobj := &newDB{db}
	tests := []struct {
		name     string
		args     args
		want     []Student
		wantErr  bool
		mockcall *sqlmock.ExpectedQuery
	}{
		{
			name:     "ReadBaseCase",
			args:     args{dbobj},
			want:     st,
			wantErr:  false,
			mockcall: (mock.ExpectQuery(`SELECT .* FROM`).WithArgs().WillReturnRows(mock.NewRows([]string{"rollno", "name", "age"}).AddRow(2, "TEST1", 1).AddRow(3, "TEST2", 2).AddRow("4", "TEST3", "3"))),
		},
		{
			name:     "ReadErrorCase",
			args:     args{dbobj},
			want:     []Student{},
			wantErr:  true,
			mockcall: (mock.ExpectQuery(`SELECT .* FROM`).WithArgs().WillReturnError(fmt.Errorf("READALL ERROR"))),
		},
		{
			name:     "ReadRowsErrorCase",
			args:     args{dbobj},
			want:     []Student{},
			wantErr:  true,
			mockcall: (mock.ExpectQuery(`SELECT .* FROM`).WithArgs().WillReturnRows(mock.NewRows([]string{"rollno", "name", "age", "extra"}).AddRow(2, "TEST1", 1, 4)).WillReturnError(fmt.Errorf("SCAN ERROR"))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readAllRecord(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("readAllRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readAllRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Server()
		})
	}
}
