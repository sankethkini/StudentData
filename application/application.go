package application

import (
	"fmt"
	"os"

	"github.com/sankethkini/StudentData/adapters"
	"github.com/sankethkini/StudentData/domain/course"
	"github.com/sankethkini/StudentData/domain/user"
)

//alisaing map
type data map[string]interface{}

//init function called one time
func init() {

	//intializing file adapter
	fileAdapter, err := adapters.Init(adapters.FILE)
	if err != nil {
		fmt.Println(err)
		Exit(nil)
	}

	//intializing memory adapter
	memoryAdapter, err := adapters.Init(adapters.MEMORY)
	if err != nil {
		fmt.Println(err)
		Exit(nil)
	}

	//retriving all data from disk
	diskdata, err := fileAdapter.RetriveAll("", 1)
	if err != nil {
		fmt.Println(err)
		Exit(nil)
	}

	//adding into memory
	err = memoryAdapter.Save(diskdata)
	if err != nil {
		fmt.Println(err)
		Exit(nil)
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
	//intializing memory adapter
	memoryAdapter, err := adapters.Init(adapters.MEMORY)
	if err != nil {
		return err
	}

	//checking existence of simialr rollnum
	isExists := memoryAdapter.Retrive("rollnum", rollnum)
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

	//saving in memory
	memoryAdapter, err := adapters.Init(adapters.MEMORY)
	if err != nil {
		return nil, err
	}

	//adding data into user struct
	curuser := user.User{}
	curuser.Fname = userdata["fname"].(string)
	curuser.RollNo = userdata["rollnum"].(string)
	curuser.Age = userdata["age"].(int)
	curuser.Adress = userdata["address"].(string)
	curuser.Courses = userdata["courses"].([]course.Course)

	//adding user
	err = memoryAdapter.Save(curuser)
	if err != nil {
		return nil, err
	}

	//message
	msg := createMessage("user added successfuly")
	return msg, err
}

//function to display
func Display(userdata data) ([]user.User, error) {
	//intializing memory adapter
	memoryAdapter, err := adapters.Init(adapters.MEMORY)
	if err != nil {
		return nil, err
	}

	field := userdata["field"].(string)
	order := userdata["order"].(int)

	//retriving all data
	items, err := memoryAdapter.RetriveAll(field, order)
	if err != nil {
		return nil, err
	}
	allitems := items.([]user.User)
	return allitems, nil
}

func Delete(userdata data) ([]data, error) {
	roll := userdata["rollnum"].(string)

	//intializing memory adapter
	memoryAdapter, err := adapters.Init(adapters.MEMORY)
	if err != nil {
		return nil, err
	}

	//deleting user
	err = memoryAdapter.Delete("rollnum", roll)
	if err != nil {
		return nil, err
	}

	msg := createMessage("user deleted successfuly")
	return msg, err
}

func Save(data) ([]data, error) {
	//intializing file adapter
	fileadapter, err := adapters.Init(adapters.FILE)
	if err != nil {
		return nil, err
	}

	//intializing memory adapter
	memoryadapter, err := adapters.Init(adapters.MEMORY)
	if err != nil {
		return nil, err
	}

	//retirving all by name
	alldata, err := memoryadapter.RetriveAll("name", 1)
	if err != nil {
		return nil, err
	}

	//saving on disk
	err = fileadapter.Save(alldata)
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
