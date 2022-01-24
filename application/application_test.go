package application

import (
	"fmt"
	"testing"

	"github.com/sankethkini/StudentData/constants"
	"github.com/sankethkini/StudentData/domain/user"
)

var testForValidator = []struct {
	fname    string
	rollnum  string
	adress   string
	age      int
	expected bool
}{
	{fname: "sanket", rollnum: "t77", adress: "abcd, defgh, jiop", age: 17, expected: false},
	{fname: "", rollnum: "t77", adress: "abcd, defgh, jiop", age: 17, expected: true},
	{fname: "sanket3", rollnum: "t77", adress: "abcd, defgh, jiop", age: 18, expected: true},
}

//testing add function
func TestAdd(t *testing.T) {

	for _, val := range testForValidator {

		usr := user.NewUser(val.fname, val.age, val.adress, val.rollnum, constants.AllCourses[:4])

		_, err := Add(usr)
		if val.expected == false && err != nil {
			t.Errorf("got unexpected error got %v", err)
		}
		if val.expected == true && err == nil {
			t.Errorf("didn't get any error got %v", err)
		}

	}
}

//check for age sorting
func isDataAgesorted(data []user.User) bool {
	for i := 1; i < len(data); i++ {
		if data[i-1].Age > data[i].Age {
			return false
		}
	}
	return true
}

//check for rollnum sorting
func isDataRollnumsorted(data []user.User) bool {
	for i := 1; i < len(data); i++ {
		if data[i-1].RollNo > data[i].RollNo {
			return false
		}
	}
	return true
}

//check for namewise sorting
func isDataNamesorted(data []user.User) bool {
	for i := 1; i < len(data); i++ {
		if data[i-1].Fname > data[i].Fname {
			return false
		}
	}
	return true
}

//testing display with required sort
func TestDisplay(t *testing.T) {

	var testsForDisplay = []struct {
		fname   string
		rollnum string
		adress  string
		age     int
	}{
		{fname: "sanket", rollnum: "s77", adress: "abcd, defgh, jiop", age: 17},
		{fname: "sanket", rollnum: "s78", adress: "abcd, defgh, jiop", age: 18},
		{fname: "sumanth", rollnum: "s79", adress: "abcd, defgh, jiop", age: 28},
		{fname: "sanket", rollnum: "s90", adress: "abcd, defgh, jiop", age: 18},
		{fname: "sanket22", rollnum: "s88", adress: "abcd, defgh, jiop", age: 16},
	}

	for i, val := range testsForDisplay {

		usr := user.NewUser(val.fname, val.age, val.adress, val.rollnum, constants.AllCourses[:4])

		data, err := Add(usr)
		if err != nil {
			t.Error(2, i, err, data)
		}
	}

	//check for age
	mp := make(map[string]interface{})
	mp["field"] = "age"
	mp["order"] = 1
	data, err := Display(mp)
	if err != nil {
		t.Error(err)
	}
	if !isDataAgesorted(data) {
		t.Error("not sorted according to age")
	}

	//check for rollnum
	mp["field"] = "rollnum"
	mp["order"] = 1
	data, err = Display(mp)
	if err != nil {
		t.Error(err)
	}
	if !isDataRollnumsorted(data) {
		t.Error("not sorted according to rollnum")
	}

	//check for name
	mp["field"] = "name"
	mp["order"] = 1
	data, err = Display(mp)
	if err != nil {
		t.Error(err)
	}
	if !isDataNamesorted(data) {
		t.Error("not sorted according to rollnum")
	}

}

//testing delete function
func TestForDelete(t *testing.T) {
	mp := make(map[string]interface{})
	mp["rollnum"] = "s77"

	msg, err := Delete(mp)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(msg)

	_, err = Delete(mp)
	if err == nil {
		t.Error("record is not deleted properly")
	}
}

func TestSave(t *testing.T) {
	_, err := Save()

	if err != nil {
		t.Errorf("error")
	}

}
