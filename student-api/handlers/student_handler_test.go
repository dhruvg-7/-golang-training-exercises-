package handlers

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"

	models "github.com/student-api/models"
)

func Test_handler_InsertStudent(t *testing.T) {
	type fields struct {
		d datastore
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler{
				d: tt.fields.d,
			}
			h.InsertStudent(tt.args.w, tt.args.r)
		})
	}
}

func getRequestResponse(w httptest.ResponseRecorder) (result string) {
	res := w.Result()
	defer res.Body.Close()
	data, _ := io.ReadAll(res.Body)
	formattedData := string(data)
	result = strings.TrimSpace(formattedData)
	return
}

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	vars := map[string]string{
		"id": "8",
	}
	s := models.Student{Name: "Ankit", Age: 21, RollNo: 2}
	mockdatastore := NewMockdatastore(ctrl)
	h := handler{mockdatastore}
	w := httptest.NewRecorder()
	want := `{"Name":"Ankit","Age":21,"RollNo":2}`
	r := mux.SetURLVars(httptest.NewRequest(http.MethodGet, "/student", nil), vars)
	mockdatastore.EXPECT().Read(vars["id"]).Return(s, nil)
	h.GetStudent(w, r)
	result := getRequestResponse(*w)
	if !reflect.DeepEqual(want, result) {
		t.Errorf("TestGet Failed...Expected %v and Got %v", want, result)
	}
}

func TestGetError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	vars := map[string]string{
		"id": "8",
	}
	mockdatastore := NewMockdatastore(ctrl)
	h := handler{mockdatastore}
	want := "Got an error"
	w := httptest.NewRecorder()
	r := mux.SetURLVars(httptest.NewRequest(http.MethodGet, "/student", nil), vars)
	mockdatastore.EXPECT().Read(vars["id"]).Return(models.Student{}, errors.New("Got an error"))
	h.GetStudent(w, r)
	result := getRequestResponse(*w)
	if !reflect.DeepEqual(want, result) {
		t.Errorf("TestGet Failed...Expected %v and Got %v", want, result)
	}
}

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdatastore := NewMockdatastore(ctrl)
	w :=
		httptest.NewRecorder()
	r :=
		httptest.NewRequest(http.MethodGet, "/student", nil)
	want := `[{"Name":"Ankit","Age":21,"RollNo":1},{"Name":"Amit","Age":21,"RollNo":2}]`
	student := []models.Student{
		{Name: "Ankit", Age: 21, RollNo: 1},
		{Name: "Amit", Age: 21, RollNo: 2},
	}
	mockdatastore.EXPECT().ReadAll().Return(student, nil)

	h := handler{mockdatastore}
	h.GetAllStudent(w, r)
	result := getRequestResponse(*w)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("TestGet Failed...Expected %v and Got %v", want, result)
	}
}

func TestGetAllError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdatastore := NewMockdatastore(ctrl)
	w :=
		httptest.NewRecorder()
	r :=
		httptest.NewRequest(http.MethodGet, "/student", nil)
	want := "Error in getting response"
	student := []models.Student{
		{Name: "Ankit", Age: 21, RollNo: 1},
		{Name: "Amit", Age: 21, RollNo: 2},
	}
	mockdatastore.EXPECT().ReadAll().Return(student, errors.New("Error in getting response"))

	h := handler{mockdatastore}
	h.GetAllStudent(w, r)
	result := getRequestResponse(*w)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("TestGet Failed...Expected %v and Got %v", want, result)
	}
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdatastore := NewMockdatastore(ctrl)
	payload := `{"Name":"Ankit","Age":21,"RollNo":1}`
	want := "Succesfully Inserted"
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/student", bytes.NewBuffer([]byte(payload)))
	mockdatastore.EXPECT().Insert(models.Student{Name: "Ankit", Age: 21, RollNo: 1}).Return(nil)
	h := handler{mockdatastore}
	h.InsertStudent(w, r)
	result := getRequestResponse(*w)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("TestGet Failed...Expected %v and Got %v", want, result)
	}
}

func TestCreateErr2(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdatastore := NewMockdatastore(ctrl)
	payload := `{"Name":"Ankit","Age":21,"RollNo":1}`
	want := "Some Error"
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/student", bytes.NewBuffer([]byte(payload)))
	mockdatastore.EXPECT().Insert(models.Student{Name: "Ankit", Age: 21, RollNo: 1}).Return(errors.New("Some Error"))
	h := handler{mockdatastore}
	h.InsertStudent(w, r)
	result := getRequestResponse(*w)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("TestGet Failed...Expected %v and Got %v", want, result)
	}
}

func TestCreateErr1(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdatastore := NewMockdatastore(ctrl)
	payload := `{"Name":"Ankit","Age":21,"RollNo":1`
	want := "Invalid JSON Format"
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/student", bytes.NewBuffer([]byte(payload)))
	h := handler{mockdatastore}
	h.InsertStudent(w, r)
	result := getRequestResponse(*w)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("TestGet Failed...Expected %v and Got %v", want, result)
	}
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdatastore := NewMockdatastore(ctrl)
	vars := map[string]string{
		"id": "8",
	}
	w := httptest.NewRecorder()
	r := mux.SetURLVars(httptest.NewRequest(http.MethodDelete, "/student", nil), vars)
	mockdatastore.EXPECT().Delete(vars["id"]).Return(nil)
	h := handler{mockdatastore}
	h.DeleteStudent(w, r)

	want := "Successfully Deleted"
	result := getRequestResponse(*w)

	if !reflect.DeepEqual(want, result) {
		t.Errorf("TestGet Failed...Expected %v and Got %v", want, result)
	}
}

func TestDeleteErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdatastore := NewMockdatastore(ctrl)
	vars := map[string]string{
		"id": "8",
	}
	w := httptest.NewRecorder()
	r := mux.SetURLVars(httptest.NewRequest(http.MethodDelete, "/student", nil), vars)
	mockdatastore.EXPECT().Delete(vars["id"]).Return(errors.New("Deletion Failed"))
	h := handler{mockdatastore}
	h.DeleteStudent(w, r)

	want := "Deletion Failed"
	result := getRequestResponse(*w)

	if !reflect.DeepEqual(want, result) {
		t.Errorf("TestGet Failed...Expected %v and Got %v", want, result)
	}
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdatastore := NewMockdatastore(ctrl)
	vars := map[string]string{
		"id": "8",
	}
	payload := `{"Name":"Ankit","Age":21}`
	want := `SuccessFull Put Request`
	w := httptest.NewRecorder()
	r := mux.SetURLVars(httptest.NewRequest(http.MethodPut, "/student", bytes.NewBuffer([]byte(payload))), vars)
	mockdatastore.EXPECT().Update(models.Student{Name: "Ankit", Age: 21, RollNo: 8}).Return(nil)
	h := handler{mockdatastore}
	h.UpdateStudent(w, r)
	result := getRequestResponse(*w)
	if !reflect.DeepEqual(want, result) {
		t.Errorf("TestGet Failed...Expected %v and Got %v", want, result)
	}
}

func TestUpdateErr1(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdatastore := NewMockdatastore(ctrl)
	vars := map[string]string{
		"id": "8",
	}
	payload := `{"Name":"Ankit","Age":21`
	want := `Invalid JSON Format`
	w := httptest.NewRecorder()
	r := mux.SetURLVars(httptest.NewRequest(http.MethodPut, "/student", bytes.NewBuffer([]byte(payload))), vars)
	h := handler{mockdatastore}
	h.UpdateStudent(w, r)
	result := getRequestResponse(*w)
	if !reflect.DeepEqual(want, result) {
		t.Errorf("TestGet Failed...Expected %v and Got %v", want, result)
	}
}

func TestUpdateErr2(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdatastore := NewMockdatastore(ctrl)
	vars := map[string]string{
		"id": "8",
	}
	payload := `{"Name":"Ankit","Age":21}`
	want := `Updation Failed`
	w := httptest.NewRecorder()
	r := mux.SetURLVars(httptest.NewRequest(http.MethodPut, "/student", bytes.NewBuffer([]byte(payload))), vars)
	mockdatastore.EXPECT().Update(models.Student{Name: "Ankit", Age: 21, RollNo: 8}).Return(errors.New("Updation Failed"))
	h := handler{mockdatastore}
	h.UpdateStudent(w, r)
	result := getRequestResponse(*w)
	if !reflect.DeepEqual(want, result) {
		t.Errorf("TestGet Failed...Expected %v and Got %v", want, result)
	}
}
