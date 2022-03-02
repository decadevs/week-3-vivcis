package main

import (
	"fmt"
	scl "github.com/vivcis/my-SCL"
)

func main() {

	a := scl.Applicants{
		Name:   "Olanma",
		Age:    16,
		Gender: "female",
		Class:  "Jss2",
	}

	p := scl.Principal{}
	p.AdmitApplicant(a)
	//fmt.Println(scl.StudentDB)
	//s := scl.Students{}
	//s.TakeACourse("Olanma", "Chemistry")
	//fmt.Println(scl.StudentDB)

	t := scl.Teachers{}
	t.GradeCourses("Franklyn Omonade", "Chisom", "Chemistry", 60)
	fmt.Println(scl.StudentDB)
	//p.ExpelStudent("Olanma")
	//fmt.Println(scl.StudentDB)
	//p.SackTeacher("Franklyn Omonade")
	//fmt.Println(scl.TeacherDB)
}
