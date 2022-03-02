package my

import "testing"

func TestPrincipal_AdmitApplicant(t *testing.T) {

	var scl = []struct{
       input []interface{}
	   expectedOutput []string
	}{
		{[]interface{"Cecilia", 20, "male", 10000, Classes{Classrooms: map[string]int{a.Class: +1}}}},
	}
}