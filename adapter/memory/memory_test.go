package memory

import (
	"testing"

	"github.com/sankethkini/StudentData/domain/course"
	"github.com/sankethkini/StudentData/domain/user"
)

var alltests = []user.User{
	{
		Fname:  "sujan",
		RollNo: "44s",
		Age:    12,
		Adress: "2 232 ",
	},
	{
		Fname:  "sumanth",
		RollNo: "44q",
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
		Fname:   "sumanth",
		RollNo:  "44p",
		Age:     16,
		Adress:  "2 232 ",
		Courses: nil,
	},
}
var retrivetests = []struct {
	rollnum string
	check   bool
}{
	{
		rollnum: "44p",
		check:   true,
	},
	{
		rollnum: "46q",
		check:   false,
	},
}
var singleTest = user.User{
	Fname:   "amr",
	RollNo:  "ka31",
	Age:     13,
	Adress:  "2 sad 232 ",
	Courses: nil,
}

func TestSave(t *testing.T) {
	adpt := adapter{}
	err := adpt.Save(alltests)
	if err != nil {
		t.Error(err)
	}
	err = adpt.Save(singleTest)
	if err != nil {
		t.Error(err)
	}

	if len(adpt.Items) != len(alltests)+1 {
		t.Errorf("not all record inserted len: %v exp: %v", len(adpt.Items), len(alltests)+1)
	}
	for i, val := range adpt.Items {
		if i == 0 {
			continue
		}
		if adpt.Items[i-1].Fname > adpt.Items[i].Fname {
			t.Errorf("not sorted among records %v and %v", adpt.Items[i-1].Fname, val.Fname)
		} else if adpt.Items[i-1].Fname == adpt.Items[i].Fname && i != 0 && adpt.Items[i-1].RollNo > adpt.Items[i].RollNo {
			t.Errorf("not sorted among records %v and %v", adpt.Items[i-1].RollNo, val.RollNo)
		}
	}

}

func TesrRetrive(t *testing.T) {
	adpt := adapter{}
	err := adpt.Save(alltests)
	if err != nil {
		t.Error(err)
	}
	err = adpt.Save(singleTest)
	if err != nil {
		t.Error(err)
	}
	for _, val := range retrivetests {
		got := adpt.Retrive("rollnum", val.rollnum)
		if got != val.check {
			t.Errorf("not equal got : %v exp: %v", got, val.check)
		}
	}
}

func TestDelete(t *testing.T) {
	adpt := adapter{}

	err := adpt.Save(alltests)
	if err != nil {
		t.Error(err)
	}

	err = adpt.Save(singleTest)
	if err != nil {
		t.Error(err)
	}

	err = adpt.Delete("rollnum", "44p")
	if err != nil {
		t.Error(err)
	}

	err = adpt.Delete("rollnum", "44p")
	if err == nil {
		t.Error("record is not deleted properly")
	}
}
