package service

import (
	"fmt"
	"reflect"

	"github.com/student-api/models"
)

type subjectDatastore interface {
	InsertSubject(models.Subject) error
	GetSubject(int) (models.Subject, error)
	GetAllSubject() ([]models.Subject, error)
}

type subjectService struct {
	d subjectDatastore
}

func NewSubjectService(d subjectDatastore) subjectService {
	return subjectService{d}
}

//SERVICE FUNCTION TO ADD A NEW SUBJECT
func (sh subjectService) InsertSubjectService(subject models.Subject) error {

	if subject.Id == 0 {
		return fmt.Errorf("Invalid id error")
	} else if reflect.DeepEqual(subject.Name, "") || reflect.DeepEqual(subject.Name, " ") {
		return fmt.Errorf("Invalid name error")
	}
	err := sh.d.InsertSubject(subject)
	if err != nil {
		return err
	}
	return nil
}

//SERVICE FUNCTION TO CHECK WEATHER GIVEN SUBJECT ID IS VALID OR NOT
func (sh subjectService) GetSubjectService(id int) (models.Subject, error) {
	empty := models.Subject{}
	sub := models.Subject{}
	if id == 0 {
		return sub, fmt.Errorf("Invalid id error")
	}
	sub, err := sh.d.GetSubject(id)
	if err != nil {
		return sub, err
	}
	if sub == empty {
		return sub, fmt.Errorf("Subject Record doesn't exist")
	}
	return sub, nil

}

func (sh subjectService) GetAllSubjectService() ([]models.Subject, error) {

	sub, err := sh.d.GetAllSubject()
	if err != nil {
		return sub, err
	}
	return sub, nil
}

// TODO : CROSS CHECK FUNCTION NAME IN SUBJECT HANDLER INTERFACE
