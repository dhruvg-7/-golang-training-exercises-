package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/student-api/models"
)

type studentService interface {
	InsertService(models.Student) error
	Update(models.Student) error
	Delete(string) error
	ReadByIdService(int) (models.Student, error)
	ReadByDetailService(string, string) ([]models.Student, error)
	ReadAllService() ([]models.Student, error)
	EnrolStudentSvs(stuId, subId string) error
	GetstudentBySubSvs(stuId string) ([]models.EnrolmentList, error)
}
type studentServiceHandler struct {
	s studentService
}

func NewStudentServicehandler(s studentService) studentServiceHandler {
	return studentServiceHandler{s}
}

// FUNCTION TO INSERT http
func (sh studentServiceHandler) InsertStudent(w http.ResponseWriter, r *http.Request) {
	stu := models.Student{}
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &stu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = sh.s.InsertService(stu)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		fmt.Fprintf(w, "Insert Done")
	}
}

//TODO: FUNCTION TO UPDATE http
func (sh studentServiceHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	st := models.Student{}
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &st)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = sh.s.Update(st)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {

		fmt.Fprintf(w, "Update Sucessful")
	}
}

//TODO: FUNCTION TO DELETE http
func (sh studentServiceHandler) DeleteStudent(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	err := sh.s.Delete(id)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {

		fmt.Fprintf(w, "Delete Sucessfull")
	}
}

//TODO: FUNCTION TO READ http
func (sh studentServiceHandler) GetStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	rollno, err := strconv.Atoi(string(id))
	if err != nil {
		println(err.Error())
	}

	ans, err := sh.s.ReadByIdService(rollno)

	if err != nil {
		println(err.Error())
	}
	json.NewEncoder(w).Encode(ans)
}

func (sh studentServiceHandler) GetAllStudent(w http.ResponseWriter, r *http.Request) {

	ans, err := sh.s.ReadAllService()
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("read error reason: %s", err))
		return
	}
	json.NewEncoder(w).Encode(ans)
}

func (sh studentServiceHandler) GetStudentByDetail(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	fmt.Printf(id)
	ans, err := sh.s.ReadByDetailService(id, name)

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	json.NewEncoder(w).Encode(ans)

}
func (sh studentServiceHandler) EnrolStudentHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	stuId := params["stuId"]
	subId := params["subId"]
	err := sh.s.EnrolStudentSvs(stuId, subId)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		fmt.Fprintf(w, "Enroled Student %s to Subject %s", stuId, subId)
	}
}

func (sh studentServiceHandler) GetSubjectsByStuSvs(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	stuId := params["stuId"]
	ans, err := sh.s.GetstudentBySubSvs(stuId)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	json.NewEncoder(w).Encode(ans)

}
