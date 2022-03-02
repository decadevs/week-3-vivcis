package my

import "log"

var ApplicantDB []Applicants

type Principal struct {
	Name   string
	Age    int
	Gender string
}

type Applicants struct {
	Name   string
	Age    int
	Gender string
	Class  string
}

//func init() {
//	app := []Applicants{
//		{Name: "Frank", Age: 26, Gender: "female", Class: "JSS3"},
//		{Name: "David", Age: 10, Gender: "male", Class: "Jss1"},
//		{Name: "Loveth", Age: 20, Gender: "female", Class: "ss1"},
//		{Name: "Ogo", Age: 15, Gender: "female", Class: "Jss2"},
//		{Name: "Ebube", Age: 26, Gender: "female", Class: "SS2"},
//	}
//	for _, v := range app {
//		ApplicantDB = append(ApplicantDB, v)
//	}
//	log.Println(ApplicantDB)
//}

func (p *Principal) AdmitApplicant(a Applicants) string {
	if a.Age < 16 {
		log.Println("not old enough to be admitted")
	} else {
		student := Students{
			Name:    a.Name,
			Age:     a.Age,
			Gender:  a.Gender,
			Fees:    2000,
			Classes: Classes{Classrooms: map[string]int{a.Class: +1}},
		}
		StudentDB = append(StudentDB, student)
		log.Println("welcome to cece's high school")

	}
	return "welcome to cece's high school"
}

//expelled for poor grades'
//student is warned and advised to put in more efforts
//good student to keep it up

func (p Principal) ExpelStudent(studentName string) {
	for i, v := range StudentDB {
		if studentName == v.Name {
			StudentDB = append(StudentDB[:i], StudentDB[i+1:]...)
			log.Printf("%s has been expeled from our school", studentName)
		}
	}
}

func (p Principal) SackTeacher(teacherName string) {
	for i, v := range TeacherDB {
		if teacherName == v.Name {
			TeacherDB = append(TeacherDB[:i], TeacherDB[i+1:]...)
			log.Printf("%s has been sacked from this HONOURABLE institution", teacherName)
		}
	}

}

func (p Principal) EmployTeacher(newTeacherName string) {
	for _, v := range TeacherDB {
		if newTeacherName == v.Name {
			TeacherDB = append(TeacherDB, v)
			log.Printf("%s welcome to our school", newTeacherName)
		}
	}
}
