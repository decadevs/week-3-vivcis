package my

import (
	"fmt"
	"log"
)

type Teachers struct {
	Name    string
	Age     int
	Gender  string
	Subject string
}

var TeacherDB []Teachers

func init() {
	t := []Teachers{
		{Name: "Franklyn Omonade", Gender: "male", Age: 76, Subject: "Chemistry"},
		{Name: "Cecilia", Gender: "female", Age: 88, Subject: "yoruba"},
		{Name: "Clinton", Gender: "male", Age: 32, Subject: "physics"},
		{Name: "Lovey", Gender: "female", Age: 23, Subject: "french"},
	}
	for _, v := range t {
		TeacherDB = append(TeacherDB, v)
		fmt.Printf("")
	}
	log.Println(TeacherDB)
}

func (t Teachers) GradeCourses(teachersName string, StudentsName string, course string, grade int) {

	for _, x := range TeacherDB {
		if x.Name == teachersName {
			if x.Subject == course {

				for i, v := range StudentDB {
					if v.Name == StudentsName {
						if _, found := v.CoursesAndGrades[course]; found {
							v.CoursesAndGrades[course] = grade
							StudentDB[i] = v
							log.Println("you have successfully been graded")
						}
					}
				}

			}
		}
	}
	//log.Println("you are not to grade this course")
}

//func (t Teachers) GetAverageGrades(name string, course string, grade int) float32 {
//	sum := 0
//	for _, v := range StudentDB {
//		if v.Name == name {
//			if _, found := v.CoursesAndGrades[course]; found {
//				v.CoursesAndGrades[course] = grade
//				sum += v
//			}
//		}
//	}
//	return float32(sum) / float32(len(v.CoursesAndGrades[course]))
//}
//
//func (s Students) GetMaxGrades() int {
//	currentMax := 0
//
//	for _, v := range StudentDB {
//		if currentMax < v {
//			currentMax = v
//		}
//	}
//	return currentMax
//}
