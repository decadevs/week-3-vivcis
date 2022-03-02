package my

import (
	"github.com/vivcis/my-SCL"
	"testing"
)

func TestStudents_TakeACourse(t *testing.T) {
	//courses := []string{"Physics", "Chemistry", "Biology", "Economics", "English", "Maths", "Agric", "Yoruba", "Literature", "Commerce"}

	var registerACourse = []struct {
		input  my.Students
		input1 string
		input2 string

		expectedOutput string
	}{
		{my.Students{
			Courses: my.Courses{map[string]int{"Chemistry": 80}},
			Classes: my.Classes{map[string]int{"JSS 1": 20}},
			Name:    "Gbenle",
			Age:     49,
			Gender:  "Shemale",
			Fees:    67850,
		}, "Gbenle", "Chemistry", "you have successfully registered for a course"},
	}
	for _, v := range registerACourse {
		registered := v.input.TakeACourse(v.input1, v.input2)
		if registered != v.expectedOutput {
			t.Errorf("expected %v, got %v", v.expectedOutput, registered)
		}
	}
}
