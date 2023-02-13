package stores

import (
	"database/sql"

	"github.com/student-api/models"
)

type subjectStore struct {
	db *sql.DB
}

func NewSubjectStore(db *sql.DB) subjectStore {
	return subjectStore{db}
}

//FUNCTION TO INSERT

func (db subjectStore) InsertSubject(s models.Subject) error {

	_, err := db.db.Exec(`insert into subject (Id,Name) values (?,?)`, s.Id, s.Name)
	if err != nil {
		return err
	}
	return nil
}

//TODO: FUNCTION TO UPDATE
func (db subjectStore) UpdateSubject(s models.Subject) error {

	_, err := db.db.Exec(`UPDATE subject SET Name=? where Id= ?`, s.Name, s.Id)
	if err != nil {
		return err
	}
	return nil
}

//TODO: FUNCTION TO DELETE
func (db subjectStore) DeleteSubject(Id int) error {
	_, err := db.db.Exec(`DELETE FROM subject WHERE Id=?`, Id)
	if err != nil {
		return err
	}
	return nil
}

//TODO: FUNCTION TO READ
func (db subjectStore) GetSubject(Id int) (models.Subject, error) {
	s := models.Subject{}

	res, err := db.db.Query(`SELECT * FROM subject WHERE Id=?`, Id)
	if err != nil {
		return s, err
	}
	defer res.Close()
	for res.Next() {
		err = res.Scan(&s.Id, &s.Name)
	}
	return s, err

}

func (db subjectStore) GetAllSubject() ([]models.Subject, error) {
	sub := []models.Subject{}

	rows, err := db.db.Query(`SELECT * FROM subject`)
	if err != nil {
		return sub, err
	}
	defer rows.Close()
	for rows.Next() {
		s := models.Subject{}
		err = rows.Scan(&s.Id, &s.Name)
		if err != nil {
			return sub, err
		}
		sub = append(sub, s)
	}
	return sub, err

}
