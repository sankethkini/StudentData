package constants

import (
	"github.com/sankethkini/StudentData/domain/course"
)

var AllCourses = []course.Course{
	{
		Name: "A",
		Code: "A",
	},
	{
		Name: "B",
		Code: "B",
	},
	{
		Name: "C",
		Code: "C",
	},
	{
		Name: "D",
		Code: "D",
	},
	{
		Name: "E",
		Code: "E",
	},
	{
		Name: "F",
		Code: "F",
	},
}

const (
	NoOfCourses  int = 4
	Filelocation     = "data.json"
)
