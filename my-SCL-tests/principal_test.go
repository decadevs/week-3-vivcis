package my

import (
	"github.com/vivcis/my-SCL"
	"testing"
)

func TestPrincipal_AdmitApplicant(t *testing.T) {

	var admitApplicant = []struct {
		input  my.Principal
		input1 string
		input2 int
		input3 string
		input4 string

		expectedOutput string
	}{
		{my.Principal{}, "Alexis", 18, "female", "JSS2", "welcome to cece's high school"},
		{my.Principal{}, "Asuquo", 14, "male", "JSS1", "not old enough to be admitted"},
	}
	for _, val := range admitApplicant {
		admitted := val.input.AdmitApplicant(val.input1, val.input4)
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

func TestPrincipal_SackTeacher(t *testing.T) {
	var sackTeacher = []struct {
		input  my.Principal
		input1 my.Teachers

		expectedOutput string
	}{
		{my.Principal{}, my.Teachers{
			Name:    "Micheal",
			Age:     60,
			Gender:  "male",
			Subject: "Yoruba",
		}, "you have been sacked from this HONOURABLE institution"},
	}

	for _, v := range sackTeacher {
		sacked := v.input.SackTeacher(v.input1)
		if sacked != v.expectedOutput {
			t.Errorf("expected output to be %v but got output %v", v.expectedOutput, sacked)
		}
	}
}
