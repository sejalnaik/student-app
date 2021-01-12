package utility

import "github.com/sejalnaik/student-app/model"

func AddStudentEmptyStringToNull(student *model.Student) {
	if student.DOB != nil {
		if *student.DOB == "" {
			student.DOB = nil
		}
	}
	if student.DOBTIME != nil {
		if *student.DOBTIME == "" {
			student.DOBTIME = nil
		}
	}
	if student.PhoneNumber != nil {
		if *student.PhoneNumber == "" {
			student.PhoneNumber = nil
		}
	}
}
