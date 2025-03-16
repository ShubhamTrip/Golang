package main

import (
	"fmt"
)

func main() {
	studentMarks := make(map[string]int)

	studentMarks["Alice"] = 90

	studentMarks["Shubh"] = 80

	studentMarks["Shubham"] = 80

	studentMarks["Sumit"] = 80

	fmt.Println(studentMarks)

	fmt.Println(len(studentMarks))

	studentsGrades := map[string]int{
		"Shubham": 30,

		"Rohit": 50,
	}

	fmt.Println(studentsGrades)

}
