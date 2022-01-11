package adapters

import (
	"reflect"
	"testing"

	"github.com/sankethkini/StudentData/domain/course"
	"github.com/sankethkini/StudentData/domain/user"
)

var tests = []user.User{
	{
		Fname:  "sanket",
		RollNo: "44s",
		Age:    12,
		Adress: "2 232 ",
	},
	{
		Fname:  "sanket",
		RollNo: "44s",
		Age:    12,
		Adress: "2 232 ",
		Courses: []course.Course{
			{
				Name: "A",
				Code: "A",
			},
			{
				Name: "B",
				Code: "B",
			},
		},
	},
	{
		Fname:   "sanket",
		RollNo:  "44s",
		Age:     12,
		Adress:  "2 232 ",
		Courses: nil,
	},
}

func TestSaveAndRetrive(t *testing.T) {
	fileAdapter := FileAdapter{}
	err := fileAdapter.Save(tests)
	if err != nil {
		t.Error(err)
	}
	data, err := fileAdapter.RetriveAll("")
	if err != nil {
		t.Error(err)
	}
	if reflect.TypeOf(data) != reflect.TypeOf(tests) {
		t.Error("not of right type", reflect.TypeOf(data))
	}
	for i, val := range tests {
		if data[i].RollNo != tests[i].RollNo {
			t.Error("records not mathcing", val, data)
		}
	}
}
