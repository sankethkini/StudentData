package user

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/sankethkini/StudentData/domain/course"
)

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

func (user User) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Fname, validation.Required, validation.Length(2, 50)),
		validation.Field(&user.Age, validation.Required, validation.Min(1)),
		validation.Field(&user.Adress, validation.Required, validation.Length(3, 100)),
		validation.Field(&user.RollNo, validation.Length(3, 20)),
		validation.Field(&user.Courses, validation.Required),
	)
}
