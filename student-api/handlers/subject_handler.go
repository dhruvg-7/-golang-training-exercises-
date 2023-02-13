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

type subjectService interface {
	GetAllSubjectService() ([]models.Subject, error)
	GetSubjectService(id int) (models.Subject, error)
	InsertSubjectService(subject models.Subject) error
}

type subjectHandler struct {
	s subjectService
}

func NewSubjectHandler(s subjectService) subjectHandler {
	return subjectHandler{s}
}

func (sh subjectHandler) GetSubject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(string(params["id"]))
	if err != nil {
		println(err.Error())
	}
	ans, err := sh.s.GetSubjectService(id)

	if err != nil {
		println(err.Error())
	}
	json.NewEncoder(w).Encode(ans)
}

func (sh subjectHandler) NewSubject(w http.ResponseWriter, r *http.Request) {
	sub := models.Subject{}

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &sub)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = sh.s.InsertSubjectService(sub)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {

		fmt.Fprintf(w, "Insert Done")
	}
}

func (sh subjectHandler) GetAllSubject(w http.ResponseWriter, r *http.Request) {

	ans, err := sh.s.GetAllSubjectService()
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("read error reason: %s", err))
		return
	}
	json.NewEncoder(w).Encode(ans)
}
