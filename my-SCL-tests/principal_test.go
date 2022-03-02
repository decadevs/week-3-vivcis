package my

import (
	"github.com/vivcis/my-SCL"
	"testing"
)

func TestPrincipal_AdmitApplicant(t *testing.T) {
	//s:= my.Students{
	//	Courses: my.Courses{map[string]int{"Chemistry": 54}},
	//	Classes: my.Classes{map[string]int{"JSS 1B": 32}},
	//	Name:    "Joe B",
	//	Age:     14,
	//	Gender:  "Male",
	//	Fees:    0,
	//}
	//
	//c:= my.Courses{CoursesAndGrades: map[string]int{"Chemistry": 54}}

	var admitApplicant = []struct {
		input  my.Principal
		input1 my.Applicants

		expectedOutput string
	}{
		{my.Principal{Name: "Okon", Age: 43, Gender: "Male"}, my.Applicants{"Alexis", 18, "Female", "JSS 2A"}, "welcome to cece's high school"},
		{my.Principal{Name: "Judith", Age: 18, Gender: "Female"}, my.Applicants{"Joe B", 14, "Male", "JSS 1C"}, "not old enough to be admitted"},
	}
	for _, val := range admitApplicant {
		admitted := val.input.AdmitApplicant(val.input1)
		if admitted != val.expectedOutput {
			t.Errorf("expected output to be %v but got output %v", val.expectedOutput, admitted)
		}
	}
}

func TestPrincipal_ExpelStudent(t *testing.T) {
	var expelStudent = []struct {
		input  my.Principal
		input1 my.Students

		expectedOutput string
	}{
		{my.Principal{}, my.Students{
			Courses: my.Courses{CoursesAndGrades: map[string]int{"Physics": 40}},
			Classes: my.Classes{Classrooms: map[string]int{"JSS2": 20}},
			Name:    "Victor",
			Age:     50,
			Gender:  "male",
			Fees:    0,
		}, "you have been expelled from this school"},
	}

	for _, val := range expelStudent {
		expelled := val.input.ExpelStudent(val.input1)
		if expelled != val.expectedOutput {
			t.Errorf("expected output to be %v but got output %v", val.expectedOutput, expelled)
		}
	}
}

//func TestPrincipal_SackTeacher(t *testing.T) {
//	var sackTeacher = []struct {
//		input  my.Principal
//		input1 my.Teachers
//
//		expectedOutput string
//	}{
//		{my.Principal{}, my.Teachers{
//			Name:    "Micheal",
//			Age:     60,
//			Gender:  "male",
//			Subject: "Yoruba",
//		}, "you have been sacked from this HONOURABLE institution"},
//	}
//
//	for _, v := range sackTeacher {
//		sacked := v.input.SackTeacher(v.input1)
//		if sacked != v.expectedOutput {
//			t.Errorf("expected output to be %v but got output %v", v.expectedOutput, sacked)
//		}
//	}
//}
