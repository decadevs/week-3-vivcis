package my

import (
	"github.com/vivcis/my-SCL"
	"testing"
)

func TestTeachers_GradeCourse(t *testing.T) {
	testGrade := []struct {
		input my.Teachers

		expectedOutput string
	}{
		{my.Teachers{"Cecilia", 40, "female", "English"}, "course has been graded"},
		{my.Teachers{"Joseph", 66, "male", "Economics"}, "you are not to grade this course"},
	}

	for _, v := range testGrade {
		got := v.input.GradeCourses(v.input)
		if got != v.expectedOutput {
			t.Errorf("expected %v, got %v", v.expectedOutput, got)
		}
	}
}
