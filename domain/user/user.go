package user

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/sankethkini/StudentData/domain/course"
)

type User struct {
	Fname   string
	Age     int
	Address string
	RollNo  string
	Courses []course.Course
}

func NewUser(fname string, age int, address string, rollNo string, courses []course.Course) User {
	curUser := User{}
	curUser.Fname = fname
	curUser.RollNo = rollNo
	curUser.Age = age
	curUser.Address = address
	curUser.Courses = courses
	return curUser
}

func (user User) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Fname, validation.Required, validation.Length(2, 50)),
		validation.Field(&user.Age, validation.Required, validation.Min(1)),
		validation.Field(&user.Address, validation.Required, validation.Length(3, 100)),
		validation.Field(&user.RollNo, validation.Length(3, 20)),
		validation.Field(&user.Courses, validation.Required),
	)
}
