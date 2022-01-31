package application

import (
	"testing"

	"github.com/sankethkini/StudentData/constants"
	"github.com/sankethkini/StudentData/domain/user"
)

var testForValidator = []struct {
	fname    string
	rollnum  string
	address  string
	age      int
	expected bool
}{
	{fname: "sanket", rollnum: "t77", address: "abcd, defgh, jiop", age: 17, expected: false},
	{fname: "", rollnum: "t77", address: "abcd, defgh, jiop", age: 17, expected: true},
	{fname: "sanket3", rollnum: "t77", address: "abcd, defgh, jiop", age: 18, expected: true},
}

// testing add function.
func TestAdd(t *testing.T) {
	app, err := NewApp()
	if err != nil {
		t.Errorf("cannot start app %v", err)
	}
	for _, val := range testForValidator {
		usr := user.NewUser(val.fname, val.age, val.address, val.rollnum, constants.AllCourses[:4])

		_, err := app.Add(usr)
		if val.expected == false && err != nil {
			t.Errorf("got unexpected error got %v", err)
		}
		if val.expected == true && err == nil {
			t.Errorf("didn't get any error got %v", err)
		}
	}
}

// check for age sorting.
func isDataAgesorted(data []user.User) bool {
	for i := 1; i < len(data); i++ {
		if data[i-1].Age > data[i].Age {
			return false
		}
	}
	return true
}

// check for rollnum sorting.
func isDataRollnumsorted(data []user.User) bool {
	for i := 1; i < len(data); i++ {
		if data[i-1].RollNo > data[i].RollNo {
			return false
		}
	}
	return true
}

// check for namewise sorting.
func isDataNamesorted(data []user.User) bool {
	for i := 1; i < len(data); i++ {
		if data[i-1].Fname > data[i].Fname {
			return false
		}
	}
	return true
}

// testing display with required sort.
func TestDisplay(t *testing.T) {
	testsForDisplay := []struct {
		fname   string
		rollnum string
		address string
		age     int
	}{
		{fname: "sanket", rollnum: "s77", address: "abcd, defgh, jiop", age: 17},
		{fname: "sanket", rollnum: "s78", address: "abcd, defgh, jiop", age: 18},
		{fname: "sumanth", rollnum: "s79", address: "abcd, defgh, jiop", age: 28},
		{fname: "sanket", rollnum: "s90", address: "abcd, defgh, jiop", age: 18},
		{fname: "sanket22", rollnum: "s88", address: "abcd, defgh, jiop", age: 16},
	}

	app, err := NewApp()
	if err != nil {
		t.Errorf("cannot start app %v", err)
	}

	for i, val := range testsForDisplay {
		usr := user.NewUser(val.fname, val.age, val.address, val.rollnum, constants.AllCourses[:4])

		_, err1 := app.Add(usr)
		if err1 != nil {
			t.Error(i, err)
		}
	}

	// check for age.
	mp := make(map[string]interface{})
	mp["field"] = "age"
	mp["order"] = 1

	data, err := app.Display(mp)
	if err != nil {
		t.Error(err)
	}
	if !isDataAgesorted(data) {
		t.Error("not sorted according to age")
	}

	// check for rollnum.
	mp["field"] = "rollnum"
	mp["order"] = 1
	data, err = app.Display(mp)
	if err != nil {
		t.Error(err)
	}
	if !isDataRollnumsorted(data) {
		t.Error("not sorted according to rollnum")
	}

	// check for name.
	mp["field"] = "name"
	mp["order"] = 1
	data, err = app.Display(mp)
	if err != nil {
		t.Error(err)
	}
	if !isDataNamesorted(data) {
		t.Error("not sorted according to rollnum")
	}
}

// testing delete function.
func TestForDelete(t *testing.T) {
	app, err := NewApp()
	if err != nil {
		t.Errorf("cannot start app %v", err)
	}

	usr := user.NewUser("name", 18, "asda sdas dasd asd a", "rs002", constants.AllCourses[:4])

	_, err = app.Add(usr)
	if err != nil {
		t.Errorf("cannot add user %v", err)
	}

	mp := make(map[string]interface{})
	mp["rollnum"] = "rs002"

	_, err = app.Delete(mp)
	if err != nil {
		t.Error(err)
	}

	_, err = app.Delete(mp)
	if err == nil {
		t.Error("record is not deleted properly")
	}
}

func TestSave(t *testing.T) {
	app, err := NewApp()
	if err != nil {
		t.Errorf("cannot start app %v", err)
	}

	_, err = app.Save()
	if err != nil {
		t.Errorf("error")
	}
}
