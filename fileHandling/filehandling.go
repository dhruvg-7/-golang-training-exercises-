package fileHandle

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

type Student struct {
	Roll  int
	Name  string
	Age   float64
	Phone []string
}

func ReadFromJson(fileName string) bool {
	student := []Student{}

	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	json.NewDecoder(f).Decode(&student)
	status := filterRecords(student)
	return status

}

func filterRecords(student []Student) bool {

	primaryStudent := []Student{}
	secondaryStudent := []Student{}

	for i := range student {
		if student[i].Age < 11 {
			primaryStudent = append(primaryStudent, student[i])
		} else {
			secondaryStudent = append(secondaryStudent, student[i])
		}
	}
	status1 := ExportFile(secondaryStudent, "secondaryStudent.json")
	status2 := ExportFile(primaryStudent, "PrimaryStudent.json")
	return status1 && status2

}
func ExportFile(student []Student, fileName string) bool {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(student)

	f, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, fileErr := io.Copy(f, buf)
	if fileErr != nil {
		println(fileErr)
		return false
	}
	return true

}
