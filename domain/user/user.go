package user

import "github.com/sankethkini/StudentData/domain/course"

type User struct {
	Fname   string
	Age     int
	Adress  string
	RollNo  string
	Courses []course.Course
}
