package stores

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/student-api/models"
	"github.com/student-api/myerrors"
)

func Test_storeCon_Insert(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		s models.Student
	}
	st := models.Student{
		RollNo: 2,
		Name:   "testStudent",
		Age:    99,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf(err.Error())
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantErr  bool
		myerr    string
		err      error
		mockcall *sqlmock.ExpectedExec
	}{
		{
			name:     "InsertStudentBaseCase",
			fields:   fields{db},
			args:     args{st},
			wantErr:  false,
			mockcall: mock.ExpectExec(`INSERT INTO`).WithArgs(st.RollNo, st.Name, st.Age).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			name:     "InsertStudentErrorCase",
			fields:   fields{db},
			args:     args{st},
			wantErr:  true,
			myerr:    "sql_i",
			err:      errors.New(myerrors.NewError("sql_i")),
			mockcall: mock.ExpectExec(`INSERT INTO`).WithArgs(st.RollNo, st.Name, st.Age).WillReturnError(errors.New(myerrors.NewError("sql_i"))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := storeCon{
				db: tt.fields.db,
			}
			if err := db.Insert(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("storeCon.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_storeCon_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		s models.Student
	}
	st := models.Student{
		RollNo: 2,
		Name:   "testStudent",
		Age:    99,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf(err.Error())
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantErr  bool
		myerr    string
		err      error
		mockcall *sqlmock.ExpectedExec
	}{
		{
			name:     "InsertStudentBaseCase",
			fields:   fields{db},
			args:     args{st},
			wantErr:  false,
			mockcall: mock.ExpectExec(`UPDATE`).WithArgs(st.Name, st.Age, st.RollNo).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			name:     "InsertStudentErrorCase",
			fields:   fields{db},
			args:     args{st},
			wantErr:  true,
			myerr:    "sql_u",
			err:      errors.New(myerrors.NewError("sql_u")),
			mockcall: mock.ExpectExec(`UPDATE`).WithArgs(st.Name, st.Age, st.RollNo).WillReturnError(errors.New(myerrors.NewError("sql_u"))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := storeCon{
				db: tt.fields.db,
			}
			if err := db.Update(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("storeCon.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_storeCon_Delete(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		rollno int
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf(err.Error())
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantErr  bool
		myerr    string
		err      error
		mockcall *sqlmock.ExpectedExec
	}{
		{
			name:     "DeleteStudentBaseCase",
			fields:   fields{db},
			args:     args{5},
			wantErr:  false,
			mockcall: mock.ExpectExec(`DELETE`).WithArgs(5).WillReturnResult(sqlmock.NewResult(0, 1)),
		},
		{
			name:     "DeleteStudentErrorCase",
			fields:   fields{db},
			args:     args{5},
			wantErr:  true,
			myerr:    "sql_d",
			err:      errors.New(myerrors.NewError("sql_d")),
			mockcall: mock.ExpectExec(`DELETE`).WithArgs(5).WillReturnError(errors.New(myerrors.NewError("sql_d"))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := storeCon{
				db: tt.fields.db,
			}
			if err := db.Delete(tt.args.rollno); (err != nil) != tt.wantErr {
				t.Errorf("storeCon.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_storeCon_Read(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		rollno int
	}
	st := models.Student{
		RollNo: 2,
		Name:   "testStudent",
		Age:    99,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf(err.Error())
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantErr  bool
		want     models.Student
		myerr    string
		err      error
		mockcall *sqlmock.ExpectedQuery
	}{
		{
			name:     "GetStudentBaseCase",
			fields:   fields{db},
			args:     args{5},
			want:     st,
			wantErr:  false,
			mockcall: mock.ExpectQuery(`SELECT`).WithArgs(5).WillReturnRows(mock.NewRows([]string{"RollNo", "Name", "Age"}).AddRow(2, "testStudent", 99)),
		},
		{
			name:     "GetStudentErrorCase",
			fields:   fields{db},
			args:     args{5},
			wantErr:  true,
			myerr:    "sql_r",
			err:      errors.New(myerrors.NewError("sql_r")),
			mockcall: mock.ExpectQuery(`SELECT`).WithArgs(5).WillReturnError(errors.New(myerrors.NewError("sql_r"))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := storeCon{
				db: tt.fields.db,
			}
			got, err := db.Read(tt.args.rollno)
			if (err != nil) != tt.wantErr {
				t.Errorf("storeCon.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("storeCon.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_storeCon_ReadByName(t *testing.T) {

	type fields struct {
		db *sql.DB
	}
	type args struct {
		name string
	}
	st := []models.Student{{

		RollNo: 2,
		Name:   "testStudent",
		Age:    99,
	},
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf(err.Error())
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     []models.Student
		wantErr  bool
		mockcall *sqlmock.ExpectedQuery
	}{
		{
			name:     "GetByNameBaseCase",
			fields:   fields{db},
			args:     args{st[0].Name},
			want:     st,
			wantErr:  false,
			mockcall: mock.ExpectQuery("SELECT").WithArgs(st[0].Name).WillReturnRows(mock.NewRows([]string{"RollNo", "Name", "Age"}).AddRow(2, "testStudent", 99)),
		},
		{
			name:     "GetByNameErrorCase",
			fields:   fields{db},
			args:     args{st[0].Name},
			want:     []models.Student{},
			wantErr:  true,
			mockcall: mock.ExpectQuery("SELECT").WithArgs(st[0].Name).WillReturnRows(mock.NewRows([]string{"RollNo", "Name"}).AddRow(2, "testStudent")).WillReturnError(errors.New(myerrors.NewError("sql_r"))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := storeCon{
				db: tt.fields.db,
			}
			got, err := db.ReadByName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("storeCon.ReadByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("storeCon.ReadByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_storeCon_ReadAll(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	st := []models.Student{{

		RollNo: 2,
		Name:   "testStudent",
		Age:    99,
	},
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		println(err.Error()) //
	}
	tests := []struct {
		name     string
		fields   fields
		want     []models.Student
		wantErr  bool
		mockcall *sqlmock.ExpectedQuery
	}{
		{
			name:     "ReadAllBaseCase",
			fields:   fields{db},
			want:     st,
			wantErr:  false,
			mockcall: mock.ExpectQuery("SELECT").WithArgs().WillReturnRows(mock.NewRows([]string{"RollNo", "Name", "Age"}).AddRow(2, "testStudent", 99)),
		},
		{
			name:     "ReadAllErrorCase",
			fields:   fields{db},
			want:     []models.Student{},
			wantErr:  true,
			mockcall: mock.ExpectQuery("SELECT").WithArgs().WillReturnRows(mock.NewRows([]string{"RollNo", "Name"}).AddRow(2, "testStudent")).WillReturnError(errors.New(myerrors.NewError("sql_r"))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := storeCon{
				db: tt.fields.db,
			}
			got, err := db.ReadAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("storeCon.ReadAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("storeCon.ReadAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
