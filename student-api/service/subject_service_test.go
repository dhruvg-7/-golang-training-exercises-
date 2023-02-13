package service

import (
	"reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/student-api/models"
)

func Test_subjectService_InsertSubjectService(t *testing.T) {
	type fields struct {
		d subjectDatastore
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSubjectDatastore := NewMocksubjectDatastore(ctrl)
	st := models.Subject{
		Id:   5,
		Name: "Test1",
	}
	type args struct {
		subject models.Subject
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantErr  bool
		mockcall []interface{}
	}{
		{
			name:   "InsertSubjectSvcBaseCase",
			fields: fields{d: mockSubjectDatastore},
			args:   args{st},
			mockcall: []interface{}{
				mockSubjectDatastore.EXPECT().InsertSubject(st).Return(nil),
			},
		},
		// {
		// 	name:    "InsertSubjectSvcErrorCase",
		// 	fields:  fields{d: mockSubjectDatastore},
		// 	wantErr: true,
		// 	mockcall: []interface{}{
		// 		mockSubjectDatastore.EXPECT().InsertSubject(gomock.Any()).Return(errors.New("someError")),
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sh := subjectService{
				d: tt.fields.d,
			}
			if err := sh.InsertSubjectService(tt.args.subject); (err != nil) != tt.wantErr {
				t.Errorf("subjectService.InsertSubjectService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_subjectService_GetSubjectService(t *testing.T) {
	type fields struct {
		d subjectDatastore
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSubjectDatastore := NewMocksubjectDatastore(ctrl)
	st := models.Subject{
		Id:   5,
		Name: "Test1",
	}
	// st2 := models.Subject{}

	type args struct {
		id int
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     models.Subject
		wantErr  bool
		mockcall []interface{}
	}{
		{
			name:   "GetSubjectSvcBaseCase",
			fields: fields{d: mockSubjectDatastore},
			args:   args{st.Id},
			want:   st,
			mockcall: []interface{}{
				mockSubjectDatastore.EXPECT().GetSubject(st.Id).Return(st, nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sh := subjectService{
				d: tt.fields.d,
			}
			got, err := sh.GetSubjectService(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("subjectService.GetSubjectService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subjectService.GetSubjectService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subjectService_GetAllSubjectService(t *testing.T) {
	type fields struct {
		d subjectDatastore
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSubjectDatastore := NewMocksubjectDatastore(ctrl)

	st := []models.Subject{
		{
			Id:   5,
			Name: "Test1",
		},
	}
	tests := []struct {
		name     string
		fields   fields
		want     []models.Subject
		wantErr  bool
		mockcall []interface{}
	}{
		{
			name:   "GetAllSubjectSvcBaseCase",
			fields: fields{d: mockSubjectDatastore},
			want:   st,
			mockcall: []interface{}{
				mockSubjectDatastore.EXPECT().GetAllSubject().Return(st, nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sh := subjectService{
				d: tt.fields.d,
			}
			got, err := sh.GetAllSubjectService()
			if (err != nil) != tt.wantErr {
				t.Errorf("subjectService.GetAllSubjectService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subjectService.GetAllSubjectService() = %v, want %v", got, tt.want)
			}
		})
	}
}
