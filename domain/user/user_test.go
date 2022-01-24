package user

import (
	"testing"

	"github.com/sankethkini/StudentData/constants"
)

var tests = []struct {
	fname    string
	rollnum  string
	adress   string
	age      int
	expected bool
}{
	{fname: "sanket", rollnum: "t77", adress: "abcd, defgh, jiop", age: 17, expected: false},
	{fname: "", rollnum: "t787", adress: "abcd, defgh, jiop", age: 18, expected: true},
	{fname: "sanket1", rollnum: "", adress: "", age: 18, expected: true},
	{fname: "sanket2", rollnum: "t776", adress: "abcd, defgh, jiop", age: 0, expected: true},
	{fname: "sanket3", rollnum: "t777", adress: "abcd, defgh, jiop", age: -1, expected: true},
	{fname: "sanket3", rollnum: "t77", adress: "abcd, defgh, jiop", age: 18, expected: false},
	{fname: "sanket4", rollnum: "t78", adress: "abcd, defgh, jiop", age: 18, expected: false},
	{fname: "sanket22", rollnum: "t79", adress: "abcd, defgh, jiop", age: 16, expected: false},
	{fname: "sanket4", rollnum: "t788", adress: "", age: 18, expected: true},
}

func TestValidator(t *testing.T) {
	for i, val := range tests {
		usr := NewUser(val.fname, val.age, val.adress, val.rollnum, constants.AllCourses[:4])
		err := usr.Validate()

		if val.expected == false && err != nil {
			t.Errorf("#%d got unexpected error got %v", i, err)
		}
		if val.expected == true && err == nil {
			t.Errorf("#%d didn't get any error got %v", i, err)
		}
	}
}
