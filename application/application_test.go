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
	expected error
}{
	{fname: "sanket", rollnum: "t77", adress: "abcd, defgh, jiop", age: 17, expected: nil},
	{fname: "", rollnum: "t77", adress: "abcd, defgh, jiop", age: 18, expected: NoNameErr},
	{fname: "sanket1", rollnum: "", adress: "", age: 18, expected: NoRollNumErr},
	{fname: "sanket2", rollnum: "t77", adress: "abcd, defgh, jiop", age: 0, expected: AgeErr},
	{fname: "sanket3", rollnum: "t77", adress: "abcd, defgh, jiop", age: -18, expected: AgeErr},
	{fname: "sanket4", rollnum: "t78", adress: "abcd, defgh, jiop", age: 18, expected: nil},
	{fname: "sanket22", rollnum: "t79", adress: "abcd, defgh, jiop", age: 16, expected: nil},
	{fname: "sanket4", rollnum: "t788", adress: "", age: 18, expected: NoAddressErr},
}

//testing validators
func TestDataValidator(t *testing.T) {

	for _, val := range testForValidator {
		userdata := make(map[string]interface{})
		userdata["fname"] = val.fname
		userdata["rollnum"] = val.rollnum
		userdata["address"] = val.adress
		userdata["age"] = val.age

		err := inputValidator(userdata)
		if err != val.expected {
			t.Errorf("error in validation exp:%v got: %v", val.expected, err)
		}
	}
}

//testing add function
func TestAdd(t *testing.T) {

	for _, val := range testForValidator {

		userdata := make(map[string]interface{})
		userdata["fname"] = val.fname
		userdata["rollnum"] = val.rollnum
		userdata["address"] = val.adress
		userdata["age"] = val.age
		userdata["courses"] = constants.AllCourses[:4]

		_, err := Add(userdata)
		if err != nil {
			if err != val.expected {
				t.Errorf("error in adding exp:%v got %v", val.expected, err)
			}
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

		userdata := make(map[string]interface{})
		userdata["fname"] = val.fname
		userdata["rollnum"] = val.rollnum
		userdata["address"] = val.adress
		userdata["age"] = val.age
		userdata["courses"] = constants.AllCourses[:4]

		data, err := Add(userdata)
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
