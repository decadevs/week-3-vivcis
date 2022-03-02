package my

import (
	"github.com/vivcis/my-SCL"
	"testing"
)

func TestStudents_RegisterACourse(t *testing.T) {
	courses := []string{"Physics", "Chemistry", "Biology", "Economics", "English", "Maths", "Agric", "Yoruba", "Literature", "Commerce"}

	var registerACourse = []struct {
		input  my.Students
		input1 my.Courses

		expectedOutput string
	}{
		{my.Students{}, my.Courses{}, "you have successfully registered for a course"},
	}
	for _, v := range registerACourse {
		registered := v.input.TakeACourse(registerACourse.)
		if registered != v.expectedOutput{
			t.Errorf("expected %v, got %v", v.expectedOutput, registered)
		}
	}
 }

