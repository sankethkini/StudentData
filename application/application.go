package application

import (
	"fmt"
	"os"

	"github.com/sankethkini/StudentData/adapter"
	"github.com/sankethkini/StudentData/adapter/file"
	"github.com/sankethkini/StudentData/adapter/memory"
	"github.com/sankethkini/StudentData/domain/user"
)

//alisaing map
type data map[string]interface{}

//global access
var adpt *adapter.Adapter

//init function called one time
func init() {

	//intializing file adapter
	fileAdapter, err := file.NewFileAdapter()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//intializing memory adapter
	memoryAdapter, err := memory.NewMemory()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//intializing adapter
	adpt, err = adapter.NewAdapter(fileAdapter, memoryAdapter)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileToMemory()
}

func fileToMemory() {
	users, err := adpt.FileAdapter.RetriveAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = adpt.MemoryAdapter.Save(users...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//function to return message
func createMessage(m string) []data {
	var msg []data
	mp := make(map[string]interface{})
	mp["message"] = m
	msg = append(msg, mp)
	return msg
}

//validations for rollnumber exists
func checkForRoll(rollnum string) error {

	//checking existence of simialr rollnum
	isExists := adpt.MemoryAdapter.Retrive("rollnum", rollnum)
	if isExists {
		return RollExistsErr
	}
	return nil
}

//function to add user
func Add(user user.User) ([]data, error) {

	//checking validity of user input
	validationErr := user.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	//checking for rollnum existence
	isRollexists := checkForRoll(user.RollNo)
	if isRollexists != nil {
		return nil, RollExistsErr
	}

	//adding user
	err := adpt.MemoryAdapter.Save(user)
	if err != nil {
		return nil, err
	}

	//message
	msg := createMessage("user added successfuly")
	return msg, err
}

//function to display
func Display(userData data) ([]user.User, error) {

	field := userData["field"].(string)
	order := userData["order"].(int)

	//retriving all data
	items, err := adpt.MemoryAdapter.RetriveAll(field, order)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func Delete(userData data) ([]data, error) {
	roll := userData["rollnum"].(string)

	//deleting user
	err := adpt.MemoryAdapter.Delete("rollnum", roll)
	if err != nil {
		return nil, err
	}

	msg := createMessage("user deleted successfuly")
	return msg, err
}

func Save() ([]data, error) {

	//retirving all by name
	allData, err := adpt.MemoryAdapter.RetriveAll("name", 1)
	if err != nil {
		return nil, err
	}

	//saving on disk
	err = adpt.FileAdapter.Save(allData)
	if err != nil {
		return nil, err
	}
	msg := createMessage("users saved to disk successfuly")
	return msg, err

}

//exit function
func Exit() {
	os.Exit(1)

}
