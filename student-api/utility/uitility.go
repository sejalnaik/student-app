package utility

import (
	"github.com/sejalnaik/student-app/model"
)

func ConvertStudentTimeToDate(student *model.Student) {
	student.DOB = student.DOB[:10]
	student.DOBTIME = student.DOBTIME[:19]
}

func ConvertStudentsTimeToDate(students *[]model.Student) {
	tempStudents := *students
	for i := 0; i < len(tempStudents); i++ {
		tempStudents[i].DOB = (tempStudents[i].DOB[:10])
		tempStudents[i].DOBTIME = (tempStudents[i].DOBTIME[:19])
	}
}
