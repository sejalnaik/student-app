package utility

import (
	"github.com/sejalnaik/student-app/model"
)

func ConvertStudentTimeToDate(student *model.Student) {
	tempDOB := *student.DOB
	tempDOB = tempDOB[:10]
	student.DOB = &tempDOB

	tempDOBTIME := *student.DOBTIME
	tempDOBTIME = tempDOBTIME[:19]
	student.DOBTIME = &tempDOBTIME
}

func ConvertStudentsTimeToDate(students *[]model.Student) {
	tempStudents := *students
	for i := 0; i < len(tempStudents); i++ {
		tempDOB := *tempStudents[i].DOB
		tempDOB = tempDOB[:10]
		tempStudents[i].DOB = &tempDOB

		tempDOBTIME := *tempStudents[i].DOBTIME
		tempDOBTIME = tempDOBTIME[:19]
		tempStudents[i].DOBTIME = &tempDOBTIME
	}
}

/*func IfZeroRollNOConvertToNull(student *model.Student) {
	if student.RollNo == 0 {
		student.RollNo = null
	}
}*/
