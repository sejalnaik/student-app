package utility

import "github.com/sejalnaik/student-app/model"

func ConvertStudentTimeToDate(student *model.Student) {
	if student.DOB != nil {
		tempDOB := *student.DOB
		tempDOB = tempDOB[:10]
		student.DOB = &tempDOB
	}

	if student.DOBTIME != nil {
		tempDOBTIME := *student.DOBTIME
		tempDOBTIME = tempDOBTIME[:19]
		student.DOBTIME = &tempDOBTIME
	}
}

func ConvertStudentsTimeToDate(students *[]model.Student) {
	tempStudents := *students

	for i := 0; i < len(tempStudents); i++ {
		if tempStudents[i].DOB != nil {
			tempDOB := *tempStudents[i].DOB
			tempDOB = tempDOB[:10]
			tempStudents[i].DOB = &tempDOB
		}
	}

	for i := 0; i < len(tempStudents); i++ {
		if tempStudents[i].DOBTIME != nil {
			tempDOBTIME := *tempStudents[i].DOBTIME
			tempDOBTIME = tempDOBTIME[:19]
			tempStudents[i].DOBTIME = &tempDOBTIME
		}
	}
}
