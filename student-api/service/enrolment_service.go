package service

import (
	"github.com/student-api/models"
)

type enrolmentDatastore interface {
	EnrolStudent(stuId, subId int) error
	GetAllSubjectByStudent(stuId int) ([]models.EnrolmentList, error)
}

type enrolmentService struct {
	d enrolmentDatastore
}

func NewEnrolmentService(d enrolmentDatastore) enrolmentService {
	return enrolmentService{d}
}
func (e enrolmentService) EnrolStudentSvs(stuId, subId int) error {
	err := e.d.EnrolStudent(stuId, subId)
	if err != nil {
		return err
	}
	return nil
}

func (e enrolmentService) GetSubjectsByStuSvs(stuId int) ([]models.EnrolmentList, error) {
	ans, err := e.d.GetAllSubjectByStudent(stuId)
	if err != nil {
		return ans, err
	}
	return ans, nil
}
