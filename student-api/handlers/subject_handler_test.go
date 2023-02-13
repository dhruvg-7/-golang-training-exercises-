package handlers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	reflect "reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/student-api/models"
)

func Test_subjectHandler_GetSubject(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSubjectService := NewMocksubjectService(ctrl)

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	st := models.Subject{
		Id:   1,
		Name: "TestSubject1",
	}
	st2 := models.Subject{}
	tests := []struct {
		name     string
		serv     subjectService
		args     args
		mockcall []interface{}
		want     string
	}{
		{
			name: "GetSubjectHandlerBaseCase",
			serv: mockSubjectService,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/subject/5", nil),
			},
			mockcall: []interface{}{
				mockSubjectService.EXPECT().GetSubjectService(gomock.Any()).Return(st, nil),
			},
			want: `{"Id":1,"Name":"TestSubject1"}`,
		},
		{
			name: "GetSubjectHandlerErrorCase",
			serv: mockSubjectService,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/subject/9", nil),
			},
			mockcall: []interface{}{
				mockSubjectService.EXPECT().GetSubjectService(gomock.Any()).Return(st2, errors.New("Invalid id error")),
			},
			want: `{"Id":0,"Name":""}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sh := subjectHandler{
				s: tt.serv,
			}
			sh.GetSubject(tt.args.w, tt.args.r)
			result := getRequestResponse(*tt.args.w)
			if !reflect.DeepEqual(tt.want, result) {
				t.Errorf("TestGet Failed...Expected %v and Got %v", tt.want, result)
			}
		})
	}
}

func Test_subjectHandler_NewSubject(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSubjectService := NewMocksubjectService(ctrl)

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name     string
		serv     subjectService
		args     args
		mockcall []interface{}
		want     string
	}{
		{
			name: "NewSubjectHandlerBaseCase",
			serv: mockSubjectService,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/subject", bytes.NewBuffer([]byte(`{"Id":1,"Name":"TestSubject1"}`))),
			},
			mockcall: []interface{}{
				mockSubjectService.EXPECT().InsertSubjectService(gomock.Any()).Return(nil),
			},
			want: `Insert Done`,
		},
		{
			name: "NewSubjectHandlerJsonErrorCase",
			serv: mockSubjectService,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/subject", bytes.NewBuffer([]byte(`{"Id": 0,"Name":" "}`))),
			},
			mockcall: []interface{}{
				mockSubjectService.EXPECT().InsertSubjectService(gomock.Any()).Return(errors.New("unexpected error")),
			},
			want: `unexpected error`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sh := subjectHandler{
				s: tt.serv,
			}
			sh.NewSubject(tt.args.w, tt.args.r)
			result := getRequestResponse(*tt.args.w)
			if !reflect.DeepEqual(tt.want, result) {
				t.Errorf("TestGet Failed...Expected %v and Got %v", tt.want, result)
			}
		})
	}
}

func Test_subjectHandler_GetAllSubject(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSubjectService := NewMocksubjectService(ctrl)
	st := []models.Subject{{

		Id:   1,
		Name: "TestSubject1",
	},
	}
	st2 := []models.Subject{}
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name     string
		serv     subjectService
		args     args
		mockcall []interface{}
		want     string
	}{
		{
			name: "GetSubjectHandlerBaseCase",
			serv: mockSubjectService,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/subject", nil),
			},
			mockcall: []interface{}{
				mockSubjectService.EXPECT().GetAllSubjectService().Return(st, nil),
			},
			want: `[{"Id":1,"Name":"TestSubject1"}]`,
		},
		{
			name: "GetSubjectHandlerErrorCase",
			serv: mockSubjectService,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/subject", nil),
			},
			mockcall: []interface{}{
				mockSubjectService.EXPECT().GetAllSubjectService().Return(st2, errors.New("Invalid id error")),
			},
			want: `read error reason: Invalid id error`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sh := subjectHandler{
				s: tt.serv,
			}
			sh.GetAllSubject(tt.args.w, tt.args.r)
			result := getRequestResponse(*tt.args.w)
			if !reflect.DeepEqual(tt.want, result) {
				t.Errorf("TestGet Failed...Expected %v and Got %v", tt.want, result)
			}
		})
	}
}
