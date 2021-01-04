package utility

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"

	"github.com/sejalnaik/student-app/model"
)

const key string = "the-key-has-to-be-32-bytes-long!"

func ConvertStudentTimeToDate(student *model.Student) {
	if student.DOB != nil {
		tempDOB := *student.DOB
		tempDOB = tempDOB[:10]
		student.DOB = &tempDOB
	}

	tempStudent := *student
	for i := 0; i < len(tempStudent.BookIssues); i++ {
		tempIssueDate := tempStudent.BookIssues[i].IssueDate
		tempIssueDate = tempIssueDate[:19]
		student.BookIssues[i].IssueDate = tempIssueDate
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

func EncryptUserPassword(password []byte) ([]byte, error) {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, password, nil), nil
}

func DecryptUserPassword(password []byte) ([]byte, error) {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(password) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := password[:nonceSize], password[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
