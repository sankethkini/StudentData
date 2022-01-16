package application

import (
	"fmt"
	"os"

	"github.com/sankethkini/StudentData/adapter"
	"github.com/sankethkini/StudentData/adapter/file"
	"github.com/sankethkini/StudentData/adapter/memory"
	"github.com/sankethkini/StudentData/domain/course"
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
	err = adpt.MemoryAdapter.Save(users)
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

//validators of input
func inputValidator(userdata data) error {

	//checking name
	curname := userdata["fname"].(string)
	if len(curname) < 2 {
		return NoNameErr
	}

	//checking rollnumber
	curroll := userdata["rollnum"].(string)
	if len(curroll) < 2 {
		return NoRollNum
	}

	//checking address
	curadd := userdata["address"].(string)
	if len(curadd) < 2 {
		return NoAddress
	}

	//checking age
	curage := userdata["age"].(int)
	if curage <= 0 || curage >= 120 {
		return AgeErr
	}

	return nil
}

//validations for rollnumber exists
func checkForRoll(rollnum string) error {

	//checking existence of simialr rollnum
	isExists := adpt.MemoryAdapter.Retrive("rollnum", rollnum)
	if isExists {
		return RollExists
	}
	return nil
}

//function to add user
func Add(userdata data) ([]data, error) {

	//checking validity of user input
	validationerr := inputValidator(userdata)
	if validationerr != nil {
		return nil, validationerr
	}

	//checking for rollnum existence
	rollnum := userdata["rollnum"].(string)
	isRollexists := checkForRoll(rollnum)
	if isRollexists != nil {
		return nil, isRollexists
	}

	//adding data into user struct
	curuser := user.User{}
	curuser.Fname = userdata["fname"].(string)
	curuser.RollNo = userdata["rollnum"].(string)
	curuser.Age = userdata["age"].(int)
	curuser.Adress = userdata["address"].(string)
	curuser.Courses = userdata["courses"].([]course.Course)

	//adding user
	err := adpt.MemoryAdapter.Save(curuser)
	if err != nil {
		return nil, err
	}

	//message
	msg := createMessage("user added successfuly")
	return msg, err
}

//function to display
func Display(userdata data) ([]user.User, error) {

	field := userdata["field"].(string)
	order := userdata["order"].(int)

	//retriving all data
	items, err := adpt.MemoryAdapter.RetriveAll(field, order)
	if err != nil {
		return nil, err
	}
	allitems := items.([]user.User)
	return allitems, nil
}

func Delete(userdata data) ([]data, error) {
	roll := userdata["rollnum"].(string)

	//deleting user
	err := adpt.MemoryAdapter.Delete("rollnum", roll)
	if err != nil {
		return nil, err
	}

	msg := createMessage("user deleted successfuly")
	return msg, err
}

func Save(data) ([]data, error) {

	//retirving all by name
	alldata, err := adpt.MemoryAdapter.RetriveAll("name", 1)
	if err != nil {
		return nil, err
	}

	//saving on disk
	err = adpt.FileAdapter.Save(alldata)
	if err != nil {
		return nil, err
	}
	msg := createMessage("users saved to disk successfuly")
	return msg, err

}

//exit function
func Exit(data) ([]data, error) {
	os.Exit(1)
	return nil, nil
}
