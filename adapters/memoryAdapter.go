package adapters

import (
	"sort"

	"github.com/sankethkini/StudentData/domain/user"
)

//memory adapter struct
//mapper for stroring existence of records
//items for saving data in a datastructure
type MemoryAdapter struct {
	mapper map[string]bool
	Items  []user.User
}

//swap function
func swap(a *user.User, b *user.User) {
	temp := *a
	*a = *b
	*b = temp
}

//checking if a perticular record exists or not
func (adapter *MemoryAdapter) Retrive(field string, value string) bool {
	return adapter.mapper[value]
}

//insertion sort for entry of single record
func (adapter *MemoryAdapter) insertionSort(a interface{}) {

	cur := a.(user.User)
	if len(adapter.Items) == 0 {
		adapter.Items = append(adapter.Items, cur)
	} else {
		adapter.Items = append(adapter.Items, cur)
		last := len(adapter.Items) - 1
		for last > 0 {
			if adapter.Items[last-1].Fname < adapter.Items[last].Fname {
				break
			} else if adapter.Items[last-1].Fname > adapter.Items[last].Fname {
				swap(&adapter.Items[last-1], &adapter.Items[last])
			} else {
				if adapter.Items[last-1].RollNo > adapter.Items[last].RollNo {
					swap(&adapter.Items[last-1], &adapter.Items[last])
				} else {
					break
				}
			}
			last -= 1
		}
	}

	if adapter.mapper == nil {
		adapter.mapper = make(map[string]bool)
	}
	//mapping rollnum
	adapter.mapper[cur.RollNo] = true
}

//arraysorting all records entered once
func (adapter *MemoryAdapter) arraysort() {

	sort.Slice(adapter.Items, func(i int, j int) bool {
		if adapter.Items[i].Fname == adapter.Items[j].Fname {
			return adapter.Items[i].RollNo < adapter.Items[j].RollNo
		} else {
			return adapter.Items[i].Fname < adapter.Items[j].Fname
		}
	})

	if adapter.mapper == nil {
		adapter.mapper = make(map[string]bool)
	}

	for _, val := range adapter.Items {
		adapter.mapper[val.RollNo] = true
	}
}

//save function
//insertion sort for single entry
//arraysorting if more than one record entry
func (adapter *MemoryAdapter) Save(item interface{}) error {

	switch item.(type) {
	case user.User:
		adapter.insertionSort(item)
	case []user.User:
		curItems := item.([]user.User)
		adapter.Items = append(adapter.Items, curItems...)
		adapter.arraysort()
	default:
		return NotARightType
	}

	return nil
}

//reversing whole slice
func reverse(arr []user.User) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

//function to retrive all records
//sorting based on field
func (adapter *MemoryAdapter) RetriveAll(field string, order int) (interface{}, error) {

	switch field {
	case "name":
		if order == 1 {
			adapter.arraysort()
			return adapter.Items, nil
		} else {
			reverse(adapter.Items)
			return adapter.Items, nil
		}

	case "rollnum":
		sort.Slice(adapter.Items, func(i, j int) bool {
			return adapter.Items[i].RollNo < adapter.Items[j].RollNo
		})
		if order == 1 {
			return adapter.Items, nil
		} else {
			reverse(adapter.Items)
			return adapter.Items, nil
		}

	case "address":
		sort.Slice(adapter.Items, func(i, j int) bool {
			return adapter.Items[i].Adress < adapter.Items[j].Adress
		})
		if order == 1 {
			return adapter.Items, nil
		} else {
			reverse(adapter.Items)
			return adapter.Items, nil
		}

	case "age":
		sort.Slice(adapter.Items, func(i, j int) bool {
			return adapter.Items[i].Age < adapter.Items[j].Age
		})
		if order == 1 {
			return adapter.Items, nil
		} else {
			reverse(adapter.Items)
			return adapter.Items, nil
		}
	}

	return adapter.Items, nil
}

//checks for a record if exists then deletes it
func (adapter *MemoryAdapter) Delete(field string, value string) error {
	if !adapter.mapper[value] {
		return RecordNotFound
	} else {

		index := -1
		for i := 0; i < len(adapter.Items); i++ {
			if adapter.Items[i].RollNo == value {
				index = i
			}
		}

		if index == len(adapter.Items)-1 {
			adapter.Items = adapter.Items[:index-1]
		} else {
			adapter.Items = append(adapter.Items[:index], adapter.Items[index+1:]...)
		}
		adapter.mapper[value] = false
	}
	return nil
}
