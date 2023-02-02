package dbcrud

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Student struct {
	RollNo int
	Name   string
	Age    float64
}
type newDB struct {
	*sql.DB
}

func OpenDbConection() newDB {
	con, _ := sql.Open("mysql", "root:pass@tcp(localhost:3306)/student")
	return newDB{con}
}

//todo: CRUD

//TODO: FUNCTION TO INSERT
func insertRecord(db *newDB, s Student) error {

	_, err := db.Exec(`insert into Student (RollNo,Name,Age) values (?,?,?)`, s.RollNo, s.Name, s.Age)
	if err != nil {
		return err
	}
	return nil
}

//TODO: FUNCTION TO UPDATE
func updateRecord(db *newDB, s Student) error {

	_, err := db.Exec(`UPDATE Student SET Name=?,Age=? where RollNo=?`, s.Name, s.Age, s.RollNo)
	if err != nil {
		return err
	}
	return nil
}

//TODO: FUNCTION TO DELETE
func deleteRecord(db *newDB, rollno int) error {
	_, err := db.Exec(`DELETE FROM Student WHERE RollNo=?`, rollno)
	if err != nil {
		return err
	}
	return nil
}

//TODO: FUNCTION TO READ
func readRecord(db *newDB, rollno int) (Student, error) {
	s := Student{}

	res, err := db.Query(`SELECT * FROM Student WHERE RollNo=?`, rollno)
	if err != nil {
		return s, err
	}
	defer res.Close()
	for res.Next() {
		err = res.Scan(&s.RollNo, &s.Name, &s.Age)
	}
	return s, err

}

//FUNCTION TO READALL DATABASE
func readAllRecord(db *newDB) ([]Student, error) {
	st := []Student{}

	rows, err := db.Query(`SELECT * FROM Student`)
	if err != nil {
		return st, err
	}
	defer rows.Close()
	for rows.Next() {
		s := Student{}
		err = rows.Scan(&s.RollNo, &s.Name, &s.Age)
		if err != nil {
			return st, err
		}
		st = append(st, s)
	}
	return st, err
}

//todo: http
//TODO: FUNCTION TO INSERT  http
func (db *newDB) insertRecordHandle(w http.ResponseWriter, r *http.Request) {
	s := Student{}

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = insertRecord(db, s)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
}

//TODO: FUNCTION TO UPDATE http
func (db *newDB) updateRecordHandle(w http.ResponseWriter, r *http.Request) {
	s := Student{}
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = updateRecord(db, s)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
}

//TODO: FUNCTION TO DELETE http
func (db *newDB) deleteRecordHandle(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	rollno, err := strconv.Atoi(string(id))
	if err != nil {
		println(err.Error())
	}

	err = deleteRecord(db, rollno)
	if err != nil {
		println(err.Error())

	}
	fmt.Fprintf(w, err.Error())
}

//TODO: FUNCTION TO READ http
func (db *newDB) readRecordHandle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	rollno, err := strconv.Atoi(string(id))
	if err != nil {
		println(err.Error())
	}

	ans, err := readRecord(db, rollno)
	if err != nil {
		println(err.Error())
	}
	fmt.Fprintf(w, "%v", ans)
}
func (db *newDB) readAllRecordHandle(w http.ResponseWriter, r *http.Request) {

	ans, err := readAllRecord(db)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("read error reason: %s", err))
		return
	}
	fmt.Fprintf(w, "%v", ans)
}

func Server() {
	dbobj := OpenDbConection()
	r := mux.NewRouter()
	r.HandleFunc("/student", dbobj.insertRecordHandle).Methods("POST")
	r.HandleFunc("/student", dbobj.readAllRecordHandle).Methods("GET")
	r.HandleFunc("/student", dbobj.updateRecordHandle).Methods("PUT")
	r.HandleFunc("/student/{id}", dbobj.readRecordHandle).Methods("GET")
	// r.HandleFunc("/student/{id}", dbobj.deleteRecordHandle).Methods("dELEY")

	err := http.ListenAndServe(":8080", r)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
func main() {
	Server()
}
