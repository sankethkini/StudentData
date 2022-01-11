package adapters

import (
	"sort"

	"github.com/sankethkini/StudentData/domain/user"
)

type MemoryAdapter struct {
	mapper map[string]bool
	Items  []user.User
}

func swap(a *user.User, b *user.User) {
	temp := *a
	*a = *b
	*b = temp
}
func (adapter *MemoryAdapter) Retrive(field string, value string) bool {
	return adapter.mapper[value]
}

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
	adapter.mapper[cur.RollNo] = true
}
func (adapter *MemoryAdapter) arraysort(item interface{}) {
	curItems := item.([]user.User)
	for _, val := range curItems {
		adapter.Items = append(adapter.Items, val)
	}
	sort.Slice(adapter.Items, func(i int, j int) bool {
		if adapter.Items[i].Fname == adapter.Items[j].Fname {
			return adapter.Items[i].RollNo < adapter.Items[j].RollNo
		} else {
			return adapter.Items[i].Fname < adapter.Items[j].Fname
		}
	})
}
func (adapter *MemoryAdapter) Save(item interface{}) error {
	switch item.(type) {
	case user.User:
		adapter.insertionSort(item)
	case []user.User:
		adapter.arraysort(item)
	default:
		return NotARightType
	}
	return nil
}
