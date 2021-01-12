package utility

import (
	"github.com/sejalnaik/student-app/model"
)

func ConvertStructStudentToMap(student *model.Student) map[string]interface{} {
	studentMap := make(map[string]interface{})

	//rollno
	if student.RollNo == nil {
		studentMap["RollNo"] = nil
	} else {
		studentMap["RollNo"] = *student.RollNo
	}

	//name
	studentMap["Name"] = student.Name

	//age
	if student.Age == nil {
		studentMap["Age"] = nil
	} else {
		studentMap["Age"] = *student.Age
	}

	//email
	studentMap["Email"] = student.Email

	//isMale
	if student.IsMale == nil {
		studentMap["IsMale"] = nil
	} else {
		studentMap["IsMale"] = *student.IsMale
	}

	//dob
	if student.DOB == nil || *student.DOB == "" {
		studentMap["DOB"] = nil
	} else {
		studentMap["DOB"] = *student.DOB
	}

	//dobTime
	if student.DOBTIME == nil || *student.DOBTIME == "" {
		studentMap["DOBTIME"] = nil
	} else {
		studentMap["DOBTIME"] = *student.DOBTIME
	}

	//phoneNumber
	if student.PhoneNumber == nil || *student.PhoneNumber == "" {
		studentMap["PhoneNumber"] = nil
	} else {
		studentMap["PhoneNumber"] = *student.PhoneNumber
	}

	return studentMap
}

func ConvertStructBookIssueToMap(bookIssue *model.BookIssue) map[string]interface{} {
	bookIssueMap := make(map[string]interface{})

	//penalty
	bookIssueMap["Penalty"] = bookIssue.Penalty

	//Returned
	bookIssueMap["Returned"] = bookIssue.Returned

	return bookIssueMap
}
