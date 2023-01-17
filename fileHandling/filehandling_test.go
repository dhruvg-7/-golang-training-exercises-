package fileHandle

import (
	"encoding/json"
	"os"
	"testing"
)

func checkFileExist(fileName string) bool {
	f, err := os.Open(fileName)
	if err != nil {
		// panic(err)
		println(err)
		return false
	}

	defer f.Close()
	return true
}

func CheckRecordsFromJson(fileName string) []Student {
	student := []Student{}

	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	json.NewDecoder(f).Decode(&student)
	return student

}

// checking weather record is under creteria or not?
func lessThanComparative(student []Student, age float64) bool {

	for i := range student {
		if student[i].Age > age {
			return false
		}
	}
	return true

}
func greaterThanComparative(student []Student, age float64) bool {

	for i := range student {
		if student[i].Age < age {
			return false
		}
	}
	return true
}

func TestReadFromJsonBaseCase(t *testing.T) {

	fileName := "Student.json"
	fileName2 := "primaryStudent.json"
	fileName3 := "secondaryStudent.json"
	student := []Student{
		{1, "Dhruv", 12, []string{"9890000010", "9999999999"}},
		{2, "Ratri", 10, []string{"9890000010", "9999999999"}},
		{3, "dash", 23, []string{"9890000010", "9999999999"}},
		{4, "dsaa", 9, []string{"9890000010", "9999999999"}},
		{5, "fwwfi", 16, []string{"9890000010", "9999999999"}},
		{6, "fdsfi", 14, []string{"9890000010", "9999999999"}},
		{7, "Ratri", 12, []string{"9890000010", "9999999999"}},
		{8, "Rat", 8, []string{"9890000010", "9999999999"}},
		{9, "MAt", 4, []string{"9890000010", "9999999999"}},
		{10, "May", 6, []string{"9890000010", "9999999999"}},
	}
	status := ExportFile(student, fileName)
	if !checkFileExist(fileName) && status {
		t.Errorf("file not found")
	}
	if status {
		status = ReadFromJson(fileName)
	}
	var result1, result2 bool
	if status {
		result1 = checkFileExist(fileName2)
		result2 = checkFileExist(fileName3)
	}
	if result1 && result2 {
		records1 := CheckRecordsFromJson(fileName2)
		lessThanComparative(records1, 10)
		records2 := CheckRecordsFromJson(fileName3)
		greaterThanComparative(records2, 10)
	}

}
