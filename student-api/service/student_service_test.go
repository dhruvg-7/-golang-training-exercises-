package service

import (
	"errors"
	"reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/student-api/models"
)

func Test_studentService_InsertService(t *testing.T) {
	type fields struct {
		d      studentDatastore
		subSvc subjectService
		enrSvc enrolmentService
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockstudentDatastore := NewMockstudentDatastore(ctrl)

	st := []models.Student{
		{

			RollNo: 5,
			Name:   "Test1",
			Age:    55,
		},
		{
			RollNo: 0,
			Name:   "Test1",
			Age:    55,
		},
		{
			RollNo: 5,
			Name:   "",
			Age:    55,
		},
	}
	type args struct {
		student models.Student
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		mockcall []interface{}
		wantErr  bool
	}{
		{

			name:     "InsertServiceBaseCase",
			fields:   fields{d: mockstudentDatastore},
			args:     args{st[0]},
			wantErr:  false,
			mockcall: []interface{}{mockstudentDatastore.EXPECT().Insert(st[0]).Return(nil)},
		},
		{
			name:     "InsertServiceErrorCase",
			fields:   fields{d: mockstudentDatastore},
			args:     args{st[0]},
			wantErr:  true,
			mockcall: []interface{}{mockstudentDatastore.EXPECT().Insert(st[0]).Return(errors.New("Insert service error"))},
		},
		// {
		// 	name:     "InsertServiceIdErrorCase",
		// 	fields:   fields{d: mockstudentDatastore},
		// 	args:     args{st[1]},
		// 	wantErr:  true,
		// 	mockcall: []interface{}{mockstudentDatastore.EXPECT().Insert(st[1]).Return(errors.New("Insert service error"))},
		// },
		// {
		// 	name:     "InsertServiceNameErrorCase",
		// 	fields:   fields{d: mockstudentDatastore},
		// 	args:     args{st2},
		// 	wantErr:  true,
		// 	mockcall: []interface{}{mockstudentDatastore.EXPECT().Insert(st).Return(errors.New("Insert service error"))},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stuSvc := studentService{
				d:      tt.fields.d,
				subSvc: tt.fields.subSvc,
				enrSvc: tt.fields.enrSvc,
			}
			if err := stuSvc.InsertService(tt.args.student); (err != nil) != tt.wantErr {
				t.Errorf("studentService.InsertService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_studentService_ReadByIdService(t *testing.T) {
	type fields struct {
		d      studentDatastore
		subSvc subjectService
		enrSvc enrolmentService
	}
	type args struct {
		id int
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockstudentDatastore := NewMockstudentDatastore(ctrl)

	st := []models.Student{
		{
			RollNo: 5,
			Name:   "Test1",
			Age:    55,
		},
		{
			RollNo: 0,
			Name:   "Test1",
			Age:    55,
		},
		{
			Name:   "",
			RollNo: 0,
			Age:    0,
		},
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     models.Student
		mockcall []interface{}
		wantErr  bool
	}{
		{

			name:     "ReadServiceBaseCase",
			fields:   fields{d: mockstudentDatastore},
			args:     args{st[0].RollNo},
			want:     st[0],
			wantErr:  false,
			mockcall: []interface{}{mockstudentDatastore.EXPECT().Read(st[0].RollNo).Return(st[0], nil)},
		},
		{
			name:     "ReadServiceErrorCase",
			fields:   fields{d: mockstudentDatastore},
			args:     args{st[0].RollNo},
			want:     st[2],
			wantErr:  true,
			mockcall: []interface{}{mockstudentDatastore.EXPECT().Read(st[0].RollNo).Return(st[2], errors.New("Read service error"))},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stuSvc := studentService{
				d:      tt.fields.d,
				subSvc: tt.fields.subSvc,
				enrSvc: tt.fields.enrSvc,
			}
			got, err := stuSvc.ReadByIdService(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("studentService.ReadByIdService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("studentService.ReadByIdService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_studentService_ReadAllService(t *testing.T) {
	type fields struct {
		d      studentDatastore
		subSvc subjectService
		enrSvc enrolmentService
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockstudentDatastore := NewMockstudentDatastore(ctrl)

	st := []models.Student{
		{
			RollNo: 5,
			Name:   "Test1",
			Age:    55,
		},
		{
			RollNo: 1,
			Name:   "Test1",
			Age:    55,
		},
	}
	tests := []struct {
		name     string
		fields   fields
		want     []models.Student
		mockcall []interface{}
		wantErr  bool
	}{
		{
			name:     "ReadServiceBaseCase",
			fields:   fields{d: mockstudentDatastore},
			want:     st,
			wantErr:  false,
			mockcall: []interface{}{mockstudentDatastore.EXPECT().ReadAll().Return(st, nil)},
		},
		{
			name:     "ReadServiceErrorCase",
			fields:   fields{d: mockstudentDatastore},
			wantErr:  true,
			mockcall: []interface{}{mockstudentDatastore.EXPECT().ReadAll().Return(nil, errors.New("ReadALL error"))},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stuSvc := studentService{
				d:      tt.fields.d,
				subSvc: tt.fields.subSvc,
				enrSvc: tt.fields.enrSvc,
			}
			got, err := stuSvc.ReadAllService()
			if (err != nil) != tt.wantErr {
				t.Errorf("studentService.ReadAllService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("studentService.ReadAllService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_studentService_Update(t *testing.T) {
	type fields struct {
		d      studentDatastore
		subSvc subjectService
		enrSvc enrolmentService
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockstudentDatastore := NewMockstudentDatastore(ctrl)

	st := []models.Student{
		{
			RollNo: 5,
			Name:   "Test1",
			Age:    55,
		},
		{
			RollNo: 0,
			Name:   "Test1",
			Age:    55,
		},
		{
			Name:   "",
			RollNo: 0,
			Age:    0,
		},
	}
	type args struct {
		st models.Student
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantErr  bool
		mockcall []interface{}
	}{
		{
			name:    "UpdateServiceBaseCase",
			fields:  fields{d: mockstudentDatastore},
			args:    args{st[0]},
			wantErr: false,
			mockcall: []interface{}{
				mockstudentDatastore.EXPECT().Read(st[0].RollNo).Return(st[0], nil),
				mockstudentDatastore.EXPECT().Update(st[0]).Return(nil)},
		},
		// {
		// 	name:    "UpdateServiceReadErrorCase",
		// 	fields:  fields{d: mockstudentDatastore},
		// 	args:    args{st[0]},
		// 	wantErr: true,
		// 	mockcall: []interface{}{
		// 		mockstudentDatastore.EXPECT().Read(st[0].RollNo).Return(nil, errors.New("Read Error")),
		// 		mockstudentDatastore.EXPECT().Update(st[0]).Return(nil)},
		// },
		{
			name:    "UpdateServiceErrorCase",
			fields:  fields{d: mockstudentDatastore},
			args:    args{st[0]},
			wantErr: true,
			mockcall: []interface{}{
				mockstudentDatastore.EXPECT().Read(st[0].RollNo).Return(st[0], nil),
				mockstudentDatastore.EXPECT().Update(st[0]).Return(errors.New("Update error"))},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stuSvc := studentService{
				d:      tt.fields.d,
				subSvc: tt.fields.subSvc,
				enrSvc: tt.fields.enrSvc,
			}
			if err := stuSvc.Update(tt.args.st); (err != nil) != tt.wantErr {
				t.Errorf("studentService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_studentService_Delete(t *testing.T) {
	type fields struct {
		d      studentDatastore
		subSvc subjectService
		enrSvc enrolmentService
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockstudentDatastore := NewMockstudentDatastore(ctrl)

	st := []models.Student{
		{
			RollNo: 5,
			Name:   "Test1",
			Age:    55,
		},
		{
			RollNo: 0,
			Name:   "Test1",
			Age:    55,
		},
		{
			RollNo: 0,
			Name:   "",
			Age:    0,
		},
	}
	type args struct {
		Id string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantErr  bool
		mockcall []interface{}
	}{
		{
			name:    "DeleteServiceBaseCase",
			fields:  fields{d: mockstudentDatastore},
			args:    args{"5"},
			wantErr: false,
			mockcall: []interface{}{
				mockstudentDatastore.EXPECT().Read(st[0].RollNo).Return(st[0], nil),
				mockstudentDatastore.EXPECT().Delete(st[0].RollNo).Return(nil)},
		},
		{
			name:    "DeleteServiceErrorCase",
			fields:  fields{d: mockstudentDatastore},
			args:    args{"5"},
			wantErr: true,
			mockcall: []interface{}{
				mockstudentDatastore.EXPECT().Read(st[0].RollNo).Return(st[0], nil),
				mockstudentDatastore.EXPECT().Delete(st[0].RollNo).Return(errors.New("Delete Error"))},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stuSvc := studentService{
				d:      tt.fields.d,
				subSvc: tt.fields.subSvc,
				enrSvc: tt.fields.enrSvc,
			}
			if err := stuSvc.Delete(tt.args.Id); (err != nil) != tt.wantErr {
				t.Errorf("studentService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
