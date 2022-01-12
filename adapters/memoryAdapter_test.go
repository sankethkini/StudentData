package adapters

import (
	"fmt"
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
	adapter := MemoryAdapter{}
	err := adapter.Save(alltests)
	if err != nil {
		t.Error(err)
	}
	err = adapter.Save(singleTest)
	if err != nil {
		t.Error(err)
	}
	for i, val := range adapter.Items {
		if i == 0 {
			continue
		}
		if adapter.Items[i-1].Fname > adapter.Items[i].Fname {
			t.Error("not sorted", adapter.Items[i-1].Fname, val.Fname)
		} else if adapter.Items[i-1].Fname == adapter.Items[i].Fname && i != 0 && adapter.Items[i-1].RollNo > adapter.Items[i].RollNo {
			t.Error("not sorted", adapter.Items[i-1].RollNo, val.RollNo)
		}
	}
	fmt.Println(adapter.Items)
}

func TesrRetrive(t *testing.T) {
	adapter := MemoryAdapter{}
	err := adapter.Save(alltests)
	if err != nil {
		t.Error(err)
	}
	err = adapter.Save(singleTest)
	if err != nil {
		t.Error(err)
	}
	for _, val := range retrivetests {
		got := adapter.Retrive("rollnum", val.rollnum)
		if got != val.check {
			t.Error("not equal", got, val.check)
		}
	}
}

func TestSort(t *testing.T) {
	adapter := MemoryAdapter{}

	err := adapter.Save(alltests)
	if err != nil {
		t.Error(err)
	}

	err = adapter.Save(singleTest)
	if err != nil {
		t.Error(err)
	}

	items, err := adapter.RetriveAll("address", 1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(items)
}

func TestDelete(t *testing.T) {
	adapter := MemoryAdapter{}

	err := adapter.Save(alltests)
	if err != nil {
		t.Error(err)
	}

	err = adapter.Save(singleTest)
	if err != nil {
		t.Error(err)
	}

	err = adapter.Delete("rollnum", "44p")
	if err != nil {
		t.Error(err)
	}

	err = adapter.Delete("rollnum", "44p")
	if err == nil {
		t.Error("record is not deleted properly")
	} else if err != RecordNotFound {
		t.Error("not a proper error message")
	}
}
