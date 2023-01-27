package dbtocsv

import (
	"database/sql"
	"fmt"
	"io"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	Name   string
	Age    float64
	Rollno int
}

func readDB(myDb *sql.DB) ([]student, error) {

	rows, err := myDb.Query("SELECT * FROM student")
	if err != nil {
		return nil, err
	}
	var st []student
	for rows.Next() {
		var s student
		err = rows.Scan(&s.Rollno, &s.Name, &s.Age)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		st = append(st, s)

	}
	return st, nil
}

func writeio(st []student, file io.Writer) (io.Writer, error) {
	// file, err := os.Create("./" + filename)

	// if err != nil {
	// 	panic(err)
	// }
	for i := range st {
		line := fmt.Sprintf("%v,%v,%v", st[i].Rollno, st[i].Name, st[i].Age)
		line += fmt.Sprintln()
		_, err := io.WriteString(file, line)
		if err != nil {
			return file, err
		}
	}
	return file, nil
}
