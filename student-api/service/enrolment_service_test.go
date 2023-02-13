package service

import (
	"errors"
	"reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/student-api/models"
)

func Test_enrolmentService_EnrolStudentSvs(t *testing.T) {
	type fields struct {
		d enrolmentDatastore
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockEnrolmentDatastore := NewMockenrolmentDatastore(ctrl)

	type args struct {
		stuId int
		subId int
	}
	tests := []struct {
		name     string
		fields   fields
		sub      models.EnrolmentList
		args     args
		wantErr  bool
		mockcall []interface{}
	}{
		{
			name:   "EnrolStudentSvcBaseCase",
			fields: fields{d: mockEnrolmentDatastore},
			sub:    models.EnrolmentList{StudentId: 1, SubjectId: 1},
			mockcall: []interface{}{
				mockEnrolmentDatastore.EXPECT().EnrolStudent(gomock.Any(), gomock.Any()).Return(nil),
			},
		},
		{
			name:    "EnrolStudentSvcErrorCase",
			fields:  fields{d: mockEnrolmentDatastore},
			sub:     models.EnrolmentList{StudentId: 1, SubjectId: 1},
			wantErr: true,
			mockcall: []interface{}{
				mockEnrolmentDatastore.EXPECT().EnrolStudent(gomock.Any(), gomock.Any()).Return(errors.New("someError")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := enrolmentService{
				d: tt.fields.d,
			}
			if err := e.EnrolStudentSvs(tt.args.stuId, tt.args.subId); (err != nil) != tt.wantErr {
				t.Errorf("enrolmentService.EnrolStudentSvs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_enrolmentService_GetSubjectsByStuSvs(t *testing.T) {
	type args struct {
		stuId int
	}
	type fields struct {
		d enrolmentDatastore
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockEnrolmentDatastore := NewMockenrolmentDatastore(ctrl)

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantErr  bool
		want     []models.EnrolmentList
		mockcall []interface{}
	}{
		{
			name:   "GetSubjectSvcBaseCase",
			fields: fields{d: mockEnrolmentDatastore},
			mockcall: []interface{}{
				mockEnrolmentDatastore.EXPECT().GetAllSubjectByStudent(gomock.Any()).Return(nil, nil),
			},
		},
		{
			name:    "GetSubjectSvcErrorCase",
			fields:  fields{d: mockEnrolmentDatastore},
			wantErr: true,
			mockcall: []interface{}{
				mockEnrolmentDatastore.EXPECT().GetAllSubjectByStudent(gomock.Any()).Return(nil, errors.New("someError")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := enrolmentService{
				d: tt.fields.d,
			}
			got, err := e.GetSubjectsByStuSvs(tt.args.stuId)
			if (err != nil) != tt.wantErr {
				t.Errorf("enrolmentService.GetAllSubjectByStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("enrolmentService.GetAllSubjectByStudent() = %v, want %v", got, tt.want)
			}
		})
	}
}
