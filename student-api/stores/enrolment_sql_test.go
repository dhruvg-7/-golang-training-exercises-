package stores

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/student-api/models"
	"github.com/student-api/myerrors"
)

func Test_enrolmentStore_EnrolStudent(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		stuId int
		subId int
	}
	st := models.EnrolmentList{
		StudentId: 1,
		SubjectId: 2,
	}
	db, mock, _ := sqlmock.New()

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
			name:     "EnrolStudentBaseCase",
			fields:   fields{db},
			args:     args{st.StudentId, st.SubjectId},
			wantErr:  false,
			mockcall: mock.ExpectExec(`INSERT INTO`).WithArgs(st.StudentId, st.SubjectId).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			name:     "EnrolStudentErrorCase",
			fields:   fields{db},
			args:     args{st.StudentId, st.SubjectId},
			wantErr:  true,
			myerr:    "sql_i",
			err:      errors.New(myerrors.NewError("sql_i")),
			mockcall: mock.ExpectExec(`INSERT INTO`).WithArgs(st.StudentId, st.SubjectId).WillReturnError(errors.New(myerrors.NewError("sql_i"))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := enrolmentStore{
				db: tt.fields.db,
			}
			if err := db.EnrolStudent(tt.args.stuId, tt.args.subId); (err != nil) != tt.wantErr {
				t.Errorf("enrolmentStore.EnrolStudent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_enrolmentStore_GetAllSubjectByStudent(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		stuId int
	}
	st := models.EnrolmentList{
		StudentId: 1,
		SubjectId: 2,
	}
	db, mock, _ := sqlmock.New()
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     []models.EnrolmentList
		myerr    string
		err      error
		mockcall *sqlmock.ExpectedQuery
		wantErr  bool
	}{
		{
			name:     "GetSubjectbySudentBaseCase",
			fields:   fields{db},
			args:     args{st.StudentId},
			wantErr:  false,
			want:     []models.EnrolmentList{st},
			mockcall: mock.ExpectQuery(`SELECT`).WithArgs(st.StudentId).WillReturnRows(mock.NewRows([]string{"StudentId", "SubjectId"}).AddRow(1, 2)),
		},
		{
			name:     "GetSubjectbySudentErrorCase",
			fields:   fields{db},
			args:     args{st.StudentId},
			wantErr:  true,
			myerr:    "sql_r",
			want:     []models.EnrolmentList{},
			err:      errors.New(myerrors.NewError("sql_r")),
			mockcall: mock.ExpectQuery(`SELECT`).WithArgs(st.StudentId).WillReturnError(errors.New(myerrors.NewError("sql_r"))),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := enrolmentStore{
				db: tt.fields.db,
			}
			got, err := db.GetAllSubjectByStudent(tt.args.stuId)
			if (err != nil) != tt.wantErr {
				t.Errorf("enrolmentStore.GetAllSubjectByStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("enrolmentStore.GetAllSubjectByStudent() = %v, want %v", got, tt.want)
			}
		})
	}
}
