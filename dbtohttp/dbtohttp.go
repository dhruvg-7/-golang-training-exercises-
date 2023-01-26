package dbtohttp

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	Name  string
	Age   float64
	Phone string
}

func readDB(myDb *sql.DB) ([]Person, error) {

	rows, err := myDb.Query("SELECT * FROM person")
	if err != nil {
		return nil, err
	}
	var person []Person
	for rows.Next() {
		var p Person
		err = rows.Scan(&p.Name, &p.Age, &p.Phone)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		person = append(person, p)

	}
	return person, nil
}

func personhandler() []Person {
	db, err := sql.Open("mysql", "root:pass@tcp(localhost:3306)/Person")
	if err != nil {
		fmt.Println("can't make connection to the database")
	}
	defer db.Close()

	list, err2 := readDB(db)

	if err2 != nil {
		fmt.Printf("%v", err2)
	}
	return list

}
func rootHandler(w http.ResponseWriter, req *http.Request) {

	if req.URL.Path == ("/") {

		fmt.Fprintln(w, "404")
	} else if req.URL.Path == "/ping" {

		fmt.Fprintln(w, "Pong")
	} else if req.URL.Path == "/person" {
		p := personhandler()
		fmt.Fprintln(w, p)
	} else {
		fmt.Fprintln(w, req.URL.Path[1:])
	}

}
