package memory

import (
	"testing"

	"github.com/sankethkini/StudentData/constants"
	"github.com/sankethkini/StudentData/domain/course"
	"github.com/sankethkini/StudentData/domain/user"
)

var alltests = []user.User{
	{
		Fname:   "sujan",
		RollNo:  "44s",
		Age:     12,
		Address: "2 232 ",
	},
	{
		Fname:   "sumanth",
		RollNo:  "44q",
		Age:     12,
		Address: "2 232 ",
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
		Address: "2 232 ",
		Courses: nil,
	},
}

var singleTest = user.User{
	Fname:   "amr",
	RollNo:  "ka31",
	Age:     13,
	Address: "2 sad 232 ",
	Courses: nil,
}

var singleTest2 = user.User{
	Fname:   "kmr",
	RollNo:  "ka33",
	Age:     13,
	Address: "2 sad 232 ",
	Courses: nil,
}

func TestSave(t *testing.T) {
	adpt, err := NewMemory()
	if err != nil {
		t.Error(err)
	}
	err = adpt.Save(alltests...)
	if err != nil {
		t.Error(err)
	}

	err = addSingles(adpt)
	if err != nil {
		t.Error(err)
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

func addSingles(adpt *Adapter) error {
	err := adpt.Save(singleTest2)
	if err != nil {
		return err
	}

	err = adpt.Save(singleTest)
	if err != nil {
		return err
	}
	return nil
}

func TesrRetrieve(t *testing.T) {
	retrivetests := []struct {
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

	adpt, err := NewMemory()
	if err != nil {
		t.Error(err)
	}

	err = adpt.Save(alltests...)
	if err != nil {
		t.Error(err)
	}
	err = adpt.Save(singleTest)
	if err != nil {
		t.Error(err)
	}
	for _, val := range retrivetests {
		got := adpt.Retrieve("rollnum", val.rollnum)
		if got != val.check {
			t.Errorf("not equal got : %v exp: %v", got, val.check)
		}
	}
}

func TestDelete(t *testing.T) {
	adpt, err := NewMemory()
	if err != nil {
		t.Error(err)
	}

	err = adpt.Save(alltests...)
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
		if data[i-1].RollNo < data[i].RollNo {
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

func isDataaddressorted(data []user.User) bool {
	for i := 1; i < len(data); i++ {
		if data[i-1].Address > data[i].Address {
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

	mem, _ := NewMemory()
	for i, val := range testsForDisplay {
		usr := user.User{}
		usr.Fname = val.fname
		usr.Address = val.address
		usr.Age = val.age
		usr.Courses = constants.AllCourses[:4]
		usr.RollNo = val.rollnum
		err := mem.Save(usr)
		if err != nil {
			t.Error(i, err)
		}
	}

	// check for age.
	data, err := mem.RetriveAll("age", 1)
	if err != nil {
		t.Error(err)
	}
	if !isDataAgesorted(data) {
		t.Error("not sorted according to age")
	}

	// check for rollnum.
	data, err = mem.RetriveAll("rollnum", 2)
	if err != nil {
		t.Error(err)
	}
	if !isDataRollnumsorted(data) {
		t.Error("not sorted according to rollnum")
	}

	// check for name.
	data, err = mem.RetriveAll("name", 1)
	if err != nil {
		t.Error(err)
	}
	if !isDataNamesorted(data) {
		t.Error("not sorted according to rollnum")
	}

	// check for address.
	data, err = mem.RetriveAll("address", 1)
	if err != nil {
		t.Error(err)
	}
	if !isDataaddressorted(data) {
		t.Error("not sorted according to rollnum")
	}
}

func TestRetrive(t *testing.T) {
	tests := []struct {
		fname   string
		rollnum string
		address string
		age     int
		res     bool
	}{
		{fname: "sanket", rollnum: "t137", address: "abcd, defgh, jiop", age: 17, res: true},
	}
	mem, _ := NewMemory()
	for _, val := range tests {
		u := user.NewUser(val.fname, val.age, val.address, val.rollnum, nil)
		_ = mem.Save(u)
		got := mem.Retrieve("rollnum", val.rollnum)
		if got != val.res {
			t.Errorf("exp: %v got %v", val.res, got)
		}
	}
}
