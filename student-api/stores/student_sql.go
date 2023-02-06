package stores

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/student-api/models"
)

type sqldb struct {
	db *sql.DB
}

func OpenDbConection() sqldb {
	con, _ := sql.Open("mysql", "root:pass@tcp(localhost:3306)/student")
	return sqldb{con}
}

func NewSqldb(db *sql.DB) sqldb {
	return sqldb{db} //
}

//FUNCTION TO INSERT
func (db sqldb) Insert(s models.Student) error {

	_, err := db.db.Exec(`insert into Student (RollNo,Name,Age) values (?,?,?)`, s.RollNo, s.Name, s.Age)
	if err != nil {
		return err
	}
	return nil
}

//TODO: FUNCTION TO UPDATE
func (db sqldb) Update(s models.Student) error {

	_, err := db.db.Exec(`UPDATE Student SET Name=?,Age=? where RollNo=?`, s.Name, s.Age, s.RollNo)
	if err != nil {
		return err
	}
	return nil
}

//TODO: FUNCTION TO DELETE
func (db sqldb) Delete(rollno int) error {
	_, err := db.db.Exec(`DELETE FROM Student WHERE RollNo=?`, rollno)
	if err != nil {
		return err
	}
	return nil
}

//TODO: FUNCTION TO READ
func (db sqldb) Read(rollno int) (models.Student, error) {
	s := models.Student{}

	res, err := db.db.Query(`SELECT * FROM Student WHERE RollNo=?`, rollno)
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
func (db sqldb) ReadAll() ([]models.Student, error) {
	st := []models.Student{}

	rows, err := db.db.Query(`SELECT * FROM Student`)
	if err != nil {
		return st, err
	}
	defer rows.Close()
	for rows.Next() {
		s := models.Student{}
		err = rows.Scan(&s.RollNo, &s.Name, &s.Age)
		if err != nil {
			return st, err
		}
		st = append(st, s)
	}
	return st, err
}
