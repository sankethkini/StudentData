package memory

import (
	"fmt"
	"sort"

	"github.com/sankethkini/StudentData/domain/user"
)

//mapper for stroring existence of records
//items for saving data in a datastructure
type adapter struct {
	mapper map[string]bool
	Items  []user.User
}

//checking if a perticular record exists or not
func (adpt *adapter) Retrive(field string, value string) bool {
	return adpt.mapper[value]
}

//Search for right index
func (adpt *adapter) searchForIndex(user user.User) int {
	return sort.Search(len(adpt.Items), func(i int) bool {
		if adpt.Items[i].Fname > user.Fname {
			return true
		} else if adpt.Items[i].Fname == user.Fname {
			return adpt.Items[i].Fname < user.Fname
		} else {
			return false
		}
	})
}

//insert into map that record exists
func (adpt *adapter) insertIntoMap(nums ...string) {

	if adpt.mapper == nil {
		adpt.mapper = make(map[string]bool)
	}

	for _, val := range nums {
		adpt.mapper[val] = true
	}
}

//insertUser for entry of single record
func (adpt *adapter) insertUser(a interface{}) {

	cur := a.(user.User)
	if len(adpt.Items) == 0 {
		adpt.Items = append(adpt.Items, cur)
	} else {
		i := adpt.searchForIndex(cur)
		if i >= len(adpt.Items) {
			adpt.Items = append(adpt.Items, cur)
		} else {
			adpt.Items = append(adpt.Items[:i+1], adpt.Items[i:]...)
			adpt.Items[i] = cur
		}

	}
	adpt.insertIntoMap(cur.RollNo)
}

//arraysorting all records entered once
func (adpt *adapter) arraysort() {

	sort.Slice(adpt.Items, func(i int, j int) bool {
		if adpt.Items[i].Fname == adpt.Items[j].Fname {
			return adpt.Items[i].RollNo < adpt.Items[j].RollNo
		} else {
			return adpt.Items[i].Fname < adpt.Items[j].Fname
		}
	})

	var nums []string
	for _, val := range adpt.Items {
		nums = append(nums, val.RollNo)
	}
	adpt.insertIntoMap(nums...)
}

//save function
//insertion sort for single entry
//arraysorting if more than one record entry
func (adpt *adapter) Save(item ...user.User) error {

	if len(item) == 1 {
		adpt.insertUser(item[0])
	} else {
		adpt.Items = append(adpt.Items, item...)
		adpt.arraysort()
	}

	return nil
}

//reversing whole slice
func reverse(arr []user.User) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

//RetriveAll function to retrive all records
//sorting based on field
func (adpt *adapter) RetriveAll(field string, order int) ([]user.User, error) {

	var allusers []user.User
	allusers = append(allusers, adpt.Items...)

	switch field {

	case "name":
		sort.Slice(allusers, func(i, j int) bool {
			if allusers[i].Fname < allusers[j].Fname {
				return true
			} else if allusers[i].Fname == allusers[j].Fname {
				return allusers[i].RollNo < allusers[j].RollNo
			}
			return false
		})

	case "rollnum":
		sort.Slice(allusers, func(i, j int) bool {
			return allusers[i].RollNo < allusers[j].RollNo
		})

	case "address":
		sort.Slice(allusers, func(i, j int) bool {
			return allusers[i].Adress < allusers[j].Adress
		})

	case "age":
		sort.Slice(allusers, func(i, j int) bool {
			return allusers[i].Age < allusers[j].Age
		})

	}

	if order == 1 {
		return allusers, nil
	}
	reverse(allusers)
	return allusers, nil
}

//Delete function deletes a perticular record
//checks for a record if exists then deletes it

func (adpt *adapter) Delete(field string, value string) error {
	if !adpt.mapper[value] {
		return fmt.Errorf("cannot delete the record %w", RecordNotFoundErr)
	} else {

		index := -1
		for i := 0; i < len(adpt.Items); i++ {
			if adpt.Items[i].RollNo == value {
				index = i
			}
		}
		if len(adpt.Items) == 1 {
			adpt.Items = []user.User{}
		} else if index == len(adpt.Items)-1 {
			adpt.Items = adpt.Items[:index-1]
		} else {
			adpt.Items = append(adpt.Items[:index], adpt.Items[index+1:]...)
		}
		delete(adpt.mapper, value)
	}
	return nil
}

//constructor
func NewMemory() (*adapter, error) {
	return &adapter{}, nil
}
