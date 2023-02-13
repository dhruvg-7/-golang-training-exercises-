package handlers

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	reflect "reflect"
	"strings"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/student-api/models"
)

func getRequestResponse(w httptest.ResponseRecorder) (result string) {
	res := w.Result()
	defer res.Body.Close()
	data, _ := io.ReadAll(res.Body)
	formattedData := string(data)
	result = strings.TrimSpace(formattedData)
	return
}
func Test_studentServiceHandler_InsertStudent(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStudentService := NewMockstudentService(ctrl)

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	st := []models.Student{
		{

			Name:   "Test1",
			RollNo: 1,
			Age:    99,
		},
		{

			Name:   "Test1",
			RollNo: 0,
			Age:    99,
		},
	}
	tests := []struct {
		name     string
		serv     studentService
		mockcall []interface{}
		args     args
		want     string
	}{
		{
			name: "InsertHandlerBaseCase",
			serv: mockStudentService,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/student", bytes.NewBuffer([]byte(`{"Name":"Test1","RollNo":1,"Age":99}`))),
			},
			mockcall: []interface{}{
				mockStudentService.EXPECT().InsertService(st[0]).Return(nil),
			},
			want: "Insert Done!",
		},
		{
			name: "InsertStudentErrorCase",
			serv: mockStudentService,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/student", bytes.NewBuffer([]byte(`{"Name":"Test1","RollNo":0,"Age":99}`))),
			},
			mockcall: []interface{}{
				mockStudentService.EXPECT().InsertService(st[1]).Return(errors.New("Invalid rollnumber error")),
			},
			want: "Insert fail",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sh := studentServiceHandler{
				tt.serv,
			}
			sh.InsertStudent(tt.args.w, tt.args.r)
		})
	}
}

func Test_studentServiceHandler_UpdateStudent(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStudentService := NewMockstudentService(ctrl)

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	st := []models.Student{
		{

			Name:   "Test1",
			RollNo: 1,
			Age:    99,
		},
		{

			Name:   "Test1",
			RollNo: 0,
			Age:    99,
		},
	}
	tests := []struct {
		name     string
		serv     studentService
		args     args
		mockcall []interface{}
	}{
		{
			name: "UpdateHandlerBaseCase",
			serv: mockStudentService,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPut, "/student", bytes.NewBuffer([]byte(`{"Name":"Test1","RollNo":1,"Age":99}`))),
			},
			mockcall: []interface{}{
				mockStudentService.EXPECT().Update(st[0]).Return(nil),
			},
		},
		{
			name: "UpdateHandlerErrorCase",
			serv: mockStudentService,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPut, "/student", bytes.NewBuffer([]byte(`{"Name":"Test1","RollNo":0,"Age":99}`))),
			},
			mockcall: []interface{}{
				mockStudentService.EXPECT().Update(st[1]).Return(errors.New("Invalid rollnumber error")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sh := studentServiceHandler{
				s: mockStudentService,
			}
			sh.UpdateStudent(tt.args.w, tt.args.r)
		})
	}
}

func Test_studentServiceHandler_DeleteStudent(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStudentService := NewMockstudentService(ctrl)

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name     string
		serv     studentService
		args     args
		mockcall []interface{}
		want     string
	}{
		{
			name: "DeleteHandlerBaseCase",
			serv: mockStudentService,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodDelete, "/student/5", nil),
			},
			mockcall: []interface{}{
				mockStudentService.EXPECT().Delete(gomock.Any()).Return(nil),
			},
			want: "Delete Sucessfull",
		},
		{
			name: "DeleteHandlerErrorCase",
			serv: mockStudentService,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodDelete, "/student/5", nil),
			},
			mockcall: []interface{}{
				mockStudentService.EXPECT().Delete(gomock.Any()).Return(errors.New("Invalid rollnumber error")),
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sh := studentServiceHandler{
				s: mockStudentService,
			}
			sh.DeleteStudent(tt.args.w, tt.args.r)
		})
	}
}

func Test_studentServiceHandler_GetStudent(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStudentService := NewMockstudentService(ctrl)

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	st := []models.Student{
		{
			RollNo: 1,
			Name:   "Test1",
			Age:    99,
		},
		{

			RollNo: 0,
			Name:   "Test1",
			Age:    99,
		},
	}
	tests := []struct {
		name     string
		serv     studentService
		args     args
		mockcall []interface{}
		want     string
	}{
		{
			name: "GetHandlerBaseCase",
			serv: mockStudentService,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/student/5", nil),
			},
			mockcall: []interface{}{
				mockStudentService.EXPECT().ReadByIdService(gomock.Any()).Return(st[0], nil),
			},
			want: `{"RollNo":1,"Name":"Test1","Age":19}`,
		},
		{
			name: "GetHandlerErrorCase",
			serv: mockStudentService,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/student/9", nil),
			},
			mockcall: []interface{}{
				mockStudentService.EXPECT().ReadByIdService(gomock.Any()).Return(st[1], errors.New("Invalid rollnumber error")),
			},
			want: `{"RollNo":0,"Name":"","Age":0}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sh := studentServiceHandler{
				s: tt.serv,
			}
			sh.GetStudent(tt.args.w, tt.args.r)
		})
	}
}

func Test_studentServiceHandler_GetAllStudent(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStudentService := NewMockstudentService(ctrl)

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	st := []models.Student{
		{

			Name:   "Test1",
			RollNo: 1,
			Age:    99,
		},
		{

			Name:   "Test2",
			RollNo: 2,
			Age:    99,
		},
	}
	tests := []struct {
		name     string
		serv     studentService
		args     args
		mockcall []interface{}
		want     string
	}{
		{
			name: "GetAllHandlerBaseCase",
			serv: mockStudentService,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/student/all", nil),
			},
			mockcall: []interface{}{
				mockStudentService.EXPECT().ReadAllService().Return(st, nil),
			},
			want: `[{"RollNo":1,"Name":"Test1","Age":99},{"RollNo":2,"Name":"Test2","Age":99}]`,
		},
		{
			name: "GetAllHandlerErrorCase",
			serv: mockStudentService,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/student/all", nil),
			},
			mockcall: []interface{}{
				mockStudentService.EXPECT().ReadAllService().Return(st, errors.New("Invalid rollnumber error")),
			},
			want: `read error reason: Invalid rollnumber error`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sh := studentServiceHandler{
				s: tt.serv,
			}
			sh.GetAllStudent(tt.args.w, tt.args.r)
			result := getRequestResponse(*tt.args.w)
			if !reflect.DeepEqual(tt.want, result) {
				t.Errorf("TestGet Failed...Expected %v and Got %v", tt.want, result)
			}
		})
	}
}

func Test_studentServiceHandler_GetStudentByDetail(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStudentService := NewMockstudentService(ctrl)

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	st := []models.Student{
		{

			Name:   "Test1",
			RollNo: 1,
			Age:    99,
		},
		{

			Name:   "Test1",
			RollNo: 2,
			Age:    99,
		},
	}
	tests := []struct {
		name     string
		serv     studentService
		args     args
		mockcall []interface{}
		want     string
	}{
		{
			name: "GetStudentByDetailHandlerBaseCase",
			serv: mockStudentService,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/student?id=1", nil),
			},
			mockcall: []interface{}{
				mockStudentService.EXPECT().ReadByDetailService(gomock.Any(), gomock.Any()).Return([]models.Student{st[0]}, nil),
			},
			want: `[{"RollNo":1,"Name":"Test1","Age":99}]`,
		},
		{
			name: "GetStudentByDetailHandlerNameCase",
			serv: mockStudentService,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/student?Name=test1", nil),
			},
			mockcall: []interface{}{
				mockStudentService.EXPECT().ReadByDetailService(gomock.Any(), gomock.Any()).Return(st, nil),
			},
			want: `[{"RollNo":1,"Name":"Test1","Age":99},{"RollNo":2,"Name":"Test1","Age":99}]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sh := studentServiceHandler{
				s: tt.serv,
			}
			sh.GetStudentByDetail(tt.args.w, tt.args.r)
			result := getRequestResponse(*tt.args.w)
			if !reflect.DeepEqual(tt.want, result) {
				t.Errorf("TestGet Failed...Expected %v and Got %v", tt.want, result)
			}
		})
	}
}

func Test_studentServiceHandler_EnrolStudentHandler(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStudentService := NewMockstudentService(ctrl)

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name     string
		serv     studentService
		args     args
		mockcall []interface{}
		want     string
	}{
		{
			name: "EnrolStudentHandlerBaseCase",
			serv: mockStudentService,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/student/1/subject2", nil),
			},
			mockcall: []interface{}{
				mockStudentService.EXPECT().EnrolStudentSvs(gomock.Any(), gomock.Any()).Return(nil),
			},
			want: `Enroled Student  to Subject`,
		},
		{
			name: "EnrolStudentHandlerErrorCase",
			serv: mockStudentService,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/student/3/subject/2", nil),
			},
			mockcall: []interface{}{
				mockStudentService.EXPECT().EnrolStudentSvs(gomock.Any(), gomock.Any()).Return(errors.New("Invalid rollnumber error")),
			},
			want: `Invalid rollnumber error`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sh := studentServiceHandler{
				s: tt.serv,
			}
			sh.EnrolStudentHandler(tt.args.w, tt.args.r)
			result := getRequestResponse(*tt.args.w)
			if !reflect.DeepEqual(tt.want, result) {
				t.Errorf("TestGet Failed...Expected %v and Got %v", tt.want, result)
			}
		})
	}
}

func Test_studentServiceHandler_GetSubjectsByStuSvs(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStudentService := NewMockstudentService(ctrl)

	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	st := []models.EnrolmentList{
		{StudentId: 2, SubjectId: 9},
		{StudentId: 2, SubjectId: 1},
	}
	tests := []struct {
		name     string
		serv     studentService
		args     args
		mockcall []interface{}
		wantErr  bool
		want     string
	}{
		{
			name: "GetSubjectByStudentIdHandlerBaseCase",
			serv: mockStudentService,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/student/2/subject", nil),
			},
			mockcall: []interface{}{
				mockStudentService.EXPECT().GetstudentBySubSvs(gomock.Any()).Return(st, nil),
			},
			want: `[{"StudentId":2,"SubjectId":9},{"StudentId":2,"SubjectId":1}]`,
		},
		{
			name: "GetSubjectByStudentIdHandlerErrorCase",
			serv: mockStudentService,
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/student/2/subject", nil),
			},
			mockcall: []interface{}{
				mockStudentService.EXPECT().GetstudentBySubSvs(gomock.Any()).Return(nil, errors.New("Invalid rollnumber error")),
			},
			wantErr: true,
			want:    `Invalid rollnumber errornull`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sh := studentServiceHandler{
				s: tt.serv,
			}
			sh.GetSubjectsByStuSvs(tt.args.w, tt.args.r)
			result := getRequestResponse(*tt.args.w)
			if !reflect.DeepEqual(tt.want, result) {
				t.Errorf("TestGet Failed...Expected %v and Got %v", tt.want, result)
			}
		})
	}
}
