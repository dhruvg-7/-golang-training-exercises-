package models

type Subject struct {
	Id   int
	Name string
}

type EnrolmentList struct {
	StudentId int
	SubjectId int
}
