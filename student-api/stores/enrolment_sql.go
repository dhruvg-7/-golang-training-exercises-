package stores

import (
	"database/sql"
	"fmt"

	"github.com/student-api/models"
	"github.com/student-api/myerrors"
)

type enrolmentStore struct {
	db *sql.DB
}

func NewEnrolmentStore(db *sql.DB) enrolmentStore {
	return enrolmentStore{db}
}

func (db enrolmentStore) EnrolStudent(stuId, subId int) error {
	_, err := db.db.Exec(`INSERT INTO EnrolmentList (StudentId,SubjectId) values (?,?)`, stuId, subId)
	if err != nil {
		return err
	}
	return nil
}

func (db enrolmentStore) GetAllSubjectByStudent(stuId int) ([]models.EnrolmentList, error) {

	ans := []models.EnrolmentList{}
	row, err := db.db.Query(`SELECT * FROM EnrolmentList where StudentId =?`, stuId)
	if err != nil {
		return ans, fmt.Errorf(myerrors.NewError("sql_r"), err)
	}
	for row.Next() {
		list := models.EnrolmentList{}
		err = row.Scan(&list.StudentId, &list.SubjectId)
		if err != nil {
			return nil, fmt.Errorf(myerrors.NewError("sql_r"), err)
		}
		ans = append(ans, list)
	}
	return ans, nil

}
