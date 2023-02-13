package service

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/student-api/models"
)

type studentDatastore interface {
	Insert(models.Student) error
	Read(int) (models.Student, error)
	ReadByName(string) ([]models.Student, error)
	ReadAll() ([]models.Student, error)
	Update(models.Student) error
	Delete(int) error
}

type studentService struct {
	d      studentDatastore
	subSvc subjectService
	enrSvc enrolmentService
}

func NewStudentService(d studentDatastore, subSvc subjectService, enrSvc enrolmentService) studentService {
	return studentService{
		d,
		subSvc,
		enrSvc,
	}
}

func (stuSvc studentService) InsertService(student models.Student) error {
	if reflect.DeepEqual(student.Name, "") || reflect.DeepEqual(student.Name, " ") {
		return fmt.Errorf("Null Name error")
	}
	if student.RollNo == 0 {
		return fmt.Errorf("Invalid Roll Number error")
	}
	err := stuSvc.d.Insert(student)
	if err != nil {
		return err
	}
	return nil
}

func (stuSvc studentService) ReadByIdService(id int) (models.Student, error) {
	st := models.Student{}
	empty := models.Student{}
	if id == 0 {
		return st, fmt.Errorf("Invalid Roll Number error")
	}
	st, err := stuSvc.d.Read(id)
	if st == empty {
		return st, fmt.Errorf("student Recored doesn't exist")
	}
	if err != nil {
		return st, err
	}

	return st, nil
}

func (stuSvc studentService) Update(st models.Student) error {
	empty := models.Student{}
	stuId := st.RollNo
	res, err := stuSvc.ReadByIdService(stuId)
	if res == empty {
		return fmt.Errorf("Record doesn't exist, GOT: %v", err)
	}

	err = stuSvc.d.Update(st)
	if err != nil {
		return err
	}
	return nil
}

func (stuSvc studentService) Delete(Id string) error {
	stuId, err := strconv.Atoi(Id)
	if err != nil {
		return fmt.Errorf("At Delete service, Parse error, Got: %v ", err)
	}
	if _, err := stuSvc.ReadByIdService(stuId); err != nil {
		return fmt.Errorf("Record Error, GOT: %v", err)
	}

	err = stuSvc.d.Delete(stuId)
	if err != nil {
		return fmt.Errorf("At Delete service, Store error, Got: %v ", err)
	}
	return nil
}

func (stuSvc studentService) ReadByDetailService(rollno, name string) ([]models.Student, error) {
	st := []models.Student{}
	fmt.Println(len(rollno))

	if len(rollno) > 0 {
		r, err := strconv.Atoi(rollno)
		if err != nil {
			return nil, err
		}
		if r > 0 {
			fmt.Println(r)
			ans, err := stuSvc.d.Read(r)
			if err != nil {
				return st, err
			}
			st = append(st, ans)
		}

	} else if len(name) > 0 {
		st, err := stuSvc.d.ReadByName(name)
		if err != nil {
			return st, err
		}
		return st, err
	}

	return st, nil
}

func (stuSvc studentService) ReadAllService() ([]models.Student, error) {
	ans, err := stuSvc.d.ReadAll()
	if err != nil {
		return nil, err
	}
	return ans, err
}

func (stuSvc studentService) EnrolStudentSvs(stuId, subId string) error {
	rollno, err := strconv.Atoi(stuId)
	if err != nil {
		println(err.Error())
	}
	if _, err := stuSvc.ReadByIdService(rollno); err != nil {
		return err
	}
	subjId, err := strconv.Atoi(subId)
	if err != nil {
		return err
	}
	if _, err := stuSvc.subSvc.GetSubjectService(subjId); err != nil {
		return err
	}
	err = stuSvc.enrSvc.EnrolStudentSvs(rollno, subjId)
	if err != nil {
		return err
	}
	return nil
}

func (stuSvc studentService) GetstudentBySubSvs(stuId string) ([]models.EnrolmentList, error) {
	id, err := strconv.Atoi(stuId)
	if err != nil {
		return []models.EnrolmentList{}, err
	}
	ans, err := stuSvc.enrSvc.GetSubjectsByStuSvs(id)
	return ans, err
}
