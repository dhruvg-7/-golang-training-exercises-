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

type datastore interface {
	Insert(models.Student) error
	Update(models.Student) error
	Delete(int) error
	Read(int) (models.Student, error)
	ReadAll() (models.Student, error)
}

type handler struct {
	d datastore
}

func NewHandler(d datastore) handler {
	return handler{d}
}

//FUNCTION TO INSERT  http
func (h handler) InsertStudent(w http.ResponseWriter, r *http.Request) {
	s := models.Student{}

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.d.Insert(s)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
}

//TODO: FUNCTION TO UPDATE http
func (h handler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	s := models.Student{}
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.d.Update(s)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
}

//TODO: FUNCTION TO DELETE http
func (h handler) DeleteStudent(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	rollno, err := strconv.Atoi(string(id))
	if err != nil {
		println(err.Error())
	}

	err = h.d.Delete(rollno)
	if err != nil {
		println(err.Error())

	}
	fmt.Fprintf(w, err.Error())
}

//TODO: FUNCTION TO READ http
func (h handler) GetStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	rollno, err := strconv.Atoi(string(id))
	if err != nil {
		println(err.Error())
	}

	ans, err := h.d.Read(rollno)
	if err != nil {
		println(err.Error())
	}
	fmt.Fprintf(w, "%v", ans)
}

func (h handler) GetAllStudent(w http.ResponseWriter, r *http.Request) {

	ans, err := h.d.ReadAll()
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("read error reason: %s", err))
		return
	}
	fmt.Fprintf(w, "%v", ans)
}
