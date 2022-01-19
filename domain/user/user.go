package user

import "github.com/sankethkini/StudentData/domain/course"

type User struct {
	Fname   string
	Age     int
	Adress  string
	RollNo  string
	Courses []course.Course
}

func NewUser(Fname string, Age int, Adress string, RollNo string, Courses []course.Course) User {
	curUser := User{}
	curUser.Fname = Fname
	curUser.RollNo = RollNo
	curUser.Age = Age
	curUser.Adress = Adress
	curUser.Courses = Courses
	return curUser
}
