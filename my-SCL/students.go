package my

import "log"

type Students struct {
	Courses
	Classes
	Name   string
	Age    int
	Gender string
	Fees   int
}

type Classes struct {
	Classrooms map[string]int
}

type Courses struct {
	CoursesAndGrades map[string]int
}

var StudentDB []Students

func init() {
	st := []Students{
		{Name: "Chisom", Age: 14, Gender: "female", Fees: 10750, Classes: Classes{Classrooms: map[string]int{"JSS1": +1}},
			Courses: Courses{CoursesAndGrades: map[string]int{"Physics": 0, "Chemistry": 0, "Economics": 0, "Biology": 0, "Maths": 0, "English": 0,
				"Agric": 0}}},
		//{Name: "Timi", Age: 13, Gender: "male", Fees: 5000, Classes: Classes{Classrooms: map[string]int{"JSS2": +1}},
		//	Courses: Courses{CoursesAndGrades: map[string]int{"Physics": 0, "Chemistry": 0, "Economics": 0, "Biology": 0, "Maths": 0, "English": 0,
		//		"Agric": 0}}},
		//{Name: "Isioma", Age: 15, Gender: "female", Fees: 10750, Classes: Classes{Classrooms: map[string]int{"JSS3": +1}},
		//	Courses: Courses{CoursesAndGrades: map[string]int{"Physics": 0, "Chemistry": 0, "Economics": 0, "Biology": 0, "Maths": 0, "English": 0,
		//		"Agric": 0}}},
		//{Name: "Asuquo", Age: 30, Gender: "male", Fees: 9000, Classes: Classes{Classrooms: map[string]int{"SSS1": +1}},
		//	Courses: Courses{CoursesAndGrades: map[string]int{"Physics": 0, "Chemistry": 0, "Economics": 0, "Biology": 0, "Maths": 0, "English": 0,
		//		"Agric": 0}}},
		//{Name: "Ebube", Age: 15, Gender: "female", Fees: 9500, Classes: Classes{Classrooms: map[string]int{"SSS2": +1}},
		//	Courses: Courses{CoursesAndGrades: map[string]int{"Physics": 0, "Chemistry": 0, "Economics": 0, "Biology": 0, "Maths": 0, "English": 0,
		//		"Agric": 0}}},
		//{Name: "Chukwuma", Age: 50, Gender: "female", Fees: 0, Classes: Classes{Classrooms: map[string]int{"SSS3": +1}},
		//	Courses: Courses{CoursesAndGrades: map[string]int{"Physics": 0, "Chemistry": 0, "Economics": 0, "Biology": 0, "Maths": 0, "English": 0,
		//		"Agric": 0}}},
	}
	for _, v := range st {
		StudentDB = append(StudentDB, v)
	}
	log.Println(StudentDB)
}

var listOfCourses = []string{"Physics", "Chemistry", "Biology", "Economics", "English", "Maths", "Agric", "Yoruba", "Literature", "Commerce"}

func (s Students) TakeACourse(name string, course string) string {
	for i, v := range StudentDB {
		if v.Name == name {
			for _, x := range listOfCourses {
				if x == course {
					v.CoursesAndGrades = map[string]int{course: 0}
					StudentDB[i] = v
					log.Println("you have successfully registered for a course")
				}
			}
		}
	}
	return "you have successfully registered for a course"
}
