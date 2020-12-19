package utility

import (
	"github.com/sejalnaik/student-app/model"
)

func ConvertStudentTimeToDate(student *model.Student) {
	student.DOB = student.DOB[:10]
}

func ConvertStudentsTimeToDate(students *[]model.Student) {
	tempStudents := *students
	for i := 0; i < len(tempStudents); i++ {
		tempStudents[i].DOB = (tempStudents[i].DOB[:10])
	}
}
