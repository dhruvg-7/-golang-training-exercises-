package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/student-api/handlers"
	"github.com/student-api/service"
	"github.com/student-api/stores"
)

func main() {
	r := mux.NewRouter()
	con, err := sql.Open("mysql", "root:pass@tcp(localhost:3306)/student")
	if err != nil {
		fmt.Printf(err.Error())
	}

	subjectCon := stores.NewSubjectStore(con)
	subSvs := service.NewSubjectService(subjectCon)
	subh := handlers.NewSubjectHandler(subSvs)

	enrolmentCon := stores.NewEnrolmentStore(con)
	enrSvs := service.NewEnrolmentService(enrolmentCon)

	studentCon := stores.NewStoreCon(con)
	stuSvs := service.NewStudentService(studentCon, subSvs, enrSvs)
	stuh := handlers.NewStudentServicehandler(stuSvs)

	//student handlers
	r.HandleFunc("/student", stuh.InsertStudent).Methods("POST")
	r.HandleFunc("/student", stuh.UpdateStudent).Methods("PUT")
	r.HandleFunc("/student/{id}", stuh.DeleteStudent).Methods("DELETE")
	r.HandleFunc("/student/all", stuh.GetAllStudent).Methods("GET")
	r.HandleFunc("/student/{id}", stuh.GetStudent).Methods("GET")
	r.HandleFunc("/student", stuh.GetStudentByDetail).Methods("GET")

	//subject handlers
	r.HandleFunc("/subject/{id}", subh.GetSubject).Methods("GET")
	r.HandleFunc("/subject", subh.NewSubject).Methods("POST")
	r.HandleFunc("/subject", subh.GetAllSubject).Methods("GET")

	//enrolment handler
	r.HandleFunc("/student/{stuId}/subject/{subId}", stuh.EnrolStudentHandler).Methods("POST")
	r.HandleFunc("/student/{stuId}/subject", stuh.GetSubjectsByStuSvs).Methods("GET")

	err = http.ListenAndServe(":8080", r)

	//error handling
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
