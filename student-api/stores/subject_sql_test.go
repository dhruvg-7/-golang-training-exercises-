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

func Test_subjectStore_InsertSubject(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		s models.Subject
	}
	st := models.Subject{
		Id:   2,
		Name: "testSubject",
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
			name:     "InsertSubjectBaseCase",
			fields:   fields{db},
			args:     args{st},
			wantErr:  false,
			mockcall: mock.ExpectExec(`insert into subject`).WithArgs(st.Id, st.Name).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			name:     "InsertSubjectErrorCase",
			fields:   fields{db},
			args:     args{st},
			wantErr:  true,
			myerr:    "sql_i",
			err:      errors.New(myerrors.NewError("sql_i")),
			mockcall: mock.ExpectExec(`insert into subject`).WithArgs(st.Id, st.Name).WillReturnError(errors.New(myerrors.NewError("sql_i"))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := subjectStore{
				db: tt.fields.db,
			}
			if err := db.InsertSubject(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("subjectStore.InsertSubject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_subjectStore_UpdateSubject(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		s models.Subject
	}
	st := models.Subject{
		Id:   2,
		Name: "testSubject",
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
			name:     "UpdateSubjectErrorCase",
			fields:   fields{db},
			args:     args{st},
			wantErr:  true,
			myerr:    "sql_u",
			err:      errors.New(myerrors.NewError("sql_u")),
			mockcall: mock.ExpectExec(`UPDATE subject SET`).WithArgs(st.Name, st.Id).WillReturnError(errors.New(myerrors.NewError("sql_u"))),
		},
		{
			name:     "UpdateSubjectBaseCase",
			fields:   fields{db},
			args:     args{st},
			wantErr:  false,
			mockcall: mock.ExpectExec(`UPDATE subject SET`).WithArgs(st.Name, st.Id).WillReturnResult(sqlmock.NewResult(0, 1)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := subjectStore{
				db: tt.fields.db,
			}

			if err := db.UpdateSubject(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("subjectStore.UpdateSubject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_subjectStore_DeleteSubject(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		Id int
	}
	st := models.Subject{
		Id:   2,
		Name: "testSubject",
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
			name:     "DeleteSubjectErrorCase",
			fields:   fields{db},
			args:     args{st.Id},
			wantErr:  true,
			myerr:    "sql_d",
			err:      errors.New(myerrors.NewError("sql_d")),
			mockcall: mock.ExpectExec(`DELETE FROM subject `).WithArgs(st.Id).WillReturnError(errors.New(myerrors.NewError("sql_d"))),
		},
		{
			name:     "DeleteSubjectBaseCase",
			fields:   fields{db},
			args:     args{st.Id},
			wantErr:  false,
			mockcall: mock.ExpectExec(`DELETE FROM subject`).WithArgs(st.Id).WillReturnResult(sqlmock.NewResult(0, 1)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := subjectStore{
				db: tt.fields.db,
			}
			if err := db.DeleteSubject(tt.args.Id); (err != nil) != tt.wantErr {
				t.Errorf("subjectStore.DeleteSubject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_subjectStore_GetSubject(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		Id int
	}
	st := models.Subject{
		Id:   2,
		Name: "testSubject",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf(err.Error())
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     models.Subject
		wantErr  bool
		mockcall *sqlmock.ExpectedQuery
	}{
		{
			name:     "GetSubjectBaseCase",
			fields:   fields{db},
			args:     args{st.Id},
			wantErr:  false,
			want:     st,
			mockcall: mock.ExpectQuery(`SELECT .* FROM subject`).WithArgs(st.Id).WillReturnRows(mock.NewRows([]string{"Id", "Name"}).AddRow(2, "testSubject")),
		},
		{
			name:     "GetSubjectErrorCase",
			fields:   fields{db},
			args:     args{st.Id},
			wantErr:  true,
			mockcall: mock.ExpectQuery(`SELECT .* FROM subject`).WithArgs(st.Id).WillReturnError(errors.New(myerrors.NewError("sql_r"))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := subjectStore{
				db: tt.fields.db,
			}
			got, err := db.GetSubject(tt.args.Id)
			if (err != nil) != tt.wantErr {
				t.Errorf("subjectStore.GetSubject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subjectStore.GetSubject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subjectStore_GetAllSubject(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	st := []models.Subject{
		{
			Id:   2,
			Name: "testSubject",
		},
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf(err.Error())
	}
	tests := []struct {
		name   string
		fields fields

		want     []models.Subject
		wantErr  bool
		mockcall *sqlmock.ExpectedQuery
	}{
		{
			name:     "GetSubjectBaseCase",
			fields:   fields{db},
			wantErr:  false,
			want:     st,
			mockcall: mock.ExpectQuery(`SELECT .* FROM subject`).WithArgs().WillReturnRows(mock.NewRows([]string{"Id", "Name"}).AddRow(2, "testSubject")),
		},
		{
			name:     "GetSubjectErrorCase",
			fields:   fields{db},
			wantErr:  true,
			want:     []models.Subject{},
			mockcall: mock.ExpectQuery(`SELECT .* FROM subject`).WithArgs().WillReturnError(errors.New(myerrors.NewError("sql_r"))),
		},
		{
			name:     "GetSubjectRowErrorCase",
			fields:   fields{db},
			wantErr:  true,
			want:     []models.Subject{},
			mockcall: mock.ExpectQuery(`SELECT .* FROM subject`).WithArgs().WillReturnRows(mock.NewRows([]string{"Id", "Name", "extra"}).AddRow(2, "testSubject", "x")).WillReturnError(errors.New(myerrors.NewError("sql_r"))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := subjectStore{
				db: tt.fields.db,
			}
			got, err := db.GetAllSubject()
			if (err != nil) != tt.wantErr {
				t.Errorf("subjectStore.GetAllSubject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subjectStore.GetAllSubject() = %v, want %v", got, tt.want)
			}
		})
	}
}
