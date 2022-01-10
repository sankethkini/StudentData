package application

import (
	"testing"
)

var testForValidator = []struct {
	fname    string
	rollnum  string
	adress   string
	age      int
	expected error
}{
	{fname: "sanket", rollnum: "t77", adress: "abcd, defgh, jiop", age: 18, expected: nil},
	{fname: "", rollnum: "t77", adress: "abcd, defgh, jiop", age: 18, expected: NoNameErr},
	{fname: "sanket", rollnum: "", adress: "", age: 18, expected: NoRollNum},
	{fname: "sanket", rollnum: "t77", adress: "abcd, defgh, jiop", age: 0, expected: AgeErr},
	{fname: "sanket", rollnum: "t77", adress: "abcd, defgh, jiop", age: -18, expected: AgeErr},
}

func TestDataValidator(t *testing.T) {
	for _, val := range testForValidator {
		userdata := make(map[string]interface{})
		userdata["fname"] = val.fname
		userdata["rollnum"] = val.rollnum
		userdata["address"] = val.adress
		userdata["age"] = val.age
		err := InputValidator(userdata)
		if err != val.expected {
			t.Error("error in testcase #1", val.expected, err)
		}
	}
}
