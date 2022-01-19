package file

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
		Fname:  "sanket1",
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
		Fname:   "sanket2",
		RollNo:  "44s",
		Age:     12,
		Adress:  "2 232 ",
		Courses: nil,
	},
}

func TestSaveAndRetrive(t *testing.T) {

	adapter, _ := NewFileAdapter()
	err := adapter.Save(tests)
	if err != nil {
		t.Error(err)
	}

	data, err := adapter.RetriveAll()
	if err != nil {
		t.Error(err)
	}
	if reflect.TypeOf(data) != reflect.TypeOf(tests) {
		t.Errorf("not of right type %v", reflect.TypeOf(data))
	}

	for i, val := range tests {
		if data[i].RollNo != tests[i].RollNo {
			t.Errorf("records not mathcing %v and %v", val, data)
		}
	}
}
